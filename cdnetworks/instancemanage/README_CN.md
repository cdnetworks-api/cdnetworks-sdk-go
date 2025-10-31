# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/instancemanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/instancemanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &instancemanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := instancemanage.{ActionName}Response{}
    _, err := auth.Invoke(config, request, response)

    // Handle response
    if err != nil {
        log.Printf("error: %s\n", err)
        return
    }

    log.Printf("response body: %s\n", response.String())
}
```

## 错误处理

始终检查 API 调用返回的错误：

```go
_, err := auth.Invoke(config, request, response)
if err != nil {
    log.Printf("error: %s\n", err)
    // Handle the error appropriately
    return
}
```

## API列表
有关详细的 API 文档和可用方法，请参阅[官方 Cdnetworks API 文档](https://docs.cdnetworks.com/en/cdn/apidocs)。

| ActionName | description | client_methods | uri |
| --- | --- | --- | --- |
| Vmpremoveinstance | 销毁指定的实例，支持批量销毁。销毁后实例所使用的物理资源都将被回收，包括磁盘及快照，相关数据全部丢失且永久不可恢复。<br>执行后可再次调用查询实例接口，确认实例是否成功删除。 | POST | /vmp/servers/delete |
| Vmpqueryinstance | 查询某个客户拥有的实例。 | GET | /vmp/servers |
| Vmpinstanceoperation | 假如您发现某个实例不能正常使用（如可以ping但不能登录），可调用该接口尝试强制重启该机器。强制重启等同于传统服务器的断电重启，可能丢失实例操作系统中未写入磁盘的数据。正常关机就是普通shutdown操作。该接口调用成功后需再调用实例查询接口，确认实例的最新状态，以验证是否重启成功。 | POST | /vmp/servers/*/action |
| Vmpcreateinstance | 通过该接口您可以在某个区域申请创建指定规格的云主机实例，实例创建完毕之后，可通过使用实例查询接口获得实例的最新状态。<br>1）如果虚拟机需要使用内网网络并且未指定节点名称，则一次请求的虚拟机都将调度到同一个Cluster（只支持Cluster内的内网互通），非同一次请求的虚拟机无法保证虚拟机间内网互通，因为可能调度到不同的Cluster。不同节点、区域、省份或运营商的虚拟机间不支持内网互通，因为他们必定被调度到不同的Cluster。<br>2）如果创建请求携带了cidr的参数，则要求节点上或者没有虚拟机（可以是创建过但是都销毁了）或者已有虚拟机的cidr与当前请求一致，这样的节点才能创建虚拟机，如果找不到这样的节点，则虚拟机创建失败。如果不指定privateIPv4，cidr的第一个ip地址默认不会分配给虚拟机。例如，cidr=192.168.10.129/25，则虚拟机的ip分配范围从192.168.10.130开始，除非指定参数privateIPv4=192.168.10.129，才能创建ip为129的虚拟机。<br>3）如果请求参数中携带的是裸机模板，则表示创建裸机实例，此时不支持内网、不支持ipv6。 | POST | /vmp/servers |
| Editinstance | 用于修改实例信息，当前仅支持实例名称修改。 | PUT | /vmp/servers |
| Instanceipv6management | 为存量云主机实例分配/移除IPv6地址，裸机实例不支持。<br>分配ipv6说明：<br>1）action= ALLOCATION；<br>2）主机状态必须为运行或停机，主机所在节点必须支持IPv6；<br>3）如果主机当前已存在ipv6则接口会直接返回已有ipv6地址。<br>移除ipv6说明：<br>1）action= REMOVE；<br>2）主机状态必须为运行或停机。 | POST | /vmp/servers/ipv6 |
| Instancediskscaling | 支持实例在线添加磁盘。 | POST | /vmp/servers/attachDisk |
| Manageinstancetags | 管理云主机实例标签，支持修改或删除实例标签。 | PUT | /vmp/server-tags |
| Convertfreetypeinstancetochargetype | 对指定的免费实例进行转正，避免免费实例过期被删除，支持批量转正。转正后实例将转为计费实例。 | POST | /vmp/servers/charge |
| InstanceRebuild | 通过该接口可以对已创建的实例进行重装系统 | POST | /vmp/servers/rebuild |