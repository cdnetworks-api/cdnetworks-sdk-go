package security

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type EditSecurityGroupRulesRequest struct {
  SecurityGroupRule []*EditSecurityGroupRulesSecurityGroupRuleParam `json:"securityGroupRule,omitempty" xml:"securityGroupRule,omitempty" require:"true" type:"Repeated"`
}

func (s EditSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupRulesRequest) SetSecurityGroupRule(v []*EditSecurityGroupRulesSecurityGroupRuleParam) *EditSecurityGroupRulesRequest {
  s.SecurityGroupRule = v
  return s
}

type EditSecurityGroupRulesSecurityGroupRuleParam struct {
  // {"en":"Security Group id", "zh_CN":"规则id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS", "zh_CN":"规则方向： INGRESS-流入 EGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP", "zh_CN":"授权策略： ACCEPT-允许 DROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP.
  // When selecting ICMP, the maximum/minimum port numbers below cannot be specified.
  // ", "zh_CN":"协议类型：TCP/UDP/ICMP 选择ICMP时，不能指定下方的最大/最小端口号"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:4-IPv4  6-IPv6", "zh_CN":"IP协议： 4-IPv4 6-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority: 1-100, higher the value, higher the priority", "zh_CN":"优先级：取值范围1-100，值越大优先级越高"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"The maximum port number.
  // 1) Value 1-65535 and greater than or equal to the minimum port number.
  // 2) When specifying a specific port, the maximum and minimum port numbers should be equal.
  // ", "zh_CN":"最大的端口号
  // 1）取值1-65535且大于等于最小端口号
  // 2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty"`
  // {"en":"The minimum port number.
  // 1) Value 1-65535 and less than or equal to the maximum port number.
  // 2) When specifying a specific port, the maximum and minimum port numbers should be equal.
  // ", "zh_CN":"最小的端口号
  // 1）取值1-65535且小于等于最大端口号
  // 2）当要指定某个特定端口时，最大和最小端口号取值相等即可'"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty"`
  // {"en":"Authorization object: The IP or CIDR that matches the security group rules.
  // 1) Only a single IP or CIDR can be specified.
  // 2) The specified IP or CIDR must match the IP protocol type mentioned above.
  // ", "zh_CN":"授权对象：匹配该安全组规则的ip或cidr
  // 1）只能指定单个IP或CIDR
  // 2）指定的IP或CIDR必须和上述的IP协议类型一致'"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes,cannot exceed 255 characters", "zh_CN":"备注信息（可选），不能超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s EditSecurityGroupRulesSecurityGroupRuleParam) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesSecurityGroupRuleParam) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetId(v string) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.Id = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetDirect(v string) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.Direct = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetPolicy(v string) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.Policy = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetProtocol(v string) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.Protocol = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetEthertype(v string) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.Ethertype = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetPriority(v int) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.Priority = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetPortRangeMax(v int) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.PortRangeMax = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetPortRangeMin(v int) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.PortRangeMin = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetRemoteIpPerfix(v string) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.RemoteIpPerfix = &v
  return s
}

func (s *EditSecurityGroupRulesSecurityGroupRuleParam) SetRemark(v string) *EditSecurityGroupRulesSecurityGroupRuleParam {
  s.Remark = &v
  return s
}

type EditSecurityGroupRulesResponse struct {
}

