# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/certificateapplication
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/certificateapplication"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &certificateapplication.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := certificateapplication.{ActionName}Response{}
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
| Createcertificateapplyingorder | 创建证书申请订单. | POST | /api/certificate/order/create |
| Getcertificateapplyingorderlist | 查询证书申请订单列表。 | POST | /api/certificate/order/list |
| Getdomaincontrolvalidationcontent | 查询CA机构进行域名所有权验证时的验证内容。 | POST | /api/certificate/order/domain/validate/info |
| Cancelcertificateapplyingorder | 取消证书申请订单。 | POST | /api/certificate/order/cancel |
| Querycertificatesalesorderdetailforwplus | 查询证书申请订单详情 | POST | /api/certificate/order/detail |