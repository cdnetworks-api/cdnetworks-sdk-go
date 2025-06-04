package instancemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type ConvertFreeTypeInstanceToChargeTypeRequest struct {
  // {"en":"Unique cloud host identity.Up to 100 IDs can be sent at a time, separated by the half comma character ', '.", "zh_CN":"云主机唯一标识。单次最多可发送100 条ID，ID 之间用半角逗号字符','隔开。"}
  Servers *string `json:"servers,omitempty" xml:"servers,omitempty" require:"true"`
}

func (s ConvertFreeTypeInstanceToChargeTypeRequest) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeRequest) GoString() string {
  return s.String()
}

func (s *ConvertFreeTypeInstanceToChargeTypeRequest) SetServers(v string) *ConvertFreeTypeInstanceToChargeTypeRequest {
  s.Servers = &v
  return s
}

type ConvertFreeTypeInstanceToChargeTypeResponse struct {
  BatchErrorMsg []*ConvertFreeTypeInstanceToChargeTypeErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s ConvertFreeTypeInstanceToChargeTypeResponse) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeResponse) GoString() string {
  return s.String()
}

func (s *ConvertFreeTypeInstanceToChargeTypeResponse) SetBatchErrorMsg(v []*ConvertFreeTypeInstanceToChargeTypeErrorMsg) *ConvertFreeTypeInstanceToChargeTypeResponse {
  s.BatchErrorMsg = v
  return s
}

type ConvertFreeTypeInstanceToChargeTypeErrorMsg struct {
  // {"en":"error code", "zh_CN":"错误编码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"instance id", "zh_CN":"实例id"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"error msg", "zh_CN":"错误信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s ConvertFreeTypeInstanceToChargeTypeErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeErrorMsg) GoString() string {
  return s.String()
}

func (s *ConvertFreeTypeInstanceToChargeTypeErrorMsg) SetCode(v string) *ConvertFreeTypeInstanceToChargeTypeErrorMsg {
  s.Code = &v
  return s
}

func (s *ConvertFreeTypeInstanceToChargeTypeErrorMsg) SetKey(v string) *ConvertFreeTypeInstanceToChargeTypeErrorMsg {
  s.Key = &v
  return s
}

func (s *ConvertFreeTypeInstanceToChargeTypeErrorMsg) SetMsg(v string) *ConvertFreeTypeInstanceToChargeTypeErrorMsg {
  s.Msg = &v
  return s
}

type ConvertFreeTypeInstanceToChargeTypePaths struct {
}

func (s ConvertFreeTypeInstanceToChargeTypePaths) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypePaths) GoString() string {
  return s.String()
}

type ConvertFreeTypeInstanceToChargeTypeParameters struct {
}

func (s ConvertFreeTypeInstanceToChargeTypeParameters) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeParameters) GoString() string {
  return s.String()
}

type ConvertFreeTypeInstanceToChargeTypeRequestHeader struct {
}

func (s ConvertFreeTypeInstanceToChargeTypeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeRequestHeader) GoString() string {
  return s.String()
}

type ConvertFreeTypeInstanceToChargeTypeResponseHeader struct {
}

func (s ConvertFreeTypeInstanceToChargeTypeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ConvertFreeTypeInstanceToChargeTypeResponseHeader) GoString() string {
  return s.String()
}




type ManageInstanceTagsRequest struct {
  // {"en":"The cloud host ID, with multiple IDs separated by the half corner comma character ', '.", "zh_CN":"云主机id，多个ID 之间用半角逗号字符','隔开。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"New instance label value;If the value is null, the instance label of the specified cloud host is deleted.", "zh_CN":"新的实例标签值；如果值为空，则表示删除指定云主机的实例标签。"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty" require:"true"`
}

func (s ManageInstanceTagsRequest) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsRequest) GoString() string {
  return s.String()
}

func (s *ManageInstanceTagsRequest) SetId(v string) *ManageInstanceTagsRequest {
  s.Id = &v
  return s
}

func (s *ManageInstanceTagsRequest) SetTag(v string) *ManageInstanceTagsRequest {
  s.Tag = &v
  return s
}

type ManageInstanceTagsResponse struct {
}

func (s ManageInstanceTagsResponse) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsResponse) GoString() string {
  return s.String()
}

type ManageInstanceTagsPaths struct {
}

func (s ManageInstanceTagsPaths) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsPaths) GoString() string {
  return s.String()
}

type ManageInstanceTagsParameters struct {
}

func (s ManageInstanceTagsParameters) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsParameters) GoString() string {
  return s.String()
}

type ManageInstanceTagsRequestHeader struct {
}

func (s ManageInstanceTagsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsRequestHeader) GoString() string {
  return s.String()
}

type ManageInstanceTagsResponseHeader struct {
}

func (s ManageInstanceTagsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ManageInstanceTagsResponseHeader) GoString() string {
  return s.String()
}




type InstanceRebuildRequest struct {
  // {"en":"vm id", "zh_CN":"云主机ID"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Image ID", "zh_CN":"镜像ID（指定了镜像ID则使用指定的镜像重装，否则使用原镜像重装）"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
  // {"en":"password", "zh_CN":"密码（使用公共镜像重装必须指定密码，使用自定义镜像可不指定）"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"Retain Data Disk", "zh_CN":"是否保留数据盘（1：是；-1：否）"}
  RetainDataDisk *int `json:"retainDataDisk,omitempty" xml:"retainDataDisk,omitempty"`
}

func (s InstanceRebuildRequest) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildRequest) GoString() string {
  return s.String()
}

func (s *InstanceRebuildRequest) SetId(v string) *InstanceRebuildRequest {
  s.Id = &v
  return s
}

func (s *InstanceRebuildRequest) SetImageId(v string) *InstanceRebuildRequest {
  s.ImageId = &v
  return s
}

func (s *InstanceRebuildRequest) SetPassword(v string) *InstanceRebuildRequest {
  s.Password = &v
  return s
}

func (s *InstanceRebuildRequest) SetRetainDataDisk(v int) *InstanceRebuildRequest {
  s.RetainDataDisk = &v
  return s
}

type InstanceRebuildResponse struct {
}

func (s InstanceRebuildResponse) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildResponse) GoString() string {
  return s.String()
}

type InstanceRebuildPaths struct {
}

func (s InstanceRebuildPaths) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildPaths) GoString() string {
  return s.String()
}

