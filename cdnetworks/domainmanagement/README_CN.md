# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/domainmanagement
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/domainmanagement"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &domainmanagement.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := domainmanagement.{ActionName}Response{}
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
| Enablesingledomainservice | 启用某个状态为“禁用”的加速域名，使用已有的配置提供加速服务。 | PUT | /api/domain/*/enable |
| Disablesingledomainservice | 禁用指定加速域名，禁用后加速域名的请求将被直接拒绝，不会回源。 | PUT | /api/domain/*/disable |
| Deleteapidomainservice | 删除已添加的某个加速域名。删除后不能启用，只能重新创建加速域名。 | DELETE | /api/domain/* |
| Getfuzzypagingdomainlist | 查询用户账号下所有的、或者指定的加速域名和状态，每个加速域名包含概要信息，返回的加速域名列表按照首字母顺序排序。 | POST | /api/domain/domainList |
| Querydomainbyoriginip | 查询用户账号下，源站IP对应的所有域名名称列表。 | GET | /api/originaldomainlist |
| Addcdndomain | 为指定的域名申请加速服务 | POST | /cdnw/api/domain |
| Updatecdndomain | 修改指定加速域名的配置。 | PUT | /cdnw/api/domain/* |
| Querycdndomainservice | Get domain basic config. | GET | /cdnw/api/domain/* |
| Queryapidomainlistservice | 查询用户账号下所有的、或者指定cname-label的加速域名和状态，每个加速域名包含概要信息，返回的加速域名列表按照首字母顺序排序。<br>请注意：禁用的域名（即enabled值为false）不可修改。 | GET | /api/domain |
| Adddomainalias | 通过别名方式创建加速域名。别名的配置与所关联的域名相同（合同信息，服务类型，加速区域，证书）。 | POST | /api/domain/alias/ |
| Queryapidomainattribution | 查询域名列表 | GET | /api/domainlist |
| Queryrelationshipbetweendomainandaliasbasedonalias | 查询主域名及别名 | GET | /api/domain/domain-and-alias/* |
| Deletedomainalias | 删除域名别名接口 | DELETE | /api/domain-alias |
| Queryrelationshipbetweendomainandaliasbasedondomain | 查询域名别名关系 | POST | /api/domain-alias |
| Getpagingdomainlist | 查询用户账号下所有的、或者指定的加速域名和状态，每个加速域名包含概要信息，返回的加速域名列表按照版本排序。<br> | POST | /api/domain/list |
| Querycustomeranycastipforwplus | 查询客户AnycastIP | GET | /api/anycast-ips |
| Predeploychangeserverconfig | 通过接口自助预部署【接入域名跳转】配置。接口url的*可为域名名称或域名id。 | PUT | /api/predeploy/changeserver/* |
| Updatechangeserver | 通过接口自助修改【接入域名跳转】配置。接口url的*可为域名名称或域名id。 | PUT | /api/config/changeserver/* |
| Querychangeserver | 通过接口自助查询【接入域名跳转】配置。接口url的*可为域名名称或域名id。 | GET | /api/config/changeserver/* |
| Updatecustomeranycastiprecordstatus | 修改Anycast IP的record status | POST | /api/anycast-ips/record-status |