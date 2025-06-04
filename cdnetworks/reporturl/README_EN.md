# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reporturl
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reporturl"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reporturl.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reporturl.{ActionName}Response{}
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
| Querytoprankingurl | This interface is used to count the number of URL requests and traffic of a domain name within each hour, and generate a ranking list of up to TOP500. Users need to provide the time return and domain name to obtain the total number of URL requests, total traffic, number of hit requests, number of failed requests, and details of different status codes. The interface supports sorting by total number of requests or total traffic to quickly identify high-frequency access or high-traffic URLs. This interface helps users monitor URL access frequency and traffic to adjust corresponding network policies. | POST | /api/report/url/top |
| Queryfailedrequestsurlnumbersrankhourlygranularity | Query URL rankings for multiple domains within a specified time range, sorted by failed requests (up to TOP5000 URLs). Users provide start and end times and a domain list. It returns URLs ranked by failed requests, along with the number of failed and successful requests and total successful traffic, ideal for identifying URLs with high failure rates. | POST | /api/report/url/fail/rank |
| Querysuccessfulrequestsurlnumbersrankhourlygranularity | Query the Top URLs ranked according to the number of successful requests at the specified time, as well as the total successful traffic and the number of successful requests of the TOP URLs. It supports querying top500 URLs at most | POST | /api/report/url/success/rank |
| Reportdomainrefererwebsiteservice | Query the website referer ranking of multiple domains | POST | /api/report/domain/referer-website |