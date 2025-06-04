# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/certificationmanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/certificationmanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &certificationmanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := certificationmanage.{ActionName}Response{}
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
| Querycertificatelist | ​查看SSL证书的列表及信息，包括证书ID、证书名称、是否共享、使用当前证书的域名信息等。 | GET | /api/ssl/certificate |
| Getcertificatecontent | 下载证书内容，包括：证书crt文件内容，证书key文件内容，证书ca文件内容。加密方式和上传方式一致。 | GET | /api/ssl/content/*/download |
| Addcertificateservicev2 | 新增证书接口，包括证书名称、证书公钥（crt和ca内容合并）、证书密钥、csrid、备注<br> | POST | /api/certificate |
| Deletecertificate | 删除SSL证书，如果查询证书还有关联域名在使用，则无法删除，需要使用【修改域名配置】接口解除域名与证书的关联。 | DELETE | /api/certificate/* |
| Querycertificateinfo | 根据证书ID查看证书详情。<br> | GET | /api/certificate/* |
| Editcertificatev2 | 修改客户证书，可重新上传证书。 | PUT | /api/certificate/* |
| Querycertificatecontent | 查看证书文件内容 | GET | /api/certificate/*/content |
| Querycertificaterelateddomains | 查看证书关联的加速域名 | GET | /api/certificate/*/domain |
| Querydomainmulticertconfig | 查询域名证书配置 | GET | /api/config/certificate/v2/* |
| Reissuecertificateforwplus | 该接口用于重颁发证书。可以通过提供证书ID、证书描述、证书算法、验证方式、是否自动验证、是否自动部署、通用名称、主体备用名称重颁发证书，调用成功时接口会返回销售订单id。 | POST | /api/certificate/reissue |