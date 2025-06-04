package resourcemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type CreateServerGroupRequest struct {
  // {"en":"server group's name", "zh_CN":"错误信息或Success"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"region id(obtain the corresponding ID in the getRegionList interface)", "zh_CN":"地理信息（在获取区域列表接口获取对应的ID）"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"is enabled, true/false", "zh_CN":"是否可用, true/false"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
}

func (s CreateServerGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateServerGroupRequest) GoString() string {
  return s.String()
}

func (s *CreateServerGroupRequest) SetServergroupName(v string) *CreateServerGroupRequest {
  s.ServergroupName = &v
  return s
}

func (s *CreateServerGroupRequest) SetRegion(v string) *CreateServerGroupRequest {
  s.Region = &v
  return s
}

func (s *CreateServerGroupRequest) SetIsEnabled(v bool) *CreateServerGroupRequest {
  s.IsEnabled = &v
  return s
}

type CreateServerGroupResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *CreateServerGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateServerGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateServerGroupResponse) GoString() string {
  return s.String()
}

func (s *CreateServerGroupResponse) SetData(v *CreateServerGroupResponseData) *CreateServerGroupResponse {
  s.Data = v
  return s
}

func (s *CreateServerGroupResponse) SetCode(v int) *CreateServerGroupResponse {
  s.Code = &v
  return s
}

func (s *CreateServerGroupResponse) SetMessage(v string) *CreateServerGroupResponse {
  s.Message = &v
  return s
}

type CreateServerGroupResponseData struct {
  // {"en":"servergroup's ID", "zh_CN":"servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"错误信息或Success"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"region id", "zh_CN":"地理信息ID"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"region name", "zh_CN":"地理信息名称"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"is enabled, true/false", "zh_CN":"是否可用, true/false"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
}

func (s CreateServerGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateServerGroupResponseData) GoString() string {
  return s.String()
}

func (s *CreateServerGroupResponseData) SetServergroupId(v int) *CreateServerGroupResponseData {
  s.ServergroupId = &v
  return s
}

func (s *CreateServerGroupResponseData) SetServergroupName(v string) *CreateServerGroupResponseData {
  s.ServergroupName = &v
  return s
}

func (s *CreateServerGroupResponseData) SetRegion(v string) *CreateServerGroupResponseData {
  s.Region = &v
  return s
}

func (s *CreateServerGroupResponseData) SetRegionName(v string) *CreateServerGroupResponseData {
  s.RegionName = &v
  return s
}

func (s *CreateServerGroupResponseData) SetIsEnabled(v bool) *CreateServerGroupResponseData {
  s.IsEnabled = &v
  return s
}

type CreateServerGroupPaths struct {
}

func (s CreateServerGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateServerGroupPaths) GoString() string {
  return s.String()
}

type CreateServerGroupParameters struct {
}

func (s CreateServerGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateServerGroupParameters) GoString() string {
  return s.String()
}

type CreateServerGroupRequestHeader struct {
}

func (s CreateServerGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateServerGroupRequestHeader) GoString() string {
  return s.String()
}

type CreateServerGroupResponseHeader struct {
}

func (s CreateServerGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateServerGroupResponseHeader) GoString() string {
  return s.String()
}




type GetCvListRequest struct {
}

func (s GetCvListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCvListRequest) GoString() string {
  return s.String()
}

type GetCvListResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetCvListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetCvListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCvListResponse) GoString() string {
  return s.String()
}

func (s *GetCvListResponse) SetData(v []*GetCvListResponseData) *GetCvListResponse {
  s.Data = v
  return s
}

func (s *GetCvListResponse) SetCode(v int) *GetCvListResponse {
  s.Code = &v
  return s
}

func (s *GetCvListResponse) SetMessage(v string) *GetCvListResponse {
  s.Message = &v
  return s
}

type GetCvListResponseData struct     {
  // {"en":"Custom View id", "zh_CN":"Custom View的ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Custom View name", "zh_CN":"Custom Views的NAME"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetCvListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetCvListResponseData) GoString() string {
  return s.String()
}

func (s *GetCvListResponseData) SetId(v string) *GetCvListResponseData {
  s.Id = &v
  return s
}

func (s *GetCvListResponseData) SetName(v string) *GetCvListResponseData {
  s.Name = &v
  return s
}

type GetCvListPaths struct {
}

func (s GetCvListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCvListPaths) GoString() string {
  return s.String()
}

type GetCvListParameters struct {
}

func (s GetCvListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCvListParameters) GoString() string {
  return s.String()
}

type GetCvListRequestHeader struct {
}

func (s GetCvListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCvListRequestHeader) GoString() string {
  return s.String()
}

type GetCvListResponseHeader struct {
}

func (s GetCvListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCvListResponseHeader) GoString() string {
  return s.String()
}




type EnableHealthCheckWhitelistRequest struct {
  // {"en":"Dictionary value[disabled, auto, manual]. disabled: Disable the whitelist, all agents can perform detection tasks; auto: Automatically manage the whitelist, and only the agents in the whitelist will perform probing tasks. In this mode, when the effective time is reached, the new agent will be automatically added to the whitelist. So you need to implement this feature at your side; manual: Manually manage the whitelist, and only the agents in the whitelist will perform probing tasks. The agent list needs to be maintained manually.", "zh_CN":"字典值: disabled,auto,manual。disabled: 禁用白名单，所有探测器都可以执行探测任务；auto:自动管理白名单，只有白名单内的探测器才会执行探测任务。 在该模式下，到达有效时间后，新添加的探测器会自动加入白名单。所以需要在您的被探测机器上实现对应的逻辑；manual: 手动添加白名单，只有白名单内的探测器才会执行探测任务。 探测器列表需要您手动维护。"}
  Enable *string `json:"enable,omitempty" xml:"enable,omitempty" require:"true"`
  // {"en":"The probe IP list is only useful when enable is manual, and the IP must be the value in the <GetHealthCheckAgentList> interface", "zh_CN":"探测器IP列表仅在enable为manual时有用，IP需为<获取健康检查探测器列表>接口中的值"}
  Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s EnableHealthCheckWhitelistRequest) String() string {
  return tea.Prettify(s)
}

func (s EnableHealthCheckWhitelistRequest) GoString() string {
  return s.String()
}

func (s *EnableHealthCheckWhitelistRequest) SetEnable(v string) *EnableHealthCheckWhitelistRequest {
  s.Enable = &v
  return s
}

func (s *EnableHealthCheckWhitelistRequest) SetIps(v []*string) *EnableHealthCheckWhitelistRequest {
  s.Ips = v
  return s
}

type EnableHealthCheckWhitelistResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EnableHealthCheckWhitelistResponse) String() string {
  return tea.Prettify(s)
}

func (s EnableHealthCheckWhitelistResponse) GoString() string {
  return s.String()
}

func (s *EnableHealthCheckWhitelistResponse) SetData(v string) *EnableHealthCheckWhitelistResponse {
  s.Data = &v
  return s
}

func (s *EnableHealthCheckWhitelistResponse) SetCode(v int) *EnableHealthCheckWhitelistResponse {
  s.Code = &v
  return s
}

func (s *EnableHealthCheckWhitelistResponse) SetMessage(v string) *EnableHealthCheckWhitelistResponse {
  s.Message = &v
  return s
}

type EnableHealthCheckWhitelistPaths struct {
}

func (s EnableHealthCheckWhitelistPaths) String() string {
  return tea.Prettify(s)
}

func (s EnableHealthCheckWhitelistPaths) GoString() string {
  return s.String()
}

type EnableHealthCheckWhitelistParameters struct {
}

func (s EnableHealthCheckWhitelistParameters) String() string {
  return tea.Prettify(s)
}

func (s EnableHealthCheckWhitelistParameters) GoString() string {
  return s.String()
}

type EnableHealthCheckWhitelistRequestHeader struct {
}

func (s EnableHealthCheckWhitelistRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableHealthCheckWhitelistRequestHeader) GoString() string {
  return s.String()
}

type EnableHealthCheckWhitelistResponseHeader struct {
}

func (s EnableHealthCheckWhitelistResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EnableHealthCheckWhitelistResponseHeader) GoString() string {
  return s.String()
}




type GetSpecifiedServergroupRequest struct {
}

func (s GetSpecifiedServergroupRequest) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupRequest) GoString() string {
  return s.String()
}

type GetSpecifiedServergroupResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *GetSpecifiedServergroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetSpecifiedServergroupResponse) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupResponse) GoString() string {
  return s.String()
}

