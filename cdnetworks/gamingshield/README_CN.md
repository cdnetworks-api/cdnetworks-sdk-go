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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/gamingshield"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &gamingshield.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := gamingshield.{ActionName}Response{}
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
| Listgroupdetails | 用于查询网站分组详情列表，用户可以通过提供网站分组ID、名称、域名来查询对应网站分组详情信息，包括协议、端口、域名等。 | POST | /api/v1/hijack/group-info/get-detail-list |
| Createwebsitegroup | 用于新增网站分组，将相同内容的网站添加至一个分组中，用户访问时将自动导航至组内的健康域名，确保访问体验不受影响，网站分组中包含域名、协议和端口信息。 | POST | /api/v1/hijack/group-info/add |
| Updatewebsitegroup | 用于修改网站分组信息，用户可以通过该接口修改网站分组的协议、端口、关联域名等信息。 | POST | /api/v1/hijack/group-info/update |
| Deletewebsitegroup | 用于删除网站分组，用户可以通过提供网站分组ID来删除不存在关联的网站分组。 | POST | /api/v1/hijack/group-info/delete |
| Listentrydomains | 用于查询入口域名详情列表，用户可以通过提供入口域名ID等信息来获取详情信息，包括入口域名协议和端口，以及关联的网站分组信息等。 | POST | /api/v1/hijack/entry-domain/get-detail-list |
| Addentrydomain | 用于新增入口域名，用户可以提供入口域名的协议和端口和关联的网站分组ID来新增入口域名。 | POST | /api/v1/hijack/entry-domain/add |
| Updateentrydomain | 用于修改入口域名，当入口域名信息有变动时，用户可以提供入口域名ID来修改相关信息，包括关联的网站分组ID、端口、协议等。 | POST | /api/v1/hijack/entry-domain/update |
| Deleteentrydomain | 用于删除入口域名，当入口域名不需要时，用户可以提供入口域名来删除。 | POST | /api/v1/hijack/entry-domain/delete |
| Restartentrydomain | 用于重启入口域名，当入口域名为禁用时，用户可以通过提供入口域名来重启入口域名以恢复服务。 | POST | /api/v1/hijack/entry-domain/restart |
| Listnavigationrules | 用于查询导航规则列表，可以通过提供网站分组ID列表或者网站分组名称列表来获取指定导航规则的完整信息，包括规则开关、匹配条件等信息。 | POST | /api/v1/hijack/group-rule/get-list |
| Updatedistributionrule | 用于修改导航规则，当业务需要导航的用户请求或目标有更新时，可通过此接口进行更新。 | POST | /api/v1/hijack/group-rule/save |