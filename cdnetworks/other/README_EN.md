# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/other
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/other"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &other.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := other.{ActionName}Response{}
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
| Querydailylivetranscodingduration | Query the live transcoding length at 1 day granularity. | POST | /myview/LiveTranscodingPerDayV2 |
| Query5minlivetranscodingduration | Query the live transcoding length of 5 minutes granularity. | POST | /myview/LiveTranscodingV2 |
| Queryliverecordingduration | Query the live recording length of minute granularity. | POST | /myview/RecordingTime |
| Queryddosmitigatedbandwidthbysuiteorproduct | DDoS mitigatedBandwidth for query | POST | /soc/api/report/QueryMitigatedBandwidth |
| Picprocessstatistics | Output the image processing times by domain at 5-minute granularity. | POST | /myview/PicFlowHit |
| Httpdnsstatistics | Output the customers's resolution times at 5-minute granularity | POST | /myview/HttpdnsStatistics |
| Vodtranscodingapi | Query 5-minute granularity statistics for VoD transcoding across various types and profiles. | POST | /myview/wct-ov |
| Querypopihmszone | Query POP and ihms zone info. | POST | /myview/pop-ihms-zone |
| Querylivestreamstatus | The current interface non-real-time data content includes: stream status information, Bitrate information, frame rate information, number of online users, header information, and Domain bandwidth; the real-time data content includes: stream status information, Bitrate information, frame rate information, number of online users (excluding HLS), and header information. | POST | /api/quality/stream-status-statistic |
| Getlivestreampushingstatus | Get the frame rate, bit rate, frame loss rate and other information of all channels under the specified domain name (supports multiple domain names and stream names) at the specified time. | POST | /api/quality/frame-rate |
| Queryonlineviewercount | This interface provides access to the online count of live streaming domain names or channels. It supports only the GET request method. | POST | /QOSS/api/onlineViewers |
| Queryedgeactiveconnectionsformultidomains | Query the number of active edge connections of multiple domain. By providing the start time, end time and domain list, you can query the number of active connections of multiple domain in every 5 minutes during the time period. It is suitable for scenarios where you need to analyze the trend of active website connections. | POST | /api/report/edge-active-conn |
| Reportdomainstreamdurationservice | Query the push duration of multi-domain names and multi-stream names | POST | /api/report/domain/stream/duration |
| Querydirectoryrankbysuccessrequesthourly | This interface is used to query the top 500 directories with the most successful requests in a specified time. Status codes 2xx and 3xx are defined as success. Users need to provide a time range and domain name to obtain a directory list. The interface returns the number of successful requests and total successful traffic data corresponding to the directory. This interface helps users analyze website performance, traffic trends, and optimize directory structure. | POST | /api/report/directory/success/rank |
| Querydirectoryrankbyfailedrequesthourly | This interface is used to query the directory information ranked by the number of failed requests per hour under the domain name. Users can obtain a list of directories with the highest number of failed requests according to the specified time range. The returned information includes the number of failed requests, the number of successful requests, and the total traffic successfully transmitted corresponding to the directory. Users can use this interface to help analyze possible problems in the website directory and optimize the use of network resources. | POST | /api/report/directory/fail/rank |
| Querydirectoryrankbytotaltraffichourly | This interface is used to query the directory list of multiple domain names sorted by total traffic, and return at most the top 500. Users need to provide a time range and domain name to obtain a directory list, and the returned content includes the total traffic and number of requests for each directory. This interface helps users analyze and optimize website traffic, helping users identify high-traffic directories and allocate resources. | POST | /api/report/directory/hourly/total-traffic/rank |
| Antihijackiplist | Retrieve the anti-hijack IP list. Customers can use this API to obtain information related to anti-hijack IPs, including region, deployment status, and other details. | POST | /api/v1/dms/antiHijackIP/list |
| Resourcegrouplist | This interface is used to query the list of all resource groups in the account granularity. You can query a resource group by its name. The value includes the resource group name, number of domains Help customers allocate resources properly. | POST | /api/v1/dms/resourceGroup/list |
| Querydomainsforresourcegroup | This interface is used to query the domains associated with a resource group. Users can obtain the complete information about the associated domains from the resource group name. The domain and domain status are included. Users can manage the relationship between domains and resource groups as required to prevent mutual impact on core service domains when they are attacked. | POST | /api/v1/dms/resourceGroup/queryDomains |
| Resourcegroupdomainchange | This interface is used to change the resource group used by domain. The user updates the specified domain from the current resource group to the target resource group. Help clients allocate resources for domain. | POST | /api/v1/dms/resourceGroup/changeGroup |
| Streamtrafficservice | This interface is used to query stream name bandwidth. Users can query data by specifying time, domain, region, publishing point, protocol, etc.<br>Data is returned by aggregating traffic dimensions based on region, domain, stream name, publishing point, protocol, etc. | POST | /api/gdp/stream-traffic-service |
| Querydomainresourcegroup | This API is used to query the resource group to which a domain belongs. Users can query related resource group information through the domain, including resource group name and resource group ID. | POST | /api/v1/dms/domainResource |