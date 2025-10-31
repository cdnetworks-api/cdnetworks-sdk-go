# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/resourcemanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/resourcemanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &resourcemanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := resourcemanage.{ActionName}Response{}
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
| Vmpqueryflavor | 您可以通过此接口查询得到某个区域能提供的实例的规格列表，然后通过这个规格来创建实例。 | GET | /vmp/flavors |
| Vmpquerynode | 查询边缘计算平台有哪些节点可以提供服务。 | GET | /vmp/nodes |
| Vmpquerybandwidth | 查询您所有在用的节点的实时冗余带宽，有冗余带宽的节点就返回True，无冗余带宽的节点就返回False，如果该节点的带宽数据采集结果不足以支撑该接口的时效性要求，即带宽统计数据较旧暂不满足时效性要求，则返回Undefined。节点是否有冗余带宽的判断方法：节点上限带宽大于节点的实时使用带宽。数据粒度：5分钟。<br>备注：<br>1）只返回客户有使用的节点的带宽冗余情况，有使用的前提是该客户在这个节点至少有1台云主机；<br>2）如果某个节点返回的result是Undefined，表示该节点无10分钟内的实时带宽统计数据，为了保证冗余带宽的时效性，返回Undefined，建议稍后再次调用该接口以查询该节点的最新情况；<br>3）如果节点是双线或多线节点，则分别返回每个运营商的带宽冗余情况；<br>4）节点的实时冗余带宽是动态变化的，若通过该接口的返回结果做业务调度，建议周期获取数据。 | POST | /vmp/redundant-bandwidth |
| Getservergrouplist | 获取servergroup列表 | GET | /api/clouddns/servergroups |
| Getservergroupbyid | 根据ID获取指定servergroup | GET | /api/clouddns/servergroups/* |
| Createservergroup | 创建新的servergroup | POST | /api/clouddns/servergroups |
| Editservergroupbyid | 修改已有的servergroup | PUT | /api/clouddns/servergroups/* |
| Deleteservergoupbyid | 删除指定的servergroup | DELETE | /api/clouddns/servergroups/* |
| Getserverlistbyservergroup | 获取指定servergroup下的server列表 | GET | /api/clouddns/servergroups/*/servers |
| Createserver | 为指定的servergroup创建server | POST | /api/clouddns/servergroups/*/servers |
| Updateserver | 根据servergroupId和serverId更新指定的server | PUT | /api/clouddns/servergroups/*/servers/* |
| Deleteserver | 删除指定的server | DELETE | /api/clouddns/servergroups/*/servers/* |
| Gethealthcheckerlist | 获取健康检查方法的列表 | GET | /api/clouddns/infra/healthcheckers |
| Getcvlist | 获取CV(自定义区域)列表 | GET | /api/clouddns/cvs |
| Getregionlist | 获取区域列表 | GET | /api/clouddns/regions |
| Getisplist | 获取运营商列表 | GET | /api/clouddns/isps |
| Gethealthcheckagentlist | 获取健康检查探测器列表 | GET | /api/clouddns/infra/healthcheckers/iplist |
| Enablehealthcheckwhitelist | 健康检查探测器是否可用白名单配置。当配置为auto模式时，需要在您的被探测机器上实现功能，将过了生效时间的探测器IP添加到白名单上。 | POST | /api/clouddns/infra/healthcheckers/whitelist |
| Gethealthcheckwhitelist | 获取健康检查白名单列表 | GET | /api/clouddns/infra/healthcheckers/whitelist |