func (s EditSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

type EditSecurityGroupRulesPaths struct {
}

func (s EditSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

type EditSecurityGroupRulesParameters struct {
}

func (s EditSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type EditSecurityGroupRulesRequestHeader struct {
}

func (s EditSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type EditSecurityGroupRulesResponseHeader struct {
}

func (s EditSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type AddSecurityGroupRulesRequest struct {
  SecurityGroupRules []*AddSecurityGroupRulesSecurityGroupRuleParam `json:"securityGroupRules,omitempty" xml:"securityGroupRules,omitempty" require:"true" type:"Repeated"`
}

func (s AddSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesRequest) SetSecurityGroupRules(v []*AddSecurityGroupRulesSecurityGroupRuleParam) *AddSecurityGroupRulesRequest {
  s.SecurityGroupRules = v
  return s
}

type AddSecurityGroupRulesSecurityGroupRuleParam struct {
  // {"en":"Security Group id", "zh_CN":"安全组id"}
  SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS", "zh_CN":"规则方向：
  // INGRESS-流入
  // EGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP", "zh_CN":"授权策略：
  // ACCEPT-允许
  // DROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP.
  // When selecting ICMP, the maximum/minimum port numbers below cannot be specified.", "zh_CN":"协议类型：TCP/UDP/ICMP
  // 选择ICMP时，不能指定下方的最大/最小端口号"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:
  // 4-IPv4
  // 6-IPv6", "zh_CN":"IP协议：
  // 4-IPv4
  // 6-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority: 1-100, higher the value, higher the priority", "zh_CN":"优先级：取值范围1-100，值越大优先级越高"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"The maximum port number.
  // 1) Value 1-65535 and greater than or equal to the minimum port number.
  // 2) When specifying a specific port, the maximum and minimum port numbers should be equal.",
  // "zh_CN":"最大的端口号
  // 1）取值1-65535且大于等于最小端口号
  // 2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty"`
  // {"en":"The minimum port number.
  // 1) Value 1-65535 and less than or equal to the maximum port number.
  // 2) When specifying a specific port, the maximum and minimum port numbers should be equal.
  // ", "zh_CN":"最小的端口号
  // 1）取值1-65535且小于等于最大端口号
  // 2）当要指定某个特定端口时，最大和最小端口号取值相等即可"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty"`
  // {"en":"Authorization object: The IP or CIDR that matches the security group rules.
  // 1) Only a single IP or CIDR can be specified.
  // 2) The specified IP or CIDR must match the IP protocol type mentioned above.", "zh_CN":"授权对象：匹配该安全组规则的ip或cidr
  // 1）只能指定单个IP或CIDR
  // 2）指定的IP或CIDR必须和上述的IP协议类型一致"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes,cannot exceed 255 characters", "zh_CN":"备注信息，不能超过255个字符"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s AddSecurityGroupRulesSecurityGroupRuleParam) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesSecurityGroupRuleParam) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetSecurityGroupId(v string) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.SecurityGroupId = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetDirect(v string) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.Direct = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetPolicy(v string) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.Policy = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetProtocol(v string) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.Protocol = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetEthertype(v string) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.Ethertype = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetPriority(v int) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.Priority = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetPortRangeMax(v int) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.PortRangeMax = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetPortRangeMin(v int) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.PortRangeMin = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetRemoteIpPerfix(v string) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.RemoteIpPerfix = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRuleParam) SetRemark(v string) *AddSecurityGroupRulesSecurityGroupRuleParam {
  s.Remark = &v
  return s
}

type AddSecurityGroupRulesResponse struct {
  SecurityGroupRules []*AddSecurityGroupRulesSecurityGroupRule `json:"securityGroupRules,omitempty" xml:"securityGroupRules,omitempty" require:"true" type:"Repeated"`
}

func (s AddSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesResponse) SetSecurityGroupRules(v []*AddSecurityGroupRulesSecurityGroupRule) *AddSecurityGroupRulesResponse {
  s.SecurityGroupRules = v
  return s
}

type AddSecurityGroupRulesSecurityGroupRule struct {
  // {"en":"Rule id", "zh_CN":"规则id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Rule direction:INGRESS;EGRESS", "zh_CN":"规则方向：
  // INGRESS-流入
  // EGRESS-流出"}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {"en":"Authorization policy:ACCEPT;DROP", "zh_CN":"授权策略：
  // ACCEPT-允许
  // DROP-拒绝"}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {"en":"Protocol type: TCP/UDP/ICMP", "zh_CN":"协议类型：TCP/UDP/ICMP"}
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  // {"en":"IP protocol:
  // 4-IPv4
  // 6-IPv6
  // ", "zh_CN":"IP协议：
  // 4-IPv4
  // 6-IPv6"}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {"en":"Priority, the higher the value, the higher the priority", "zh_CN":"优先级，值越大优先级越高"}
  Priority *int `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {"en":"Maximum Port Number", "zh_CN":"最大的端口号"}
  PortRangeMax *int `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty" require:"true"`
  // {"en":"Minimum Port Number", "zh_CN":"最小的端口号"}
  PortRangeMin *int `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty" require:"true"`
  // {"en":"Authorization object: IP or CIDR that matches the security group rules", "zh_CN":"授权对象：匹配该安全组规则的IP或CIDR"}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
  // {"en":"Notes", "zh_CN":"备注信息"}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {"en":"Create Time", "zh_CN":"创建时间"}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"Modify Time", "zh_CN":"修改时间"}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s AddSecurityGroupRulesSecurityGroupRule) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesSecurityGroupRule) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetId(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.Id = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetDirect(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.Direct = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetPolicy(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.Policy = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetProtocol(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.Protocol = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetEthertype(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.Ethertype = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetPriority(v int) *AddSecurityGroupRulesSecurityGroupRule {
  s.Priority = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetPortRangeMax(v int) *AddSecurityGroupRulesSecurityGroupRule {
  s.PortRangeMax = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetPortRangeMin(v int) *AddSecurityGroupRulesSecurityGroupRule {
  s.PortRangeMin = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetRemoteIpPerfix(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.RemoteIpPerfix = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetRemark(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.Remark = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetCreateTime(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.CreateTime = &v
  return s
}

func (s *AddSecurityGroupRulesSecurityGroupRule) SetModifyTime(v string) *AddSecurityGroupRulesSecurityGroupRule {
  s.ModifyTime = &v
  return s
}

type AddSecurityGroupRulesPaths struct {
}

func (s AddSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

type AddSecurityGroupRulesParameters struct {
}

func (s AddSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type AddSecurityGroupRulesRequestHeader struct {
}

func (s AddSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type AddSecurityGroupRulesResponseHeader struct {
}

func (s AddSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type VMPBindSecurityGroupRequest struct {
  BindInfo []*VMPBindSecurityGroupParam `json:"bindInfo,omitempty" xml:"bindInfo,omitempty" require:"true" type:"Repeated"`
}

func (s VMPBindSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *VMPBindSecurityGroupRequest) SetBindInfo(v []*VMPBindSecurityGroupParam) *VMPBindSecurityGroupRequest {
  s.BindInfo = v
  return s
}

type VMPBindSecurityGroupParam struct {
  // {"en":"Instance ID", "zh_CN":"实例id"}
  ServerId *string `json:"serverId,omitempty" xml:"serverId,omitempty" require:"true"`
  // {"en":"Security Group ID", "zh_CN":"安全组id"}
  SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
}

func (s VMPBindSecurityGroupParam) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupParam) GoString() string {
  return s.String()
}

func (s *VMPBindSecurityGroupParam) SetServerId(v string) *VMPBindSecurityGroupParam {
  s.ServerId = &v
  return s
}

func (s *VMPBindSecurityGroupParam) SetSecurityGroupIds(v []*string) *VMPBindSecurityGroupParam {
  s.SecurityGroupIds = v
  return s
}

type VMPBindSecurityGroupResponse struct {
}

func (s VMPBindSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupResponse) GoString() string {
  return s.String()
}

type VMPBindSecurityGroupPaths struct {
}

func (s VMPBindSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupPaths) GoString() string {
  return s.String()
}

type VMPBindSecurityGroupParameters struct {
}

func (s VMPBindSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupParameters) GoString() string {
  return s.String()
}

type VMPBindSecurityGroupRequestHeader struct {
}

func (s VMPBindSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type VMPBindSecurityGroupResponseHeader struct {
}

func (s VMPBindSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPBindSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type DeleteSecurityGroupRequest struct {
}

func (s DeleteSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
  return s.String()
}

type DeleteSecurityGroupResponse struct {
}

func (s DeleteSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
  return s.String()
}

type DeleteSecurityGroupPaths struct {
  // {"en":"Security Group id,multiple values separated by commas", "zh_CN":"安全组id，多个用逗号分隔"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupPaths) GoString() string {
  return s.String()
}

func (s *DeleteSecurityGroupPaths) SetId(v string) *DeleteSecurityGroupPaths {
  s.Id = &v
  return s
}

type DeleteSecurityGroupParameters struct {
}

func (s DeleteSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupParameters) GoString() string {
  return s.String()
}

type DeleteSecurityGroupRequestHeader struct {
}

func (s DeleteSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type DeleteSecurityGroupResponseHeader struct {
}

func (s DeleteSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type DeletionSecurityGroupRulesRequest struct {
}

func (s DeletionSecurityGroupRulesRequest) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesRequest) GoString() string {
  return s.String()
}

type DeletionSecurityGroupRulesResponse struct {
}

func (s DeletionSecurityGroupRulesResponse) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesResponse) GoString() string {
  return s.String()
}

type DeletionSecurityGroupRulesPaths struct {
  // {"en":"Security group rule id, multiple separated by commas", "zh_CN":"安全组规则id，多个用逗号分隔"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeletionSecurityGroupRulesPaths) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesPaths) GoString() string {
  return s.String()
}

func (s *DeletionSecurityGroupRulesPaths) SetId(v string) *DeletionSecurityGroupRulesPaths {
  s.Id = &v
  return s
}

type DeletionSecurityGroupRulesParameters struct {
}

func (s DeletionSecurityGroupRulesParameters) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesParameters) GoString() string {
  return s.String()
}

type DeletionSecurityGroupRulesRequestHeader struct {
}

func (s DeletionSecurityGroupRulesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesRequestHeader) GoString() string {
  return s.String()
}

type DeletionSecurityGroupRulesResponseHeader struct {
}

func (s DeletionSecurityGroupRulesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletionSecurityGroupRulesResponseHeader) GoString() string {
  return s.String()
}




type VMPCreateSSHKeyRequest struct {
  // {"en":"Creating array objects for SSH key", "zh_CN":"创建SSH密钥对的数组对象"}
  Keypair []*VMPCreateSSHKeyKeypairCreate `json:"keypair,omitempty" xml:"keypair,omitempty" require:"true" type:"Repeated"`
}

func (s VMPCreateSSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyRequest) GoString() string {
  return s.String()
}

func (s *VMPCreateSSHKeyRequest) SetKeypair(v []*VMPCreateSSHKeyKeypairCreate) *VMPCreateSSHKeyRequest {
  s.Keypair = v
  return s
}

type VMPCreateSSHKeyKeypairCreate struct {
  // {"en":"Key pair name, must be unique, otherwise creation will fail", "zh_CN":"密钥对名称，必须保持唯一，否则会创建失败"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Imported public key information. The maximum length of the import public key is 1024 bytes. If this parameter is not specified, it will be generated automatically.", "zh_CN":"导入的公钥信息。导入公钥最大长度为1024字节。如果未指定该参数，系统将自动生成。"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty"`
}

func (s VMPCreateSSHKeyKeypairCreate) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyKeypairCreate) GoString() string {
  return s.String()
}

func (s *VMPCreateSSHKeyKeypairCreate) SetName(v string) *VMPCreateSSHKeyKeypairCreate {
  s.Name = &v
  return s
}

func (s *VMPCreateSSHKeyKeypairCreate) SetPublicKey(v string) *VMPCreateSSHKeyKeypairCreate {
  s.PublicKey = &v
  return s
}

type VMPCreateSSHKeyKeypairResult struct {
  // {"en":"Key pair name, must be unique, otherwise creation will fail", "zh_CN":"密钥对名称，必须保持唯一，否则会创建失败"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Imported public key information. The maximum length of the import public key is 1024 bytes. If this parameter is not specified, it will be generated automatically.", "zh_CN":"导入的公钥信息。导入公钥最大长度为1024字节。如果未指定该参数，系统将自动生成。"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty" require:"true"`
  // {"en":"The key corresponds to the privatekey information. If publickey is given in the request parameter, the privatekey will not be returned. Otherwise, the system generated private key will be returned.", "zh_CN":"密钥对应privateKey信息，如果请求参数中给了publicKey，则不会返回privateKey，否则返回系统生成的私钥。"}
  PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty" require:"true"`
}

func (s VMPCreateSSHKeyKeypairResult) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyKeypairResult) GoString() string {
  return s.String()
}

func (s *VMPCreateSSHKeyKeypairResult) SetName(v string) *VMPCreateSSHKeyKeypairResult {
  s.Name = &v
  return s
}

func (s *VMPCreateSSHKeyKeypairResult) SetPublicKey(v string) *VMPCreateSSHKeyKeypairResult {
  s.PublicKey = &v
  return s
}

func (s *VMPCreateSSHKeyKeypairResult) SetPrivateKey(v string) *VMPCreateSSHKeyKeypairResult {
  s.PrivateKey = &v
  return s
}

type VMPCreateSSHKeyResponse struct {
  Keypair []*VMPCreateSSHKeyKeypairResult `json:"keypair,omitempty" xml:"keypair,omitempty" require:"true" type:"Repeated"`
}

func (s VMPCreateSSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyResponse) GoString() string {
  return s.String()
}

func (s *VMPCreateSSHKeyResponse) SetKeypair(v []*VMPCreateSSHKeyKeypairResult) *VMPCreateSSHKeyResponse {
  s.Keypair = v
  return s
}

type VMPCreateSSHKeyPaths struct {
}

func (s VMPCreateSSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyPaths) GoString() string {
  return s.String()
}

type VMPCreateSSHKeyParameters struct {
}

func (s VMPCreateSSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyParameters) GoString() string {
  return s.String()
}

type VMPCreateSSHKeyRequestHeader struct {
}

func (s VMPCreateSSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyRequestHeader) GoString() string {
  return s.String()
}

type VMPCreateSSHKeyResponseHeader struct {
}

func (s VMPCreateSSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateSSHKeyResponseHeader) GoString() string {
  return s.String()
}




type EditSecurityGroupRequest struct {
  SecurityGroup []*EditSecurityGroupSecurityGroupParam `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" require:"true" type:"Repeated"`
}

func (s EditSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupRequest) SetSecurityGroup(v []*EditSecurityGroupSecurityGroupParam) *EditSecurityGroupRequest {
  s.SecurityGroup = v
  return s
}

type EditSecurityGroupSecurityGroupParam struct {
  // {'en':'Security Group id', 'zh_CN':'安全组id'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Security Group name.
  //     Constraint: can only consist of letters, numbers, and underscores, with a length of no more than 64 characters.
  //     It cannot be named default and must be unique.',
  //     'zh_CN':'安全组名称 约束：只能由字母/数字/下划线组成，长度不超过64个字符，不能命名为default且名称唯一'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Notes,not exceeding 255 characters in length.', 'zh_CN':'备注信息，长度不超过255个字符'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s EditSecurityGroupSecurityGroupParam) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupSecurityGroupParam) GoString() string {
  return s.String()
}

func (s *EditSecurityGroupSecurityGroupParam) SetId(v string) *EditSecurityGroupSecurityGroupParam {
  s.Id = &v
  return s
}

func (s *EditSecurityGroupSecurityGroupParam) SetName(v string) *EditSecurityGroupSecurityGroupParam {
  s.Name = &v
  return s
}

func (s *EditSecurityGroupSecurityGroupParam) SetRemark(v string) *EditSecurityGroupSecurityGroupParam {
  s.Remark = &v
  return s
}

type EditSecurityGroupResponse struct {
}

func (s EditSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupResponse) GoString() string {
  return s.String()
}

type EditSecurityGroupPaths struct {
}

func (s EditSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupPaths) GoString() string {
  return s.String()
}

type EditSecurityGroupParameters struct {
}

func (s EditSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupParameters) GoString() string {
  return s.String()
}

type EditSecurityGroupRequestHeader struct {
}

func (s EditSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type EditSecurityGroupResponseHeader struct {
}

func (s EditSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type VMPRemoveSSHKeyRequest struct {
  // {"en":"ssh key name", "zh_CN":"ssh key 名称"}
  Keyname *string `json:"keyname,omitempty" xml:"keyname,omitempty" require:"true"`
}

func (s VMPRemoveSSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyRequest) GoString() string {
  return s.String()
}

func (s *VMPRemoveSSHKeyRequest) SetKeyname(v string) *VMPRemoveSSHKeyRequest {
  s.Keyname = &v
  return s
}

type VMPRemoveSSHKeyResponse struct {
}

func (s VMPRemoveSSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyResponse) GoString() string {
  return s.String()
}

type VMPRemoveSSHKeyPaths struct {
  // {"en":"ssh key name", "zh_CN":"ssh key 名称"}
  KeyName *string `json:"keyName,omitempty" xml:"keyName,omitempty" require:"true"`
}

func (s VMPRemoveSSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyPaths) GoString() string {
  return s.String()
}

func (s *VMPRemoveSSHKeyPaths) SetKeyName(v string) *VMPRemoveSSHKeyPaths {
  s.KeyName = &v
  return s
}

type VMPRemoveSSHKeyParameters struct {
}

func (s VMPRemoveSSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyParameters) GoString() string {
  return s.String()
}

type VMPRemoveSSHKeyRequestHeader struct {
}

func (s VMPRemoveSSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyRequestHeader) GoString() string {
  return s.String()
}

type VMPRemoveSSHKeyResponseHeader struct {
}

func (s VMPRemoveSSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveSSHKeyResponseHeader) GoString() string {
  return s.String()
}




type VMPQuerySecurityGroupRequest struct {
}

func (s VMPQuerySecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupRequest) GoString() string {
  return s.String()
}

type VMPQuerySecurityGroupResponse struct {
  SecurityGroups []*VMPQuerySecurityGroupSecurityGroup `json:"securityGroups,omitempty" xml:"securityGroups,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQuerySecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupResponse) SetSecurityGroups(v []*VMPQuerySecurityGroupSecurityGroup) *VMPQuerySecurityGroupResponse {
  s.SecurityGroups = v
  return s
}

type VMPQuerySecurityGroupSecurityGroup struct {
  // {'en':'Security Group ID', 'zh_CN':'安全组id'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Security Group name', 'zh_CN':'安全组名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Notes', 'zh_CN':'备注信息'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {'en':'Create Time:yyyy-MM-dd HH:mm:ss', 'zh_CN':'创建时间（yyyy-MM-dd HH:mm:ss）'}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {'en':'Modify Time:yyyy-MM-dd HH:mm:ss', 'zh_CN':'修改时间（yyyy-MM-dd HH:mm:ss）'}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
  Rules []*VMPQuerySecurityGroupSecurityGroupRule `json:"rules,omitempty" xml:"rules,omitempty" require:"true" type:"Repeated"`
}

func (s VMPQuerySecurityGroupSecurityGroup) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupSecurityGroup) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupSecurityGroup) SetId(v string) *VMPQuerySecurityGroupSecurityGroup {
  s.Id = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroup) SetName(v string) *VMPQuerySecurityGroupSecurityGroup {
  s.Name = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroup) SetRemark(v string) *VMPQuerySecurityGroupSecurityGroup {
  s.Remark = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroup) SetCreateTime(v string) *VMPQuerySecurityGroupSecurityGroup {
  s.CreateTime = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroup) SetModifyTime(v string) *VMPQuerySecurityGroupSecurityGroup {
  s.ModifyTime = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroup) SetRules(v []*VMPQuerySecurityGroupSecurityGroupRule) *VMPQuerySecurityGroupSecurityGroup {
  s.Rules = v
  return s
}

type VMPQuerySecurityGroupSecurityGroupRule struct {
  // {'en':'Rule direction', 'zh_CN':'规则方向'}
  Direct *string `json:"direct,omitempty" xml:"direct,omitempty" require:"true"`
  // {'en':'Authorization policy', 'zh_CN':'授权策略'}
  Policy *string `json:"policy,omitempty" xml:"policy,omitempty" require:"true"`
  // {'en':'IP protocol', 'zh_CN':'ip协议'}
  Ethertype *string `json:"ethertype,omitempty" xml:"ethertype,omitempty" require:"true"`
  // {'en':'Priority', 'zh_CN':'优先级'}
  Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
  // {'en':'Maximum Port Number', 'zh_CN':'最大端口号'}
  PortRangeMax *int32 `json:"portRangeMax,omitempty" xml:"portRangeMax,omitempty" require:"true"`
  // {'en':'Minimum Port Number', 'zh_CN':'最小端口号'}
  PortRangeMin *int32 `json:"portRangeMin,omitempty" xml:"portRangeMin,omitempty" require:"true"`
  // {'en':'IP or CIDR matching rules', 'zh_CN':'匹配规则的ip或cidr'}
  RemoteIpPerfix *string `json:"remoteIpPerfix,omitempty" xml:"remoteIpPerfix,omitempty" require:"true"`
}

func (s VMPQuerySecurityGroupSecurityGroupRule) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupSecurityGroupRule) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupSecurityGroupRule) SetDirect(v string) *VMPQuerySecurityGroupSecurityGroupRule {
  s.Direct = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroupRule) SetPolicy(v string) *VMPQuerySecurityGroupSecurityGroupRule {
  s.Policy = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroupRule) SetEthertype(v string) *VMPQuerySecurityGroupSecurityGroupRule {
  s.Ethertype = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroupRule) SetPriority(v int32) *VMPQuerySecurityGroupSecurityGroupRule {
  s.Priority = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroupRule) SetPortRangeMax(v int32) *VMPQuerySecurityGroupSecurityGroupRule {
  s.PortRangeMax = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroupRule) SetPortRangeMin(v int32) *VMPQuerySecurityGroupSecurityGroupRule {
  s.PortRangeMin = &v
  return s
}

func (s *VMPQuerySecurityGroupSecurityGroupRule) SetRemoteIpPerfix(v string) *VMPQuerySecurityGroupSecurityGroupRule {
  s.RemoteIpPerfix = &v
  return s
}

type VMPQuerySecurityGroupPaths struct {
}

func (s VMPQuerySecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupPaths) GoString() string {
  return s.String()
}

type VMPQuerySecurityGroupParameters struct {
  // {'en':'Security Group ID', 'zh_CN':'根据安全组id查询'}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {'en':'Security Group name', 'zh_CN':'根据安全组名称查询'}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s VMPQuerySecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupParameters) GoString() string {
  return s.String()
}

func (s *VMPQuerySecurityGroupParameters) SetId(v string) *VMPQuerySecurityGroupParameters {
  s.Id = &v
  return s
}

func (s *VMPQuerySecurityGroupParameters) SetName(v string) *VMPQuerySecurityGroupParameters {
  s.Name = &v
  return s
}

type VMPQuerySecurityGroupRequestHeader struct {
}

func (s VMPQuerySecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type VMPQuerySecurityGroupResponseHeader struct {
}

func (s VMPQuerySecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type AddSecurityGroupRequest struct {
  AddSecurityGroupSecurityGroup []*AddSecurityGroupSecurityGroupParam `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" require:"true" type:"Repeated"`
}

func (s AddSecurityGroupRequest) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRequest) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupRequest) SetSecurityGroup(v []*AddSecurityGroupSecurityGroupParam) *AddSecurityGroupRequest {
  s.AddSecurityGroupSecurityGroup = v
  return s
}

type AddSecurityGroupSecurityGroupParam struct {
  // {'en':'Security group name. 
  //     Constraint: can only consist of letters, numbers, and underscores, with a length of no more than 64 characters. 
  //     It cannot be named default and must be unique.',
  //     'zh_CN':'安全组名称。约束：只能由字母/数字/下划线组成，长度不超过64个字符，不能命名为default且名称唯一。'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Notes (optional), no more than 255 characters in length', 'zh_CN':'备注（可选），长度不超过255个字符'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s AddSecurityGroupSecurityGroupParam) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupSecurityGroupParam) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupSecurityGroupParam) SetName(v string) *AddSecurityGroupSecurityGroupParam {
  s.Name = &v
  return s
}

func (s *AddSecurityGroupSecurityGroupParam) SetRemark(v string) *AddSecurityGroupSecurityGroupParam {
  s.Remark = &v
  return s
}

type AddSecurityGroupResponse struct {
  AddSecurityGroupSecurityGroup []*AddSecurityGroupSecurityGroup `json:"securityGroup,omitempty" xml:"securityGroup,omitempty" require:"true" type:"Repeated"`
}

func (s AddSecurityGroupResponse) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupResponse) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupResponse) SetSecurityGroup(v []*AddSecurityGroupSecurityGroup) *AddSecurityGroupResponse {
  s.AddSecurityGroupSecurityGroup = v
  return s
}

type AddSecurityGroupSecurityGroup struct {
  // {'en':'Security Group ID', 'zh_CN':'安全组id'}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {'en':'Security Group name', 'zh_CN':'安全组名称'}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {'en':'Notes', 'zh_CN':'备注信息'}
  Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
  // {'en':'Create Time', 'zh_CN':'创建时间（yyyy-MM-dd HH:mm:ss）'}
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {'en':'Modify Time', 'zh_CN':'修改时间（yyyy-MM-dd HH:mm:ss）'}
  ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty" require:"true"`
}

func (s AddSecurityGroupSecurityGroup) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupSecurityGroup) GoString() string {
  return s.String()
}

func (s *AddSecurityGroupSecurityGroup) SetId(v string) *AddSecurityGroupSecurityGroup {
  s.Id = &v
  return s
}

func (s *AddSecurityGroupSecurityGroup) SetName(v string) *AddSecurityGroupSecurityGroup {
  s.Name = &v
  return s
}

func (s *AddSecurityGroupSecurityGroup) SetRemark(v string) *AddSecurityGroupSecurityGroup {
  s.Remark = &v
  return s
}

func (s *AddSecurityGroupSecurityGroup) SetCreateTime(v string) *AddSecurityGroupSecurityGroup {
  s.CreateTime = &v
  return s
}

func (s *AddSecurityGroupSecurityGroup) SetModifyTime(v string) *AddSecurityGroupSecurityGroup {
  s.ModifyTime = &v
  return s
}

type AddSecurityGroupPaths struct {
}

func (s AddSecurityGroupPaths) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupPaths) GoString() string {
  return s.String()
}

type AddSecurityGroupParameters struct {
}

func (s AddSecurityGroupParameters) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupParameters) GoString() string {
  return s.String()
}

type AddSecurityGroupRequestHeader struct {
}

func (s AddSecurityGroupRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupRequestHeader) GoString() string {
  return s.String()
}

type AddSecurityGroupResponseHeader struct {
}

func (s AddSecurityGroupResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s AddSecurityGroupResponseHeader) GoString() string {
  return s.String()
}




type VMPQuerySSHKeyRequest struct {
  // {"en":"The number of items displayed on each page is 20 by default", "zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the virtual machine ID specified by the marker", "zh_CN":"从marker指定的实例id开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s VMPQuerySSHKeyRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyRequest) GoString() string {
  return s.String()
}

func (s *VMPQuerySSHKeyRequest) SetLimit(v int) *VMPQuerySSHKeyRequest {
  s.Limit = &v
  return s
}

func (s *VMPQuerySSHKeyRequest) SetMarker(v string) *VMPQuerySSHKeyRequest {
  s.Marker = &v
  return s
}

type VMPQuerySSHKeyResponse struct {
  // {"en":"Key pair", "zh_CN":"密钥对"}
  Keypairs []*string `json:"keypairs,omitempty" xml:"keypairs,omitempty" require:"true" type:"Repeated"`
  // {"en":"Key pair name", "zh_CN":"密钥对名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Public key information corresponding to key", "zh_CN":"密钥对应publicKey信息"}
  PublicKey *string `json:"publicKey,omitempty" xml:"publicKey,omitempty" require:"true"`
}

func (s VMPQuerySSHKeyResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyResponse) GoString() string {
  return s.String()
}

func (s *VMPQuerySSHKeyResponse) SetKeypairs(v []*string) *VMPQuerySSHKeyResponse {
  s.Keypairs = v
  return s
}

func (s *VMPQuerySSHKeyResponse) SetName(v string) *VMPQuerySSHKeyResponse {
  s.Name = &v
  return s
}

func (s *VMPQuerySSHKeyResponse) SetPublicKey(v string) *VMPQuerySSHKeyResponse {
  s.PublicKey = &v
  return s
}

type VMPQuerySSHKeyPaths struct {
}

func (s VMPQuerySSHKeyPaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyPaths) GoString() string {
  return s.String()
}

type VMPQuerySSHKeyParameters struct {
}

func (s VMPQuerySSHKeyParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyParameters) GoString() string {
  return s.String()
}

type VMPQuerySSHKeyRequestHeader struct {
}

func (s VMPQuerySSHKeyRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyRequestHeader) GoString() string {
  return s.String()
}

type VMPQuerySSHKeyResponseHeader struct {
}

func (s VMPQuerySSHKeyResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQuerySSHKeyResponseHeader) GoString() string {
  return s.String()
}




