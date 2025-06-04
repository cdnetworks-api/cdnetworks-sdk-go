# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportstatuscode
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportstatuscode"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportstatuscode.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportstatuscode.{ActionName}Response{}
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
| Querystatuscodedistributionofeachispandprovince | 该接口查询多个域名在各运营商和省份的状态码分布。用户提供时间范围、域名列表，并可选运营商和省份。结果按域名、运营商、省份展示状态码分布，支持1分钟/5分钟/1小时粒度查询。支持`Accept-Language`请求头，仅支持`zh-CN`和`en-US`，默认`zh-CN`。`en-US`时，省份及运营商使用代码，否则返回中文。 | POST | /api/report/status-code/isp-province |
| Querystatuscodedistribution | 该接口用于统计多个域名的边缘状态码数据。用户可选择查询时间范围和域名列表获取数据。返回内容包括每个域名的状态码分布及其对应的请求数。有助于用户分析域名访问状态和优化内容分发策略的用户。 | POST | /api/report/status-code |
| Queryoriginstatuscodedistribution | 多域名的回源状态码统计，针对的是所有节点的回源数据。 | POST | /api/report/status-code/origin |
| Queryipv6statusofeachispandprovince | 该接口用于查询多域名在各省各ISP的IPV6状态码。用户可通过传递域名及选择具体的省份和运营商、IP协议类型来获取详细的状态码请求数信息。查询时间粒度可选5分钟/1小时。返回数据包括各域名在不同省份和ISP的详细状态码及请求数量。此接口适用于需要监控和分析网站在不同区域的指定IP协议的访问情况的场景。<br>支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商入参及返回都为code，否则返回的为中文。 | POST | /api/report/statusCode/isp-province/ipv6 |
| Queryrealtimeoriginstatuscode | 该接口用于查询指定时间范围内每分钟的回源请求状态码。用户需提供时间区间和域名列表，并可选择查询维度，如具体状态码或状态码类型（成功、重定向、服务器错误等）。返回数据包括请求时间和请求总数的详细信息。此接口用于监控网站健康状态，优化源站响应，帮助分析和调整网络架构，以提升服务稳定性和访客体验。 | POST | /api/report/status-code/real-time/origin/total |
| Queryispprovincestatuscode | 该接口查询特定省份和运营商的分钟级状态码统计。用户可设置起止时间、域名、状态码、省份和运营商等参数。返回每分钟状态码请求数及总请求数。支持`Accept-Language`请求头，仅支持`zh-CN`和`en-US`，默认`zh-CN`。若`Accept-Language`为`en-US`，省份及运营商参数和返回值为代码，否则为中文。 | POST | /api/report/status-code/isp-province/total |
| Reportstatuscoderealtimeedgeservice | 该接口用于查询边缘状态码的分钟粒度数据，用户可指定开始时间和结束时间以及域名查询状态码对应请求数数据。请求中可选数据粒度1分钟或5分钟。返回结果包括各状态码在每分钟请求数。 | POST | /api/report/status-code/real-time/edge |
| Querymultidomainsipv6oripv4statuscode | 查询多域名IPV6/IPV4状态码请求数 | POST | /api/report/status-code/ipv6 |
| Reportstatuscodelogservice | 查询各频道每小时粒度状态码分布 | POST | /api/report/status-code/log |
| Querystatuscodedistributionincountries | 按国家粒度查询状态码分布情况(按服务IP归属) | POST | /api/report/status-code/country |