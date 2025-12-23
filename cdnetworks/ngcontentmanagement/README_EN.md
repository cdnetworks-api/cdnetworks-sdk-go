# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngcontentmanagement
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngcontentmanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &ngcontentmanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ngcontentmanagement.{ActionName}Response{}
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
| CreateAPurgeRequest | CreateAPurgeRequest to force a refresh of the content stored in the CDN Pro cache servers.  You may wish to do this if you just updated content on your origin server and want visitors to see the changes right away rather than wait for the updates to propagate according to the schedule defined in your property configuration. | POST | /cdn/purges |
| GetListOfPurgeRequests | Get a list of purge requests. The API only returns purge requests that were created in the past 15 days. | GET | /cdn/purges |
| GetPurgeSummary | Returns a summary of purge requests that were made during a time period. Query parameters let you specify the timeframe and target environment. | GET | /cdn/purges/purgeSummary |
| GetPurgeRequestStatus | Gets information about a purge request including its associated hostnames and status. | GET | /cdn/purges/* |
| CreateAPrefetchRequest | CreateAPrefetchRequest to warm up the CDN Pro cache with contents from your origin in anticipation of visitors. This is helpful to prevent a sudden influx of requests to your origin servers. The URLs you prefetch must be for hostnames of properties deployed to production.<br> | POST | /cdn/prefetches |
| GetListOfPrefetchRequests | Get a list of prefetch requests. The API only returns prefetch requests that were created in the past 15 days. | GET | /cdn/prefetches |
| GetPrefetchRequestStatus | Get details on the status of a prefetch request. | GET | /cdn/prefetches/* |
| GetAvailablePurgeRequests | This API indicates the number of directory and file purge requests you can make in the staging or production environments. Your service quota specifies a daily limit on file and directory purges for production and staging environments. You can temporarily exceed the limit during a day, but it will reduce the number of purges available during the next day. | GET | /cdn/purges/purgeTokens |