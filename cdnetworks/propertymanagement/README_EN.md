# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/propertymanagement
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/propertymanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &propertymanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := propertymanagement.{ActionName}Response{}
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
| GetAPropertyVersion | Get the detailed configuration of a property version. | GET | /cdn/properties/*/versions/* |
| UpdateAPropertyVersion | This API lets you update a property version. Please note that a property that has already been deployed to production or staging is 'frozen' and cannot be updated or validated again. The only exception is that the version description can be updated. | PATCH | /cdn/properties/*/versions/* |
| DeleteAPropertyVersion | DeleteAPropertyVersion. | DELETE | /cdn/properties/*/versions/* |
| GetAProperty | Returns a summary about a property including the number of versions that have been created and which versions are deployed. | GET | /cdn/properties/* |
| UpdateAProperty | This endpoint changes the name and description of a property. Its versions are unaffected. <br><br>If you wish to change a property's configuration, use the <a href="#operation/updatePropertyVersion">UpdateAProperty version API</a>. | PATCH | /cdn/properties/* |
| DeleteAProperty | Delete a property by its ID. | DELETE | /cdn/properties/* |
| GetListOfPropertyVersions | Get a list of all property versions. A summary of each version including its status and associated hostnames is returned. | GET | /cdn/properties/*/versions |
| CreateAPropertyVersion | Create a new version of a property. | POST | /cdn/properties/*/versions |
| CreateAProperty | Creates a property to define the configuration of one or more hostnames (domains) to deploy to the CDN servers. | POST | /cdn/properties |
| GetListOfProperties | The API returns a summary of properties including the ID, latest version number, comments, staging version number, production version number, and last update time of each one. Use query parameters to filter the list of returned properties. | GET | /cdn/properties |