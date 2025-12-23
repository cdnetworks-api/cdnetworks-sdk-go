# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportbandwidth
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportbandwidth"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportbandwidth.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportbandwidth.{ActionName}Response{}
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
| Queryp2pbandwidth | 查询CDN带宽、P2P带宽、总带宽和计费值 | POST | /myview/BandWidthP2p |
| Bandwidthmiddle | 查询中间缓存带宽，输出分钟粒度的中间缓存带宽、动态hit、动态misss、静态hit、静态miss带宽 | POST | /myview/bandwidth-middle |
| Querydomainbandwidth | 该接口用于统计多个域名在指定时间内的带宽汇总信息。用户需提供加速域名及时间参数来生成汇总数据，返回内容包括每个时间片的带宽汇总数据。此接口有助于用户监控和分析域名的流量使用情况，以便更好地进行网络优化和成本控制。<br> | POST | /api/report/domainbandwidth |
| Queryipv6bandwidthofeachispandprovince | 根据访客访问日志，查询域名在各ISP各省份不同IP类型的带宽<br>支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。 | POST | /api/report/bandwidth/isp-province/ipv6 |
| Getbandwidthlog | 查询客户的日志带宽（域名、分钟粒度的日志带宽） | POST | /myview/bandwidth-log |
| Bandwidthpeakranking | 对频道进行带宽峰值排行，输出按照峰值排行后的频道、峰值时间、峰值、总流量 | POST | /myview/bandwidth-peak-ranking |
| Bandwidthchannel | 查询域名带宽（域名、分钟粒度的带宽） | POST | /myview/bandwidth-channel |
| Querybandwidthbyispprovince | 该接口用于查询中国大陆访客IP在各省和ISP的带宽信息。用户可指定时间段、域名、地区和ISP等参数，获取5分钟粒度的带宽数据，以Mbps为单位。接口返回包括时间戳和相应带宽值，帮助用户分析不同地区和ISP的带宽使用情况，从而优化网络资源配置和提升用户体验。 | POST | /api/report/bandwidth/area/isp-province |
| Querybandwidthoforiginminutely | 该接口用于查询多域名的分钟级别回源带宽。用户需提供开始和结束时间，指定域名和数据粒度。返回结果包括每个域名的带宽峰值、回源总流量及每分钟的回源带宽。 | POST | /api/report/bandwidth/multi-domain/real-time/origin |
| Querycpsbandwidth | 专线带宽接口 | POST | /myview/bandwidth-HK |
| Queryrealtimebandwidthformultidomain | 该接口用于查询域名分钟级别的边缘带宽，用户需提供时间范围和域名来获取详细的带宽使用数据。返回内容包括每个域名的总流量、分钟级别带宽数据，以及带宽峰值等。有助于用户分析网站或应用的流量情况，从而优化资源管理和提升性能。<br> | POST | /api/report/bandwidth/multi-domain/real-time/edge |
| Querymultidomainsipv6oripv4bandwidth | 该接口用于查询多个域名的IPV6或IPV4带宽数据。用户提供时间范围、域名、IP类型及地区等参数，以获取相关域名的带宽详情。返回结果包括每个域名在指定时间段内的带宽数据。该接口有助于用户监测和分析不同域名在特定IP类型下的网络流量情况，从而优化资源分配和提高网络性能。 | POST | /api/report/bandwidth/ipv6 |
| Perzonebilling | 查询perzone计费数据 | POST | /myview/perzone-billing |
| Querybandwidthformultidomain | 该接口用于查询动态加速产品的域名带宽情况。用户需提供时间参数来获取指定时间段内的带宽数据。返回内容包括带宽每个时间片的详细数据。此接口有助于用户全面了解多个域名的带宽消耗，从而优化网络资源和规划带宽。 | POST | /api/report/domainbandwidth/dwa |
| Reportcountryserverbandwidthservice | 该接口用于查询服务器IP所属国家的带宽明细。用户需提供时间范围、域名、国家地区代码以及数据粒度。返回内容包含指定时间段内的每个时间点的边缘流量和带宽值。此接口有助于用户了解当前在全球不同国家的服务流量分布和使用情况。 | POST | /api/report/server/country-bandwidth |
| Getcdnrelaytraffic | 该接口用于查询指定维度的cdn中间流量数据，用户可以通过该接口来查询对应客户的详细频道中间流量报表，包括输出日期、峰值时间、带宽峰值、总流量等。这对客户了解自身的中间流量使用情况有很直接的帮助。 | POST | /cdn/traffic/relay |