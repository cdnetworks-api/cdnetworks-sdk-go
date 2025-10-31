# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportrequest
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportrequest"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportrequest.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportrequest.{ActionName}Response{}
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
| Querydomaintotalrequest | 该接口用于查询多个域名的总请求数。用户通过提供时间，域名等信息，接口将返回每个时间片段内的请求数，支持不同粒度的数据（如每五分钟、每小时或每天的数据）。该接口有助于用户监控网站访问情况，及时对异常情况进行跟进和优化。 | POST | /api/report/domainhit |
| Queryrequestnumbersundershieldpop | 该接口用于查询多个域名在父节点的请求数。用户提供查询时间范围和域名列表，并可选择是否按域名分组输出。返回结果包括总请求数、请求数峰值，以及每5分钟的请求数。此接口有助于用户需要分析域名在父节点上的访问量，以调整缓存策略。 | POST | /api/report/request/parent-node |
| Queryrequestbyspecificprotocol | 查询多域名的指定传输协议的请求数，针对的是所有边缘节点的数据。 | POST | /api/report/request/protocol |
| Queryipv6requestofeachispandprovince | 根据访客访问日志，查询多域名访客IP归属各ISP各省份不同IP类型的请求数<br>支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。 | POST | /api/report/request/isp-province/ipv6 |
| Queryrequesthitratio | 该接口用于查询特定时间段内的分钟级别请求数缓存命中率，用户需提供开始时间、结束时间和域名信息，以获取请求命中率的数据。可选数据粒度1分钟、5分钟、1小时和1天。接口返回的数据信息包括每分钟的缓存命中率，帮助用户评估其域名命中CDN缓存的情况。 | POST | /api/report/request/hit-ratio/total |
| Querymultidomainsipv6oripv4requests | 该接口用于查询指定域名的IPV6/IPV4请求数。用户可以指定时间范围、域名、IP类型和区域进行数据筛选。返回的数据包括每个域名在指定时间段内的每个时间点的请求数。该接口有助于用户分析网络流量的组成和趋势，以便更好地了解不同地域和IP类型的请求分布情况。 | POST | /api/report/request/ipv6 |
| Reportuserrequestcountryservice | 查询多域名访客IP归属各国家区域请求数 | POST | /api/report/request/country |
| Queryedgerequesthitratioservice | 查询多域名分钟级别的在边缘节点命中缓存的请求数命中率 | POST | /api/report/request/edge-hit-ratio/total |