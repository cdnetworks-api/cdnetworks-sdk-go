# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/propertymanagement
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/propertymanagement"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &propertymanagement.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := propertymanagement.{ActionName}Response{}
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
| GetAPropertyVersion | 获取加速项目版本的详细配置信息。 | GET | /cdn/properties/*/versions/* |
| UpdateAPropertyVersion | 该接口可用来更新加速项目版本。已部署到生产或演练环境的加速项目处于“冻结”状态，无法再次被更新或验证，该状态下的加速项目版本只有版本描述字段可被更新。 | PATCH | /cdn/properties/*/versions/* |
| DeleteAPropertyVersion | 删除一个加速项目版本。 | DELETE | /cdn/properties/*/versions/* |
| GetAProperty | 查询某个加速项目的信息，包括已创建的版本数以及哪些版本已部署等。 | GET | /cdn/properties/* |
| UpdateAProperty | 该接口用于更新加速项目的名称和描述。该操作不会创建加速项目新版本，也不会更新现有的加速项目版本。<br>如果要更改加速项目的配置，请调用“更新加速项目版本”接口。 | PATCH | /cdn/properties/* |
| DeleteAProperty | 根据ID删除加速项目。 | DELETE | /cdn/properties/* |
| GetListOfPropertyVersions | 获取加速项目版本列表。返回每个版本的摘要信息，包括其状态及关联的加速域名等。 | GET | /cdn/properties/*/versions |
| CreateAPropertyVersion | 创建一个新的加速项目版本。 | POST | /cdn/properties/*/versions |
| CreateAProperty | 创建加速项目，在加速项目中定义需要部署到CDN Pro服务器的一个或多个加速域名的配置。 | POST | /cdn/properties |
| GetListOfProperties | 该接口返回加速项目列表，包括每个加速项目的ID、最新版本号、注释、演练环境版本号、生产环境版本号，以及每个版本的最后更新时间等信息。可使用查询参数筛选加速项目。 | GET | /cdn/properties |