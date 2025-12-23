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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
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
| Createcertificateapplicationorder | 创建证书申请订单. | POST | /api/certificate/order/create |
| Listcertificateapplicationorders | 查询证书申请订单列表。 | POST | /api/certificate/order/list |
| Getdcvcontent | 查询CA机构进行域名所有权验证时的验证内容。 | POST | /api/certificate/order/domain/validate/info |
| Cancelcertificateapplicationorder | 取消证书申请订单。 | POST | /api/certificate/order/cancel |
| Getcertificateapplicationorder | 查询证书申请订单详情 | POST | /api/certificate/order/detail |
| Reissuecertificateforwplus | 该接口用于重颁发证书。可以通过提供证书ID、证书描述、证书算法、验证方式、是否自动验证、是否自动部署、通用名称、主体备用名称重颁发证书，调用成功时接口会返回销售订单id。 | POST | /api/certificate/reissue |
| Createcertificateapplicationorderforterraform | 该接口专为Terraform场景设计，用于创建证书申请订单。用户需提供订单的技术联系人和管理联系人信息，以及证书签名请求（CSR）相关信息，包括通用名称、备用域名（SANs）、国家、城市等。同时，还需指定证书的品牌、类型、算法、有效期等规格，并可选择是否开启自动验证、自动部署和自动续签功能。针对域名控制权验证，支持配置DNS托管商信息。接口成功创建订单后，将返回订单ID和相关的关联信息。 | POST | /api/terraform/certificate/order/create |
| Listcertificateapplicationordersforterraform | 该接口用于Terraform场景查询证书申请订单列表。用户可以通过订单ID、域名、订单状态、证书名称、订单创建起止时间等条件进行筛选，并支持分页查询。响应结果包含订单列表、分页信息等。 | POST | /api/terraform/certificate/order/list |
| Getcertificateapplicationorderforterraform | 该接口用于在Terraform场景下，查询指定证书申请订单的详细信息，包括订单状态、证书详情等。用户可通过`purchaseRecordId`或`orderId`进行查询。 | POST | /api/terraform/certificate/order/detail |
| Cancelcertificateapplicationorderforterraform | 该接口用于Terraform场景取消已提交的证书申请订单。用户需提供采购记录ID或订单ID来指定要取消的订单。取消成功后，订单状态将更新。 | POST | /api/terraform/certificate/order/cancel |
| Batchgetdcvcontent | 该接口用于批量查询CA机构在进行域名所有权验证时所需的验证内容。用户需提供采购记录ID列表或域名列表进行查询，接口将返回相应的证书域名验证信息。 | POST | /api/certificate/orders/domain/validate/info |