func (s *GetSpecifiedServergroupResponse) SetData(v *GetSpecifiedServergroupResponseData) *GetSpecifiedServergroupResponse {
  s.Data = v
  return s
}

func (s *GetSpecifiedServergroupResponse) SetCode(v int) *GetSpecifiedServergroupResponse {
  s.Code = &v
  return s
}

func (s *GetSpecifiedServergroupResponse) SetMessage(v string) *GetSpecifiedServergroupResponse {
  s.Message = &v
  return s
}

type GetSpecifiedServergroupResponseData struct {
  // {"en":"servergroup's ID", "zh_CN":"servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"错误信息或Success"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"region id", "zh_CN":"地理信息ID"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"region name", "zh_CN":"地理信息名称"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"is enabled, true/false", "zh_CN":"是否可用, true/false"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"server count", "zh_CN":"server的数量"}
  ServerCount *int `json:"serverCount,omitempty" xml:"serverCount,omitempty" require:"true"`
  // {"en":"server list", "zh_CN":"server的列表"}
  Servers []*GetSpecifiedServergroupResponseDataServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s GetSpecifiedServergroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupResponseData) GoString() string {
  return s.String()
}

func (s *GetSpecifiedServergroupResponseData) SetServergroupId(v int) *GetSpecifiedServergroupResponseData {
  s.ServergroupId = &v
  return s
}

func (s *GetSpecifiedServergroupResponseData) SetServergroupName(v string) *GetSpecifiedServergroupResponseData {
  s.ServergroupName = &v
  return s
}

func (s *GetSpecifiedServergroupResponseData) SetRegion(v string) *GetSpecifiedServergroupResponseData {
  s.Region = &v
  return s
}

func (s *GetSpecifiedServergroupResponseData) SetRegionName(v string) *GetSpecifiedServergroupResponseData {
  s.RegionName = &v
  return s
}

func (s *GetSpecifiedServergroupResponseData) SetIsEnabled(v bool) *GetSpecifiedServergroupResponseData {
  s.IsEnabled = &v
  return s
}

func (s *GetSpecifiedServergroupResponseData) SetServerCount(v int) *GetSpecifiedServergroupResponseData {
  s.ServerCount = &v
  return s
}

func (s *GetSpecifiedServergroupResponseData) SetServers(v []*GetSpecifiedServergroupResponseDataServers) *GetSpecifiedServergroupResponseData {
  s.Servers = v
  return s
}

type GetSpecifiedServergroupResponseDataServers struct     {
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*GetSpecifiedServergroupResponseDataServersHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
}

func (s GetSpecifiedServergroupResponseDataServers) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupResponseDataServers) GoString() string {
  return s.String()
}

func (s *GetSpecifiedServergroupResponseDataServers) SetIp(v string) *GetSpecifiedServergroupResponseDataServers {
  s.Ip = &v
  return s
}

func (s *GetSpecifiedServergroupResponseDataServers) SetIsEnabled(v bool) *GetSpecifiedServergroupResponseDataServers {
  s.IsEnabled = &v
  return s
}

func (s *GetSpecifiedServergroupResponseDataServers) SetServerName(v string) *GetSpecifiedServergroupResponseDataServers {
  s.ServerName = &v
  return s
}

func (s *GetSpecifiedServergroupResponseDataServers) SetServerId(v int) *GetSpecifiedServergroupResponseDataServers {
  s.ServerId = &v
  return s
}

func (s *GetSpecifiedServergroupResponseDataServers) SetHealthcheckers(v []*GetSpecifiedServergroupResponseDataServersHealthcheckers) *GetSpecifiedServergroupResponseDataServers {
  s.Healthcheckers = v
  return s
}

type GetSpecifiedServergroupResponseDataServersHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetSpecifiedServergroupResponseDataServersHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupResponseDataServersHealthcheckers) GoString() string {
  return s.String()
}

func (s *GetSpecifiedServergroupResponseDataServersHealthcheckers) SetId(v int) *GetSpecifiedServergroupResponseDataServersHealthcheckers {
  s.Id = &v
  return s
}

func (s *GetSpecifiedServergroupResponseDataServersHealthcheckers) SetName(v string) *GetSpecifiedServergroupResponseDataServersHealthcheckers {
  s.Name = &v
  return s
}

type GetSpecifiedServergroupPaths struct {
  // {"en":"servergroup's ID", "zh_CN":"servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
}

func (s GetSpecifiedServergroupPaths) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupPaths) GoString() string {
  return s.String()
}

func (s *GetSpecifiedServergroupPaths) SetServergroupId(v int) *GetSpecifiedServergroupPaths {
  s.ServergroupId = &v
  return s
}

type GetSpecifiedServergroupParameters struct {
}

func (s GetSpecifiedServergroupParameters) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupParameters) GoString() string {
  return s.String()
}

type GetSpecifiedServergroupRequestHeader struct {
}

func (s GetSpecifiedServergroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupRequestHeader) GoString() string {
  return s.String()
}

type GetSpecifiedServergroupResponseHeader struct {
}

func (s GetSpecifiedServergroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetSpecifiedServergroupResponseHeader) GoString() string {
  return s.String()
}




type CreateServerRequest struct {
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*CreateServerRequestHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
}

func (s CreateServerRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateServerRequest) GoString() string {
  return s.String()
}

func (s *CreateServerRequest) SetServerName(v string) *CreateServerRequest {
  s.ServerName = &v
  return s
}

func (s *CreateServerRequest) SetIsEnabled(v bool) *CreateServerRequest {
  s.IsEnabled = &v
  return s
}

func (s *CreateServerRequest) SetIp(v string) *CreateServerRequest {
  s.Ip = &v
  return s
}

func (s *CreateServerRequest) SetHealthcheckers(v []*CreateServerRequestHealthcheckers) *CreateServerRequest {
  s.Healthcheckers = v
  return s
}

type CreateServerRequestHealthcheckers struct     {
  // {"en":"health checker id, obtained from interface 'getHealthcheckerList'", "zh_CN":"监控检查ID，从”获取健康检查方法列表“接口获取"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s CreateServerRequestHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s CreateServerRequestHealthcheckers) GoString() string {
  return s.String()
}

func (s *CreateServerRequestHealthcheckers) SetId(v int) *CreateServerRequestHealthcheckers {
  s.Id = &v
  return s
}

type CreateServerResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *CreateServerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateServerResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateServerResponse) GoString() string {
  return s.String()
}

func (s *CreateServerResponse) SetData(v *CreateServerResponseData) *CreateServerResponse {
  s.Data = v
  return s
}

func (s *CreateServerResponse) SetCode(v int) *CreateServerResponse {
  s.Code = &v
  return s
}

func (s *CreateServerResponse) SetMessage(v string) *CreateServerResponse {
  s.Message = &v
  return s
}

