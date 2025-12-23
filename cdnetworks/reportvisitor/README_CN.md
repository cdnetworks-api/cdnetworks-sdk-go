# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportvisitor
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportvisitor"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &reportvisitor.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportvisitor.{ActionName}Response{}
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
| Querytotalnumberofuniqueipundersingledomain | 该接口用于查询单个域名内的各流名的独立IP数量。用户提供具体的域名和时间，获取在此期间指定流的信息（以天为粒度）。接口返回的结果包括域名及其下对应流名的独立IP统计信息。有助于用户了解每路流的访问者分布，从而优化流量配置或评估流媒体服务的使用情况。 | POST | /api/report/visitor/total/stream |
| Statisticalanalysisofchinamainlandprovincevisitors | 该接口用于分析中国大陆各省份的用户访问情况。用户提供时间范围、域名来获取详细的统计数据。接口返回每个地区的请求数、成功和失败请求数、成功率、总流量、独立访客数和页面浏览量等信息。这些数据可以帮助网站管理员分析流量来源，提高区域性市场策略，改进用户体验。 | POST | /api/report/visit/analysis/combine/province |
| Statisticalanalysisofcountries | 该接口用于分析不同国家和地区的访问统计数据，用户可以通过提供起止时间、域名、地区代码等参数进行查询。响应结果包括各地区的请求数、成功请求数、失败请求数、成功率、总流量、独立访客数、页面浏览量和IP数量等详细统计信息。此接口适用于需要了解网站在全球不同地区的流量和用户行为的用户，可帮助企业优化其市场策略和技术资源分配。<br> | POST | /api/report/visit/analysis/combine/country |
| Reportvisitorcustomtopdailyservice | 查询天粒度访客IP的自定义TOP排行 | POST | /api/report/visitor/custom-top/daily |