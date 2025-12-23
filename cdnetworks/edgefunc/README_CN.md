# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/edgefunc
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/edgefunc"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &edgefunc.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := edgefunc.{ActionName}Response{}
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
| Uploadfunccode | 函数代码上传接口，针对不使用 CloudIDE ，而是直接调用接口上传函数代码的场景 | POST | /edgefunc/upload |
| Createedgefunctrigger | 该接口用于创建函数触发器。用户需要提供域名和路由列表作为请求参数。成功创建后，接口将返回响应码和创建成功的触发器列表。 | POST | /api/v2/cdn_triggers |
| Queryedgefunctrigger | 该接口用于查询函数触发器列表。用户可以根据域名或函数名称进行筛选，并支持分页查询。响应中将返回触发器的详细信息列表。 | GET | /api/v2/cdn_triggers |
| Deletefuncdomaintrigger | 该接口用于删除一个已存在的函数触发器。用户需要通过REST参数 `id`指定待删除的函数触发器的唯一标识。 | DELETE | /api/v2/cdn_triggers/* |