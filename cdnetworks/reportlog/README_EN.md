# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportlog
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportlog"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportlog.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportlog.{ActionName}Response{}
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
| Getbotattackincidentlogdata | Get bot attack event log data | POST | /api/bot/report/event-log |
| Getoverallrequests | Get Overall Requests. When querying all, the report will include the data of the deleted zone | GET | /api/clouddns/reports/requests/overall |
| Getrequestsbyresponsetype | Get Requests By Response Type. When querying all, the report will include the data of the deleted zone. | GET | /api/clouddns/reports/requests/response_type |
| Getrequestsbyprocesstime | Get Requests By Process Time. When querying all, the report will include the data of the deleted zone. | GET | /api/clouddns/reports/requests/process_time |
| Getrequestsbylocationdistribution | Get Requests By Location. When querying all, the report will include the data of the deleted zone. | GET | /api/clouddns/reports/requests/location |
| Getrequestsbyrecordtype | Get Requests By Record Type. When querying all, the report will include the data of the deleted zone. | GET | /api/clouddns/reports/requests/record_type |