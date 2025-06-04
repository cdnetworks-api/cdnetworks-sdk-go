package templatemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type VMPQueryNodeServersRequest struct {
}

func (s VMPQueryNodeServersRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersRequest) GoString() string {
  return s.String()
}

type VMPQueryNodeServersResponse struct {
  // {"en":"Instance quantity information", "zh_CN":"实例数量信息列表"}
  Servers []*VMPQueryNodeServersServer `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQueryNodeServersResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeServersResponse) SetServers(v []*VMPQueryNodeServersServer) *VMPQueryNodeServersResponse {
  s.Servers = v
  return s
}

type VMPQueryNodeServersServer struct {
  // {"en":"Province", "zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty" require:"true"`
  // {"en":"Carrier", "zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty" require:"true"`
  // {"en":"Node Name", "zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Number of instances that can be created", "zh_CN":"可创建的实例的数量"}
  Count *string `json:"count,omitempty" xml:"count,omitempty" require:"true"`
}

func (s VMPQueryNodeServersServer) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersServer) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeServersServer) SetProvince(v string) *VMPQueryNodeServersServer {
  s.Province = &v
  return s
}

func (s *VMPQueryNodeServersServer) SetCarrier(v string) *VMPQueryNodeServersServer {
  s.Carrier = &v
  return s
}

func (s *VMPQueryNodeServersServer) SetNodeName(v string) *VMPQueryNodeServersServer {
  s.NodeName = &v
  return s
}

func (s *VMPQueryNodeServersServer) SetCount(v string) *VMPQueryNodeServersServer {
  s.Count = &v
  return s
}

type VMPQueryNodeServersPaths struct {
}

func (s VMPQueryNodeServersPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersPaths) GoString() string {
  return s.String()
}

type VMPQueryNodeServersParameters struct {
  // {"en":"Province", "zh_CN":"节点所属省份（详见附录2：https://www.wangsu.com/document/18204/isp-list?rsr=ws）"}
  Province *string `json:"province,omitempty" xml:"province,omitempty"`
  // {"en":"Carrier", "zh_CN":"节点所属运营商：dx-电信；wt-网通；yd-移动"}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {"en":"Node Name", "zh_CN":"节点名称，表示指定节点创建实例"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
  // {"en":"FlavorId", "zh_CN":"实例规格标识（可通过实例规格查询接口进行查询：https://www.wangsu.com/document/api-doc/22811?rsr=ws）"}
  FlavorId *string `json:"flavorId,omitempty" xml:"flavorId,omitempty" require:"true"`
  // {"en":"Protocols:
  // 4:only ipv4
  // 6:only ipv6
  // 0:both ipv4 & ipv6", "zh_CN":"是否需要多ip协议地址
  // 4：只需要ipv4地址，默认值
  // 6：只需要ipv6地址
  // 0：ipv4、ipv6都需要"}
  Protocols *int `json:"protocols,omitempty" xml:"protocols,omitempty"`
  // {"en":"Virtual machine network type:
  // ANY - the default;
  // DPDK - DPDK network.", "zh_CN":"虚拟机网络类型：
  // ANY-默认；
  // DPDK-DPDK网络。"}
  NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
  // {"en":"Use unique IP segment", "zh_CN":"是否使用唯一网段 
  // 1：是 
  // -1：否"}
  UseUniqueIpSegment *int `json:"useUniqueIpSegment,omitempty" xml:"useUniqueIpSegment,omitempty"`
}

func (s VMPQueryNodeServersParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryNodeServersParameters) SetProvince(v string) *VMPQueryNodeServersParameters {
  s.Province = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetCarrier(v string) *VMPQueryNodeServersParameters {
  s.Carrier = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetNodeName(v string) *VMPQueryNodeServersParameters {
  s.NodeName = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetFlavorId(v string) *VMPQueryNodeServersParameters {
  s.FlavorId = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetProtocols(v int) *VMPQueryNodeServersParameters {
  s.Protocols = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetNetworkType(v string) *VMPQueryNodeServersParameters {
  s.NetworkType = &v
  return s
}

func (s *VMPQueryNodeServersParameters) SetUseUniqueIpSegment(v int) *VMPQueryNodeServersParameters {
  s.UseUniqueIpSegment = &v
  return s
}

type VMPQueryNodeServersRequestHeader struct {
}

func (s VMPQueryNodeServersRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryNodeServersResponseHeader struct {
}

func (s VMPQueryNodeServersResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryNodeServersResponseHeader) GoString() string {
  return s.String()
}




