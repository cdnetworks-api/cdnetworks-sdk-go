# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngreports
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngreports"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &ngreports.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ngreports.{ActionName}Response{}
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
| GetEdgeTrafficVolume | 获取一段时间内的边缘流量数据（以MB为单位）。可以使用查询参数指定时间段和协议。通过在请求体中传递filters对象可进一步指定查询范围。该报表的数据延迟小于两分钟。<br> | POST | /cdn/report/vol |
| GetEdgeRequests | 获取一段时间内的边缘请求数数据。可以使用查询参数指定时间段和协议。通过在请求体中传递filters对象可进一步指定查询范围。该报表的数据延迟小于两分钟。 | POST | /cdn/report/req |
| GetEdgeBandwidth | 获取一段时间内的边缘带宽数据（以Mbps为单位）。可以使用查询参数指定时间段和协议。通过在请求中传递filters对象可进一步指定查询范围。此报表的数据延迟小于两分钟。 | POST | /cdn/report/bandwidth |
| GetOriginFastRouteTrafficVolume | 查询通过快速回源功能传输的回源流量。 | POST | /cdn/report/fastOriginVol |
| GetASummaryOfRequests | 获取一段时间内的请求数汇总数据。可以按加速域名或serverGroups（节点组）进行查询和分组。 | POST | /cdn/report/reqSummary |
| GetASummaryOfTrafficBandwidth | 获取一段时间内的带宽汇总数据。可以按加速域名或serverGroups（节点组）进行查询和分组。 | POST | /cdn/report/bandwidthSummary |
| GetOriginTrafficVolume | 获取一段时间内的回源流量数据（以MB为单位）。可以使用查询参数指定时间段和协议。通过在请求体中传递filters对象可进一步指定查询范围。此报表的数据延迟小于两分钟。<br><br>注意：该接口所返回的统计数据不包括TCP、IP和MAC报头的开销。 | POST | /cdn/report/volOrigin |
| GetEdgeStatusCodeDetails | 获取一段时间内的边缘状态码统计信息。可以使用查询参数指定时间段、协议和数据粒度。通过在请求体中传递filters对象可进一步指定查询范围。此报表的数据延迟最长可达30分钟。 | POST | /cdn/report/statusCodeDetails |
| GetTheNumberOfRequestsToOrigin | 获取一段时间内的回源请求数数据。可以使用查询参数指定时间段和协议。通过在请求体中传递filters对象可进一步指定查询范围。此报表的数据延迟小于两分钟。 | POST | /cdn/report/reqOrigin |
| GetOriginStatusCodeDetails | 获取一段时间内的回源状态码统计信息。可以使用查询参数指定时间段、协议和数据粒度。通过在请求体中传递filters对象可进一步指定查询范围。此报表的数据延迟最长可达30分钟。 | POST | /cdn/report/statusCodeDetailsOrigin |
| GetTheIntermediateTrafficVolume | 获取一段时间内的中间层流量数据（以MB为单位）。中间层流量是指CDN Pro缓存服务器之间的流量。可以使用查询参数指定时间段和协议。通过在请求体中传递filters对象可进一步指定查询范围。此报表的数据延迟小于两分钟。 | POST | /cdn/report/volInterm |
| GetTheEdgeUploadTrafficVolume | 获取一段时间内上传到CDN Pro边缘服务器的流量数据（以MB为单位）。可以使用查询参数指定时间段和协议。通过在请求中传递filters对象可进一步指定查询范围。此报表的数据延迟小于两分钟。<br><br>注意：该接口返回的数据中不包括TCP、IP和MAC报头的开销。 | POST | /cdn/report/upVol |
| GetASummaryOfStatusCodesReturnedByEdgeServers | 获取一段时间内CDN Pro边缘服务器返回的HTTP状态码的汇总数据。可以使用查询参数指定时间段和协议。在请求体中传递filters对象可进一步指定查询范围。此报告的数据延迟最长可达30分钟。 | POST | /cdn/report/edgeStatusSummary |
| GetASummaryOfStatusCodesReturnedByOriginServers | 获取一段时间内源站返回的HTTP状态码的汇总数据。可以使用查询参数指定时间段和协议。在请求体中传递filters对象可进一步指定查询范围。此报告的数据延迟最长可达30分钟。 | POST | /cdn/report/originStatusSummary |
| GetASummaryOfCpuUsage | 获取CPU使用时间汇总信息，包括边缘服务器和中间层服务器的数据。 | POST | /cdn/report/cpuSummary |
| GetOriginFastRouteRequests | 查询通过快速回源功能所产生的回源请求数。 | POST | /cdn/report/fastOriginReq |
| GetEdgeHostnameSummaryStatistics | 获取一段时间内一个或多个调度域名的DNS解析请求数汇总数据。可通过该接口查询每个调度域名的请求数数据。 | POST | /cdn/report/edgeHostnameReqSummary |
| GetCpuTimeUsed | 查询处理用户请求所消耗的CPU时间（以秒为单位）。 | POST | /cdn/report/cpuTime |
| GetASummaryOfTrafficVolume | 获取一段时间内的边缘、中间层（CDN Pro服务器之间）和回源流量的汇总数据。可以在请求体中指定加速域名或serverGroups（节点组）等参数进行查询和分组。<br> | POST | /cdn/report/volSummary |
| Getedgehostnamestatistics | 获取一段时间内对调度域名发起的DNS解析请求数。 | POST | /cdn/report/edgeHostnameReq |