# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/protectassets
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/protectassets"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &protectassets.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := protectassets.{ActionName}Response{}
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
| Listapidefinitions | Query a complete list of API definition basic information. | POST | /api/v1/sam/api-define/basic/query |
| Queryapidefinitiondetail | Get API definition details. | POST | /api/v1/sam/api-define/detail |
| Listapiassetdiscovery | Query list of API discovery. | POST | /api/v1/sam/api-discovery/get-list |
| Feedbackwrongapiassetdiscovery | Feedback wrong API discovery. | POST | /api/v1/sam/api-discovery/false-marking |
| Createapidefinition | Add API definition. | POST | /api/v1/sam/api-define/add |
| Updateapidefinition | Modify API definition. | POST | /api/v1/sam/api-define/update |
| Batchdeleteapidefinition | Delete API definitions in batches. | POST | /api/v1/sam/api-define/delete |