type InstanceRebuildParameters struct {
}

func (s InstanceRebuildParameters) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildParameters) GoString() string {
  return s.String()
}

type InstanceRebuildRequestHeader struct {
}

func (s InstanceRebuildRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildRequestHeader) GoString() string {
  return s.String()
}

type InstanceRebuildResponseHeader struct {
}

func (s InstanceRebuildResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceRebuildResponseHeader) GoString() string {
  return s.String()
}




type VMPRemoveInstanceRequest struct {
  // {"en":"Instance ID", "zh_CN":"实例唯一标识。单次最多可发送100 条ID，ID 之间用半角逗号字符“,”隔开。"}
  Servers *string `json:"servers,omitempty" xml:"servers,omitempty" require:"true"`
}

func (s VMPRemoveInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceRequest) GoString() string {
  return s.String()
}

func (s *VMPRemoveInstanceRequest) SetServers(v string) *VMPRemoveInstanceRequest {
  s.Servers = &v
  return s
}

type VMPRemoveInstanceResponse struct {
}

func (s VMPRemoveInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceResponse) GoString() string {
  return s.String()
}

type VMPRemoveInstancePaths struct {
}

func (s VMPRemoveInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstancePaths) GoString() string {
  return s.String()
}

type VMPRemoveInstanceParameters struct {
}

func (s VMPRemoveInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceParameters) GoString() string {
  return s.String()
}

type VMPRemoveInstanceRequestHeader struct {
}

func (s VMPRemoveInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceRequestHeader) GoString() string {
  return s.String()
}

type VMPRemoveInstanceResponseHeader struct {
}

func (s VMPRemoveInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveInstanceResponseHeader) GoString() string {
  return s.String()
}




type VMPInstanceOperationRequest struct {
  // {"en":"Instance action.
  // Values:
  // START: START the
  // SHUTDOWN: SHUTDOWN
  // REBOOT: Force a REBOOT", "zh_CN":"实例操作动作。
  // 取值：
  // START：启动
  // SHUTDOWN：正常关机
  // REBOOT：强制重启"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty" require:"true"`
}

func (s VMPInstanceOperationRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationRequest) GoString() string {
  return s.String()
}

func (s *VMPInstanceOperationRequest) SetOperation(v string) *VMPInstanceOperationRequest {
  s.Operation = &v
  return s
}

type VMPInstanceOperationResponse struct {
}

func (s VMPInstanceOperationResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationResponse) GoString() string {
  return s.String()
}

type VMPInstanceOperationPaths struct {
  // {"en":"Unique identity of virtual machine", "zh_CN":"实例唯一标识"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
}

func (s VMPInstanceOperationPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationPaths) GoString() string {
  return s.String()
}

func (s *VMPInstanceOperationPaths) SetServerId(v string) *VMPInstanceOperationPaths {
  s.ServerId = &v
  return s
}

type VMPInstanceOperationParameters struct {
}

func (s VMPInstanceOperationParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationParameters) GoString() string {
  return s.String()
}

type VMPInstanceOperationRequestHeader struct {
}

func (s VMPInstanceOperationRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationRequestHeader) GoString() string {
  return s.String()
}

type VMPInstanceOperationResponseHeader struct {
}

func (s VMPInstanceOperationResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPInstanceOperationResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryInstanceRequest struct {
}

func (s VMPQueryInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceRequest) GoString() string {
  return s.String()
}

type VMPQueryInstanceAccessIP struct {
  // {"en":"Ip address", "zh_CN":"Ip地址"}
  Address *string `json:"address,omitempty" xml:"address,omitempty" require:"true"`
  // {"en":"IP address operator", "zh_CN":"Ip地址所属运营商"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Protocol type: 4: IPv4 address; 6: IPv6 address", "zh_CN":"协议类型：4：ipv4地址；6：ipv6地址"}
  Protocol *int `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
}

func (s VMPQueryInstanceAccessIP) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceAccessIP) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceAccessIP) SetAddress(v string) *VMPQueryInstanceAccessIP {
  s.Address = &v
  return s
}

func (s *VMPQueryInstanceAccessIP) SetCarrier(v string) *VMPQueryInstanceAccessIP {
  s.Carrier = &v
  return s
}

func (s *VMPQueryInstanceAccessIP) SetProtocol(v int) *VMPQueryInstanceAccessIP {
  s.Protocol = &v
  return s
}

type VMPQueryInstanceDiskInfo struct {
  // {"en":"Disk size (GB)", "zh_CN":"磁盘大小（GB）"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Disk Purpose: System - System disk;DATA - DATA plate", "zh_CN":"磁盘用途：SYSTEM-系统盘；DATA-数据盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Disk type: HDD/SSD", "zh_CN":"磁盘类型：HDD/SSD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s VMPQueryInstanceDiskInfo) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceDiskInfo) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceDiskInfo) SetSize(v int) *VMPQueryInstanceDiskInfo {
  s.Size = &v
  return s
}

func (s *VMPQueryInstanceDiskInfo) SetType(v string) *VMPQueryInstanceDiskInfo {
  s.Type = &v
  return s
}

func (s *VMPQueryInstanceDiskInfo) SetCategory(v string) *VMPQueryInstanceDiskInfo {
  s.Category = &v
  return s
}

type VMPQueryInstanceServer struct {
  // {"en":"Unique identity of virtual machine", "zh_CN":"实例唯一标识"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Virtual Machine Area", "zh_CN":"实例所属区域"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty" require:"true"`
  // {"en":"Province of virtual machine", "zh_CN":"实例所属省份"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Operator of virtual machine", "zh_CN":"实例所属运营商"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Virtual machine image information", "zh_CN":"实例镜像信息"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Virtual machine specifications", "zh_CN":"实例规格"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Virtual machine name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Virtual machine status", "zh_CN":"实例状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"Virtual machine creation time", "zh_CN":"实例创建时间"}
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty" require:"true"`
  // {"en":"Virtual machine IP address", "zh_CN":"实例IP地址"}
  AccessIPv4 *string `json:"accessIPv4,omitempty" xml:"accessIPv4,omitempty" require:"true"`
  // {"en":"IPv4 address of virtual machine intranet. If intranet is specified when creating virtual machine, intranet IP is returned. Otherwise, it is empty.", "zh_CN":"实例内网IPv4地址，如果创建实例时指定了需要内网，则返回内网IPv4，否则是空"}
  PrivateIPv4 *string `json:"privateIPv4,omitempty" xml:"privateIPv4,omitempty" require:"true"`
  // {"en":"IPv6 address of virtual machine intranet. If intranet is specified when creating virtual machine, intranet IP is returned. Otherwise, it is empty.", "zh_CN":"实例内网IPv6地址，如果创建实例时指定了需要内网，则返回内网IPv6，否则是空"}
  PrivateIPv6 *string `json:"privateIPv6,omitempty" xml:"privateIPv6,omitempty" require:"true"`
  // {"en":"Virtual machine login SSH key pair name", "zh_CN":"实例登录SSH秘钥对名称"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty" require:"true"`
  // {"en":"Node name, the name of the node where the virtual machine is located. Through this node name, you can query the real-time redundant bandwidth of each node.", "zh_CN":"节点名称，实例所在节点名称，通过这个节点名称可以查询每个节点的实时冗余带宽情况"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Virtual machine IP address", "zh_CN":"实例ip地址"}
  VMPQueryInstanceAccessIP []*VMPQueryInstanceAccessIP `json:"accessIP,omitempty" xml:"accessIP,omitempty" require:"true" type:"Repeated"`
  // {"en":"If it is free instance, value: YES free instance, NO billing instance", "zh_CN":"是否免费实例，取值：YES 免费实例，NO 计费实例"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty" require:"true"`
  // {"en":"1 indicates that the instance is bare metal
  // -1 indicates that the instance is a virtual machine", "zh_CN":"1表示该实例是裸机
  // -1表示该实例是虚拟机"}
  IsBm *int `json:"isBm,omitempty" xml:"isBm,omitempty" require:"true"`
  // {"en":"A list of security group IDs for instance bindings", "zh_CN":"实例绑定的安全组id列表"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" require:"true" type:"Repeated"`
  // {"en":"Disk information", "zh_CN":"磁盘信息"}
  VMPQueryInstanceDiskInfo []*VMPQueryInstanceDiskInfo `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
  // {"en":"instance tag", "zh_CN":"实例标签"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty" require:"true"`
}

func (s VMPQueryInstanceServer) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceServer) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceServer) SetId(v string) *VMPQueryInstanceServer {
  s.Id = &v
  return s
}

