# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/sharedconfiguration
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/sharedconfiguration"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &sharedconfiguration.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := sharedconfiguration.{ActionName}Response{}
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
| Listsharedwafruleexceptions | 获取共享配置WAF规则例外。 | POST | /api/v1/waf/share/exception/get-list |
| Createsharedwafruleexception | 新增共享配置的WAF规则例外。 | POST | /api/v1/waf/share/exception/create |
| Updatesharedwafruleexception | 修改共享配置的WAF规则例外。 | POST | /api/v1/waf/share/exception/update |
| Deletesharedwafruleexception | 删除共享配置的WAF规则例外。 | POST | /api/v1/waf/share/exception/delete |
| Createcustomaction | 新增自定义处理动作，最多可配置5个自定义响应动作。 | POST | /api/v1/share-action/add-customize-act |
| Listcustomactions | 获取自定义处理动作列表。 | POST | /api/v1/share-action/get-customize-act-list |
| Updatecustomaction | 修改自定义处理动作。 | POST | /api/v1/share-action/update-customize-act |
| Deletecustomactions | 删除自定义处理动作。注意：无法删除已被引用的自定义处理动作。 | POST | /api/v1/share-action/delete-customize-act-batch |
| Createsharedwhitelistrule | 新增共享配置的白名单规则。 | POST | /api/v1/common/share-whitelist/add |
| Updatesharedwhitelistrule | 修改共享配置的白名单规则。 | POST | /api/v1/common/share-whitelist/update |
| Deletesharedwhitelistrule | 删除共享配置的白名单规则。 | POST | /api/v1/common/share-whitelist/delete |
| Listsharedwhitelistrules | 获取共享配置的白名单规则 | POST | /api/v1/common/share-whitelist/get-list |
| Createsharedratelimitingrule | 新增共享配置的频率限制规则。 | POST | /api/v1/share-rate-limit/add-rule |
| Updatesharedratelimitingrule | 修改共享配置的频率限制规则。 | POST | /api/v1/share-rate-limit/update-rule |
| Deletesharedratelimitingrule | 删除共享配置的频率限制规则。 | POST | /api/v1/share-rate-limit/delete-by-ids |
| Listsharedratelimitingrules | 获取共享配置的频率限制规则。 | POST | /api/v1/share-rate-limit/get-rule-list |
| Createappapiexceptionfeature | 新增共享配置-APP/API例外特征。 | POST | /api/v1/dms/service-feature/add |
| Queryappapiexceptionlist | 查询共享配置-APP/API例外的列表。 | POST | /api/v1/dms/service-feature/get-list |
| Deleteappapiexceptionfeature | 删除共享配置-APP/API例外特征。 | POST | /api/v1/dms/service-feature/delete |
| Queryappapiexceptionfeaturedetail | 查看共享配置-APP/API例外特征详情。 | POST | /api/v1/dms/service-feature/get-detail |
| Queryappapiexceptionfeaturereferencedhostnames | 查看共享配置-APP/API例外特征关联的域名列表。 | POST | /api/v1/dms/service-feature/get-relate-domain-list |
| Updateshareconfigurationsappapiexceptionfeature | 修改App/API例外（共享配置）。 | POST | /api/v1/dms/service-feature/update |
| Listsharedcustomrules | 获取共享配置的自定义规则。 | POST | /api/v1/share-customize-rule/get-list |
| Disassociateshareratelimit | 解除共享配置的频率限制规则与域名的关联。 | POST | /api/v1/common/share-rate-limit/disassociate |
| Associateshareratelimit | 将共享配置的频率限制规则与域名的关联。 | POST | /api/v1/common/share-rate-limit/associate |
| Associatesharecustomizerule | 将共享配置的自定义规则与域名的关联。 | POST | /api/v1/common/share-customize-rule/associate |
| Disassociatesharecustomizerule | 解除共享配置的自定义规则与域名的关联。 | POST | /api/v1/common/share-customize-rule/disassociate |
| Associatesharecustomizebots | 将共享配置的定义Bot与域名的关联。 | POST | /api/v1/common/share-customize-bots/associate |
| Disassociatesharecustomizebots | 解除共享配置的自定义Bot与域名的关联。 | POST | /api/v1/common/share-customize-bots/disassociate |
| Associatesharedwhitelistrule | 将共享配置的白名单规则与域名的关联。 | POST | /api/v1/common/share-whitelist/associate |
| Disassociatesharedwhitelistrule | 解除共享配置的白名单规则与域名的关联。 | POST | /api/v1/common/share-whitelist/disassociate |
| Associatedmsshareservicefeature | 将共享配置的APP/API例外与域名的关联。 | POST | /api/v1/dms/service-feature/relateDomains |
| Disassociatedmsshareservicefeature | 将共享配置的APP/API例外与域名解除关联。 | POST | /api/v1/dms/service-feature/disRelateDomains |