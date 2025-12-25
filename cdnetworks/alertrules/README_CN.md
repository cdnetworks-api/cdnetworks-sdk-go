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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/alertrules"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &alertrules.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := alertrules.{ActionName}Response{}
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
| Createcloudmonitorrealtimealarmrule | 创建云监控的实时监控告警规则。用户可传入监控规则名称、监控资源对象、告警触发条件阈值、告警联系人等信息生成实时监控规则。适用于对控制台云产品业务需要及时关注监控数据变化的场景。接口创建监控规则当前仅支持按域名监控的部分CDN监控项。<br> | POST | /api/cloudmonitor/alarm/real-time/add |
| Editcloudmonitorrealtimealarmrule | 编辑云监控的实时监控告警规则。本接口修改规则按覆盖式修改。用户可传入监控规则id、可传入要修改的监控规则名称、监控资源对象、告警触发条件阈值、告警联系人等信息将覆盖修改监控规则。<br><br>接口管理监控规则当前仅支持按域名监控的部份监控项。 | POST | /api/cloudmonitor/alarm/real-time/edit |
| Querycloudmonitorrealtimealarmrule | 	<br>查询云监控的实时监控告警规则。用户可传入监控规则id、查询监控规则名称、监控资源对象、告警触发条件阈值、告警联系人等监控规则明细信息。<br><br>接口管理监控规则当前仅支持按域名监控的部份监控项。 | POST | /api/cloudmonitor/alarm/real-time/query |
| Deletecloudmonitorrealtimealarmrule | 删除云监控的实时监控告警规则。用户可传入监控规则id删除指定监控规则。<br><br> | POST | /api/cloudmonitor/alarm/real-time/delete |