# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngcontentmanagement
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngcontentmanagement"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &ngcontentmanagement.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ngcontentmanagement.{ActionName}Response{}
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
| CreateAPurgeRequest | 创建一个刷新任务来强制刷新存储在CDN Pro缓存服务器中的内容。当您更新了源站的内容，并希望访客立即看到更新后的内容，您可以创建一个刷新任务立即刷新内容。 | POST | /cdn/purges |
| GetListOfPurgeRequests | 获取刷新请求列表。接口仅支持查询近15天内的刷新请求。 | GET | /cdn/purges |
| GetPurgeSummary | 查询某个时间范围内的刷新请求的汇总信息。可通过查询参数指定时间范围和目标环境。 | GET | /cdn/purges/purgeSummary |
| GetPurgeRequestStatus | 获取刷新请求的详细信息，包括涉及的加速域名和刷新任务的执行状态等。 | GET | /cdn/purges/* |
| CreateAPrefetchRequest | 创建一个预取请求，从您的源站预取内容来预热CDN Pro的缓存。通过内容提前预取，可避免大量请求涌入源站服务器。发起预取之前，必须先将域名对应的加速项目部署到生产环境。 | POST | /cdn/prefetches |
| GetListOfPrefetchRequests | 获取预取请求列表。接口仅支持查询近15天内的预取请求。 | GET | /cdn/prefetches |
| GetPrefetchRequestStatus | 获取预取请求的详细信息。 | GET | /cdn/prefetches/* |
| GetAvailablePurgeRequests | 该接口返回您可以在演练或生产环境进行目录和文件刷新的额度。使用限额每日有一个固定值。您可以临时超过该限额，但这会减少第二天可用的刷新额度。 | GET | /cdn/purges/purgeTokens |