func (s *VMPQueryInstanceServer) SetRegionName(v string) *VMPQueryInstanceServer {
  s.RegionName = &v
  return s
}

func (s *VMPQueryInstanceServer) SetProvince(v string) *VMPQueryInstanceServer {
  s.Province = &v
  return s
}

func (s *VMPQueryInstanceServer) SetCarrier(v string) *VMPQueryInstanceServer {
  s.Carrier = &v
  return s
}

func (s *VMPQueryInstanceServer) SetImageId(v string) *VMPQueryInstanceServer {
  s.ImageId = &v
  return s
}

func (s *VMPQueryInstanceServer) SetFlavorId(v string) *VMPQueryInstanceServer {
  s.FlavorId = &v
  return s
}

func (s *VMPQueryInstanceServer) SetName(v string) *VMPQueryInstanceServer {
  s.Name = &v
  return s
}

func (s *VMPQueryInstanceServer) SetState(v string) *VMPQueryInstanceServer {
  s.State = &v
  return s
}

func (s *VMPQueryInstanceServer) SetCreatedAt(v string) *VMPQueryInstanceServer {
  s.CreatedAt = &v
  return s
}

func (s *VMPQueryInstanceServer) SetAccessIPv4(v string) *VMPQueryInstanceServer {
  s.AccessIPv4 = &v
  return s
}

func (s *VMPQueryInstanceServer) SetPrivateIPv4(v string) *VMPQueryInstanceServer {
  s.PrivateIPv4 = &v
  return s
}

func (s *VMPQueryInstanceServer) SetPrivateIPv6(v string) *VMPQueryInstanceServer {
  s.PrivateIPv6 = &v
  return s
}

func (s *VMPQueryInstanceServer) SetKeyName(v string) *VMPQueryInstanceServer {
  s.KeyName = &v
  return s
}

func (s *VMPQueryInstanceServer) SetNodeName(v string) *VMPQueryInstanceServer {
  s.NodeName = &v
  return s
}

func (s *VMPQueryInstanceServer) SetAccessIP(v []*VMPQueryInstanceAccessIP) *VMPQueryInstanceServer {
  s.VMPQueryInstanceAccessIP = v
  return s
}

func (s *VMPQueryInstanceServer) SetIsFree(v string) *VMPQueryInstanceServer {
  s.IsFree = &v
  return s
}

func (s *VMPQueryInstanceServer) SetIsBm(v int) *VMPQueryInstanceServer {
  s.IsBm = &v
  return s
}

func (s *VMPQueryInstanceServer) SetSecurityGroupIds(v []*string) *VMPQueryInstanceServer {
  s.SecurityGroupIds = v
  return s
}

func (s *VMPQueryInstanceServer) SetDiskInfo(v []*VMPQueryInstanceDiskInfo) *VMPQueryInstanceServer {
  s.VMPQueryInstanceDiskInfo = v
  return s
}

func (s *VMPQueryInstanceServer) SetTag(v string) *VMPQueryInstanceServer {
  s.Tag = &v
  return s
}

