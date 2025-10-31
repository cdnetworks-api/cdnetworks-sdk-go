# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/uploadassets
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/uploadassets"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &uploadassets.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := uploadassets.{ActionName}Response{}
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
| Getuploadtoken | 调用getUploadToken获取视频文件上传地址和凭证 | POST | /vod/videoManage/getUploadToken |
| Getaudiouploadtoken | 调用getAudioUploadToken获取音频文件上传地址和凭证，支持批量获取多个音频上传地址和凭证（最多50个）。 | GET | /vod/audioManage/getAudioUploadToken |
| Getmaterialuploadtoken | 调用getMaterialUploadToken获取素材上传地址和凭证，支持批量获取多个素材地址和凭证（最多50个）。 | POST | /vod/material/getMaterialUploadToken |
| Pullvideo | 调用pullVideo向后台设置要拉取的视频url。后台定时自动完成第三方平台url视频拉取并保存。支持批量设置拉取任务。 | POST | /vod/videoManage/pullVideo |
| Pullvideoquery | 调用pullVideoQuery可以查询拉取任务完成情况。 | POST | /vod/videoManage/pullVideoQuery |