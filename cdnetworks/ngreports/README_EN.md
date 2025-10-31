# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngreports
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngreports"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &ngreports.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ngreports.{ActionName}Response{}
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
| GetEdgeTrafficVolume | Get the edge traffic volume in megabytes during a time period. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of less than two minutes.<br> | POST | /cdn/report/vol |
| GetEdgeRequests | Gets the number of requests to CDN Pro edge servers during a time period. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of less than two minutes. | POST | /cdn/report/req |
| GetEdgeBandwidth | Get the bandwidth in Mbps to CDN Pro edge servers during a time period. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of less than two minutes.<br> | POST | /cdn/report/bandwidth |
| GetOriginFastRouteTrafficVolume | Return the traffic to your origin server that uses our <a href="/cdn/docs/edge-logic/supported-directives#origin_fast_route">fast route feature</a> to make the connections faster and more reliable. | POST | /cdn/report/fastOriginVol |
| GetASummaryOfRequests | GetASummaryOfRequests during a time period. You can filter and group results by hostnames or server groups. If you are a reseller, you can also filter and group by customerIds to distinguish your child customers' requests. | POST | /cdn/report/reqSummary |
| GetASummaryOfTrafficBandwidth | GetASummaryOfTrafficBandwidth during a time period. You can filter and group results by hostnames or server groups. If you are a reseller, you can also filter and group by customerIds to distinguish your child customers' traffic. | POST | /cdn/report/bandwidthSummary |
| GetOriginTrafficVolume | Gets the origin traffic volume in megabytes during a time period. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of less than two minutes.<br><br>Note: The overhead of TCP, IP, and MAC headers is not included in this API's results. | POST | /cdn/report/volOrigin |
| GetEdgeStatusCodeDetails | Gets details about HTTP status codes returned by requests to CDN Pro edge servers during a time period. Query parameters let you specify the time period, scheme, and granularity so you can monitor changes during the period. Filter the results further by passing a filters object in the request. This report's data has a latency of up to thirty minutes. | POST | /cdn/report/statusCodeDetails |
| GetTheNumberOfRequestsToOrigin | Get the number of requests to your origin servers during a time period. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of less than two minutes. | POST | /cdn/report/reqOrigin |
| GetOriginStatusCodeDetails | Gets details about HTTP status codes returned by requests to your origin servers during a time period. Query parameters let you specify the time period, scheme, and granularity so you can monitor changes during the period. Filter the results further by passing a filters object in the request. This report's data has a latency of up to thirty minutes. | POST | /cdn/report/statusCodeDetailsOrigin |
| GetTheIntermediateTrafficVolume | GetTheIntermediateTrafficVolume in megabytes during a time period. Intermediate traffic is traffic between CDN Pro servers. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of less than two minutes.<br> | POST | /cdn/report/volInterm |
| GetTheEdgeUploadTrafficVolume | Gets the traffic volume in megabytes uploaded to edge servers during a time period. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of less than two minutes.<br><br>Note: The overhead of TCP, IP, and MAC headers is not included in this API's results. | POST | /cdn/report/upVol |
| GetASummaryOfStatusCodesReturnedByEdgeServers | Gets a summary of HTTP status codes returned by requests to the CDN Pro edge servers during a time period. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of up to thirty minutes. | POST | /cdn/report/edgeStatusSummary |
| GetASummaryOfStatusCodesReturnedByOriginServers | Gets a summary of HTTP status codes returned by requests to your origin servers during a time period. Query parameters let you specify the time period and scheme. Filter the results further by passing a filters object in the request. This report's data has a latency of up to thirty minutes. | POST | /cdn/report/originStatusSummary |
| GetASummaryOfCpuUsage | Obtain a summary of CPU usage by edge and intermediate servers. | POST | /cdn/report/cpuSummary |
| GetOriginFastRouteRequests | Return the number of requests to your origin server that uses our <a href="/cdn/docs/edge-logic/supported-directives#origin_fast_route">fast route feature</a> to make the connections faster and more reliable. | POST | /cdn/report/fastOriginReq |
| GetEdgeHostnameSummaryStatistics | Obtain a summary of DNS requests for one or more edge hostnames during a time period. This API allows you to see the actual number of requests for each edge hostname. | POST | /cdn/report/edgeHostnameReqSummary |
| GetCpuTimeUsed | Get the amount of CPU time in seconds used to handle requests for your content. | POST | /cdn/report/cpuTime |
| GetASummaryOfTrafficVolume | Get a summary of edge, intermediate (between CDN Pro servers), and origin traffic volume during a time period. You can filter and group results by hostnames or server groups. If you are a reseller, you can also filter and group by customerIds to distinguish your child customers' traffic.<br> | POST | /cdn/report/volSummary |
| Getedgehostnamestatistics | Get the number of DNS requests to resolve your edge hostnames during a time period. | POST | /cdn/report/edgeHostnameReq |