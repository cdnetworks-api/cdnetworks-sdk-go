# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/secretmanagement
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/secretmanagement"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &secretmanagement.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := secretmanagement.{ActionName}Response{}
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
| CreateASecret | 创建“保密信息”来存放敏感内容，避免敏感内容直接暴露在加速项目的边缘逻辑中。在边缘逻辑中使用$SECRET(secretName)语法来引用保密信息。 | POST | /cdn/secrets |
| GetAListOfSecrets | 获取保密信息列表。可使用查询参数筛选保密信息。 | GET | /cdn/secrets |
| GetASecret | 获取保密信息的详情，包括其在演练环境及生产环境中被使用的情况。 | GET | /cdn/secrets/* |
| UpdateASecret | 使用该接口更新保密信息。如果您修改了保密信息的名称，则以原有名称引用保密信息的加速项目版本将无法通过验证。此类加速项目版本必须移除对该保密信息的引用或者采用更新后的名称重新引用，才能通过验证。如果修改了保密信息存放的内容，则使用到该保密信息的加速项目在重新部署前必须进行重新验证。在保密信息更新后，在已部署加速项目中已引用的保密信息不会自动被更新。 | PATCH | /cdn/secrets/* |
| DeleteASecret | 删除保密信息。如果保密信息已在加速项目中使用，在删除保密信息之前，必须先从加速项目中解除引用，并重新验证和部署加速项目。 | DELETE | /cdn/secrets/* |