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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngcertificate"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &ngcertificate.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ngcertificate.{ActionName}Response{}
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
| CreateACertificate | 创建证书及其第一个版本。您可以上传私钥/公钥对来创建证书，或者选择生成自签名证书。 | POST | /cdn/certificates |
| GetListOfCertificates | 获取证书列表。接口返回证书的ID、名称、部署到生产环境的版本、部署到演练环境的版本、最新的版本号、证书的最后更新时间以及证书的过期时间等信息。 证书列表默认按证书的最后更新时间降序排序。 | GET | /cdn/certificates |
| GetACertificate | 获取证书的详细信息，包括证书的版本以及证书使用情况。 | GET | /cdn/certificates/* |
| UpdateACertificate | 您可以使用该接口来更新证书。相关字段与”创建证书”接口中的字段相同。 如果证书已在加速项目中使用且已部署到生产环境或演练环境，我们建议您调用“创建部署任务”接口来部署更新后的证书，否则CDN Pro仍会使用旧版本的证书来响应用户请求。 | PATCH | /cdn/certificates/* |
| DeleteACertificate | 该接口用来删除证书。当证书在生产环境或演练环境中使用时不可删除。 | DELETE | /cdn/certificates/* |
| DownloadTheCsr | 下载证书签名请求 (CSR) 文件 。您可以向证书颁发机构提交CSR，申请证书。申请到证书后，请调用“更新证书”接口更新证书到CDN Pro。 | GET | /cdn/certificates/*/csr |
| GetACertificateVersionsDetails | 获取某个证书版本的详细信息，包括到期日期、加密算法、私钥长度、指纹，以及是否在生产环境和演练环境中部署等信息。 | GET | /cdn/certificates/*/versions/* |