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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
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
| Querypurgestatus | Query the purge task status. Use this API to check if the task has been completed successfully throughout the global CDN. | POST | /ccm/purge/ItemIdQuery |
| Purgefilewithtag | Push the file with the tag in the domain name with the tag switch enabled under the customer | POST | /api/content/tag/purge |
| Regexurlpurge | The content of the file cached on the CDN node is cleared according to the regular url method, and the entire network takes effect within 1 minute. | POST | /api/content/regular-url/purge |
| Purge | Refresh the files cached on the CDN nodes, which takes effect throughout the whole CDN within about 10 minutes. | POST | /ccm/purge/ItemIdReceiver |
| Querypurgelimit | query purge daily limit | POST | /ccm/purge/limit |
| Querypurgeresiduals | Query url purge , directory purge and prefetch daily limit . | GET | /ccm/upperQuery |