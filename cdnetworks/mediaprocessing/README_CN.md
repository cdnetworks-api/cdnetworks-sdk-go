# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/mediaprocessing
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/mediaprocessing"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &mediaprocessing.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := mediaprocessing.{ActionName}Response{}
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
| Videoclip | 调用videoClip对已上传的视频进行按时间区间剪切视频并转换关键参数。 | POST | /vod/videoManage/videoClip |
| Videoclipquery | 调用videoClip查询剪切任务的剪切结果。 | POST | /vod/videoManage/videoClipQuery |
| Videoconcat | 调用videoConcat对多个视频按指定的顺序进行拼接，生成新的视频文件。 | POST | /vod/videoManage/videoConcat |
| Videoconcatquery | 调用videoConcatQuery查询视频拼接任务结果 | POST | /vod/videoManage/videoConcatQuery |
| Createclearadtask | 调用createClearAdTask为指定视频创建AI去广告任务，系统会自动针对该视频进行AI去广告操作（增值服务功能，如若需要请联系客服开通“AI清除广告”增值服务）。 | POST | /vod/ai/createClearAdTask |
| Clearadtaskquery | 调用clearAdTaskQuery查询AI去广告任务处理状态和结果。 | POST | /vod/ai/clearAdTaskQuery |
| Transcode | 调用transCode对指定的视频进行转码（主要包含：分辨率转码、加水印转码） | POST | /vod/videoManage/transCode |
| Startworkflow | 调用startWorkflow对指定的视频进行工作流执行操作。主要用于批量修改视频信息 | POST | /workflow/startWorkflow |