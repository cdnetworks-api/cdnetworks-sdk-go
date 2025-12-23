# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cacherefresh
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cacherefresh"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &cacherefresh.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := cacherefresh.{ActionName}Response{}
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
| Getpurgestatus | Query the purge task status. Use this API to check if the task has been completed successfully throughout the global CDN. | POST | /ccm/purge/ItemIdQuery |
| Purgefilewithtag | Purge cache by tags. The tag purge feature should be enabled first. | POST | /api/content/tag/purge |
| Purgeurlsmatchingregex | Clear the cached file content on CDN nodes according to the regular URL method. | POST | /api/content/regular-url/purge |
| Purge | Refresh the cached file content on CDN nodes. It generally takes effect within 1 - 3 minutes across the network. The daily limit for directory pushing per account and per domain name is 500. | POST | /ccm/purge/ItemIdReceiver |
| Querypurgelimit | query purge daily limit | POST | /ccm/purge/limit |
| Querypurgeresiduals | Query url purge , directory purge and prefetch daily limit . | GET | /ccm/upperQuery |
| Getpurgequota | This interface is used to query the daily push limit and daily remaining amount for a specified customer. Users can obtain their daily push limit and remaining amount through this interface, including URL push, directory push, regular expression push, and tag push. This is of great significance for determining whether the customer can continue to submit tasks. It can help customers promptly find out if they have reached the daily limit, and when customers have a large number of tasks to submit, they can adjust the daily limit in advance.<br> | POST | /ccm/purge/quota |
| Getprefetchquota | This interface is used to query the daily upper limit and remaining daily quantity of prefetching for a specified customer. Through this interface, users can obtain their daily prefetching upper limit and remaining quantity, including the number of URLs and prefetching size. This is crucial for whether the customer can continue to submit tasks. It helps customers promptly identify whether the daily upper limit has been reached. When customers have a large number of tasks to submit, they can adjust the daily upper limit in advance. | POST | /ccm/prefetch/quota |