type VMPQueryInstanceResponse struct {
  // {"en":"Virtual machine information array", "zh_CN":"实例信息数组"}
  Servers []*VMPQueryInstanceServer `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceResponse) SetServers(v []*VMPQueryInstanceServer) *VMPQueryInstanceResponse {
  s.Servers = v
  return s
}

type VMPQueryInstancePaths struct {
}

func (s VMPQueryInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstancePaths) GoString() string {
  return s.String()
}

type VMPQueryInstanceParameters struct {
  // {"en":"Sort field name, can have more than one, value:
  // Name, regionName, createdAt, type, state, nodeName, etc", "zh_CN":"排序的字段名称，可以有多个，取值：
  // name、regionName、createdAt、province、state、nodeName等"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"There can be multiple field names for sorting, with values:
  // 
  // Name, regionname, createdat, province, state, nodeName, etc.'", "zh_CN":"排序方向，必须跟在sortKey后面出现，取值：
  // desc：降序，默认值
  // asc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default", "zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the virtual machine ID specified by the marker", "zh_CN":"从marker指定的实例id开始查询，升序查询（若要分页查询，则不能指定sortKey参数）"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Virtual machine ID. A maximum of 100 IDS can be queried at a time. The IDs are separated by a half angle comma character ','.", "zh_CN":"实例ID。单次最多查询 100 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"Name of the region, for example, South China region is Huanan, refer to Appendix 3", "zh_CN":"区域名称（区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Province of virtual machine", "zh_CN":"实例所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Operator of virtual machine", "zh_CN":"实例所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Virtual machine image identity", "zh_CN":"实例镜像标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
  // {"en":"Virtual machine specification ID", "zh_CN":"实例规格标识"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty"`
  // {"en":"Virtual machine name", "zh_CN":"实例名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Virtual machine ips. A maximum of 100 IPS can be queried at a time. The IPs are separated by a half angle comma character ','.", "zh_CN":"实例IP。单次最多查询 100 条 IP，IP 之间用半角逗号字符','隔开。"}
  Ips *string `json:"ips,omitempty" xml:"ips,omitempty"`
  // {"en":"Virtual machine status, value (meaning to query virtual machines in the following status)
  // 
  // Running status
  // 
  // Building new status
  // 
  // Stopped stop
  // 
  // ERROR error
  // 
  // Deleting destroying
  // 
  // Restarting
  // 
  // Starting
  // 
  // Stopping", "zh_CN":"实例状态，取值（意为查询处于下述状态的虚拟机）
  // RUNNING 运行状态
  // BUILDING 新建状态
  // STOPPED 停机
  // ERROR 错误
  // DELETING 销毁中
  // RESTARTING  重启中
  // STARTING  启动中
  // STOPPING  停止中"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
  // {"en":"Whether it is a free instance, value:
  // Yes free instances, NO billed instances", "zh_CN":"是否免费实例，取值:
  // YES 免费实例，NO 计费实例"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty"`
  // {"en":"1 means to query only bare-metal instances,
  // -1 means that only virtual machine instances are queried,
  // Without this parameter all cloud hosts are queried", "zh_CN":"1表示只查询裸机实例，
  // -1表示只查询虚拟机实例，
  // 不带这个参数表示查询所有云主机"}
  IsBm *string `json:"isBm,omitempty" xml:"isBm,omitempty"`
  // {"en":"Cloud Host Label
  // Multiple values are separated by a half corner comma, and the relationship between multiple values is or, that is, the instance label equals any one of these multiple values", "zh_CN":"云主机标签
  // 多个值用半角逗号隔开，多个值是或者的关系，即实例标签等于这多个中的任意一个就满足条件"}
  Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
  // {"en":"ipv6 formate: 1: Zero Compressed(default); 6: With Leading Zero Suppression; 3: Full Address", "zh_CN":"ipv6格式：1：省略零压缩格式(默认)；2：省略前导零格式; 3: 完整格式"}
  Ipv6Format *string `json:"ipv6Format,omitempty" xml:"ipv6Format,omitempty"`
}

