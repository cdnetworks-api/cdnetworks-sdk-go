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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/toolservice"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &toolservice.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := toolservice.{ActionName}Response{}
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
| Bandwidthlimitservice | 设置/取消指定域名的带宽限制。 | POST | /api/tools/setBandwidthLimit |
| Querybandwidthlimittasklistservice | 此接口用于查询账号下的带宽限制任务列表，返回所有有带宽的任务详细信息。返回内容包括域名、任务名称以及设置的最大带宽值。该接口适用于需要评估和管理流量控制策略的场景，帮助用户快速识别并管理当前设置的带宽限制任务。 | POST | /api/tools/queryBandwidthLimitTaskList |
| Icpqueryservice | 查询所指定的域名是否在中国大陆工信部进行备案。 | GET | /api/icp |
| Ipdomainservice | 该接口用于根据IP地址查询正在使用该IP的域名。用户输入IP地址来获取与该IP相关联的域名列表。接口返回的信息包含IP当前使用状态，以及使用该IP的域名列表。在实际应用中，此接口可帮助用户检测特定IP的域名使用情况，适用于网络监控和管理。 | POST | /api/tools/ip/domain-list |
| Queryallbandwidthlimittasklistservice | 该接口用于查询用户账号下配置的所有带宽限制任务。用户在调用时可以选择是否包含任务所涉及的所有客户域名以及决定是否返回所有任务状态的信息。返回的数据以清单方式展示每个带宽限制任务的详细信息，包括任务名称、类别、状态以及相关控制策略和参数。此接口有助于用户管理带宽设置，可以及时对特定的流量和请求进行有效的控制和处理。 | POST | /api/tools/queryAllBandwidthLimitTaskList |