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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/domainconfigurtion"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &domainconfigurtion.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := domainconfigurtion.{ActionName}Response{}
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
| Batchupdateapidomainservice | 修改指定的多个加速域名的配置信息，禁用的域名不可修改（慎用！） | POST | /api/domain/batchupdate |
| Editdomainconfig | 修改指定加速域名的基础配置。接口调用url的*可为域名名称或域名ID | PUT | /api/domain/* |
| Updatetimecontrolservice | 修改时间戳防盗链 | PUT | /api/config/timecontrol/* |
| Querytimecontrolservice | 查看时间戳防盗链 | GET | /api/config/timecontrol/* |
| Editdomainredirectconfig | 通过接口自助实现，url重定向功能 | PUT | /api/config/InnerRedirect/* |
| Edithttpheaderconfig | 通过接口自助实现http头部增删改功能，可通过在cdn层对客户实现个性化http头部控制，让客户在不需要修改源站的情况下，实现定制化的http头部和加速效果的提升。接口url的*可为域名名称或域名id。 | PUT | /api/config/headermodify/* |
| Queryhttpheaderconfig | 通过接口自助查询http头部配置。接口url的*可为域名名称或域名id。 | GET | /api/config/headermodify/* |
| Editcachetimeconfig | 通过接口自助实现修改域名缓存时间配置，实现针对客户的请求实现定制缓存功能。节点缓存分为常规缓存和带查询串URl缓存，在常规缓存里您可以设置缓存时间以及对某些影响缓存的头进行忽略处理，以及是否缓存空文件等，带查询串Url可以设置缓存成多份还是缓存成去掉问号后的url（增加命中率）。接口调用url的*可为域名名称或域名id | PUT | /api/config/cachetime/* |
| Editantihotlinkingconfig | 通过接口自助实现修改域名防盗链下IP白名单配置、referer防盗链配置，实现对具体访问IP做白名单或黑名单控制或referer控制。接口url的*可为域名名称或域名id。 | PUT | /api/config/visitcontrol/* |
| Queryantihotlinkingconfig | 通过接口自助查询防盗链配置。接口url的*可为域名名称或域名id。 | GET | /api/config/visitcontrol/* |
| Updatesourceverificationconfig | 通过接口自助实现转推回源带参数鉴权，实现rtmp转推回源，边缘能够根据加密规则，模拟客户端带上时间戳防盗链的参数进行转推回源 | PUT | /api/config/sourceverification/* |
| Querysourceverificationconfig | 通过接口自助查询转推回源带参数鉴权配置 | GET | /api/config/sourceverification/* |
| Updatedomainsrcinfo | 通过接口自助修改高级源配置。允许客户指定区域配置，不同的IP回源，以及多源时候，可以指定策略回源。接口url的*可为域名名称或域名id。 | PUT | /api/basicconfig/advancedsource/* |
| Querybanurlunderdomain | 通过接口自助查询域名URL屏蔽内容 | GET | /api/basicconfig/illegalinformation/* |
| Addbanurltodomian | 通过接口自助新增URL屏蔽的功能。客户可以根据自己的需求调用接口，新增URL屏蔽内容。 | PUT | /api/basicconfig/illegalinformation |
| Predeploysrcinfo | 通过接口自助修改高级源配置。允许客户指定区域配置，不同的IP回源，以及多源时候，可以指定策略回源。 | PUT | /api/predeploy/advancedsource/* |
| Queryapideployservice | 对于域名的新增、修改、启用、禁用、取消、恢复、删除等请求 | POST | /api/request/* |
| Updatecompressionconfig | 通过接口自助实现压缩响应功能 | PUT | /api/config/compresssetting/* |
| Querycompressionconfig | 通过接口自助查询压缩响应配置 | GET | /api/config/compresssetting/* |
| Queryhttpcodecasheconfig | 通过接口自助查询状态码缓存配置 | GET | /api/config/httpcodecache/* |
| Queryipv6config | 通过接口自助查询域名是否使用ipv6资源 | GET | /api/domain/ipv6/* |
| Querysrcinfo | 通过接口查询客户高级源配置 | GET | /api/basicconfig/advancedsource/* |
| Queryinnerredirect | 查看内部重定向配置 | GET | /api/config/InnerRedirect/* |
| Querycachetime | 通过接口自助查询节点缓存配置配置。接口调用url的*可为域名名称或域名id。 | GET | /api/config/cachetime/* |
| Editquerystringurlconfig | 通过接口自助实现查询串设置功能 | PUT | /api/config/querystring/* |
| Queryquerystringurlconfig | 通过接口自助实现修改域名的查询串设置，实现针对客户的请求实现定制缓存功能。带查询串Url可以设置缓存成多份还是缓存成去掉问号后的url（增加命中率），可以设置是否用原始请求回源等。接口调用url的*可为域名名称或域名id。 | GET | /api/config/querystring/* |
| Edithttpcodecache | 通过接口自助实现状态码缓存设置功能 | PUT | /api/config/httpcodecache/* |
| Queryamazons3authorizationconfig | 通过接口自助查询Amazon S3鉴权配置。接口url的*可为域名名称或域名id。 | GET | /api/config/amazons3access/* |
| Updateamazons3authorizationconfig | 通过接口自助实现Amazon S3鉴权配置功能。接口url的*可为域名名称或域名id。 | PUT | /api/config/amazons3access/* |
| Editdomainproperty | 修改域名属性，如回源IP、回源host、回源端口。 | POST | /api/domain/property/* |
| Queryignoreprotocol | 通过接口自助查询忽略协议缓存和推送配置 | GET | /api/config/ignoreprotocol/* |
| Editignoreprotocol | 通过接口自助修改忽略协议缓存和推送配置 | PUT | /api/config/ignoreprotocol/* |
| Editoriginuriandhost | 修改回源uri和host改写 | PUT | /api/config/originrulesrewrites/* |
| Queryoriginuriandhost | 查询回源uri和host改写 | GET | /api/config/originrulesrewrites/* |
| Editwebsocketconfig | 通过接口自助修改websocket开关配置。接口url的*可为域名名称或域名id。 | PUT | /api/config/websocket/* |
| Querywebsocketconfig | 通过接口自助查询websocket开关配置。接口url的*可为域名名称或域名id。 | GET | /api/config/websocket/* |
| Deletebanurls | 删除非法屏蔽url接口 | DELETE | /api/basicconfig/illegalinformation |
| Querylivestreamingantihotlinkingconfig | 通过接口自助查询流媒体普通防盗链配置。接口url的*可为域名名称或域名id。 | GET | /api/config/live-visitcontrol/* |
| Updatelivevisitcontrolconfig | 通过接口自助修改流媒体普通防盗链配置。接口url的*可为域名名称或域名id。 | PUT | /api/config/live-visitcontrol/* |
| Querylivestreamingtimestampantihotlinkingconfig | 通过接口自助查询流媒体时间戳防盗链配置。接口url的*可为域名名称或域名id。 | GET | /api/config/live-timestampvisitcontrol/* |
| Editlivestreamingtimestampantihotlinkingconfig | 通过接口自助修改流媒体防盗链配置。接口url的*可为域名名称或域名id。 | PUT | /api/config/live-timestampvisitcontrol/* |
| Getbasicconfigurationofdomain | 查看指定加速域名的基础配置。 | GET | /api/domain/* |
| Queryhttp2settingsconfigforwplus | 查询Http2.0配置接口。接口url的*可为域名名称或域名id。 | GET | /api/config/http2/* |
| Updatehttp2settingsconfigforwplus | 修改http2.0开关配置接口。接口url的*可为域名名称或域名id。 | PUT | /api/config/http2/* |
| Enableordisablewafprotection | 开启或关闭域名的WAF服务 | PUT | /api/domain/wafsecurity |
| Enableordisablebotprotection | 开启或关闭域名Bot防护 | PUT | /api/domain/botsecurity |
| UpdateCacheKeyConfiguration | 修改缓存key配置。接口url的*可为域名名称或域名id。 | PUT | /api/config/cachekey/* |
| QueryCacheKeyConfiguration | 查询缓存key配置。接口url的*可为域名名称或域名id。 | GET | /api/config/cachekey/* |
| Editaccessspeedlimit | 修改访问限速配置 | POST | /api/config/accessspeed/* |
| Enableordisabledmsprotection | 开启或关闭域名Dms防护 | PUT | /api/domain/dmssecurity |
| Editback2originprotocolrewriteconfig | 支持修改回源协议和端口	 | PUT | /api/config/back2originrewrite/* |
| Queryback2originprotocolrewriteconfig | 支持查询回源协议和端口	 | GET | /api/config/back2originrewrite/* |
| Updatestreamnotificationconfig | 修改流状态反馈配置。接口url的*可为域名名称或域名id。 | PUT | /api/config/streamnotification/* |
| Querystreamnotificationconfig | 查询流状态反馈配置。接口url的*可为域名名称或域名id。 | GET | /api/config/streamnotification/* |
| Querydomainbillingarea | 查询域名的计费区域 | GET | /api/domain/billingarea/* |
| Updatedomaincertconfig | 为CDN加速域名配置证书，仅支持SNI。 | PUT | /api/config/certificate/* |
| Querydomaincertconfig | 为CDN加速域名查询配置证书，仅支持SNI。 | GET | /api/config/certificate/* |
| Updatedomainmulticertconfig | 为CDN加速域名配置证书，支持SNI和多SNI。 | PUT | /api/config/certificate/v2/* |
| Updateipversionconfig | 域名默认使用V4资源，可以设置域名V6或V4+V6。 | PUT | /api/config/ipversion/* |
| Antihijackipadd | 该接口用于新增缓解IP，用户可以通过调用该接口来新增劫持缓解IP。 | POST | /api/v1/dms/antiHijackIP/add |
| Antihijackipupdate | 修改劫持缓解IP，客户调用该接口可以修改缓解IP关联的业务域名信息。 | POST | /api/v1/dms/antiHijackIP/update |
| Antihijackipdelete | 删除劫持缓解IP，客户调用该接口可以删除对应的缓解IP。 | POST | /api/v1/dms/antiHijackIP/delete |
| Batchupdatedomaincertconfig | 该接口用于修改多个CDN加速域名的配置证书。用户可以通过提供证书ID与待修改的域名名称列表修改域名关联的证书。 | PUT | /api/config/certificate/batch |
| Batchupdateapidomainforwplus | 该接口用于批量修改指定加速域名的基础配置。用户需在请求体中传入一个domainConfigs列表，每个元素包含待修改域名的`domain-name`，以及可选的其他域名配置。 | PUT | /api/batch/domain |