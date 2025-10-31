# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/toolservice
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/toolservice"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &toolservice.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := toolservice.{ActionName}Response{}
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
| Icpqueryservice | 查询所指定的域名是否在中国大陆工信部进行备案。 | GET | /api/icp |
| Queryforbiddingvisitoripsbydomainservice | 查询域名粒度的封禁IP信息，支持分页查询。 | POST | /api/spider/ip-forbid/query |
| Ipdomainservice | 该接口用于根据IP地址查询正在使用该IP的域名。用户输入IP地址来获取与该IP相关联的域名列表。接口返回的信息包含IP当前使用状态，以及使用该IP的域名列表。在实际应用中，此接口可帮助用户检测特定IP的域名使用情况，适用于网络监控和管理。 | POST | /api/tools/ip/domain-list |
| Queryforbiddingvisitoripsbylabelcodeservice | 查询标签粒度的封禁IP信息，支持分页查询。 | POST | /api/spider/label-ip-forbid/query |