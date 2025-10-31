# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/webhooks
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/webhooks"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &webhooks.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := webhooks.{ActionName}Response{}
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
| GetAListOfWebhooks | GetAListOfWebhooks that have been defined. You can reference a webhook when creating a deployment task, creating a purge request, creating a prefetch request, or creating a validation task. | GET | /cdn/webhooks |
| GetAWebhook | This API returns details about a webhook including its URL and metadata about its use. | GET | /cdn/webhooks/* |
| DeleteAWebhook | DeleteAWebhook. If an asychronous task using it is still in progress, the webhook may not be called when the task completes. | DELETE | /cdn/webhooks/* |
| UpdateAWebhook | UpdateAWebhook. Only the fields in the request body will be updated. | PATCH | /cdn/webhooks/* |
| CreateAWebhook | CreateAWebhook. It can be used to notify you when an asynchronous task completes. You can reference a webhook when creating a deployment task, creating a purge request, creating a prefetch request, or creating a validation task. | POST | /cdn/webhooks |