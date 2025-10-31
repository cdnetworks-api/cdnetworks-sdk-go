# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cgmanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cgmanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &cgmanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := cgmanage.{ActionName}Response{}
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
| Editcontrolgroup | 该接口用于修改指定的ControlGroup信息，包括ControlGroup的名称、域名列表及账号列表。用户可根据需求选择覆盖或者追加的方式进行更新。针对自定义类型的ControlGroup，用户可以修改其名称和域名列表，而客户类型和合同类型的ControlGroup则不允许此类修改。此外，还可以为ControlGroup指定有权限访问的账号列表。此接口便于用户灵活管理ControlGroup的配置信息，通过调整这些设置可以控制访问权限和适用性，从而更好地服务于业务需求。 | PUT | /user/control-groups/* |
| Deletecontrolgroup | 删除ControlGroup接口的功能是用于移除指定的服务组。用户需要传递要删除的ControlGroup标识信息作为入参，以执行删除操作。本接口的返回值包含三个关键数据：操作成功与否的标记、状态码以及提示消息，帮助用户了解操作的执行结果及错误原因。此接口尤为适用于场景中需要对系统中指定的服务组进行管理和清理时，保证系统资源的动态调整和优化。 | DELETE | /user/control-groups/* |
| Createcontrolgroup | 新增ControlGroup接口用于创建一个新的服务组，用户需要提供如服务组编号、名称等关键信息。响应参数包括指定的域名列表和有权限访问的账号列表。其中，域名字符串数组用于指定服务组所包含的域名，而账号对象数组用于指定有权限访问该服务组的账号，确保相应人员能够访问和管理此组。此接口适用于需要组织和权限管理的场景，例如在大型系统中管理多个子系统的访问控制。通过使用该接口，用户可以高效地分配和控制权限结构。 | POST | /user/control-groups |
| Querycontrolgrouplist | 该接口用于查询具有访问权限的服务组列表，帮助用户获取与其账号相关联的服务组详细信息。通过调用此接口，可以获得服务组的基本属性，包括服务组的名称、编号、ID，以及类型等信息。此外，该接口还提供成功或失败状态标志和状态消息，帮助用户判断请求的处理结果。此接口适用于需要管理和检索用户权限服务组信息的场景，协助用户高效地进行权限维护和管理。 | GET | /user/control-groups |
| Querycustomizedcontrolgroup | 该接口用于查询指定的ControlGroup信息，仅支持查询自定义CG，用户可通过其进行特定ControlGroup的详细信息获取，包括其ID、名称、所包含的域名列表以及有权限访问的账号列表。该接口返回结构中会提供状态码和消息提示，指示请求的成功与否。用户主要应用于管理和监控自定义的ControlGroup，确保对特定域名的访问控制正确，并便于权限分配管理。通过对返回信息的解析，用户能够掌控当前ControlGroup的设置及其状态。 | GET | /user/control-groups/* |
| Querysubaccountrelatedcontrolgroup | 该接口用于查询特定子用户已经关联的服务组列表。当用户提供子用户的登录名作为参数时，系统将返回该用户所关联的服务组信息，包括服务组的编号和名称。这使得用户可以方便地查看和管理不同子用户的权限与所属服务组情况。接口的响应结构还包含请求的状态码和信息，帮助用户了解请求是否成功及其详情。 | GET | /user/sub-account/control-groups/* |
| Batchassociateordetachcontrolgroupwithsubaccount | 该接口用于为指定用户批量添加或撤销关联服务组。用户需提供子用户的登录名和指定的Control Group编号列表，并选择操作类型为添加或撤销关联。该接口返回请求状态的信息，包括成功或失败的状态码和具体请求结果消息。应用场景包括快速管理多个子用户与服务组的关联关系，使管理员能够高效地进行批量操作，满足组织内复杂和动态的服务组管理需求。 | POST | /user/sub-account/control-groups |
| Getdomainlistofcontrolgroup | 该接口的功能是获取当前用户账号有权限访问的服务组（CG）列表及其关联的域名信息。用户在调用时无需提供具体的参数，系统会返回一个包含CG名称和对应域名的完整列表。返回结果包括每个服务组的编号和关联的域名。此接口特别适用于需要快速了解某账号可管理或监控的域名全貌的场景，帮助用户有效管理资源和权限。 | POST | /user/cgdomainlist |