type CreateServerResponseData struct {
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group region id", "zh_CN":"server group的地理ID"}
  ServergroupRegion *string `json:"servergroupRegion,omitempty" xml:"servergroupRegion,omitempty" require:"true"`
  // {"en":"server group region name", "zh_CN":"server group的地理名称"}
  ServergroupRegionName *string `json:"servergroupRegionName,omitempty" xml:"servergroupRegionName,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"server group的名称"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*CreateServerResponseDataHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
  // {"en":"The name of the associated ClbDomain", "zh_CN":"关联的ClbDomain的名称"}
  ClbDomains []*string `json:"clbDomains,omitempty" xml:"clbDomains,omitempty" require:"true" type:"Repeated"`
}

func (s CreateServerResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateServerResponseData) GoString() string {
  return s.String()
}

func (s *CreateServerResponseData) SetServergroupId(v int) *CreateServerResponseData {
  s.ServergroupId = &v
  return s
}

func (s *CreateServerResponseData) SetServergroupRegion(v string) *CreateServerResponseData {
  s.ServergroupRegion = &v
  return s
}

func (s *CreateServerResponseData) SetServergroupRegionName(v string) *CreateServerResponseData {
  s.ServergroupRegionName = &v
  return s
}

func (s *CreateServerResponseData) SetServergroupName(v string) *CreateServerResponseData {
  s.ServergroupName = &v
  return s
}

func (s *CreateServerResponseData) SetIp(v string) *CreateServerResponseData {
  s.Ip = &v
  return s
}

func (s *CreateServerResponseData) SetIsEnabled(v bool) *CreateServerResponseData {
  s.IsEnabled = &v
  return s
}

func (s *CreateServerResponseData) SetServerName(v string) *CreateServerResponseData {
  s.ServerName = &v
  return s
}

func (s *CreateServerResponseData) SetServerId(v int) *CreateServerResponseData {
  s.ServerId = &v
  return s
}

func (s *CreateServerResponseData) SetHealthcheckers(v []*CreateServerResponseDataHealthcheckers) *CreateServerResponseData {
  s.Healthcheckers = v
  return s
}

func (s *CreateServerResponseData) SetClbDomains(v []*string) *CreateServerResponseData {
  s.ClbDomains = v
  return s
}

type CreateServerResponseDataHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s CreateServerResponseDataHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s CreateServerResponseDataHealthcheckers) GoString() string {
  return s.String()
}

func (s *CreateServerResponseDataHealthcheckers) SetId(v int) *CreateServerResponseDataHealthcheckers {
  s.Id = &v
  return s
}

func (s *CreateServerResponseDataHealthcheckers) SetName(v string) *CreateServerResponseDataHealthcheckers {
  s.Name = &v
  return s
}

type CreateServerPaths struct {
  // {"en":"servergroup's ID", "zh_CN":"servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
}

func (s CreateServerPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateServerPaths) GoString() string {
  return s.String()
}

func (s *CreateServerPaths) SetServergroupId(v int) *CreateServerPaths {
  s.ServergroupId = &v
  return s
}

type CreateServerParameters struct {
}

func (s CreateServerParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateServerParameters) GoString() string {
  return s.String()
}

type CreateServerRequestHeader struct {
}

func (s CreateServerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateServerRequestHeader) GoString() string {
  return s.String()
}

type CreateServerResponseHeader struct {
}

func (s CreateServerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateServerResponseHeader) GoString() string {
  return s.String()
}




type GetHealthCheckAgentListRequest struct {
}

func (s GetHealthCheckAgentListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckAgentListRequest) GoString() string {
  return s.String()
}

type GetHealthCheckAgentListResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetHealthCheckAgentListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetHealthCheckAgentListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckAgentListResponse) GoString() string {
  return s.String()
}

func (s *GetHealthCheckAgentListResponse) SetData(v []*GetHealthCheckAgentListResponseData) *GetHealthCheckAgentListResponse {
  s.Data = v
  return s
}

func (s *GetHealthCheckAgentListResponse) SetCode(v int) *GetHealthCheckAgentListResponse {
  s.Code = &v
  return s
}

func (s *GetHealthCheckAgentListResponse) SetMessage(v string) *GetHealthCheckAgentListResponse {
  s.Message = &v
  return s
}

type GetHealthCheckAgentListResponseData struct     {
  // {"en":"Last update time, timestamp(second)", "zh_CN":"最近更新时间，时间戳（秒）"}
  LastUpdated *int `json:"lastUpdated,omitempty" xml:"lastUpdated,omitempty" require:"true"`
  // {"en":"data", "zh_CN":"值"}
  Ips []*GetHealthCheckAgentListResponseDataIps `json:"ips,omitempty" xml:"ips,omitempty" require:"true" type:"Repeated"`
}

func (s GetHealthCheckAgentListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckAgentListResponseData) GoString() string {
  return s.String()
}

func (s *GetHealthCheckAgentListResponseData) SetLastUpdated(v int) *GetHealthCheckAgentListResponseData {
  s.LastUpdated = &v
  return s
}

func (s *GetHealthCheckAgentListResponseData) SetIps(v []*GetHealthCheckAgentListResponseDataIps) *GetHealthCheckAgentListResponseData {
  s.Ips = v
  return s
}

type GetHealthCheckAgentListResponseDataIps struct     {
  // {"en":"IP of agent", "zh_CN":"探测器的IP"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"Effective time, timestamp(second)", "zh_CN":"生效时间，时间戳(秒)"}
  EffectiveTime *int `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty" require:"true"`
}

func (s GetHealthCheckAgentListResponseDataIps) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckAgentListResponseDataIps) GoString() string {
  return s.String()
}

func (s *GetHealthCheckAgentListResponseDataIps) SetIp(v string) *GetHealthCheckAgentListResponseDataIps {
  s.Ip = &v
  return s
}

func (s *GetHealthCheckAgentListResponseDataIps) SetEffectiveTime(v int) *GetHealthCheckAgentListResponseDataIps {
  s.EffectiveTime = &v
  return s
}

type GetHealthCheckAgentListPaths struct {
}

func (s GetHealthCheckAgentListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckAgentListPaths) GoString() string {
  return s.String()
}

type GetHealthCheckAgentListParameters struct {
}

func (s GetHealthCheckAgentListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckAgentListParameters) GoString() string {
  return s.String()
}

type GetHealthCheckAgentListRequestHeader struct {
}

func (s GetHealthCheckAgentListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckAgentListRequestHeader) GoString() string {
  return s.String()
}

type GetHealthCheckAgentListResponseHeader struct {
}