func (s VMPQueryInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryInstanceParameters) SetSortKey(v string) *VMPQueryInstanceParameters {
  s.SortKey = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetSortDir(v string) *VMPQueryInstanceParameters {
  s.SortDir = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetLimit(v int) *VMPQueryInstanceParameters {
  s.Limit = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetMarker(v string) *VMPQueryInstanceParameters {
  s.Marker = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIds(v string) *VMPQueryInstanceParameters {
  s.Ids = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetRegionName(v string) *VMPQueryInstanceParameters {
  s.RegionName = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetProvince(v string) *VMPQueryInstanceParameters {
  s.Province = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetCarrier(v string) *VMPQueryInstanceParameters {
  s.Carrier = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetImageId(v string) *VMPQueryInstanceParameters {
  s.ImageId = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetFlavorId(v string) *VMPQueryInstanceParameters {
  s.FlavorId = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetName(v string) *VMPQueryInstanceParameters {
  s.Name = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIps(v string) *VMPQueryInstanceParameters {
  s.Ips = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetState(v string) *VMPQueryInstanceParameters {
  s.State = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIsFree(v string) *VMPQueryInstanceParameters {
  s.IsFree = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIsBm(v string) *VMPQueryInstanceParameters {
  s.IsBm = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetTags(v string) *VMPQueryInstanceParameters {
  s.Tags = &v
  return s
}

func (s *VMPQueryInstanceParameters) SetIpv6Format(v string) *VMPQueryInstanceParameters {
  s.Ipv6Format = &v
  return s
}

type VMPQueryInstanceRequestHeader struct {
}

func (s VMPQueryInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryInstanceResponseHeader struct {
}

func (s VMPQueryInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryInstanceResponseHeader) GoString() string {
  return s.String()
}




type InstanceIPV6ManagementRequest struct {
  // {"en":"Operation:
  // ALLOCATION - ipv6 application
  // REMOVE - ipv6 is removed", "zh_CN":"操作：
  // ALLOCATION-ipv6申请 
  // REMOVE-ipv6移除"}
  Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
  // {"en":"Instance ID (single only)", "zh_CN":"实例id（只支持单个）"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
}

func (s InstanceIPV6ManagementRequest) String() string {
  return tea.Prettify(s)
}

func (s InstanceIPV6ManagementRequest) GoString() string {
  return s.String()
}

func (s *InstanceIPV6ManagementRequest) SetAction(v string) *InstanceIPV6ManagementRequest {
  s.Action = &v
  return s
}

func (s *InstanceIPV6ManagementRequest) SetInstanceId(v string) *InstanceIPV6ManagementRequest {
  s.InstanceId = &v
  return s
}

type InstanceIPV6ManagementResponse struct {
  // {"en":"none", "zh_CN":"实例id"}
  InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty" require:"true"`
  // {"en":"Instance IPv6 addresses, one for each operator if it is a multi-line node", "zh_CN":"实例ipv6地址，如果是多线节点，每个运营商各分配一个地址"}
  AccessIPv6 []*string `json:"accessIPv6,omitempty" xml:"accessIPv6,omitempty" require:"true" type:"Repeated"`
}

func (s InstanceIPV6ManagementResponse) String() string {
  return tea.Prettify(s)
}

func (s InstanceIPV6ManagementResponse) GoString() string {
  return s.String()
}

func (s *InstanceIPV6ManagementResponse) SetInstanceId(v string) *InstanceIPV6ManagementResponse {
  s.InstanceId = &v
  return s
}

func (s *InstanceIPV6ManagementResponse) SetAccessIPv6(v []*string) *InstanceIPV6ManagementResponse {
  s.AccessIPv6 = v
  return s
}

type InstanceIPV6ManagementPaths struct {
}

func (s InstanceIPV6ManagementPaths) String() string {
  return tea.Prettify(s)
}

func (s InstanceIPV6ManagementPaths) GoString() string {
  return s.String()
}

type InstanceIPV6ManagementParameters struct {
}

func (s InstanceIPV6ManagementParameters) String() string {
  return tea.Prettify(s)
}

func (s InstanceIPV6ManagementParameters) GoString() string {
  return s.String()
}

type InstanceIPV6ManagementRequestHeader struct {
}

func (s InstanceIPV6ManagementRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceIPV6ManagementRequestHeader) GoString() string {
  return s.String()
}

type InstanceIPV6ManagementResponseHeader struct {
}

func (s InstanceIPV6ManagementResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceIPV6ManagementResponseHeader) GoString() string {
  return s.String()
}




type VMPCreateInstanceRequest struct {
  // {"en":"Creating array objects for virtual machines", "zh_CN":"创建实例的数组对象"}
  Servers []*VMPCreateInstanceServer `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPCreateInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequest) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceRequest) SetServers(v []*VMPCreateInstanceServer) *VMPCreateInstanceRequest {
  s.Servers = v
  return s
}

type VMPCreateInstanceServer struct {
  // {"en":"Virtual machine area (see Appendix for details)", "zh_CN":"实例所属区域（节点名称nodeName和区域regionName至少需要上传一个。
  //     区域列表详见附录1：https://www.wangsu.com/document/18204/areas-list?rsr=ws）"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"Province of virtual machine (see Appendix for details)", "zh_CN":"实例所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"If the operator of the virtual machine (see the appendix for details) carries this parameter, please keep it consistent with the carrier returned from the '3.4 node list query' interface.", "zh_CN":"实例所属运营商（dx-电信；wt-网通；yd-移动）如果携带了该参数，请与'3.4节点列表查询'接口返回的carrier保持一致"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Node name, indicating that the specified node creates a virtual machine (the node name returned by interface 3.4)", "zh_CN":"节点名称，表示指定节点创建实例（节点名称可通过资源管理-节点列表查询接口获取）"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"Virtual machine image identity", "zh_CN":"实例镜像标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Virtual machine specification ID", "zh_CN":"实例规格标识"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Virtual machine name. If the created quantity is greater than 1, the real name is spliced with 3 digits after the parameter. For example, instance 0001, instance 0002", "zh_CN":"实例名称，如果创建数量大于1，则真实名称是在该参数后拼接3位数字。如instance_0001，instance_0002"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Inject user data, support to inject text, text file or gzip file. The maximum length of injected content is 32KB. For content injection, Base64 format encoding is required.", "zh_CN":"注入用户数据，支持注入文本、文本文件或gzip文件。注入内容最大长度32KB。注入内容，需要进行base64格式编码。"}
  UserData *string `json:"userData,omitempty" xml:"userData,omitempty"`
  // {"en":"Number of virtual machines applied", "zh_CN":"申请实例数量"}
  Count *int `json:"count,omitempty" xml:"count,omitempty"`
  // {"en":"Virtual machine root login password", "zh_CN":"实例root用户登录密码（如果选择的是公共镜像，则密码password必填）"}
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
  // {"en":"The name of the SSH secret key pair for virtual machine login. If this parameter is specified, the password login mode is disabled by default, and the password parameter is invalid at the same time.", "zh_CN":"实例登录SSH秘钥对名称，如果指定该参数，默认禁用密码登录方式，password参数同时失效"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty"`
  // {"en":"Whether the virtual machine needs intranet, value:
  // Yes: intranet required
  // No: no intranet is required, default value'", "zh_CN":"实例是否需要内网网络，取值：
  // YES：需要内网
  // NO：不需要内网，默认值"}
  InnerNet *string `json:"innerNet,omitempty" xml:"innerNet,omitempty"`
  // {"en":"CIDR of virtual machine intranet is meaningful only when innernet = yes", "zh_CN":"实例内网的cidr，只有innerNet=YES时才有意义"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"If IP address is specified, it must be within the scope of CIDR, otherwise creation fails.", "zh_CN":"实例内网ip地址，如果指定了ip，必须在cidr的范围内，否则创建失败"}
  PrivateIPv4 *string `json:"privateIPv4,omitempty" xml:"privateIPv4,omitempty"`
  // {"en":"Whether the virtual machine needs intranet2, value:
  // Yes: intranet2 required
  // No: no intranet2 is required, default value'", "zh_CN":"实例是否需要内网2网络，取值：
  // YES：需要内网2
  // NO：不需要内网2，默认值"}
  InnerNet2 *string `json:"innerNet2,omitempty" xml:"innerNet2,omitempty"`
  // {"en":"CIDR of virtual machine intranet2 is meaningful only when innernet = yes", "zh_CN":"实例内网2的cidr，只有innerNet2=YES时才有意义"}
  Cidr2 *string `json:"cidr2,omitempty" xml:"cidr2,omitempty"`
  // {"en":"If IP address is specified, it must be within the scope of CIDR2, otherwise creation fails.", "zh_CN":"实例内网2ip地址，如果指定了ip，必须在cidr2的范围内，否则创建失败"}
  PrivateIPv42 *string `json:"privateIPv42,omitempty" xml:"privateIPv42,omitempty"`
  // {"en":"Inner ipv6 info", "zh_CN":"内网IPv6信息"}
  PrivateIpv6Info []*VMPCreateInstancePivateIpv6Info `json:"privateIpv6Info,omitempty" xml:"privateIpv6Info,omitempty" type:"Repeated"`
  // {"en":"Whether multiple IP protocol addresses are required
  // 
  // 4: only IPv4 address is required, default value
  // 
  // 0: both IPv4 and IPv6 need'", "zh_CN":"是否需要多ip协议地址
  // 4：只需要ipv4地址，默认值
  // 0：ipv4、ipv6都需要"}
  Protocols *int `json:"protocols,omitempty" xml:"protocols,omitempty"`
  // {"en":"IPv4 native attribute, 1: non-native;-1: native;", "zh_CN":"IPv4原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  Ipv4NativeAttribute *string `json:"ipv4NativeAttribute,omitempty" xml:"ipv4NativeAttribute,omitempty"`
  // {"en":"IPv6 native attribute, 1: non-native;-1: native;", "zh_CN":"IPv6原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  Ipv6NativeAttribute *string `json:"ipv6NativeAttribute,omitempty" xml:"ipv6NativeAttribute,omitempty"`
  // {"en":"Whether the instance is free or not, the default billing instance, and the bare machine instance cannot be free, values are as follows:
  // Yes: Free instances
  // No: Billing instance
  // If you are using a free instance, you need to configure permissions in advance", "zh_CN":"是否免费实例，默认计费实例，裸机实例不能免费，取值：
  // YES：免费实例
  // NO：计费实例
  // 如果使用免费实例，需要提前配置权限"}
  IsFree *string `json:"isFree,omitempty" xml:"isFree,omitempty"`
  // {"en":"Specify a security group ID to create multiple security groups separated by commas, up to 5
  // If you are creating a bare machine, you cannot specify a security group", "zh_CN":"指定安全组id进行创建，多个安全组以逗号分隔，最多指定5个
  // 如果是创建裸机，不能指定安全组"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
  // {"en":"VMPCreateInstanceDisk information
  // If this information is carried, the disk definition on the template will be ignored and the instance disk will be created with this information, not for bare-metal instance creation", "zh_CN":"磁盘信息
  // 如果携带该信息，将忽略模板上的磁盘定义，以该信息创建实例磁盘，不适用于裸机实例创建"}
  DiskInfo []*VMPCreateInstanceDisk `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" type:"Repeated"`
  // {"en":"Anti-affinity group name
  // Virtual machines with the same ServerGroup are created on different hosts", "zh_CN":"反亲和性组名称
  // 拥有相同serverGroup的虚拟机会被创建在不同的宿主机上"}
  ServerGroup *string `json:"serverGroup,omitempty" xml:"serverGroup,omitempty"`
  // {"en":"Instance Tag", "zh_CN":"实例标签"}
  Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
  // {"en":"Use  unique ip segment", "zh_CN":"是否使用唯一网段
  // 1：是
  // -1：否"}
  UseUniqueIpSegment *int `json:"useUniqueIpSegment,omitempty" xml:"useUniqueIpSegment,omitempty"`
  // {"en":"Allocate IP randomly", "zh_CN":"是否需要随机分配IPv4
  // 1：是
  // -1：否"}
  RandomAllocateIp *int `json:"randomAllocateIp,omitempty" xml:"randomAllocateIp,omitempty"`
  // {"en":"Default Gateway", "zh_CN":"默认网关运营商如：dx-电信；yd-移动；wt-网通"}
  DefaultGateway *string `json:"defaultGateway,omitempty" xml:"defaultGateway,omitempty"`
  // {"en":"Policy routing type", "zh_CN":"策略路由类型：0-目的地址策略路由（默认）；1-源地址策略路由；"}
  PolicyRoutingType *int `json:"policyRoutingType,omitempty" xml:"policyRoutingType,omitempty"`
  // {"en":"Private gateway flag", "zh_CN":"内网网关标识：1-分配内网网关"}
  PrivateGatewayFlag *int `json:"privateGatewayFlag,omitempty" xml:"privateGatewayFlag,omitempty"`
  // {"en":"Nic allocate type", "zh_CN":"实例网卡分配方式：0-多个ip共用一张网卡（默认）；1-每个ip独立一张网卡；2-V4V6混合，同协议IP同网卡，不同线路IP不同网卡"}
  NicAllocateType *int `json:"nicAllocateType,omitempty" xml:"nicAllocateType,omitempty"`
  // {"en":"Ipv4 cidr", "zh_CN":"指定外网IPv4网段CIDR(不支持多线)"}
  SinglePublicIpv4Cidr *string `json:"singlePublicIpv4Cidr,omitempty" xml:"singlePublicIpv4Cidr,omitempty"`
}

func (s VMPCreateInstanceServer) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceServer) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceServer) SetRegionName(v string) *VMPCreateInstanceServer {
  s.RegionName = &v
  return s
}

func (s *VMPCreateInstanceServer) SetProvince(v string) *VMPCreateInstanceServer {
  s.Province = &v
  return s
}

func (s *VMPCreateInstanceServer) SetCarrier(v string) *VMPCreateInstanceServer {
  s.Carrier = &v
  return s
}

func (s *VMPCreateInstanceServer) SetNodeName(v string) *VMPCreateInstanceServer {
  s.NodeName = &v
  return s
}

func (s *VMPCreateInstanceServer) SetImageId(v string) *VMPCreateInstanceServer {
  s.ImageId = &v
  return s
}

func (s *VMPCreateInstanceServer) SetFlavorId(v string) *VMPCreateInstanceServer {
  s.FlavorId = &v
  return s
}

func (s *VMPCreateInstanceServer) SetName(v string) *VMPCreateInstanceServer {
  s.Name = &v
  return s
}

func (s *VMPCreateInstanceServer) SetUserData(v string) *VMPCreateInstanceServer {
  s.UserData = &v
  return s
}

func (s *VMPCreateInstanceServer) SetCount(v int) *VMPCreateInstanceServer {
  s.Count = &v
  return s
}

func (s *VMPCreateInstanceServer) SetPassword(v string) *VMPCreateInstanceServer {
  s.Password = &v
  return s
}

func (s *VMPCreateInstanceServer) SetKeyName(v string) *VMPCreateInstanceServer {
  s.KeyName = &v
  return s
}

func (s *VMPCreateInstanceServer) SetInnerNet(v string) *VMPCreateInstanceServer {
  s.InnerNet = &v
  return s
}

func (s *VMPCreateInstanceServer) SetCidr(v string) *VMPCreateInstanceServer {
  s.Cidr = &v
  return s
}

func (s *VMPCreateInstanceServer) SetPrivateIPv4(v string) *VMPCreateInstanceServer {
  s.PrivateIPv4 = &v
  return s
}

func (s *VMPCreateInstanceServer) SetInnerNet2(v string) *VMPCreateInstanceServer {
  s.InnerNet2 = &v
  return s
}

func (s *VMPCreateInstanceServer) SetCidr2(v string) *VMPCreateInstanceServer {
  s.Cidr2 = &v
  return s
}

func (s *VMPCreateInstanceServer) SetPrivateIPv42(v string) *VMPCreateInstanceServer {
  s.PrivateIPv42 = &v
  return s
}

func (s *VMPCreateInstanceServer) SetPrivateIpv6Info(v []*VMPCreateInstancePivateIpv6Info) *VMPCreateInstanceServer {
  s.PrivateIpv6Info = v
  return s
}

func (s *VMPCreateInstanceServer) SetProtocols(v int) *VMPCreateInstanceServer {
  s.Protocols = &v
  return s
}

func (s *VMPCreateInstanceServer) SetIpv4NativeAttribute(v string) *VMPCreateInstanceServer {
  s.Ipv4NativeAttribute = &v
  return s
}

func (s *VMPCreateInstanceServer) SetIpv6NativeAttribute(v string) *VMPCreateInstanceServer {
  s.Ipv6NativeAttribute = &v
  return s
}

func (s *VMPCreateInstanceServer) SetIsFree(v string) *VMPCreateInstanceServer {
  s.IsFree = &v
  return s
}

func (s *VMPCreateInstanceServer) SetSecurityGroupIds(v []*string) *VMPCreateInstanceServer {
  s.SecurityGroupIds = v
  return s
}

func (s *VMPCreateInstanceServer) SetDiskInfo(v []*VMPCreateInstanceDisk) *VMPCreateInstanceServer {
  s.DiskInfo = v
  return s
}

func (s *VMPCreateInstanceServer) SetServerGroup(v string) *VMPCreateInstanceServer {
  s.ServerGroup = &v
  return s
}

func (s *VMPCreateInstanceServer) SetTag(v string) *VMPCreateInstanceServer {
  s.Tag = &v
  return s
}

func (s *VMPCreateInstanceServer) SetUseUniqueIpSegment(v int) *VMPCreateInstanceServer {
  s.UseUniqueIpSegment = &v
  return s
}

func (s *VMPCreateInstanceServer) SetRandomAllocateIp(v int) *VMPCreateInstanceServer {
  s.RandomAllocateIp = &v
  return s
}

func (s *VMPCreateInstanceServer) SetDefaultGateway(v string) *VMPCreateInstanceServer {
  s.DefaultGateway = &v
  return s
}

func (s *VMPCreateInstanceServer) SetPolicyRoutingType(v int) *VMPCreateInstanceServer {
  s.PolicyRoutingType = &v
  return s
}

func (s *VMPCreateInstanceServer) SetPrivateGatewayFlag(v int) *VMPCreateInstanceServer {
  s.PrivateGatewayFlag = &v
  return s
}

func (s *VMPCreateInstanceServer) SetNicAllocateType(v int) *VMPCreateInstanceServer {
  s.NicAllocateType = &v
  return s
}

func (s *VMPCreateInstanceServer) SetSinglePublicIpv4Cidr(v string) *VMPCreateInstanceServer {
  s.SinglePublicIpv4Cidr = &v
  return s
}

type VMPCreateInstancePivateIpv6Info struct {
  // {"en":"Inner network number", "zh_CN":"内网编号（1-对应v4的内网1；2-对应v4的内网2）"}
  NetNo *int32 `json:"netNo,omitempty" xml:"netNo,omitempty"`
  // {"en":"Inner network ipv6 cidr", "zh_CN":"指定内网IPv6 CIDR"}
  PrivateCidr *string `json:"privateCidr,omitempty" xml:"privateCidr,omitempty"`
  // {"en":"Inner network ipv6 address:", "zh_CN":"指定内网IPv6地址"}
  PrivateIp *string `json:"privateIp,omitempty" xml:"privateIp,omitempty"`
}

func (s VMPCreateInstancePivateIpv6Info) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstancePivateIpv6Info) GoString() string {
  return s.String()
}

func (s *VMPCreateInstancePivateIpv6Info) SetNetNo(v int32) *VMPCreateInstancePivateIpv6Info {
  s.NetNo = &v
  return s
}

func (s *VMPCreateInstancePivateIpv6Info) SetPrivateCidr(v string) *VMPCreateInstancePivateIpv6Info {
  s.PrivateCidr = &v
  return s
}

func (s *VMPCreateInstancePivateIpv6Info) SetPrivateIp(v string) *VMPCreateInstancePivateIpv6Info {
  s.PrivateIp = &v
  return s
}

type VMPCreateInstanceDisk struct {
  // {"en":"VMPCreateInstanceDisk size (GB)", "zh_CN":"磁盘大小（GB）"}
  Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
  // {"en":"VMPCreateInstanceDisk Purpose: 
  // System - System disk;
  // DATA - DATA plate", "zh_CN":"磁盘用途：
  // SYSTEM-系统盘；
  // DATA-数据盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"VMPCreateInstanceDisk type: HDD/SSD", "zh_CN":"磁盘类型：HDD/SSD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  // {"en":"Is isolate: 1(Yes) / -1(No)", "zh_CN":"是否独立盘：1(是) / -1(否)"}
  IsIndependent *string `json:"isIndependent,omitempty" xml:"isIndependent,omitempty"`
}

func (s VMPCreateInstanceDisk) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceDisk) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceDisk) SetSize(v int32) *VMPCreateInstanceDisk {
  s.Size = &v
  return s
}

func (s *VMPCreateInstanceDisk) SetType(v string) *VMPCreateInstanceDisk {
  s.Type = &v
  return s
}

func (s *VMPCreateInstanceDisk) SetCategory(v string) *VMPCreateInstanceDisk {
  s.Category = &v
  return s
}

func (s *VMPCreateInstanceDisk) SetIsIndependent(v string) *VMPCreateInstanceDisk {
  s.IsIndependent = &v
  return s
}

type VMPCreateInstanceResponse struct {
  // {"en":"Virtual machine identity list", "zh_CN":"实例标识列表"}
  Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPCreateInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceResponse) GoString() string {
  return s.String()
}

func (s *VMPCreateInstanceResponse) SetServers(v []*string) *VMPCreateInstanceResponse {
  s.Servers = v
  return s
}

type VMPCreateInstancePaths struct {
}

func (s VMPCreateInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstancePaths) GoString() string {
  return s.String()
}

type VMPCreateInstanceParameters struct {
}

func (s VMPCreateInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceParameters) GoString() string {
  return s.String()
}

type VMPCreateInstanceRequestHeader struct {
}

func (s VMPCreateInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceRequestHeader) GoString() string {
  return s.String()
}

type VMPCreateInstanceResponseHeader struct {
}

func (s VMPCreateInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateInstanceResponseHeader) GoString() string {
  return s.String()
}




type InstanceDiskScalingRequest struct {
  // {"en":"instance id", "zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  DiskInfo []*InstanceDiskScalingDiskParam `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
}

func (s InstanceDiskScalingRequest) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingRequest) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingRequest) SetId(v string) *InstanceDiskScalingRequest {
  s.Id = &v
  return s
}

func (s *InstanceDiskScalingRequest) SetDiskInfo(v []*InstanceDiskScalingDiskParam) *InstanceDiskScalingRequest {
  s.DiskInfo = v
  return s
}

type InstanceDiskScalingDiskParam struct {
  // {"en":"disk size GB", "zh_CN":"磁盘大小，单位GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"disk type ,HDD/SDD", "zh_CN":"磁盘类型，取值：
  // HDD：普通硬盘
  // SSD：固态硬盘"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
}

func (s InstanceDiskScalingDiskParam) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingDiskParam) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingDiskParam) SetSize(v int) *InstanceDiskScalingDiskParam {
  s.Size = &v
  return s
}

func (s *InstanceDiskScalingDiskParam) SetCategory(v string) *InstanceDiskScalingDiskParam {
  s.Category = &v
  return s
}

type InstanceDiskScalingResponse struct {
  // {"en":"instance info", "zh_CN":"实例信息"}
  InstanceDiskScalingServer *InstanceDiskScalingServer `json:"server,omitempty" xml:"server,omitempty" require:"true"`
}

func (s InstanceDiskScalingResponse) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingResponse) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingResponse) SetServer(v *InstanceDiskScalingServer) *InstanceDiskScalingResponse {
  s.InstanceDiskScalingServer = v
  return s
}

type InstanceDiskScalingServer struct {
  // {"en":"server id", "zh_CN":"实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  DiskInfo []*InstanceDiskScalingDisk `json:"diskInfo,omitempty" xml:"diskInfo,omitempty" require:"true" type:"Repeated"`
}

func (s InstanceDiskScalingServer) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingServer) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingServer) SetId(v string) *InstanceDiskScalingServer {
  s.Id = &v
  return s
}

func (s *InstanceDiskScalingServer) SetDiskInfo(v []*InstanceDiskScalingDisk) *InstanceDiskScalingServer {
  s.DiskInfo = v
  return s
}

type InstanceDiskScalingDisk struct {
  // {"en":"disk category，HDD/SDD", "zh_CN":"磁盘类型，HDD/SDD"}
  Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true"`
  // {"en":"disk type ,DATA：data disk, SYSTEM：system disk", "zh_CN":"磁盘类型,DATA：数据盘, SYSTEM：系统盘"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"disk size GB", "zh_CN":"磁盘大小，单位GB"}
  Size *int `json:"size,omitempty" xml:"size,omitempty" require:"true"`
}

func (s InstanceDiskScalingDisk) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingDisk) GoString() string {
  return s.String()
}

func (s *InstanceDiskScalingDisk) SetCategory(v string) *InstanceDiskScalingDisk {
  s.Category = &v
  return s
}

func (s *InstanceDiskScalingDisk) SetType(v string) *InstanceDiskScalingDisk {
  s.Type = &v
  return s
}

func (s *InstanceDiskScalingDisk) SetSize(v int) *InstanceDiskScalingDisk {
  s.Size = &v
  return s
}

type InstanceDiskScalingPaths struct {
}

func (s InstanceDiskScalingPaths) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingPaths) GoString() string {
  return s.String()
}

type InstanceDiskScalingParameters struct {
}

func (s InstanceDiskScalingParameters) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingParameters) GoString() string {
  return s.String()
}

