# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/imagemanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/imagemanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &imagemanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := imagemanage.{ActionName}Response{}
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
| Vmpqueryimage | 查询用户可以使用的镜像列表。显示出的镜像资源列表包括用户自定义的镜像及边缘计算平台提供的公共镜像。 | GET | /vmp/images |
| Vmpcreateimage | 支持将某个虚拟机系统盘制作成镜像，之后便可将其用于创建新的虚拟机。建议在制作镜像期间关闭虚拟机或者停止虚拟机上的应用或服务，以免影响镜像数据的完整性，待镜像制作完成，再启动虚拟机及其应用。此类操作创建的镜像在镜像查询接口中返回的镜像属主是SNAPSHOT，表示用虚拟机快照做的镜像。 | POST | /vmp/images |
| Vmpremoveimage | 删除您自定义镜像，删除镜像不影响已经创建的虚拟机，只是后续不能再使用该镜像创建新的虚拟机。您只能删除自己创建的自定义镜像，其他客户的自定义镜像以及公共镜像不能删除。 | DELETE | /vmp/images/* |
| Deployvmpimagepreheating | 用于提取预热客户私有镜像，该接口为异步接口，镜像预热结果需另外查询 | PUT | /vmp/images/preHeating |
| Queryvmpimagepreheatingstate | 用于查询镜像预热状态 | GET | /vmp/images/preHeatingInfo/* |