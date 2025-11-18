# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cacherefresh
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cacherefresh"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &cacherefresh.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := cacherefresh.{ActionName}Response{}
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
| Getpurgestatus | 查询刷新缓存任务的执行状态，可查看是否已经全网执行生效。 | POST | /ccm/purge/ItemIdQuery |
| Purgefilewithtag | 针对客户下有开启tag开关的域名内的有标识的tag的文件进行推送。 | POST | /api/content/tag/purge |
| Purgeurlsmatchingregex | 根据正则url方式清理CDN节点上缓存的文件内容 | POST | /api/content/regular-url/purge |
| Purge | 刷新CDN节点上缓存的文件内容，全网一般在1~3分钟内生效，目录推送单账号单域名每日上限为500 | POST | /ccm/purge/ItemIdReceiver |
| Querypurgelimit | 查询推送剩余量接口 | POST | /ccm/purge/limit |
| Querypurgeresiduals | 推送剩余量查询 | GET | /ccm/upperQuery |
| Getpurgequota | 该接口用于查询指定客户推送日上限和日剩余量，用户通过该接口获取其推送日上限和剩余量，包括url推送，目录推送，正则推送，tag推送。这对于客户是否能在继续提交任务至关重要，可以帮助客户及时发现是否达到日上限，当客户有大量任务要提交时，提前调整日上限。 | POST | /ccm/purge/quota |
| Getprefetchquota | 该接口用于查询指定客户预取日上限和日剩余量，用户通过该接口获取其预取日上限和剩余量，包括url个数和预取大小。这对于客户是否能在继续提交任务至关重要，可以帮助客户及时发现是否达到日上限，当客户有大量任务要提交时，提前调整日上限。 | POST | /ccm/prefetch/quota |