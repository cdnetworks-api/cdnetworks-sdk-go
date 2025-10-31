# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/webhooks
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/webhooks"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &webhooks.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := webhooks.{ActionName}Response{}
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
| GetAListOfWebhooks | 获取已创建的webhook接口列表。当你创建部署任务，刷新任务，预取任务，或者配置验证任务时，都可以使用webhook接口来接收任务完成通知。 | GET | /cdn/webhooks |
| GetAWebhook | 该接口返回webhook接口的详细信息，包括webhook接口基本信息以及接口调用情况。 | GET | /cdn/webhooks/* |
| DeleteAWebhook | 删除webhook接口。删除后，如果有异步任务引用了该webhook接口，且任务仍在执行中，则任务执行完毕后不会触发webhook接口调用。 | DELETE | /cdn/webhooks/* |
| UpdateAWebhook | 更新webhook接口。只有请求体中携带的字段才会被更新。 | PATCH | /cdn/webhooks/* |
| CreateAWebhook | 创建webhook接口。当某个异步任务完成时，可用webhook接口来接收任务完成通知。当你创建部署任务，刷新任务，预取任务，或者配置验证任务时，都可以使用webhook接口来接收任务完成通知。 | POST | /cdn/webhooks |