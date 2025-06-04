# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/manageassets
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/manageassets"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &manageassets.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := manageassets.{ActionName}Response{}
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
| Getvideolist | 调用getVideoList客户获取已上传视频的信息列表（包含：上传时间、视频名称、视频ID、视频状态等） | POST | /vod/videoManage/getVideoList |
| Videoedit | 调用videoEdit可以对单个视频的基础信息进行编辑。 | POST | /vod/videoManage/videoEdit |
| Videoblock | 调用videoBlock对已上传视频进行屏蔽禁用，被禁用的时候将不能再继续观看。 | POST | /vod/videoManage/videoBlock |
| Videoenable | 调用videoEnable对已禁用的视频进行重新启用。 | POST | /vod/videoManage/videoEnable |
| Deletevideo | 调用deleteVideo对已上传的视频进行删除。 | POST | /vod/videoManage/deleteVideo |
| Getmateriallist | 调用getMaterialList获取已上传素材的信息列表里 | POST | /vod/material/getMaterialList |
| Materialedit | 调用materialEdit可以对单个素材文件的基础信息进行编辑。 | POST | /vod/material/materialEdit |
| Deletematerial | 调用deleteVideo对已上传的素材文件进行删除。 | POST | /vod/material/deleteMaterial |
| Getaudiolist | 查询音频文件列表，按照上传时间倒序排序。支持分页去获取列表 | POST | /vod/audioManage/getAudioList |
| Editaudio | 对指定单个音频频的基础信息进行编辑 | POST | /vod/audioManage/editAudio |
| Deleteaudio | 通过该接口删除指定音频。 | POST | /vod/audioManage/deleteAudio |
| Getcategorylist | 获取标签分类列表 | POST | /vod/videoCategory/getCategoryList |
| Createcategory | 创建标签分类 | POST | /vod/videoCategory/createCategory |
| Deletecategory | 删除标签分类 | POST | /vod/videoCategory/deleteCategory |
| Deletevideobycategory | 根据标签分类删除视频 | POST | /vod/videoCategory/deleteVideoByCategory |