# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportvisitor
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportvisitor"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportvisitor.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportvisitor.{ActionName}Response{}
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
| Querytotalnumberofuniqueipundersingledomain | This interface is used to query the number of independent IP addresses for each stream name within a single domain name. The user provides a specific domain name and time to obtain information about the specified stream during this period (in days). The results returned by the interface include statistics on independent IP addresses for the domain name and its corresponding stream name. This helps users understand the distribution of visitors to each stream, thereby optimizing traffic configuration or evaluating the use of streaming media services. | POST | /api/report/visitor/total/stream |
| Statisticalanalysisofchinamainlandprovincevisitors | This interface is used to analyze user access in various provinces in mainland China. Users provide a time range and domain name to obtain detailed statistics. The interface returns information such as the number of requests, successful and failed requests, success rate, total traffic, number of unique visitors, and page views for each region. These data can help website administrators analyze traffic sources, improve regional marketing strategies, and improve user experience. | POST | /api/report/visit/analysis/combine/province |
| Statisticalanalysisofcountries | This interface is used to analyze the access statistics of different countries and regions. Users can query by providing parameters such as start and end time, domain name, region code, etc. The response results include detailed statistical information such as the number of requests, number of successful requests, number of failed requests, success rate, total traffic, number of unique visitors, page views, and number of IP addresses in each region. This interface is suitable for users who need to understand the traffic and user behavior of websites in different regions around the world, which can help companies optimize their market strategies and technical resource allocation. | POST | /api/report/visit/analysis/combine/country |
| Reportvisitorcustomtopdailyservice | Reportvisitorcustomtopdailyservice | POST | /api/report/visitor/custom-top/daily |