# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/usermanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/usermanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &usermanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := usermanage.{ActionName}Response{}
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
| Usercustomerupdateservice | 修改自身创建的客户信息。 | POST | /api/user/customer/update |
| Addsubaccount | 新增子账号 | POST | /sub-account |
| Querysubaccountinfo | 根据子用户登录名查询账号信息 | GET | /sub-account/* |
| Getsubaccountlist | 通过主账号获取其下的子用户列表 | POST | /sub-account/list |
| Updatesubaccount | 修改指定子用户的账号基本信息。 | PUT | /sub-account |
| Deletesubaccount | 删除指定的子用户 | DELETE | /sub-account/* |
| Querypolicyattachedmainaccountorsubaccount | 该接口用于查询指定用户已关联的权限策略列表。用户通过输入其登录名可以获得与之关联的策略信息，该信息包括策略的ID、名称、描述内容、类型（系统策略或自定义策略）等，支持多语言国际化描述。接口返回的结果包含状态码和相关信息，以帮助用户了解策略的具体属性和分类。这在管理用户权限和制定或修改权限策略时非常有用，帮助系统管理员更高效地进行权限配置和调整。 | GET | /user/policy-attached/* |
| Batchaddorrevokepolicytosubaccount | 该接口用于为指定的子用户批量添加或撤销权限策略，通过输入子用户的登录名以及权限策略的标识列表，可以选择执行添加或撤销操作。添加权限时，系统会为子账号增加相应策略，以扩展其权限范围；而在撤销权限时，则会移除指定的策略，减少子账号的访问权限。返回值包括请求状态码和操作信息提示，用户可以通过这些信息确认批量操作的成功与否。此接口适用于需要集中管理子账号权限的场景，简化权限的批量调整过程。 | POST | /user/policies |