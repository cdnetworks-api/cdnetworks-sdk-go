# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportflow
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportflow"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportflow.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportflow.{ActionName}Response{}
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
| Querydirectorybandwidthtrafficunderlivestreamdomain | Query traffic or bandwidth statistics for each directory under the live push and pull streaming domains. Users need to provide the start and end time, domain name, and can specify a directory for querying. By default, the data is aggregated every 5 minutes. The returned results include the total traffic or bandwidth peak for each directory, along with detailed data for each time segment. | POST | /api/report/flow/dir/detail |
| Querystreamtrafficundermalbdomain | This interface is used to query the traffic or bandwidth data of live streaming under a domain name within a specified time period. The user needs to provide the time range, domain name and stream name, as well as data type and granularity options. The interface returns the traffic or bandwidth details of each stream under each domain name, including total traffic and peak bandwidth. It helps users monitor live streaming traffic to optimize resource allocation. | POST | /api/report/flow/stream/detail |
| Querystaticdynamictrafficbyprotocol | This interface is used to query the dynamic and static traffic details of different protocols under each domain name. Users need to provide a time range and a domain name list, and can query data for up to the past six months. The returned results include traffic data for each domain name, which is displayed hierarchically by specified protocol and node. This interface is suitable for users who need to analyze domain name traffic patterns and optimize network resource configuration, and supports detailed traffic management and performance evaluation. | POST | /api/report/flow/static-dynamic |
| Queryoutputtrafficundershieldpop | Query the downstream traffic data of the Relay node of multiple domain . Users can obtain detailed Traffic data reports by providing the start time, end time, and domain list. The return value includes the total Traffic of each domain and provides the Traffic value of each time slice. This interface is suitable for monitoring and analyzing the Traffic profile and change trend of CDN Relay node. It's recommended that the call frequency is no higher than 30/5min. | POST | /api/report/flow/parent-node |
| Queryspecificpoptraffic | Query the Traffic details of each POP of multiple domain in a short period of time. Users can obtain the Traffic data of different POP nodes under each domain by specifying the start time and end time, including static and dynamic inflow and outflow Traffic. The data returned by the interface is displayed in 5-minute granularity. | POST | /api/report/flow/pop |
| Querymultiplestreamtrafficandbandwidthunderthedomain | This interface is used to query the traffic and bandwidth information of live streams under multiple domain names within a specified time period. Users can enter the start time, end time, and domain name to obtain detailed data, including the total traffic of each stream under each domain name, peak bandwidth, traffic or bandwidth data at each time point, and (optional) the number of online users at a time point. This is suitable for monitoring live stream network performance and user interaction, optimizing network configuration and resource allocation through data analysis, and improving the quality of live broadcast services. | POST | /api/report/flow-bandwidth/stream/detail |
| Wctquery | Query the 5-minute granularity statistics of various types and gradients of customer on-demand transcoding | POST | /myview/Wctquery |
| Querydatatransferforalldomainbyte | ReportFlowAllForByteService | POST | /api/report/flow/byte |
| QuerytotaltrafficinbyteforalldomainsunderdwaByte | ReportFlowAllForByteForDWAService | POST | /api/report/flow/byte/dwa |
| Querytotaltrafficforalldomainsunderdwa | Query the Traffic data of all domain under dynamic acceleration. By Invoke this interface, users can obtain the Traffic data of all domain in a certain time period, including the specific Traffic of each time slice. The data granularity can be specified as 5 minutes, 1 hour or 1 day. | GET | /api/report/flow/dwa |
| Querytrafficforselectedbillingarea | Query traffic statistics for dynamic acceleration domains in a specified settlement area. Users provide start and end times, domain list, settlement area, and data granularity. Results include total traffic and detailed traffic data for each time slice, ideal for analyzing traffic by settlement area. | POST | /api/report/flow/area/dwa |
| Querytrafficformultidomain | This interface is used to query the traffic summary of multiple domain names in the dynamic acceleration service. Users need to provide the acceleration domain name and time range to obtain traffic statistics. The returned data includes detailed data on the traffic of each time slice and the total traffic. It helps users analyze the traffic usage of each domain name to better manage and monitor their network resources. | POST | /api/report/domainflow/dwa |
| Queryrequestnumberandpeakbandwidthformultidomain | This interface is used to query the traffic, number of requests and peak bandwidth corresponding to multiple domain names under the dynamic acceleration product. Users provide the query time range and domain name, and the interface returns the total traffic, total number of requests, bandwidth peak and detailed data for each time slice of the domain name. It helps users analyze the traffic pattern of the website and identify the peak period, thereby optimizing bandwidth resource allocation and improving site performance. | POST | /api/report/flow-request/dwa |
| QueryKeywordaquanEdgereport | Query the traffic with a specific protocol for multiple domains, and the traffic statistics contain data of all edge nodes. | POST | /myview/CdnwEdgeReportForKeywordAquan |
| Queryedgebytehitratioservice | Query the hit rate of cache traffic at the edge node for multiple domains at the minute level | POST | /api/report/flow/edge-hit-ratio/total |
| Gettotaltrafficforalldomains | This API is used to query the traffic summary information for all domains under a user's account. Users provide start and end time parameters to obtain detailed traffic data. Through this API, users can understand and analyze the CDN traffic usage of all domains within the specified time period. | GET | /api/report/traffic |
| Gettrafficandbandwidthbyclientcountry | Query the traffic bandwidth distribution of a domain across countries and regions. Users input a time range and domain list to view data by domain, country, or domestic vs. foreign. Results include total regional traffic, its percentage, and traffic and bandwidth per time segment, aiding in global website traffic analysis and management. Query data from 24 hours prior is recommended. | POST | /api/report/traffic/domain-country |
| Gettrafficandrequestshitratiobyispprovince1mingranularity | This API is used to query the minute-level traffic hit rate and request hit rate for specific provinces and ISPs. Users can obtain traffic or request hit rate data per minute by providing information such as start time, end time, domain name, province, and operator. This API is applicable for scenarios where network performance needs to be analyzed for different provinces and operators within a specified time period. Users can monitor the hit conditions of traffic and requests and optimize their network resource configuration through this API. | POST | /api/report/traffic/isp-province/hit-rate/total |
| Gettrafficbyprotocol | This API is used to query the Traffic data of multiple domain under the specified transmission protocol. The query is for the data of all Edge node. The user needs to specify the domain, start and end time, transmission protocol and other parameters, and can select the data granularity and grouping dimension. By default, the query will return the Traffic data of the HTTPS protocol, and the output result is the Traffic value of each domain at different time points. This interface is suitable for scenarios where you need to query and analyze the Traffic of multiple domain with different transmission protocols. | POST | /api/report/traffic/protocol |
| Gettrafficrequestsandpeakbandwidthformultidomains | This interface is used to query the traffic request count and peak bandwidth of multiple domain names. Users need to provide a time range and a domain name list to obtain the total traffic, total request count, peak bandwidth and its occurrence time of each domain name, as well as detailed traffic and request count for each time period. This helps users understand the access status of the website and effectively adjust strategies to meet different traffic requirements. | POST | /api/report/traffic-request |
| Querytotaltrafficformultidomains | This interface is used to obtain traffic summary information of multiple domain names. Users need to provide a time range and domain name, and the data granularity can be selected as five minutes, one hour, or one day. The returned content includes the total traffic and the traffic value of each time node. This interface helps users monitor and analyze the traffic of multiple domain names, so as to grasp the overall traffic trend and changes. | POST | /api/report/domaintraffic |
| Getorigintrafficandrequestsformultidomains | This interface is designed to query the back-to-source traffic and number of requests for multiple domain names. Users can obtain this data by entering information such as the start time, end time, and specified domain name. The interface returns the total traffic, total number of requests, request and bandwidth peaks for each domain name and their corresponding time points, and supports providing detailed traffic, request, and bandwidth data at a minute granularity. This feature helps users monitor and analyze website back-to-source traffic. | POST | /api/report/traffic-request/origin |
| Getbandwidthsavingratio | This API is used to query the amount of bandwidth savings over a specified time period. Users need to provide the start time, end time, and domain information to retrieve data. The returned data includes the average total bandwidth savings and the actual bandwidth savings per time slice. This API is particularly suitable for scenarios requiring monitoring and analysis of bandwidth savings resulting from CDN cache hits. | POST | /api/report/bandwidth/saving-bandwidth/total |