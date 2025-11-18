# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/other
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/other"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &other.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := other.{ActionName}Response{}
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
| Querydailylivetranscodingduration | 直播转码时长每日统计。 | POST | /myview/LiveTranscodingPerDayV2 |
| Query5minlivetranscodingduration | 输出客户分钟粒度的直播转码时长。 | POST | /myview/LiveTranscodingV2 |
| Queryliverecordingduration | 输出客户分钟粒度的直播录制时长 | POST | /myview/RecordingTime |
| Picprocessstatistics | 查询域名5分钟粒度图片处理次数。 | POST | /myview/PicFlowHit |
| Httpdnsstatistics | 查询5分钟粒度httpdns解析量。 | POST | /myview/HttpdnsStatistics |
| Querylivestreamstatus | 当前接口非实时数据内容： 流状态信息，码率信息，帧率信息，在线人数，头部信息，频道带宽；实时数据内容有： 流状态信息，码率信息，帧率信息，在线人数(不包含HLS)，头部信息。 | POST | /api/quality/stream-status-statistic |
| Getlivestreampushingstatus | 获取指定时刻，指定域名下（支持多个域名和流名）的所有频道的帧率、码率、丢帧率等信息 | POST | /api/quality/frame-rate |
| Queryonlineviewercount | 提供直播域名或者频道的在线人数接口，接口只支持GET请求方式 | POST | /QOSS/api/onlineViewers |
| Queryddosmitigatedbandwidthbysuiteorproduct | 查看套餐或产品下所有域名及转发规则所清洗的攻击带宽 | POST | /soc/api/report/QueryMitigatedBandwidth |
| Queryedgeactiveconnectionsformultidomains | 此接口用于查询多域名的边缘活跃连接数，通过提供开始时间、结束时间和域名列表，可查询该时间段内多个域名在每个整5分钟的瞬间活跃连接数。适用于需要分析网站活跃连接趋势的场景。 | POST | /api/report/edge-active-conn |
| Reportdomainstreamdurationservice | 查询多域名多流名的推流时长 | POST | /api/report/domain/stream/duration |
| Querydirectoryrankbysuccessrequesthourly | 该接口用于查询指定时间下根据成功请求数最多Top 500的目录，状态码为2xx，3xx的定义为成功。用户需提供时间范围和域名，获取目录列表。接口返回目录对应的成功请求数和成功总流量数据。该接口有助于用户分析网站性能、流量趋势和优化目录结构。 | POST | /api/report/directory/success/rank |
| Querydirectoryrankbyfailedrequesthourly | 该接口用于查询域名下每小时按失败请求数进行排名的目录信息。用户可以根据指定的时间范围获取具有最高失败请求数的目录列表。返回信息包括目录所对应的失败请求数，成功请求数和成功传输的总流量。用户通过该接口有助于分析网站目录中可能存在的问题，并优化网络资源的使用。最多可支持查询top500，失败请求定义：除状态码为3xx，2xx外的请求 | POST | /api/report/directory/fail/rank |
| Querydirectoryrankbytotaltraffichourly | 这个接口用于查询多个域名根据总流量排序的目录列表，最多返回top 500。用户需提供时间范围和域名来获取目录列表，返回内容包括每个目录的总流量和请求数。此接口有助于用户分析及优化网站流量，帮助用户识别高流量目录并进行资源分配。<br> | POST | /api/report/directory/hourly/total-traffic/rank |
| Antihijackiplist | 获取劫持缓解IP列表，客户通过该接口可以获取劫持缓解IP相关的信息，包含区域，部署状态等信息。 | POST | /api/v1/dms/antiHijackIP/list |
| Resourcegrouplist | 该接口用于查询账号粒度下所有的资源组列表。用户可以通过资源组名称来查询相关的资源组。包括资源组名称、域名数量.帮助客户合理分配资源。 | POST | /api/v1/dms/resourceGroup/list |
| Querydomainsforresourcegroup | 该接口用于查询指定资源组关联的域名详情。用户可以通过资源组名称来获取关联域名的完整信息。包括具体域名及域名状态。用户可以按需管理域名和资源组关系，避免核心业务域名受攻击时相互影响。 | POST | /api/v1/dms/resourceGroup/queryDomains |
| Resourcegroupdomainchange | 该接口用于变更域名使用的资源组。用户将指定的域名从当前归属资源项更新到目标资源组。帮助客户分配域名使用的资源。 | POST | /api/v1/dms/resourceGroup/changeGroup |
| Streamtrafficservice | 该接口用于查询流名带宽，用户可以通过指定时间，域名，区域，发布点，协议等查询数据；根据区域，域名，流名，发布点，协议等维度聚合流量维度返回数据 | POST | /api/gdp/stream-traffic-service |
| Querydomainresourcegroup | 该接口用于查询域名所属资源组。用户可以通过域名查询相关的资源组信息。包括资源组名称、资源组ID。 | POST | /api/v1/dms/domainResource |