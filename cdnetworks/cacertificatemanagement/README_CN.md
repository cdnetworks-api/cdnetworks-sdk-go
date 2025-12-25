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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cacertificatemanagement"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &cacertificatemanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := cacertificatemanagement.{ActionName}Response{}
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
| CreateACaCertificate | 新增CA证书 | POST | /api/certificate/ca |
| AssociateDomainWithCaCertificate | 创建CA证书与域名关联关系接口 | POST | /api/certificate/ca/associatedomain |
| DeleteACaCertificate | 删除CA证书 | DELETE | /api/certificate/ca/* |
| DisassociateDomainWithCaCertificate | 解除CA证书与域名关联关系接口 | POST | /api/certificate/ca/disassociatedomain |
| UpdateACaCertificate | 更新CA证书 | PUT | /api/certificate/ca/* |
| Listcacertificates | 查询ca证书列表 | POST | /api/certificate/ca/list |
| Getcacertificate | 获取ca证书详情 | GET | /api/certificate/ca/* |