func (s GetHealthCheckAgentListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckAgentListResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryFlavorRequest struct {
}

func (s VMPQueryFlavorRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorRequest) GoString() string {
  return s.String()
}

type VMPQueryFlavorDisk struct {
  // {"en":"disk type ,system disk or data disk", "zh_CN":"磁盘类型,数据盘或者系统盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"VMPQueryFlavorDisk space size in GB", "zh_CN":"磁盘空间大小，单位是GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"VMPQueryFlavorDisk type, value:HDD: ordinary hard disk,SSD: solid state drive,The default is HDD", "zh_CN":"磁盘类型，取值：HDD：普通硬盘,SSD：固态硬盘,默认是HDD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s VMPQueryFlavorDisk) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorDisk) GoString() string {
  return s.String()
}

func (s *VMPQueryFlavorDisk) SetType(v string) *VMPQueryFlavorDisk {
  s.Type = &v
  return s
}

func (s *VMPQueryFlavorDisk) SetSize(v int) *VMPQueryFlavorDisk {
  s.Size = &v
  return s
}

func (s *VMPQueryFlavorDisk) SetCategory(v string) *VMPQueryFlavorDisk {
  s.Category = &v
  return s
}

type VMPQueryFlavorResponse struct {
  // {"en":"flavors", "zh_CN":"规格"}
  Flavors []*string `json:"flavors,omitempty" xml:"flavors,omitempty" require:"true" type:"Repeated"`
  // {"en":"Unique identification of virtual machine specification, global unique", "zh_CN":"实例规格唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"form name", "zh_CN":"规格名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Number of CPUs of virtual machine", "zh_CN":"实例的cpu数"}
  Vcpus *int `json:"vcpus,omitempty" xml:"vcpus,omitempty" require:"true"`
  // {"en":"Virtual machine memory in GB", "zh_CN":"实例内存,单位是GB"}
  Ram *int `json:"ram,omitempty" xml:"ram,omitempty" require:"true"`
  // {"en":"VMPQueryFlavorDisk information of virtual machine", "zh_CN":"实例的磁盘信息"}
  Disks []*VMPQueryFlavorDisk `json:"disks,omitempty" xml:"disks,omitempty" require:"true" type:"Repeated"`
  // {"en":"Bearable bandwidth, Mbps", "zh_CN":"可承载带宽，单位是Mbps"}
  Bandwidth *int `json:"bandwidth,omitempty" xml:"bandwidth,omitempty" require:"true"`
  // {"en":"1: Yes, -1: No,1 means the template is bare metal template;,-1 indicates that the template is a cloud host template;", "zh_CN":"1：是，-1：否,1表示该模板是裸机模板；-1表示该模板是云主机模板；"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"SSD system disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template", "zh_CN":"SSD系统盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  SysSsdLimit *int `json:"sysSsdLimit,omitempty" xml:"sysSsdLimit,omitempty" require:"true"`
  // {"en":"HDD system disk quota (GB). If it is a bare-metal template or a stand-alone template, this parameter has no meaning", "zh_CN":"HDD系统盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  SysHddLimit *int `json:"sysHddLimit,omitempty" xml:"sysHddLimit,omitempty" require:"true"`
  // {"en":"SSD disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template", "zh_CN":"SSD数据盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  DataSsdLimit *int `json:"dataSsdLimit,omitempty" xml:"dataSsdLimit,omitempty" require:"true"`
  // {"en":"HDD disk quota (GB). This parameter has no meaning if it is a bare-metal template or a stand-alone disk template", "zh_CN":"HDD数据盘限额（GB），如果是裸机模板或者是独立盘模板，该参数无意义"}
  DataHddLimit *int `json:"dataHddLimit,omitempty" xml:"dataHddLimit,omitempty" require:"true"`
  // {"en":"Template type,Values: 201- public template, 202- custom template", "zh_CN":"模板类型,取值：201-公共模板、202-自定义模板"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s VMPQueryFlavorResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryFlavorResponse) SetFlavors(v []*string) *VMPQueryFlavorResponse {
  s.Flavors = v
  return s
}

func (s *VMPQueryFlavorResponse) SetId(v string) *VMPQueryFlavorResponse {
  s.Id = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetName(v string) *VMPQueryFlavorResponse {
  s.Name = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetVcpus(v int) *VMPQueryFlavorResponse {
  s.Vcpus = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetRam(v int) *VMPQueryFlavorResponse {
  s.Ram = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetDisks(v []*VMPQueryFlavorDisk) *VMPQueryFlavorResponse {
  s.Disks = v
  return s
}

func (s *VMPQueryFlavorResponse) SetBandwidth(v int) *VMPQueryFlavorResponse {
  s.Bandwidth = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetIsBm(v int) *VMPQueryFlavorResponse {
  s.IsBm = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetSysSsdLimit(v int) *VMPQueryFlavorResponse {
  s.SysSsdLimit = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetSysHddLimit(v int) *VMPQueryFlavorResponse {
  s.SysHddLimit = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetDataSsdLimit(v int) *VMPQueryFlavorResponse {
  s.DataSsdLimit = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetDataHddLimit(v int) *VMPQueryFlavorResponse {
  s.DataHddLimit = &v
  return s
}

func (s *VMPQueryFlavorResponse) SetType(v string) *VMPQueryFlavorResponse {
  s.Type = &v
  return s
}

type VMPQueryFlavorPaths struct {
}

func (s VMPQueryFlavorPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorPaths) GoString() string {
  return s.String()
}

type VMPQueryFlavorParameters struct {
  // {"en":"The virtual machine specification is a unique identifier. Multiple values are separated by commas.Can be left blank, and when left blank, all available template specifications will be returned.", "zh_CN":"实例规格唯一标识，多个值用英文逗号分隔。可放空不填，不填时返回所有可用模板规格。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
}

func (s VMPQueryFlavorParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryFlavorParameters) SetIds(v string) *VMPQueryFlavorParameters {
  s.Ids = &v
  return s
}

type VMPQueryFlavorRequestHeader struct {
}

func (s VMPQueryFlavorRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryFlavorResponseHeader struct {
}

func (s VMPQueryFlavorResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryFlavorResponseHeader) GoString() string {
  return s.String()
}




type GetHealthCheckWhitelistRequest struct {
}

func (s GetHealthCheckWhitelistRequest) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckWhitelistRequest) GoString() string {
  return s.String()
}

type GetHealthCheckWhitelistResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetHealthCheckWhitelistResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetHealthCheckWhitelistResponse) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckWhitelistResponse) GoString() string {
  return s.String()
}

func (s *GetHealthCheckWhitelistResponse) SetData(v []*GetHealthCheckWhitelistResponseData) *GetHealthCheckWhitelistResponse {
  s.Data = v
  return s
}

func (s *GetHealthCheckWhitelistResponse) SetCode(v int) *GetHealthCheckWhitelistResponse {
  s.Code = &v
  return s
}

func (s *GetHealthCheckWhitelistResponse) SetMessage(v string) *GetHealthCheckWhitelistResponse {
  s.Message = &v
  return s
}

type GetHealthCheckWhitelistResponseData struct     {
  // {"en":"Dictionary value[disabled, auto, manual]", "zh_CN":"字典值[disabled, auto, manual]"}
  Enable *string `json:"enable,omitempty" xml:"enable,omitempty" require:"true"`
  // {"en":"Health check agent whitelist, will not return when disabled", "zh_CN":"健康检查探测白名单列表，disabled时不会返回"}
  Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s GetHealthCheckWhitelistResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckWhitelistResponseData) GoString() string {
  return s.String()
}

func (s *GetHealthCheckWhitelistResponseData) SetEnable(v string) *GetHealthCheckWhitelistResponseData {
  s.Enable = &v
  return s
}

func (s *GetHealthCheckWhitelistResponseData) SetIps(v []*string) *GetHealthCheckWhitelistResponseData {
  s.Ips = v
  return s
}

type GetHealthCheckWhitelistPaths struct {
}

func (s GetHealthCheckWhitelistPaths) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckWhitelistPaths) GoString() string {
  return s.String()
}

type GetHealthCheckWhitelistParameters struct {
}

func (s GetHealthCheckWhitelistParameters) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckWhitelistParameters) GoString() string {
  return s.String()
}

type GetHealthCheckWhitelistRequestHeader struct {
}

func (s GetHealthCheckWhitelistRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckWhitelistRequestHeader) GoString() string {
  return s.String()
}

type GetHealthCheckWhitelistResponseHeader struct {
}

func (s GetHealthCheckWhitelistResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckWhitelistResponseHeader) GoString() string {
  return s.String()
}




type GetRegionListRequest struct {
}

func (s GetRegionListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRegionListRequest) GoString() string {
  return s.String()
}

type GetRegionListResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetRegionListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetRegionListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRegionListResponse) GoString() string {
  return s.String()
}

func (s *GetRegionListResponse) SetData(v []*GetRegionListResponseData) *GetRegionListResponse {
  s.Data = v
  return s
}

func (s *GetRegionListResponse) SetCode(v int) *GetRegionListResponse {
  s.Code = &v
  return s
}

func (s *GetRegionListResponse) SetMessage(v string) *GetRegionListResponse {
  s.Message = &v
  return s
}

type GetRegionListResponseData struct     {
  // {"en":"REGION id", "zh_CN":"REGION的ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"REGION name", "zh_CN":"REGION的NAME"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetRegionListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetRegionListResponseData) GoString() string {
  return s.String()
}

func (s *GetRegionListResponseData) SetId(v string) *GetRegionListResponseData {
  s.Id = &v
  return s
}

func (s *GetRegionListResponseData) SetName(v string) *GetRegionListResponseData {
  s.Name = &v
  return s
}

type GetRegionListPaths struct {
}

func (s GetRegionListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRegionListPaths) GoString() string {
  return s.String()
}

type GetRegionListParameters struct {
  // {"en":"Query the sub level region by a region ID", "zh_CN":"根据region id查询子节点的信息"}
  ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s GetRegionListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRegionListParameters) GoString() string {
  return s.String()
}

func (s *GetRegionListParameters) SetParentId(v string) *GetRegionListParameters {
  s.ParentId = &v
  return s
}

type GetRegionListRequestHeader struct {
}

func (s GetRegionListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRegionListRequestHeader) GoString() string {
  return s.String()
}

type GetRegionListResponseHeader struct {
}

func (s GetRegionListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRegionListResponseHeader) GoString() string {
  return s.String()
}




type GetHealthCheckerListRequest struct {
}

func (s GetHealthCheckerListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckerListRequest) GoString() string {
  return s.String()
}

type GetHealthCheckerListResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetHealthCheckerListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetHealthCheckerListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckerListResponse) GoString() string {
  return s.String()
}

func (s *GetHealthCheckerListResponse) SetData(v []*GetHealthCheckerListResponseData) *GetHealthCheckerListResponse {
  s.Data = v
  return s
}

func (s *GetHealthCheckerListResponse) SetCode(v int) *GetHealthCheckerListResponse {
  s.Code = &v
  return s
}

func (s *GetHealthCheckerListResponse) SetMessage(v string) *GetHealthCheckerListResponse {
  s.Message = &v
  return s
}

type GetHealthCheckerListResponseData struct     {
  // {"en":"healthchecker id", "zh_CN":"healthchecker的ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"healthchecker name", "zh_CN":"healthchecker的NAME"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetHealthCheckerListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckerListResponseData) GoString() string {
  return s.String()
}

func (s *GetHealthCheckerListResponseData) SetId(v int) *GetHealthCheckerListResponseData {
  s.Id = &v
  return s
}

func (s *GetHealthCheckerListResponseData) SetName(v string) *GetHealthCheckerListResponseData {
  s.Name = &v
  return s
}

type GetHealthCheckerListPaths struct {
}

func (s GetHealthCheckerListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckerListPaths) GoString() string {
  return s.String()
}

type GetHealthCheckerListParameters struct {
}

func (s GetHealthCheckerListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckerListParameters) GoString() string {
  return s.String()
}

type GetHealthCheckerListRequestHeader struct {
}

func (s GetHealthCheckerListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckerListRequestHeader) GoString() string {
  return s.String()
}

type GetHealthCheckerListResponseHeader struct {
}

func (s GetHealthCheckerListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetHealthCheckerListResponseHeader) GoString() string {
  return s.String()
}




type GetServerListByServerGroupRequest struct {
}

func (s GetServerListByServerGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s GetServerListByServerGroupRequest) GoString() string {
  return s.String()
}

type GetServerListByServerGroupResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetServerListByServerGroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetServerListByServerGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s GetServerListByServerGroupResponse) GoString() string {
  return s.String()
}

func (s *GetServerListByServerGroupResponse) SetData(v []*GetServerListByServerGroupResponseData) *GetServerListByServerGroupResponse {
  s.Data = v
  return s
}

func (s *GetServerListByServerGroupResponse) SetCode(v int) *GetServerListByServerGroupResponse {
  s.Code = &v
  return s
}

func (s *GetServerListByServerGroupResponse) SetMessage(v string) *GetServerListByServerGroupResponse {
  s.Message = &v
  return s
}

type GetServerListByServerGroupResponseData struct     {
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group region id", "zh_CN":"server group的地理ID"}
  ServergroupRegion *string `json:"servergroupRegion,omitempty" xml:"servergroupRegion,omitempty" require:"true"`
  // {"en":"server group region name", "zh_CN":"server group的地理名称"}
  ServergroupRegionName *string `json:"servergroupRegionName,omitempty" xml:"servergroupRegionName,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"server group的名称"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*GetServerListByServerGroupResponseDataHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
  // {"en":"The name of the associated ClbDomain", "zh_CN":"关联的ClbDomain的名称"}
  ClbDomains []*string `json:"clbDomains,omitempty" xml:"clbDomains,omitempty" require:"true" type:"Repeated"`
}

func (s GetServerListByServerGroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetServerListByServerGroupResponseData) GoString() string {
  return s.String()
}

func (s *GetServerListByServerGroupResponseData) SetServergroupId(v int) *GetServerListByServerGroupResponseData {
  s.ServergroupId = &v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetServergroupRegion(v string) *GetServerListByServerGroupResponseData {
  s.ServergroupRegion = &v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetServergroupRegionName(v string) *GetServerListByServerGroupResponseData {
  s.ServergroupRegionName = &v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetServergroupName(v string) *GetServerListByServerGroupResponseData {
  s.ServergroupName = &v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetIp(v string) *GetServerListByServerGroupResponseData {
  s.Ip = &v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetIsEnabled(v bool) *GetServerListByServerGroupResponseData {
  s.IsEnabled = &v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetServerName(v string) *GetServerListByServerGroupResponseData {
  s.ServerName = &v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetServerId(v int) *GetServerListByServerGroupResponseData {
  s.ServerId = &v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetHealthcheckers(v []*GetServerListByServerGroupResponseDataHealthcheckers) *GetServerListByServerGroupResponseData {
  s.Healthcheckers = v
  return s
}

func (s *GetServerListByServerGroupResponseData) SetClbDomains(v []*string) *GetServerListByServerGroupResponseData {
  s.ClbDomains = v
  return s
}

type GetServerListByServerGroupResponseDataHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetServerListByServerGroupResponseDataHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s GetServerListByServerGroupResponseDataHealthcheckers) GoString() string {
  return s.String()
}

func (s *GetServerListByServerGroupResponseDataHealthcheckers) SetId(v int) *GetServerListByServerGroupResponseDataHealthcheckers {
  s.Id = &v
  return s
}

func (s *GetServerListByServerGroupResponseDataHealthcheckers) SetName(v string) *GetServerListByServerGroupResponseDataHealthcheckers {
  s.Name = &v
  return s
}

type GetServerListByServerGroupPaths struct {
  // {"en":"servergroup's ID, if set to 0, query all servers", "zh_CN":"servergroup的ID，如果设置为0，则查询所有的server"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
}

func (s GetServerListByServerGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s GetServerListByServerGroupPaths) GoString() string {
  return s.String()
}

func (s *GetServerListByServerGroupPaths) SetServergroupId(v int) *GetServerListByServerGroupPaths {
  s.ServergroupId = &v
  return s
}

type GetServerListByServerGroupParameters struct {
  // {"en":"server group region id", "zh_CN":"server group的地理ID"}
  ServergroupRegion *string `json:"servergroupRegion,omitempty" xml:"servergroupRegion,omitempty"`
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Healthcheckers *int `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty"`
  // {"en":"Server filter keywords, Server name or IP", "zh_CN":"查询关键字，过滤名称或IP"}
  Search *string `json:"search,omitempty" xml:"search,omitempty"`
}

func (s GetServerListByServerGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s GetServerListByServerGroupParameters) GoString() string {
  return s.String()
}

func (s *GetServerListByServerGroupParameters) SetServergroupRegion(v string) *GetServerListByServerGroupParameters {
  s.ServergroupRegion = &v
  return s
}

func (s *GetServerListByServerGroupParameters) SetHealthcheckers(v int) *GetServerListByServerGroupParameters {
  s.Healthcheckers = &v
  return s
}

func (s *GetServerListByServerGroupParameters) SetIsEnabled(v bool) *GetServerListByServerGroupParameters {
  s.IsEnabled = &v
  return s
}

func (s *GetServerListByServerGroupParameters) SetSearch(v string) *GetServerListByServerGroupParameters {
  s.Search = &v
  return s
}

type GetServerListByServerGroupRequestHeader struct {
}

func (s GetServerListByServerGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetServerListByServerGroupRequestHeader) GoString() string {
  return s.String()
}

type GetServerListByServerGroupResponseHeader struct {
}

func (s GetServerListByServerGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetServerListByServerGroupResponseHeader) GoString() string {
  return s.String()
}




type GetServerGroupListRequest struct {
}

func (s GetServerGroupListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListRequest) GoString() string {
  return s.String()
}

type GetServerGroupListResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetServerGroupListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetServerGroupListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListResponse) GoString() string {
  return s.String()
}

func (s *GetServerGroupListResponse) SetData(v []*GetServerGroupListResponseData) *GetServerGroupListResponse {
  s.Data = v
  return s
}

func (s *GetServerGroupListResponse) SetCode(v int) *GetServerGroupListResponse {
  s.Code = &v
  return s
}

func (s *GetServerGroupListResponse) SetMessage(v string) *GetServerGroupListResponse {
  s.Message = &v
  return s
}

type GetServerGroupListResponseData struct     {
  // {"en":"servergroup's ID", "zh_CN":"servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"错误信息或Success"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"region id", "zh_CN":"地理信息ID"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"region name", "zh_CN":"地理信息名称"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"is enabled, true/false", "zh_CN":"是否可用, true/false"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"server count", "zh_CN":"server的数量"}
  ServerCount *int `json:"serverCount,omitempty" xml:"serverCount,omitempty" require:"true"`
  // {"en":"server list", "zh_CN":"server的列表"}
  Servers []*GetServerGroupListResponseDataServers `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s GetServerGroupListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListResponseData) GoString() string {
  return s.String()
}

func (s *GetServerGroupListResponseData) SetServergroupId(v int) *GetServerGroupListResponseData {
  s.ServergroupId = &v
  return s
}

func (s *GetServerGroupListResponseData) SetServergroupName(v string) *GetServerGroupListResponseData {
  s.ServergroupName = &v
  return s
}

func (s *GetServerGroupListResponseData) SetRegion(v string) *GetServerGroupListResponseData {
  s.Region = &v
  return s
}

func (s *GetServerGroupListResponseData) SetRegionName(v string) *GetServerGroupListResponseData {
  s.RegionName = &v
  return s
}

func (s *GetServerGroupListResponseData) SetIsEnabled(v bool) *GetServerGroupListResponseData {
  s.IsEnabled = &v
  return s
}

func (s *GetServerGroupListResponseData) SetServerCount(v int) *GetServerGroupListResponseData {
  s.ServerCount = &v
  return s
}

func (s *GetServerGroupListResponseData) SetServers(v []*GetServerGroupListResponseDataServers) *GetServerGroupListResponseData {
  s.Servers = v
  return s
}

type GetServerGroupListResponseDataServers struct     {
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*GetServerGroupListResponseDataServersHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
}

func (s GetServerGroupListResponseDataServers) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListResponseDataServers) GoString() string {
  return s.String()
}

func (s *GetServerGroupListResponseDataServers) SetIp(v string) *GetServerGroupListResponseDataServers {
  s.Ip = &v
  return s
}

func (s *GetServerGroupListResponseDataServers) SetIsEnabled(v bool) *GetServerGroupListResponseDataServers {
  s.IsEnabled = &v
  return s
}

func (s *GetServerGroupListResponseDataServers) SetServerName(v string) *GetServerGroupListResponseDataServers {
  s.ServerName = &v
  return s
}

func (s *GetServerGroupListResponseDataServers) SetServerId(v int) *GetServerGroupListResponseDataServers {
  s.ServerId = &v
  return s
}

func (s *GetServerGroupListResponseDataServers) SetHealthcheckers(v []*GetServerGroupListResponseDataServersHealthcheckers) *GetServerGroupListResponseDataServers {
  s.Healthcheckers = v
  return s
}

type GetServerGroupListResponseDataServersHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetServerGroupListResponseDataServersHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListResponseDataServersHealthcheckers) GoString() string {
  return s.String()
}

func (s *GetServerGroupListResponseDataServersHealthcheckers) SetId(v int) *GetServerGroupListResponseDataServersHealthcheckers {
  s.Id = &v
  return s
}

func (s *GetServerGroupListResponseDataServersHealthcheckers) SetName(v string) *GetServerGroupListResponseDataServersHealthcheckers {
  s.Name = &v
  return s
}

type GetServerGroupListPaths struct {
}

func (s GetServerGroupListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListPaths) GoString() string {
  return s.String()
}

type GetServerGroupListParameters struct {
}

func (s GetServerGroupListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListParameters) GoString() string {
  return s.String()
}

type GetServerGroupListRequestHeader struct {
}

func (s GetServerGroupListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListRequestHeader) GoString() string {
  return s.String()
}

type GetServerGroupListResponseHeader struct {
}

func (s GetServerGroupListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetServerGroupListResponseHeader) GoString() string {
  return s.String()
}




type UpdateServergroupRequest struct {
  // {"en":"server group's name", "zh_CN":"错误信息或Success"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"region id(obtain the corresponding ID in the getRegionList interface)", "zh_CN":"地理信息（在获取区域列表接口获取对应的ID）"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"is enabled, true/false", "zh_CN":"是否可用, true/false"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
}

func (s UpdateServergroupRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateServergroupRequest) GoString() string {
  return s.String()
}

func (s *UpdateServergroupRequest) SetServergroupName(v string) *UpdateServergroupRequest {
  s.ServergroupName = &v
  return s
}

func (s *UpdateServergroupRequest) SetRegion(v string) *UpdateServergroupRequest {
  s.Region = &v
  return s
}

func (s *UpdateServergroupRequest) SetIsEnabled(v bool) *UpdateServergroupRequest {
  s.IsEnabled = &v
  return s
}

type UpdateServergroupResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *UpdateServergroupResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s UpdateServergroupResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateServergroupResponse) GoString() string {
  return s.String()
}

func (s *UpdateServergroupResponse) SetData(v *UpdateServergroupResponseData) *UpdateServergroupResponse {
  s.Data = v
  return s
}

func (s *UpdateServergroupResponse) SetCode(v int) *UpdateServergroupResponse {
  s.Code = &v
  return s
}

func (s *UpdateServergroupResponse) SetMessage(v string) *UpdateServergroupResponse {
  s.Message = &v
  return s
}

type UpdateServergroupResponseData struct {
  // {"en":"servergroup's ID", "zh_CN":"servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"错误信息或Success"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"region id", "zh_CN":"地理信息ID"}
  Region *string `json:"region,omitempty" xml:"region,omitempty" require:"true"`
  // {"en":"region name", "zh_CN":"地理信息名称"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"is enabled, true/false", "zh_CN":"是否可用, true/false"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
}

func (s UpdateServergroupResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateServergroupResponseData) GoString() string {
  return s.String()
}

func (s *UpdateServergroupResponseData) SetServergroupId(v int) *UpdateServergroupResponseData {
  s.ServergroupId = &v
  return s
}

func (s *UpdateServergroupResponseData) SetServergroupName(v string) *UpdateServergroupResponseData {
  s.ServergroupName = &v
  return s
}

func (s *UpdateServergroupResponseData) SetRegion(v string) *UpdateServergroupResponseData {
  s.Region = &v
  return s
}

func (s *UpdateServergroupResponseData) SetRegionName(v string) *UpdateServergroupResponseData {
  s.RegionName = &v
  return s
}

func (s *UpdateServergroupResponseData) SetIsEnabled(v bool) *UpdateServergroupResponseData {
  s.IsEnabled = &v
  return s
}

type UpdateServergroupPaths struct {
  // {"en":"servergroup's ID", "zh_CN":"servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
}

func (s UpdateServergroupPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateServergroupPaths) GoString() string {
  return s.String()
}

func (s *UpdateServergroupPaths) SetServergroupId(v int) *UpdateServergroupPaths {
  s.ServergroupId = &v
  return s
}

type UpdateServergroupParameters struct {
}

func (s UpdateServergroupParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateServergroupParameters) GoString() string {
  return s.String()
}

type UpdateServergroupRequestHeader struct {
}

func (s UpdateServergroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateServergroupRequestHeader) GoString() string {
  return s.String()
}

type UpdateServergroupResponseHeader struct {
}

func (s UpdateServergroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateServergroupResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryNodeRequest struct {
}

func (s VMPQueryNodeRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeRequest) GoString() string {
  return s.String()
}

type VMPQueryNodeResponse struct {
  // {"en":"Node information array", "zh_CN":"节点信息数组"}
  Nodes []*string `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
  // {"en":"Node name, unique", "zh_CN":"节点名称，唯一"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Node area", "zh_CN":"节点所在区域"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"Province of node", "zh_CN":"节点所在省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"If the node is a multi line node, multiple operators will be returned, separated by '/'", "zh_CN":"节点所在运营商，如果是多线节点，则返回多个运营商，以'/'分隔"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Node status: running - node available; maintenance - node in maintenance, temporarily unavailable", "zh_CN":"节点状态：RUNNING ---节点可用；MAINTENANCE ---节点维护中，暂时不可用"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Line typenode", "zh_CN":"线路类型"}
  LineType *string `json:"lineType,omitempty" xml:"lineType,omitempty" require:"true"`
  // {"en":"IPv6 supported", "zh_CN":"是否支持ipv6"}
  Ipv6Supported *string `json:"ipv6Supported,omitempty" xml:"ipv6Supported,omitempty" require:"true"`
  // {"en":"Whether the node has bare metal resources", "zh_CN":"该节点是否有裸机资源"}
  BmSupported *string `json:"bmSupported,omitempty" xml:"bmSupported,omitempty" require:"true"`
}

func (s VMPQueryNodeResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeResponse) SetNodes(v []*string) *VMPQueryNodeResponse {
  s.Nodes = v
  return s
}

func (s *VMPQueryNodeResponse) SetName(v string) *VMPQueryNodeResponse {
  s.Name = &v
  return s
}

func (s *VMPQueryNodeResponse) SetRegionName(v string) *VMPQueryNodeResponse {
  s.RegionName = &v
  return s
}

func (s *VMPQueryNodeResponse) SetProvince(v string) *VMPQueryNodeResponse {
  s.Province = &v
  return s
}

func (s *VMPQueryNodeResponse) SetCarrier(v string) *VMPQueryNodeResponse {
  s.Carrier = &v
  return s
}

func (s *VMPQueryNodeResponse) SetState(v string) *VMPQueryNodeResponse {
  s.State = &v
  return s
}

func (s *VMPQueryNodeResponse) SetLineType(v string) *VMPQueryNodeResponse {
  s.LineType = &v
  return s
}

func (s *VMPQueryNodeResponse) SetIpv6Supported(v string) *VMPQueryNodeResponse {
  s.Ipv6Supported = &v
  return s
}

func (s *VMPQueryNodeResponse) SetBmSupported(v string) *VMPQueryNodeResponse {
  s.BmSupported = &v
  return s
}

type VMPQueryNodePaths struct {
}

func (s VMPQueryNodePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodePaths) GoString() string {
  return s.String()
}

type VMPQueryNodeParameters struct {
  // {"en":"The sorted field name can have multiple values: name, regionname, province", "zh_CN":"排序的字段名称，可以有多个，取值：name、regionName、province"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"Sorting direction must follow sortkey, value: desc: descending, default value: ASC: ascending", "zh_CN":"排序方向，必须跟在sortKey后面出现，取值：desc：降序，默认值 asc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default", "zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the name specified by marker", "zh_CN":"从marker指定的名称开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Node area (see Appendix for details)", "zh_CN":"节点所属区域（区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Node province (see Appendix for details)", "zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Node carrier (see Appendix for details)", "zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Line type: single -- single line node; double -- double line node; triple -- three line node; BGP -- BGP node", "zh_CN":"线路类型：single -- 单线节点；double -- 双线节点；triple -- 三线节点；bgp -- BGP节点"}
  LineType *string `json:"lineType,omitempty" xml:"lineType,omitempty"`
  // {"en":"IPv6 supported: true: IPv6 supported false: IPv6 not supported", "zh_CN":"是否支持ipv6：True：支持ipv6 False：不支持ipv6"}
  Ipv6Supported *string `json:"ipv6Supported,omitempty" xml:"ipv6Supported,omitempty"`
  // {"en":"Whether the node has bare metal resources
  // True: There are bare metal resources
  // False: No bare metal resources, only virtual machine resources", "zh_CN":"该节点是否有裸机资源
  // True：有裸机资源
  // False：没有裸机资源，只有虚拟机资源"}
  BmSupported *string `json:"bmSupported,omitempty" xml:"bmSupported,omitempty"`
}

func (s VMPQueryNodeParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeParameters) SetSortKey(v string) *VMPQueryNodeParameters {
  s.SortKey = &v
  return s
}

func (s *VMPQueryNodeParameters) SetSortDir(v string) *VMPQueryNodeParameters {
  s.SortDir = &v
  return s
}

func (s *VMPQueryNodeParameters) SetLimit(v int) *VMPQueryNodeParameters {
  s.Limit = &v
  return s
}

func (s *VMPQueryNodeParameters) SetMarker(v string) *VMPQueryNodeParameters {
  s.Marker = &v
  return s
}

func (s *VMPQueryNodeParameters) SetRegionName(v string) *VMPQueryNodeParameters {
  s.RegionName = &v
  return s
}

func (s *VMPQueryNodeParameters) SetProvince(v string) *VMPQueryNodeParameters {
  s.Province = &v
  return s
}

func (s *VMPQueryNodeParameters) SetCarrier(v string) *VMPQueryNodeParameters {
  s.Carrier = &v
  return s
}

func (s *VMPQueryNodeParameters) SetLineType(v string) *VMPQueryNodeParameters {
  s.LineType = &v
  return s
}

func (s *VMPQueryNodeParameters) SetIpv6Supported(v string) *VMPQueryNodeParameters {
  s.Ipv6Supported = &v
  return s
}

func (s *VMPQueryNodeParameters) SetBmSupported(v string) *VMPQueryNodeParameters {
  s.BmSupported = &v
  return s
}

type VMPQueryNodeRequestHeader struct {
}

func (s VMPQueryNodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryNodeResponseHeader struct {
}

func (s VMPQueryNodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryBandwidthRequest struct {
  // {"en":"VMPQueryBandwidthNode array", "zh_CN":"节点数组"}
  NodeNames []*string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" type:"Repeated"`
}

func (s VMPQueryBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthRequest) GoString() string {
  return s.String()
}

func (s *VMPQueryBandwidthRequest) SetNodeNames(v []*string) *VMPQueryBandwidthRequest {
  s.NodeNames = v
  return s
}

type VMPQueryBandwidthResponse struct {
  // {"en":"node", "zh_CN":"节点"}
  Nodes []*VMPQueryBandwidthNode `json:"nodes,omitempty" xml:"nodes,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryBandwidthResponse) SetNodes(v []*VMPQueryBandwidthNode) *VMPQueryBandwidthResponse {
  s.Nodes = v
  return s
}

type VMPQueryBandwidthNode struct {
  // {"en":"VMPQueryBandwidthNode name", "zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"VMPQueryBandwidthNode bandwidth", "zh_CN":"节点带宽"}
  VMPQueryBandwidthNodeBw []*VMPQueryBandwidthNodeBw `json:"nodeBw,omitempty" xml:"nodeBw,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryBandwidthNode) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthNode) GoString() string {
  return s.String()
}

func (s *VMPQueryBandwidthNode) SetNodeName(v string) *VMPQueryBandwidthNode {
  s.NodeName = &v
  return s
}

func (s *VMPQueryBandwidthNode) SetNodeBw(v []*VMPQueryBandwidthNodeBw) *VMPQueryBandwidthNode {
  s.VMPQueryBandwidthNodeBw = v
  return s
}

type VMPQueryBandwidthNodeBw struct {
  // {"en":"Operator code", "zh_CN":"运营商代码"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Whether the node has redundant bandwidth
  // True: redundant bandwidth
  // False: no redundant bandwidth
  // Undefined: node bandwidth statistics are not available. Whether there is redundant bandwidth is unknown. It is recommended to query again later.", "zh_CN":"节点是否有冗余带宽
  // True：有冗余带宽
  // False：无冗余带宽
  // Undefined：节点带宽统计数据不可用，是否有冗余带宽未知，建议稍后重新查询"}
  Result *string `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s VMPQueryBandwidthNodeBw) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthNodeBw) GoString() string {
  return s.String()
}

func (s *VMPQueryBandwidthNodeBw) SetCarrier(v string) *VMPQueryBandwidthNodeBw {
  s.Carrier = &v
  return s
}

func (s *VMPQueryBandwidthNodeBw) SetResult(v string) *VMPQueryBandwidthNodeBw {
  s.Result = &v
  return s
}

type VMPQueryBandwidthPaths struct {
}

func (s VMPQueryBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthPaths) GoString() string {
  return s.String()
}

type VMPQueryBandwidthParameters struct {
}

func (s VMPQueryBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthParameters) GoString() string {
  return s.String()
}

type VMPQueryBandwidthRequestHeader struct {
}

func (s VMPQueryBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryBandwidthResponseHeader struct {
}

func (s VMPQueryBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryBandwidthResponseHeader) GoString() string {
  return s.String()
}




type DeleteServerRequest struct {
}

func (s DeleteServerRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteServerRequest) GoString() string {
  return s.String()
}

type DeleteServerResponse struct {
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"result data", "zh_CN":"返回数据"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteServerResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteServerResponse) GoString() string {
  return s.String()
}

func (s *DeleteServerResponse) SetCode(v int) *DeleteServerResponse {
  s.Code = &v
  return s
}

func (s *DeleteServerResponse) SetMessage(v string) *DeleteServerResponse {
  s.Message = &v
  return s
}

func (s *DeleteServerResponse) SetData(v string) *DeleteServerResponse {
  s.Data = &v
  return s
}

type DeleteServerPaths struct {
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
}

func (s DeleteServerPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteServerPaths) GoString() string {
  return s.String()
}

func (s *DeleteServerPaths) SetServergroupId(v int) *DeleteServerPaths {
  s.ServergroupId = &v
  return s
}

func (s *DeleteServerPaths) SetServerId(v int) *DeleteServerPaths {
  s.ServerId = &v
  return s
}

type DeleteServerParameters struct {
}

func (s DeleteServerParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteServerParameters) GoString() string {
  return s.String()
}

type DeleteServerRequestHeader struct {
}

func (s DeleteServerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteServerRequestHeader) GoString() string {
  return s.String()
}

type DeleteServerResponseHeader struct {
}

func (s DeleteServerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteServerResponseHeader) GoString() string {
  return s.String()
}




type DeleteServergroupRequest struct {
}

func (s DeleteServergroupRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteServergroupRequest) GoString() string {
  return s.String()
}

type DeleteServergroupResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteServergroupResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteServergroupResponse) GoString() string {
  return s.String()
}

func (s *DeleteServergroupResponse) SetData(v string) *DeleteServergroupResponse {
  s.Data = &v
  return s
}

func (s *DeleteServergroupResponse) SetCode(v int) *DeleteServergroupResponse {
  s.Code = &v
  return s
}

func (s *DeleteServergroupResponse) SetMessage(v string) *DeleteServergroupResponse {
  s.Message = &v
  return s
}

type DeleteServergroupPaths struct {
  // {"en":"servergroup's ID", "zh_CN":"servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
}

func (s DeleteServergroupPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteServergroupPaths) GoString() string {
  return s.String()
}

func (s *DeleteServergroupPaths) SetServergroupId(v int) *DeleteServergroupPaths {
  s.ServergroupId = &v
  return s
}

type DeleteServergroupParameters struct {
}

func (s DeleteServergroupParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteServergroupParameters) GoString() string {
  return s.String()
}

type DeleteServergroupRequestHeader struct {
}

func (s DeleteServergroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteServergroupRequestHeader) GoString() string {
  return s.String()
}

type DeleteServergroupResponseHeader struct {
}

func (s DeleteServergroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteServergroupResponseHeader) GoString() string {
  return s.String()
}




type UpdateServerRequest struct {
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*UpdateServerRequestHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateServerRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerRequest) GoString() string {
  return s.String()
}

func (s *UpdateServerRequest) SetServerName(v string) *UpdateServerRequest {
  s.ServerName = &v
  return s
}

func (s *UpdateServerRequest) SetIsEnabled(v bool) *UpdateServerRequest {
  s.IsEnabled = &v
  return s
}

func (s *UpdateServerRequest) SetIp(v string) *UpdateServerRequest {
  s.Ip = &v
  return s
}

func (s *UpdateServerRequest) SetHealthcheckers(v []*UpdateServerRequestHealthcheckers) *UpdateServerRequest {
  s.Healthcheckers = v
  return s
}

type UpdateServerRequestHealthcheckers struct     {
  // {"en":"health checker id, obtained from interface 'getHealthcheckerList'", "zh_CN":"监控检查ID，从”获取健康检查方法列表“接口获取"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s UpdateServerRequestHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerRequestHealthcheckers) GoString() string {
  return s.String()
}

func (s *UpdateServerRequestHealthcheckers) SetId(v int) *UpdateServerRequestHealthcheckers {
  s.Id = &v
  return s
}

type UpdateServerResponse struct {
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"result data", "zh_CN":"返回数据"}
  Data *UpdateServerResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s UpdateServerResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerResponse) GoString() string {
  return s.String()
}

func (s *UpdateServerResponse) SetCode(v int) *UpdateServerResponse {
  s.Code = &v
  return s
}

func (s *UpdateServerResponse) SetMessage(v string) *UpdateServerResponse {
  s.Message = &v
  return s
}

func (s *UpdateServerResponse) SetData(v *UpdateServerResponseData) *UpdateServerResponse {
  s.Data = v
  return s
}

type UpdateServerResponseData struct {
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"server group region id", "zh_CN":"server group的地理ID"}
  ServergroupRegion *string `json:"servergroupRegion,omitempty" xml:"servergroupRegion,omitempty" require:"true"`
  // {"en":"server group region name", "zh_CN":"server group的地理名称"}
  ServergroupRegionName *string `json:"servergroupRegionName,omitempty" xml:"servergroupRegionName,omitempty" require:"true"`
  // {"en":"server group's name", "zh_CN":"server group的名称"}
  ServergroupName *string `json:"servergroupName,omitempty" xml:"servergroupName,omitempty" require:"true"`
  // {"en":"IP", "zh_CN":"ip地址"}
  Ip *string `json:"ip,omitempty" xml:"ip,omitempty" require:"true"`
  // {"en":"enable", "zh_CN":"是否可用"}
  IsEnabled *bool `json:"isEnabled,omitempty" xml:"isEnabled,omitempty" require:"true"`
  // {"en":"server name", "zh_CN":"服务名"}
  ServerName *string `json:"serverName,omitempty" xml:"serverName,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"health checkers", "zh_CN":"监控检查列表"}
  Healthcheckers []*UpdateServerResponseDataHealthcheckers `json:"healthcheckers,omitempty" xml:"healthcheckers,omitempty" require:"true" type:"Repeated"`
  // {"en":"The name of the associated ClbDomain", "zh_CN":"关联的ClbDomain的名称"}
  ClbDomains []*string `json:"clbDomains,omitempty" xml:"clbDomains,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateServerResponseData) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerResponseData) GoString() string {
  return s.String()
}

func (s *UpdateServerResponseData) SetServergroupId(v int) *UpdateServerResponseData {
  s.ServergroupId = &v
  return s
}

func (s *UpdateServerResponseData) SetServergroupRegion(v string) *UpdateServerResponseData {
  s.ServergroupRegion = &v
  return s
}

func (s *UpdateServerResponseData) SetServergroupRegionName(v string) *UpdateServerResponseData {
  s.ServergroupRegionName = &v
  return s
}

func (s *UpdateServerResponseData) SetServergroupName(v string) *UpdateServerResponseData {
  s.ServergroupName = &v
  return s
}

func (s *UpdateServerResponseData) SetIp(v string) *UpdateServerResponseData {
  s.Ip = &v
  return s
}

func (s *UpdateServerResponseData) SetIsEnabled(v bool) *UpdateServerResponseData {
  s.IsEnabled = &v
  return s
}

func (s *UpdateServerResponseData) SetServerName(v string) *UpdateServerResponseData {
  s.ServerName = &v
  return s
}

func (s *UpdateServerResponseData) SetServerId(v int) *UpdateServerResponseData {
  s.ServerId = &v
  return s
}

func (s *UpdateServerResponseData) SetHealthcheckers(v []*UpdateServerResponseDataHealthcheckers) *UpdateServerResponseData {
  s.Healthcheckers = v
  return s
}

func (s *UpdateServerResponseData) SetClbDomains(v []*string) *UpdateServerResponseData {
  s.ClbDomains = v
  return s
}

type UpdateServerResponseDataHealthcheckers struct     {
  // {"en":"health checker id", "zh_CN":"监控检查ID"}
  Id *int `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"health checker name", "zh_CN":"监控检查名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdateServerResponseDataHealthcheckers) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerResponseDataHealthcheckers) GoString() string {
  return s.String()
}

func (s *UpdateServerResponseDataHealthcheckers) SetId(v int) *UpdateServerResponseDataHealthcheckers {
  s.Id = &v
  return s
}

func (s *UpdateServerResponseDataHealthcheckers) SetName(v string) *UpdateServerResponseDataHealthcheckers {
  s.Name = &v
  return s
}

type UpdateServerPaths struct {
  // {"en":"Servergroup's ID", "zh_CN":"Servergroup的ID"}
  ServergroupId *int `json:"servergroupId,omitempty" xml:"servergroupId,omitempty" require:"true"`
  // {"en":"Server's ID", "zh_CN":"Server的ID"}
  ServerId *int `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
}

func (s UpdateServerPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerPaths) GoString() string {
  return s.String()
}

func (s *UpdateServerPaths) SetServergroupId(v int) *UpdateServerPaths {
  s.ServergroupId = &v
  return s
}

func (s *UpdateServerPaths) SetServerId(v int) *UpdateServerPaths {
  s.ServerId = &v
  return s
}

type UpdateServerParameters struct {
}

func (s UpdateServerParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerParameters) GoString() string {
  return s.String()
}

type UpdateServerRequestHeader struct {
}

func (s UpdateServerRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerRequestHeader) GoString() string {
  return s.String()
}

type UpdateServerResponseHeader struct {
}

func (s UpdateServerResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateServerResponseHeader) GoString() string {
  return s.String()
}




type GetIspListRequest struct {
}

func (s GetIspListRequest) String() string {
  return tea.Prettify(s)
}

func (s GetIspListRequest) GoString() string {
  return s.String()
}

type GetIspListResponse struct {
  // {"en":"return data", "zh_CN":"返回值"}
  Data []*GetIspListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"code,success is 0", "zh_CN":"状态码,成功为0"}
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"error message or Success", "zh_CN":"错误信息或Success"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetIspListResponse) String() string {
  return tea.Prettify(s)
}

func (s GetIspListResponse) GoString() string {
  return s.String()
}

func (s *GetIspListResponse) SetData(v []*GetIspListResponseData) *GetIspListResponse {
  s.Data = v
  return s
}

func (s *GetIspListResponse) SetCode(v int) *GetIspListResponse {
  s.Code = &v
  return s
}

func (s *GetIspListResponse) SetMessage(v string) *GetIspListResponse {
  s.Message = &v
  return s
}

type GetIspListResponseData struct     {
  // {"en":"ISP id", "zh_CN":"ISP的ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"ISP name", "zh_CN":"ISP的NAME"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetIspListResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetIspListResponseData) GoString() string {
  return s.String()
}

func (s *GetIspListResponseData) SetId(v string) *GetIspListResponseData {
  s.Id = &v
  return s
}

func (s *GetIspListResponseData) SetName(v string) *GetIspListResponseData {
  s.Name = &v
  return s
}

type GetIspListPaths struct {
}

func (s GetIspListPaths) String() string {
  return tea.Prettify(s)
}

func (s GetIspListPaths) GoString() string {
  return s.String()
}

type GetIspListParameters struct {
  // {"en":"Keyword of ISP name", "zh_CN":"isp名称的关键字"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetIspListParameters) String() string {
  return tea.Prettify(s)
}

func (s GetIspListParameters) GoString() string {
  return s.String()
}

func (s *GetIspListParameters) SetName(v string) *GetIspListParameters {
  s.Name = &v
  return s
}

type GetIspListRequestHeader struct {
}

func (s GetIspListRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetIspListRequestHeader) GoString() string {
  return s.String()
}

type GetIspListResponseHeader struct {
}

func (s GetIspListResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetIspListResponseHeader) GoString() string {
  return s.String()
}