type InstanceDiskScalingRequestHeader struct {
}

func (s InstanceDiskScalingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingRequestHeader) GoString() string {
  return s.String()
}

type InstanceDiskScalingResponseHeader struct {
}

func (s InstanceDiskScalingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s InstanceDiskScalingResponseHeader) GoString() string {
  return s.String()
}




type EditInstanceRequest struct {
  // {"en":"server", "zh_CN":"实例信息对象"}
  EditInstanceServer []*EditInstanceServer `json:"server,omitempty" xml:"server,omitempty" require:"true" type:"Repeated"`
}

func (s EditInstanceRequest) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceRequest) GoString() string {
  return s.String()
}

func (s *EditInstanceRequest) SetServer(v []*EditInstanceServer) *EditInstanceRequest {
  s.EditInstanceServer = v
  return s
}

type EditInstanceServer struct {
  // {"en":"EditInstanceServer id", "zh_CN":"要更新的实例id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"New instance name
  // Constraints:
  // 1. Length of 2-128 characters
  // 2. Must start with a letter, and can only contain letters, numbers, underlines, lines, and dots", "zh_CN":"新的实例名称
  // 约束：
  // 1. 长度2-128个字符
  // 2. 必须以字母开头，且只能包含字母、数字、下划线、横线、点号"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s EditInstanceServer) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceServer) GoString() string {
  return s.String()
}

func (s *EditInstanceServer) SetId(v string) *EditInstanceServer {
  s.Id = &v
  return s
}

func (s *EditInstanceServer) SetName(v string) *EditInstanceServer {
  s.Name = &v
  return s
}

type EditInstanceResponse struct {
}

func (s EditInstanceResponse) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceResponse) GoString() string {
  return s.String()
}

type EditInstancePaths struct {
}

func (s EditInstancePaths) String() string {
  return tea.Prettify(s)
}

func (s EditInstancePaths) GoString() string {
  return s.String()
}

type EditInstanceParameters struct {
}

func (s EditInstanceParameters) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceParameters) GoString() string {
  return s.String()
}

type EditInstanceRequestHeader struct {
}

func (s EditInstanceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceRequestHeader) GoString() string {
  return s.String()
}

type EditInstanceResponseHeader struct {
}

func (s EditInstanceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditInstanceResponseHeader) GoString() string {
  return s.String()
}




