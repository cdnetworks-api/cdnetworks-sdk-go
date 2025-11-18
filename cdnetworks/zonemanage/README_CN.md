# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/zonemanage
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/zonemanage"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &zonemanage.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := zonemanage.{ActionName}Response{}
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
| Getzonelist | 用于获取zone列表 | GET | /api/clouddns/zones |
| Getzonebyid | 根据ID获取指定zone | GET | /api/clouddns/zones/* |
| Updatazonebyid | 更新zone配置 | PUT | /api/clouddns/zones/* |
| Deployzoneconfig | 部署zone配置到staging或production | POST | /api/clouddns/zones/*/deployment |
| Getrecordlistbyzone | 获取zone下的所有解析记录 | GET | /api/clouddns/zones/*/records |
| Getrecordbyid | 根据zoneId和recordId获取指定record | GET | /api/clouddns/zones/*/records/* |
| Bulkcreaterecordsbyzone | 为指定zone批量创建record。该方法需要请求体为列表形式，OpenApi暂不支持配置调用，请使用另外的工具测试调用。 | POST | /api/clouddns/zones/*/records |
| Updaterecordbyid | 更新指定解析记录 | PUT | /api/clouddns/zones/*/records/* |
| Bulkupdaterecordsbyzoneid | 批量更新指定zone的指定record。该方法需要请求体为列表形式，OpenApi暂不支持配置调用，请使用另外的工具测试调用。 | PUT | /api/clouddns/zones/*/records |
| Deleterecordbyid | 根据zoneId和recordId删除指定解析记录 | DELETE | /api/clouddns/zones/*/records/* |
| Createzone | 用于新建zone | POST | /api/clouddns/zones |
| Deletezone | 删除指定zone | DELETE | /api/clouddns/zones/* |
| Batchcreatezone | 用于批量新建zone | POST | /api/clouddns/zones/bulk |
| Batchdeletezone | 批量删除Zone，最大可同时删除20个 | DELETE | /api/clouddns/zones/bulk/* |
| Updateztsbulk | 批量创建更新ZTS配置信息，用户可通过接口自助创建更新ZTS的配置信息，实现Zone数据的同步 | PUT | /api/clouddns/zts/bulk |
| Deleteztsbulk | ZTS配置删除接口，删除Zone上已配置的ZTS信息，删除之后将不会再自动同步Zone的配置修改 | DELETE | /api/clouddns/zones/zts/bulk |