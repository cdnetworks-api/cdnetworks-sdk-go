# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/protectassets
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/protectassets"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &protectassets.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := protectassets.{ActionName}Response{}
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
| Listapidefinitions | 获取API定义基础信息全量列表。 | POST | /api/v1/sam/api-define/basic/query |
| Queryapidefinitiondetail | 获取API定义详情。 | POST | /api/v1/sam/api-define/detail |
| Listapiassetdiscovery | 查询API发现列表。 | POST | /api/v1/sam/api-discovery/get-list |
| Feedbackwrongapiassetdiscovery | 上报错误的API发现。 | POST | /api/v1/sam/api-discovery/false-marking |
| Createapidefinition | 新增API定义。 | POST | /api/v1/sam/api-define/add |
| Updateapidefinition | 修改API定义。 | POST | /api/v1/sam/api-define/update |
| Batchdeleteapidefinition | 批量删除API定义。 | POST | /api/v1/sam/api-define/delete |