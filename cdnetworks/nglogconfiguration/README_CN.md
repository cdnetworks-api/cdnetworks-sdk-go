# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/nglogconfiguration
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/nglogconfiguration"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &nglogconfiguration.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := nglogconfiguration.{ActionName}Response{}
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
| GetAListOfLogConfigurations | 获取已创建的日志配置列表。 | GET | /cdn/report/logConfigs |
| CreateALogConfiguration | 创建日志配置，用于定义访问日志的输出格式。同一个日志配置可适用于一个或多个域名。每个域名必须有指定的日志配置，才能获取到该域名的日志。创建日志配置后，大约需要等待40分钟才会生效。 | POST | /cdn/report/logConfigs |
| GetALogConfiguration | 获取日志配置的详细信息。 | GET | /cdn/report/logConfigs/* |
| UpdateALogConfiguration | 更新日志配置。配置更新后，大约需要等待40分钟才会生效。 | PATCH | /cdn/report/logConfigs/* |
| DeleteALogConfiguration | 删除日志配置。 | DELETE | /cdn/report/logConfigs/* |
| GetAccessLogsForHostnames | 获取已部署的一个或多个加速域名的访问日志。这些日志是2小时或更早时间之前客户端请求所产生的日志。默认情况下，只能查询近14天内的日志。这些日志被切割成多个文件供下载。<br><br>要获取加速域名的日志，必须首先通过“创建日志投递配置”接口定义日志格式。注意：必须先创建日志投递配置，访问日志才会被收集。如未创建日志投递配置，则无法下载日志。<br><br>日志格式示例:<br>%cltip %rmtuser [%utctime] "%method %url %protocol" %statuscode %rspsize "%referer" "%ua" %rsptime<br><br>调用该接口时，将返回如下的日志信息和下载链接:<br>{<br>  "logs":[<br>          {<br>          "dateFrom":"2021-10-31T00:00:00Z",<br>          "dateTo":"2021-10-31T00:29:59Z",<br>          "fileSize":105878,<br>          "logUrl":"https://abc.example.com/logd/v2/download/0621c8fc885089805kea5f610797ff8ba92bc98c049c2bb308cbdb?traceId=ac6d696c657765625f74657374cf0000018dd01d89e8cd06d3",<br>          "hostname":"mydomain.domain.info"<br>          }<br>        ]<br>}<br><br>其中，logUrl字段返回日志文件（经过gzip压缩）的下载链接。下载链接有效期为24个小时。如链接已过期，需重新调用接口获取链接。日志文件内容示例：<br>9.8.7.6 - [31/Oct/2021:19:59:57 +0000] "GET http://mydomain.domain.info/i/js/tab.js HTTP/1.1" 304 529 "http://mydomain.domain.info/?q=downloads" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36" 76<br><br> | POST | /cdn/report/logDownload |