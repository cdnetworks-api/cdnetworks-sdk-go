# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportrequest
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportrequest"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportrequest.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportrequest.{ActionName}Response{}
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
| Querydomaintotalrequest | This interface is used to query the total number of requests for multiple domain names. Users provide information such as time and domain name, and the interface will return the number of requests in each time segment, supporting data of different granularities (such as data every five minutes, every hour, or every day). This interface helps users monitor website access and follow up and optimize abnormal situations in a timely manner. | POST | /api/report/domainhit |
| Queryrequestnumbersundershieldpop | This interface is used to query the number of requests for multiple domain names on the parent node. The user provides the query time range and domain name list, and can choose whether to group the output by domain name. The returned results include the total number of requests, the peak number of requests, and the number of requests every 5 minutes. This interface is helpful for users who need to analyze the number of visits to the domain name on the parent node to adjust the cache strategy. | POST | /api/report/request/parent-node |
| Queryrequestbyspecificprotocol | This interface is used to count the number of requests for a specified transport protocol on edge nodes for multiple domain names. Users need to provide a query time range and a domain name list, and select a protocol type. The returned content includes the number of requests for each domain name with a specified protocol. This interface helps users monitor and analyze traffic under different transport protocols to optimize network performance and improve service quality. | POST | /api/report/request/protocol |
| Queryipv6requestofeachispandprovince | According to the visitor visit log, query request numberof domain in each ISP each province different IP type.<br>Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese. | POST | /api/report/request/isp-province/ipv6 |
| Queryrequesthitratio | Query minute-level request cache hit ratios for a specified period. Provide start time, end time, and domain to get data. Granularity options are 1 minute, 5 minutes, 1 hour, or 1 day. It returns cache hit ratios for each minute, helping evaluate domain cache performance on the CDN. | POST | /api/report/request/hit-ratio/total |
| Querymultidomainsipv6oripv4requests | This interface is used to query the number of IPV6/IPV4 requests for a specified domain name. Users can specify the time range, domain name, IP type, and region to filter data. The returned data includes the number of requests for each domain name at each time point in the specified time period. This interface helps users analyze the composition and trend of network traffic to better understand the distribution of requests in different regions and IP types. | POST | /api/report/request/ipv6 |
| Reportrequestquicservice | This interface is used to query the number of QUIC requests and the total number of requests for multiple domain names. Users need to provide the start time, end time, and domain name list to obtain data. The data returned by the interface includes the number of QUIC requests and the total number of requests for each timestamp. It helps users analyze the usage of QUIC and adjust strategies to improve the response speed and user experience of the website. | POST | /api/report/request/quic |
| Reportuserrequestcountryservice | Query the number of requests for multiple domain names in each country and region. (Statistics by visitor IP) | POST | /api/report/request/country |
| Queryedgerequesthitratioservice | Query the hit rate of cache requests at the edge node for multiple domains at the minute level | POST | /api/report/request/edge-hit-ratio/total |