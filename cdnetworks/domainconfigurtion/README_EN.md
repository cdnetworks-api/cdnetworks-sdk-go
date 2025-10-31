# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/domainconfigurtion
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/domainconfigurtion"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
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

## Error Handling

Always check for errors returned by the API calls:

```go
_, err := auth.Invoke(config, request, response)
if err != nil {
    log.Printf("error: %s\n", err)
    // Handle the error appropriately
    return
}
```

## API List
For detailed API documentation and available methods, please refer to the [official Cdnetworks API documentation](https://docs.cdnetworks.com/en/cdn/apidocs).

| ActionName | enDescription | client_methods | uri |
| --- | --- | --- | --- |
| Editdomainconfig | To modify configuration of the specified domain; both domain name and domain name ID are supported. | PUT | /api/domain/* |
| Updatetimecontrolservice | Modify Time anti-theft chain custom configuration item content | PUT | /api/config/timecontrol/* |
| Querytimecontrolservice | Query Time anti-theft chain custom configuration item content | GET | /api/config/timecontrol/* |
| Editdomainredirectconfig | Self-service through interface, url redirection function | PUT | /api/config/InnerRedirect/* |
| Edithttpheaderconfig | Through the interface self-implementation http header additions and deletions to the function, can be achieved in the cdn layer of personalized http header control, so that customers do not need to modify the source stationin this case, a custom http header and speedup are implemented. The interface * can be domain name or domain id. | PUT | /api/config/headermodify/* |
| Queryhttpheaderconfig | Query http header configuration via interface self - service. The interface * can be domain name or domain id. | GET | /api/config/headermodify/* |
| Editcachetimeconfig | The interface self-help implementation modifies the domain name cache time configuration, realizes the custom cache function according to the customer's request. Node cache is divided into regular cache and query string URl cache, where you can set the ca | PUT | /api/config/cachetime/* |
| Editantihotlinkingconfig | Modify the IP whitelist configuration under the domain name anti-theft chain through the interface self-service, and implement whitelist or blacklist control on the specific access IP. The interface * can be domain name or domain id. | PUT | /api/config/visitcontrol/* |
| Queryantihotlinkingconfig | Self-checking IP black and white list anti-theft chain configuration through interface. The interface * can be domain name or domain id. | GET | /api/config/visitcontrol/* |
| Updatesourceverificationconfig | Retweet back to the source with parameter authentication through the interface self-service, realize rtmp retweet back to the source, the edge can simulate the client with timestamp anti-leech parameters according to the encryption rules to retweet back to the source. | PUT | /api/config/sourceverification/* |
| Querysourceverificationconfig | Query retweet back to the source with parameter authentication configuration. | GET | /api/config/sourceverification/* |
| Addbanurltodomian | xxx | PUT | /api/basicconfig/illegalinformation |
| Queryapideployservice | Requests for new, modified, enabled, disabled, cancelled, restored, deleted, etc. for domain names | POST | /api/request/* |
| Updatecompressionconfig | Modify compress setting configurations. | PUT | /api/config/compresssetting/* |
| Querycompressionconfig | Get compress setting configurations. | GET | /api/config/compresssetting/* |
| Queryhttpcodecasheconfig | Get http code cache configurations. | GET | /api/config/httpcodecache/* |
| Queryipv6config | Get whether a domain uses ipv6 resources. | GET | /api/domain/ipv6/* |
| Queryinnerredirect | View internal redirect configuration | GET | /api/config/InnerRedirect/* |
| Querycachetime | Self-service query of node cache configuration configuration through the interface. The * of the interface calling url can be the domain name or domain id. | GET | /api/config/cachetime/* |
| Editquerystringurlconfig | Modify query string setting configurations. | PUT | /api/config/querystring/* |
| Queryquerystringurlconfig | Through the interface, the query string settings can be modified by self-service, realizing the customized cache function. With the query string URL, you can set whether to cache multiple copies or cache the URL after removing the question mark (to increase the hit rate), and you can set whether to use the original request to return to the source, etc. The * of the interface calling url can be the domain name or domain id. | GET | /api/config/querystring/* |
| Edithttpcodecache | Modify http code cache configurations. | PUT | /api/config/httpcodecache/* |
| Queryamazons3authorizationconfig | Get amazon s3 access authorization  configurations. The interface * can be domain name or domain id. | GET | /api/config/amazons3access/* |
| Updateamazons3authorizationconfig | Modify amazon s3 access authorization configurations. The interface * can be domain name or domain id. | PUT | /api/config/amazons3access/* |
| Editdomainproperty | Modify domain name properties, such as back source IP, back source host, and back source port. | POST | /api/domain/property/* |
| Queryignoreprotocol | Protocol caching and push configuration are ignored through interface self-service queries | GET | /api/config/ignoreprotocol/* |
| Editignoreprotocol | Modify ignore protocol cache and push configuration. | PUT | /api/config/ignoreprotocol/* |
| Editoriginuriandhost | Modify the source URI and host rewrite | PUT | /api/config/originrulesrewrites/* |
| Queryoriginuriandhost | Query back source URI and host rewrite | GET | /api/config/originrulesrewrites/* |
| Editwebsocketconfig | Self-modify the websocket switch configuration through the interface. The interface * can be domain name or domain id. | PUT | /api/config/websocket/* |
| Querywebsocketconfig | Self-query the websocket switch configuration via the interface. The interface * can be domain name or domain id. | GET | /api/config/websocket/* |
| Deletebanurls | the interface of delete the illegal masking | DELETE | /api/basicconfig/illegalinformation |
| Querylivestreamingantihotlinkingconfig | Self-query the live-visitcontrol configuration via the interface. The interface * can be domain name or domain id. | GET | /api/config/live-visitcontrol/* |
| Updatelivevisitcontrolconfig | Self-modify the live-visitcontrol configuration through the interface. The interface * can be domain name or domain id. | PUT | /api/config/live-visitcontrol/* |
| Querylivestreamingtimestampantihotlinkingconfig | Query streaming media timestamp visit control. The interface * can be domain name or domain id. | GET | /api/config/live-timestampvisitcontrol/* |
| Editlivestreamingtimestampantihotlinkingconfig | Update streaming media timestamp visit control. The interface * can be domain name or domain id. | PUT | /api/config/live-timestampvisitcontrol/* |
| Getbasicconfigurationofdomain | View the configuration of the specified domain. Both domain name and domain name ID are supported. | GET | /api/domain/* |
| Queryhttp2settingsconfigforwplus | Queryhttp2settingsconfigforwplus. The interface * can be domain name or domain id. | GET | /api/config/http2/* |
| Updatehttp2settingsconfigforwplus | Updatehttp2settingsconfigforwplus. The interface * can be domain name or domain id. | PUT | /api/config/http2/* |
| Enableordisablewafprotection | Enable or disable WAF protection | PUT | /api/domain/wafsecurity |
| Enableordisablebotprotection | Enable or disable Bot protection | PUT | /api/domain/botsecurity |
| UpdateCacheKeyConfiguration | UpdateCacheKeyConfiguration.The interface * can be domain name or domain id. | PUT | /api/config/cachekey/* |
| QueryCacheKeyConfiguration | QueryCacheKeyConfiguration.The interface * can be domain name or domain id. | GET | /api/config/cachekey/* |
| Editaccessspeedlimit | Editaccessspeedlimit | POST | /api/config/accessspeed/* |
| Enableordisabledmsprotection | UpdateDomainDmsStatusServiceForWplus | PUT | /api/domain/dmssecurity |
| Editback2originprotocolrewriteconfig | Update Back2Origin Protocol Rewrite Config	 | PUT | /api/config/back2originrewrite/* |
| Queryback2originprotocolrewriteconfig | Update Back2Origin Protocol Rewrite Config. | GET | /api/config/back2originrewrite/* |
| Updatestreamnotificationconfig | Update Stream Notification Config (Normal) | PUT | /api/config/streamnotification/* |
| Querystreamnotificationconfig | Query Stream Notification Config (Normal) | GET | /api/config/streamnotification/* |
| Querydomainbillingarea | Get domain's billing areas. | GET | /api/domain/billingarea/* |
| Updatedomaincertconfig | Set domain's certificate config, support SNI only. | PUT | /api/config/certificate/* |
| Querydomaincertconfig | Get domain's certificate config, support SNI only. | GET | /api/config/certificate/* |
| Updatedomainmulticertconfig | Set domain's certificate config, support SNI and mutil SNI. | PUT | /api/config/certificate/v2/* |
| Updateipversionconfig | Set IPv4 and IPv6 resource for your domain. | PUT | /api/config/ipversion/* |