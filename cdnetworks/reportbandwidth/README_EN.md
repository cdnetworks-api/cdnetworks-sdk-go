# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportbandwidth
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportbandwidth"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportbandwidth.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportbandwidth.{ActionName}Response{}
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
| Queryp2pbandwidth | Query CDN bandwidth, P2P bandwidth, total bandwidth and billing value | POST | /myview/BandWidthP2p |
| Bandwidthmiddle | Query the bandwidth of middle cache. | POST | /myview/bandwidth-middle |
| Querydomainbandwidth | This interface is used to collect statistics on bandwidth summary information of multiple domain names within a specified time. Users need to provide acceleration domain names and time parameters to generate summary data, and the returned content includes bandwidth summary data for each time slice. This interface helps users monitor and analyze the traffic usage of domain names to better optimize the network and control costs. | POST | /api/report/domainbandwidth |
| Queryipv6bandwidthofeachispandprovince | According to the visitor visit log, query bandwidth of domain in each ISP each province different IP type.<br>Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese. | POST | /api/report/bandwidth/isp-province/ipv6 |
| Getbandwidthlog | Query the log bandwidth. | POST | /myview/bandwidth-log |
| Bandwidthchannel | Query the channels' bandwidth. | POST | /myview/bandwidth-channel |
| Querybandwidthbyispprovince | This interface is used to query the bandwidth information of visitor IPs in various provinces and ISPs in mainland China. Users can specify parameters such as time period, domain name, region, and ISP to obtain bandwidth data with a granularity of 5 minutes in Mbps. The interface returns a timestamp and the corresponding bandwidth value to help users analyze the bandwidth usage of different regions and ISPs, thereby optimizing network resource allocation and improving user experience. | POST | /api/report/bandwidth/area/isp-province |
| Querybandwidthoforiginminutely | Query minute-level back-to-origin bandwidth for multiple domains. Users must provide start and end times, domain(s), and data granularity.The results include bandwidth peak, total back-to-origin traffic, and per-minute back-to-origin bandwidth for each domain. | POST | /api/report/bandwidth/multi-domain/real-time/origin |
| Querycpsbandwidth | The channels' bandwidth of  dedicated line | POST | /myview/bandwidth-HK |
| Queryrealtimebandwidthformultidomain | This interface is used to query the minute-level edge bandwidth of a domain name. Users need to provide a time range and domain name to obtain detailed bandwidth usage data. The returned content includes the total traffic of each domain name, minute-level bandwidth data, and bandwidth peak, etc. It helps users analyze the traffic of websites or applications, thereby optimizing resource management and improving performance. | POST | /api/report/bandwidth/multi-domain/real-time/edge |
| Querymultidomainsipv6oripv4bandwidth | This interface is used to query the IPV6 or IPV4 bandwidth data of multiple domain names. The user provides parameters such as time range, domain name, IP type and region to obtain bandwidth details of related domain names. The returned results include bandwidth data for each domain name in the specified time period. This interface helps users monitor and analyze network traffic of different domain names under specific IP types, thereby optimizing resource allocation and improving network performance. | POST | /api/report/bandwidth/ipv6 |
| Perzonebilling | Query the perzone data | POST | /myview/perzone-billing |
| Querybandwidthformultidomain | This interface is used to query the domain bandwidth of dynamic acceleration products. Users need to provide time parameters to obtain bandwidth data within the specified time period. The return content includes detailed data for each time slice of bandwidth. This interface helps users fully understand the bandwidth consumption of multiple domain names, thereby optimizing network resources and planning bandwidth. | POST | /api/report/domainbandwidth/dwa |
| Reportcountryserverbandwidthservice | This interface is used to query the bandwidth details of the country to which the server IP belongs. The user needs to provide the time range, domain name, country code, and data granularity. The return content includes the edge traffic and bandwidth values at each time point in the specified time period. This interface helps users understand the current service traffic distribution and usage in different countries around the world. | POST | /api/report/server/country-bandwidth |