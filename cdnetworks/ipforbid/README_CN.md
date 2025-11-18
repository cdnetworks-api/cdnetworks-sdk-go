# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ipforbid
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ipforbid"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &ipforbid.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ipforbid.{ActionName}Response{}
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
| Forbidorresumevisitoripsbydomainservice | 对访问指定域名的IP，进行封禁解禁操作。 | POST | /api/spider/ip-forbid |
| Queryforbiddingvisitoripsbydomainservice | 查询域名粒度的封禁IP信息，支持分页查询。 | POST | /api/spider/ip-forbid/query |
| Queryforbiddingvisitoripsbylabelcodeservice | 查询标签粒度的封禁IP信息，支持分页查询。 | POST | /api/spider/label-ip-forbid/query |
| Forbidorresumevisitoripsbylabelcodeservice | 一个客户可以创建一个标签，该标签可以关联若干域名，通过该标签进行封禁解禁访问IP的操作，效果等同于对该标签关联的所有域名，进行封禁解禁指定的访客IP。 | POST | /api/spider/label-ip-forbid/operate |