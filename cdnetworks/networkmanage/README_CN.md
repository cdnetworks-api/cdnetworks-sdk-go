# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/networkmanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/networkmanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &networkmanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := networkmanage.{ActionName}Response{}
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
| Vmpreleaseedgeip | 用于释放独占式和漂移式额外公网IP。<br>说明：<br>1）漂移式额外IP需要先解除和云主机的绑定才能释放，如果未解除绑定会释放失败；<br>2）如果批量释放多个漂移式IP，部分IP已解除绑定，部分IP未解除，则已解除绑定的会正常释放，未解除绑定的会释放失败。<br>3）独占式IP无需解除绑定直接释放。<br> | PUT | /vmp/edgeIp/release |
| Vmpallocateedgeip | 用于申请漂移式额外公网IP。漂移模式支持同一IP为多实例同时使用，常用于主备切换场景，如LVS等。 | POST | /vmp/edgeIp/allocate |
| Vmpassignedgeip | 将申请到的漂移式额外IP绑定到指定实例。<br>漂移式额外IP支持同一IP为多实例同时使用，常用于主备切换场景，如LVS。<br> | PUT | /vmp/edgeIp/assign |
| Vmpunassignedgeip | 可用于解除漂移式额外公网IP和实例的绑定。漂移式额外公网IP支持同一IP为多实例同时使用，常用于主备切换场景，如LVS。<br>说明：<br>1）要解绑的漂移式额外公网IP必须是绑定在指定实例上的IP，不能是绑定在其他虚拟机的IP；<br>2）如果批量解绑多个漂移式额外公网IP，存在部分IP是绑定在其他实例上的，则全部IP都解绑失败，接口返回错误提示信息。 | PUT | /vmp/edgeIp/unassign |
| Vmpqueryedgeip | 用于查询已申请的额外公网IP。 | GET | /vmp/edgeIp |