# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/edgefunc
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/edgefunc"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &edgefunc.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := edgefunc.{ActionName}Response{}
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
| Uploadfunccode | Function code upload interface, for the scenario that does not use CloudIDE, but directly calls the interface to upload function code | POST | /edgefunc/upload |
| Createedgefunctrigger | This interface is used to create function triggers. Users are required to provide the domain and a list of routing rules as request parameters. Upon successful creation, the interface will return a status code and the list of successfully created triggers. | POST | /api/v2/cdn_triggers |
| Queryedgefunctrigger | This interface is used to query a list of function triggers. Users can filter by domain or function name and use pagination. The response will return a detailed list of triggers. | GET | /api/v2/cdn_triggers |
| Deletefuncdomaintrigger | This interface is used to remove an existing function trigger. The user needs to specify the unique identifier of the function trigger to be deleted via the REST parameter 'id'. | DELETE | /api/v2/cdn_triggers/* |