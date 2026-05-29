package networkmanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type VMPAllocateEdgeIPRequest struct {
  // {"en":"node name","zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"IP protocol:4-ipv4(default);6-ipv6(temporary unsupported)","zh_CN":"可选\nIP协议：4-ipv4(默认)；6-ipv6"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
  // {"en":"IP native attribute, 1: non-native;-1: native;","zh_CN":"IPv4原生属性，1：非原生；-1：原生；不指定默认随机分配原生属性"}
  NativeAttribute *string `json:"nativeAttribute,omitempty" xml:"nativeAttribute,omitempty"`
  // {"en":"cidr","zh_CN":"CIDR，一次只能传入一个CIDR"}
  Cidr *string `json:"cidr,omitempty" xml:"cidr,omitempty"`
  // {"en":"number of applications IP  (the single upper limit is 50)","zh_CN":"申请IP数（单次申请Ip数上限为50个）"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en":"Allocate IP randomly","zh_CN":"是否需要随机分配IP（仅对ipv4生效）\n1：是\n-1：否"}
  RandomAllocateIp *int `json:"randomAllocateIp,omitempty" xml:"randomAllocateIp,omitempty"`
}

func (s VMPAllocateEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *VMPAllocateEdgeIPRequest) SetNodeName(v string) *VMPAllocateEdgeIPRequest {
  s.NodeName = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetProtocol(v string) *VMPAllocateEdgeIPRequest {
  s.Protocol = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetNativeAttribute(v string) *VMPAllocateEdgeIPRequest {
  s.NativeAttribute = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetCidr(v string) *VMPAllocateEdgeIPRequest {
  s.Cidr = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetCount(v int) *VMPAllocateEdgeIPRequest {
  s.Count = &v
  return s
}

func (s *VMPAllocateEdgeIPRequest) SetRandomAllocateIp(v int) *VMPAllocateEdgeIPRequest {
  s.RandomAllocateIp = &v
  return s
}

type VMPAllocateEdgeIPRequestHeader struct {
}

func (s VMPAllocateEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPAllocateEdgeIPPaths struct {
}

func (s VMPAllocateEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPAllocateEdgeIPParameters struct {
}

func (s VMPAllocateEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPParameters) GoString() string {
  return s.String()
}

type VMPAllocateEdgeIPResponse struct {
  // {"en":"successful application for all or part of IP","zh_CN":"成功申请到的全部或部分IP\n说明：不同场景的响应说明如下\nA、所有IP都申请成功，返回申请到的所有IP\nB、只申请到部分IP，返回申请到的那部分IP\nC、未申请到任何IP，返回失败信息\nD、若出现申请失败的情况，请间隔10S之后再次申请"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPAllocateEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPAllocateEdgeIPResponse) SetEdgeIps(v []*string) *VMPAllocateEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type VMPAllocateEdgeIPResponseHeader struct {
}

func (s VMPAllocateEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPAllocateEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryEdgeIPRequest struct {
}

func (s VMPQueryEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPRequest) GoString() string {
  return s.String()
}

type VMPQueryEdgeIPRequestHeader struct {
}

func (s VMPQueryEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryEdgeIPPaths struct {
}

func (s VMPQueryEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPQueryEdgeIPParameters struct {
  // {"en":"node name","zh_CN":"可选\n节点名称，多个用英文逗号分隔"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"virtual machine ID","zh_CN":"可选\n虚拟机ID，多个用英文逗号分隔"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty"`
  // {"en":"virtual machine master IP","zh_CN":"可选\n虚拟机主IP，多个用英文逗号分隔"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty"`
  // {"en":"extra Ip","zh_CN":"可选\n额外Ip，多个用英文逗号分隔"}
  EdgeIp *string `json:"edgeIp,omitempty" xml:"edgeIp,omitempty"`
  // {"en":"IP state","zh_CN":"可选\nIP状态：FREE-空闲未绑定；ASSIGNED-已绑定"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s VMPQueryEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryEdgeIPParameters) SetNodeName(v string) *VMPQueryEdgeIPParameters {
  s.NodeName = &v
  return s
}

func (s *VMPQueryEdgeIPParameters) SetServerId(v string) *VMPQueryEdgeIPParameters {
  s.ServerId = &v
  return s
}

func (s *VMPQueryEdgeIPParameters) SetServerIp(v string) *VMPQueryEdgeIPParameters {
  s.ServerIp = &v
  return s
}

func (s *VMPQueryEdgeIPParameters) SetEdgeIp(v string) *VMPQueryEdgeIPParameters {
  s.EdgeIp = &v
  return s
}

func (s *VMPQueryEdgeIPParameters) SetState(v string) *VMPQueryEdgeIPParameters {
  s.State = &v
  return s
}

type VMPQueryEdgeIPResponse struct {
  // {"en":"additional IP details","zh_CN":"额外Ip详细信息"}
  EdgeIps []*VMPQueryEdgeIPResponseEdgeIps `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryEdgeIPResponse) SetEdgeIps(v []*VMPQueryEdgeIPResponseEdgeIps) *VMPQueryEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type VMPQueryEdgeIPResponseEdgeIps struct     {
  // {"en":"extra Ip","zh_CN":"额外Ip"}
  EdgeIp *string `json:"edgeIp,omitempty" xml:"edgeIp,omitempty" require:"true"`
  // {"en":"IP state","zh_CN":"IP状态：FREE-空闲未绑定；ASSIGNED-已绑定"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
  // {"en":"binding virtual machine ID","zh_CN":"绑定的虚拟机ID"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"binding virtual machine extranet IP","zh_CN":"绑定的虚拟机外网IP"}
  ServerIp []*string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true" type:"Repeated"`
  // {"en":"node name","zh_CN":"所属节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Is it exclusive","zh_CN":"是否独占"}
  OccupancyFlag *bool `json:"occupancyFlag,omitempty" xml:"occupancyFlag,omitempty" require:"true"`
  // {"en":"Net Mask","zh_CN":"子网掩码"}
  Netmask *string `json:"netmask,omitempty" xml:"netmask,omitempty" require:"true"`
}

func (s VMPQueryEdgeIPResponseEdgeIps) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPResponseEdgeIps) GoString() string {
  return s.String()
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetEdgeIp(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.EdgeIp = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetState(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.State = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetServerId(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.ServerId = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetServerIp(v []*string) *VMPQueryEdgeIPResponseEdgeIps {
  s.ServerIp = v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetNodeName(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.NodeName = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetOccupancyFlag(v bool) *VMPQueryEdgeIPResponseEdgeIps {
  s.OccupancyFlag = &v
  return s
}

func (s *VMPQueryEdgeIPResponseEdgeIps) SetNetmask(v string) *VMPQueryEdgeIPResponseEdgeIps {
  s.Netmask = &v
  return s
}

type VMPQueryEdgeIPResponseHeader struct {
}

func (s VMPQueryEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type VMPAssignEdgeIPRequest struct {
  // {"en":"target virtual machine external network Ip","zh_CN":"目标实例的公网Ip"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional IP to bind to the virtual machine","zh_CN":"要绑定到目标实例的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPAssignEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *VMPAssignEdgeIPRequest) SetServerIp(v string) *VMPAssignEdgeIPRequest {
  s.ServerIp = &v
  return s
}

func (s *VMPAssignEdgeIPRequest) SetEdgeIps(v []*string) *VMPAssignEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type VMPAssignEdgeIPRequestHeader struct {
}

func (s VMPAssignEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPAssignEdgeIPPaths struct {
}

func (s VMPAssignEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPAssignEdgeIPParameters struct {
}

func (s VMPAssignEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPParameters) GoString() string {
  return s.String()
}

type VMPAssignEdgeIPResponse struct {
  // {"en":"target virtual machine IP","zh_CN":"目标实例公网IP"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"Additional IP that is bound to the instance","zh_CN":"已绑定到实例的额外公网IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPAssignEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPAssignEdgeIPResponse) SetServerIp(v string) *VMPAssignEdgeIPResponse {
  s.ServerIp = &v
  return s
}

func (s *VMPAssignEdgeIPResponse) SetEdgeIps(v []*string) *VMPAssignEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type VMPAssignEdgeIPResponseHeader struct {
}

func (s VMPAssignEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPAssignEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type VMPReleaseEdgeIPRequest struct {
  // {"en":"additional IP to be released","zh_CN":"要释放的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPReleaseEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *VMPReleaseEdgeIPRequest) SetEdgeIps(v []*string) *VMPReleaseEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type VMPReleaseEdgeIPRequestHeader struct {
}

func (s VMPReleaseEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPReleaseEdgeIPPaths struct {
}

func (s VMPReleaseEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPReleaseEdgeIPParameters struct {
}

func (s VMPReleaseEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPParameters) GoString() string {
  return s.String()
}

type VMPReleaseEdgeIPResponse struct {
  // {"en":"Error Message Set","zh_CN":"错误信息集"}
  BatchErrorMsg []*VMPReleaseEdgeIPResponseBatchErrorMsg `json:"batchErrorMsg,omitempty" xml:"batchErrorMsg,omitempty" require:"true" type:"Repeated"`
}

func (s VMPReleaseEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPReleaseEdgeIPResponse) SetBatchErrorMsg(v []*VMPReleaseEdgeIPResponseBatchErrorMsg) *VMPReleaseEdgeIPResponse {
  s.BatchErrorMsg = v
  return s
}

type VMPReleaseEdgeIPResponseBatchErrorMsg struct     {
  // {"en":"Ip","zh_CN":"Ip"}
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  // {"en":"Message","zh_CN":"消息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s VMPReleaseEdgeIPResponseBatchErrorMsg) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPResponseBatchErrorMsg) GoString() string {
  return s.String()
}

func (s *VMPReleaseEdgeIPResponseBatchErrorMsg) SetKey(v string) *VMPReleaseEdgeIPResponseBatchErrorMsg {
  s.Key = &v
  return s
}

func (s *VMPReleaseEdgeIPResponseBatchErrorMsg) SetMsg(v string) *VMPReleaseEdgeIPResponseBatchErrorMsg {
  s.Msg = &v
  return s
}

type VMPReleaseEdgeIPResponseHeader struct {
}

func (s VMPReleaseEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPReleaseEdgeIPResponseHeader) GoString() string {
  return s.String()
}




type VMPUnassignEdgeIPRequest struct {
  // {"en":"target virtual machine external network Ip","zh_CN":"目标实例外网Ip"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional Ip to unbind","zh_CN":"要解除绑定的额外IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPUnassignEdgeIPRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPRequest) GoString() string {
  return s.String()
}

func (s *VMPUnassignEdgeIPRequest) SetServerIp(v string) *VMPUnassignEdgeIPRequest {
  s.ServerIp = &v
  return s
}

func (s *VMPUnassignEdgeIPRequest) SetEdgeIps(v []*string) *VMPUnassignEdgeIPRequest {
  s.EdgeIps = v
  return s
}

type VMPUnassignEdgeIPRequestHeader struct {
}

func (s VMPUnassignEdgeIPRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPRequestHeader) GoString() string {
  return s.String()
}

type VMPUnassignEdgeIPPaths struct {
}

func (s VMPUnassignEdgeIPPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPPaths) GoString() string {
  return s.String()
}

type VMPUnassignEdgeIPParameters struct {
}

func (s VMPUnassignEdgeIPParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPParameters) GoString() string {
  return s.String()
}

type VMPUnassignEdgeIPResponse struct {
  // {"en":"target virtual machine IP","zh_CN":"目标实例公网IP"}
  ServerIp *string `json:"serverIp,omitempty" xml:"serverIp,omitempty" require:"true"`
  // {"en":"additional Ip that has been bound to the virtual machine","zh_CN":"已绑定到实例的额外公网IP"}
  EdgeIps []*string `json:"edgeIps,omitempty" xml:"edgeIps,omitempty" require:"true" type:"Repeated"`
}

func (s VMPUnassignEdgeIPResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPResponse) GoString() string {
  return s.String()
}

func (s *VMPUnassignEdgeIPResponse) SetServerIp(v string) *VMPUnassignEdgeIPResponse {
  s.ServerIp = &v
  return s
}

func (s *VMPUnassignEdgeIPResponse) SetEdgeIps(v []*string) *VMPUnassignEdgeIPResponse {
  s.EdgeIps = v
  return s
}

type VMPUnassignEdgeIPResponseHeader struct {
}

func (s VMPUnassignEdgeIPResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPUnassignEdgeIPResponseHeader) GoString() string {
  return s.String()
}




