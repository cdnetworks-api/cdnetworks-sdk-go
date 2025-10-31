# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportflow
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportflow"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportflow.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportflow.{ActionName}Response{}
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
| Querydirectorybandwidthtrafficunderlivestreamdomain | 该接口用于查询直播推拉流域名下各个目录的流量或带宽统计数据。用户需提供起止时间、域名，可指定目录查询。默认情况下，数据以5分钟为粒度进行统计。返回结果包括每个目录的总流量或带宽峰值，以及时间分片的详细数据。 | POST | /api/report/flow/dir/detail |
| Querydomaintotaltraffic | 此接口用于获取多个域名的流量汇总信息。用户需提供时间范围和域名，数据粒度可选五分钟、一小时或一天。返回内容括总流量和各时间节点的流量值。此接口有助于用户监控和分析多个域名流量情况，从而掌握整体流量趋势及变化。 | POST | /api/report/domainflow |
| Querystreamtrafficundermalbdomain | 该接口用于查询指定时间段内域名下的直播流流量或带宽数据。用户需提供时间范围、域名和流名，以及数据类型和粒度选项。接口返回每个域名下各流的流量或带宽详情，包括总流量和峰值带宽。有助于用户监控直播流量情况以优化资源配置。不用于计费核对。 | POST | /api/report/flow/stream/detail |
| Querytrafficrequestintotalandpeakvalue | 该接口用于查询多个域名的流量请求数和峰值带宽。用户需提供时间范围和域名列表，可获取各域名的总流量、总请求数、峰值带宽及其发生时间，以及每个时间段的详细流量与请求数。有助于用户了解网站的访问情况，从而有效调整策略以满足不同流量需求。 | POST | /api/report/flow-request |
| Querystaticdynamictrafficbyprotocol | 该接口用于查询各域名下不同协议的动静态流量明细。用户需提供时间范围和域名列表，最多可查询过去半年的数据。返回结果包括每个域名的流量数据，按指定协议和节点分级展示。此接口适用于需要分析域名流量模式、优化网络资源配置的用户，支持细致的流量管理和性能评估。 | POST | /api/report/flow/static-dynamic |
| Queryoutputtrafficundershieldpop | 该接口的主要功能是查询多个域名的父节点下行流量数据。用户可以通过提供开始时间、结束时间以及域名列表来获得详细的流量数据报告。返回值包括每个域名的总流量，并提供每个时间片的流量值。此接口适用于需要监测和分析CDN父节点的流量概况及变化趋势。接口调用频率勿高于30/5min。 | POST | /api/report/flow/parent-node |
| Querytrafficbyspecificprotocol | 该接口用于查询多个域名在指定传输协议下的流量数据，查询的是所有边缘节点的数据。用户需指定域名、起止时间、传输协议等参数，并可选择数据粒度和分组维度。默认情况下，查询会返回HTTPS协议的流量数据，输出结果为每个域名在不同时间点的流量值。此接口适用于需要查询和分析多域名不同传输协议流量的场景。 | POST | /api/report/flow/protocol |
| Queryspecificpoptraffic | 该接口用于查询短时间内指定多个域名的各个POP的流量明细。用户可通过指定开始时间和结束时间，获取各个域名下不同POP节点的流量数据，包括静态和动态的流入和流出流量。接口返回的数据以5分钟粒度展示。 | POST | /api/report/flow/pop |
| Querybacktoorigintrafficandrequest | 该接口旨在查询多个域名的回源流量和请求数，用户可以通过输入开始时间、结束时间、指定的域名等信息来获取这些数据。接口会返回每个域名的总流量、总请求数、请求和带宽峰值及其对应的时间点，并支持以分钟粒度提供详细流量、请求和带宽数据。这一功能有助于用户监控和分析网站回源流量。<br>限制说明：一般数据延迟5~15分钟左右。建议调用频率不高于300次/5分钟。 | POST | /api/report/flow-request/origin |
| Querysumuptrafficunderaccount | 该接口用于查询用户账号下所有域名的流量汇总信息。用户提供开始结束时间参数来获取详细的流量数据，查询数据粒度可选5分钟/1小时/1天。通过该接口，用户可了解和分析各个域名在指定时间段内的CDN流量使用情况。 | GET | /api/report/flow |
| Queryispprovincehitrate | 查询省份ISP的分钟级别的流量命中率和请求数命中率。<br>支持语言请求头Accept-Language，只支持zh-CN、en-US，默认为zh-CN。Accept-Language：en-US时，省份及运营商 入参及返回都为code，否则返回的为中文。 | POST | /api/report/flow/isp-province/hit-rate/total |
| Queryrequesbandwidthtsavingratio | 该接口用于查询指定时间段内的带宽节省量，用户需提供起始和结束时间，域名。支持的数据粒度包括五分钟、一小时和一天，如未指定，默认按照天粒度查询。返回的数据包括总节省带宽平均值，按时间片对应的实际带宽节省量。该接口特别适用于需要监控分析CDN缓存命中带来的带宽节省量的场景。 | POST | /api/report/request/saving-bandwidth/total |
| Querymultiplestreamtrafficandbandwidthunderthedomain | 该接口用于查询指定时间段内多个域名下直播流的流量和带宽信息。用户输入开始时间、结束时间及域名等，即可获取详细数据，包含每个域名下各流的总流量、峰值带宽、各时间点的流量或带宽数据，以及（可选的）时间点在线人数。这适用于监控直播流网络性能和用户交互，通过数据分析优化网络配置和资源分配，提高直播服务质量。 | POST | /api/report/flow-bandwidth/stream/detail |
| Reportflowdomaincountryservice | 该接口查询指定域名在各国家和地区的访客流量带宽分布。用户输入时间范围和域名列表，可按域名、国家或国内外查询数据。返回结果包括各地区的流量总和及其百分比，以及每个时间片段的流量和带宽，帮助分析和管理全球范围内的网站流量带宽。建议查询24小时前的数据。 | POST | /api/report/flow/domain-country |
| Querydatatransferforalldomainbyte | 查询账号下所有域名的流量汇总，单位字节。 | POST | /api/report/flow/byte |
| QuerytotaltrafficinbyteforalldomainsunderdwaByte | 查询账号下所有域名的流量汇总，单位字节。 | POST | /api/report/flow/byte/dwa |
| Querytotaltrafficforalldomainsunderdwa | 该接口用于查询动态加速下所有域名的流量数据。用户通过调用此接口可以获取某一时间段内所有域名的流量数据，包括每个时间片的具体流量，数据粒度可指定5分钟、1小时或1天。 | GET | /api/report/flow/dwa |
| Querytrafficforselectedbillingarea | 该接口用于获取动态加速域名在指定结算区域的流量汇总统计数据。用户提供查询的起止时间，指定域名列表、结算区域和数据粒度。返回结果包含总流量及按照时间片对应的详细流量数据。此接口适用于希望按结算区域查询动态加速域名流量的场景。 | POST | /api/report/flow/area/dwa |
| Querytrafficformultidomain | 该接口用于查询动态加速服务中多个域名的流量汇总。用户需提供加速域名和时间范围来获取流量统计信息。返回的数据包括每个时间片的流量和总流量的详细数据。有助于用户分析各域名的流量使用情况，以更好地管理和监控其网络资源。<br> | POST | /api/report/domainflow/dwa |
| Queryrequestnumberandpeakbandwidthformultidomain | 该接口用于查询动态加速产品下多个域名对应的流量和请求数及峰值带宽。用户通过提供查询时间范围和域名，接口返回域名总流量，总请求数，带宽峰值以及每个时间片下的详细数据。有助于用户分析网站的流量模式，识别高峰期，从而优化带宽资源配置和提高站点性能。 | POST | /api/report/flow-request/dwa |
| QueryKeywordaquanEdgereport | 查询特定关键字“aquan”的域流量，流量统计包含所有边缘节点的数据 | POST | /myview/CdnwEdgeReportForKeywordAquan |
| Queryedgebytehitratioservice | 查询多域名分钟级别的在边缘节点命中缓存的字节命中率 | POST | /api/report/flow/edge-hit-ratio/total |