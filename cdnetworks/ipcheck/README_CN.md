# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ipcheck
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ipcheck"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &ipcheck.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ipcheck.{ActionName}Response{}
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
| Querycdnservicerealip | 该接口用于获取CDN服务的真实IP列表，特别适用于源站设置了白名单限制的场景。用户可以通过调用此接口获得我司提供的CDN节点用于回源的IP白名单，以便正确配置并确保数据请求能够通过CDN顺利到达源站。 | GET | /api/si/report/whiteip-list |
| Queryspecificipbelong | 该接口用于查询给定IP地址是否属于我司CDN IP。用户需提供IP地址列表。返回结果包含每个IP地址是否归属于我司CDN IP。 | POST | /api/si/tools/ipCheck |
| Ipinfoservice | 该接口用于查询特定IP地址的归属信息。用户可以通过提供一个或多个IP地址来查询它们是否为公司CDN节点，以及其归属的国家、省份、城市和运营商信息。返回结果包括是否为公司CDN节点的标识，如不是公司CDN的节点，该接口将返回未知。此接口适用于当用户查询IP是否我司IP。 | POST | /api/tools/ip-info |