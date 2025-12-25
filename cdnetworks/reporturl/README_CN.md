# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reporturl"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reporturl.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reporturl.{ActionName}Response{}
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
| Querytoprankingurl | 该接口用于统计域名在每小时内URL请求数和流量，并生成最多TOP500的排行榜。用户需提供时间返回和域名，可获取URL的总请求数、总流量、命中请求数、失败请求数及不同状态码详情。接口支持按总请求数或总流量排序，以快速识别高频访问或高流量URL，该接口有助于用户监控URL访问频率及流量，以调整对应的网络策略。 | POST | /api/report/url/top |
| Queryfailedrequestsurlnumbersrankhourlygranularity | 该接口用于查询多个域名在指定时间范围内，以失败请求数排序的URL排名（最多可查询TOP5000条URL）。用户提供开始和结束时间，域名列表。接口返回查询域名下的URL按失败请求数倒序的排名、具体失败请求数、成功请求数及成功总流量。它适用于需要分析识别高失败率的特定URL的场景。 | POST | /api/report/url/fail/rank |
| Querysuccessfulrequestsurlnumbersrankhourlygranularity | 该接口查询多个域名在指定时间段内的成功请求次数URL排行榜，提供前500个URL的成功请求数和流量（MB）。返回数据包括每个URL的排名、成功请求数和流量。适用于分析网站访问情况、识别高访问量URL，帮助优化资源管理和提高访问效率。<br> | POST | /api/report/url/success/rank |
| Reportdomainrefererwebsiteservice | 根据域名查询网站来源排行，支持传入多个域名 | POST | /api/report/domain/referer-website |