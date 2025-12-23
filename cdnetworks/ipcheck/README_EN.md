# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ipcheck
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ipcheck"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &ipcheck.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ipcheck.{ActionName}Response{}
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
| Querycdnservicerealip | This interface is used to obtain the real IP list of the CDN service, which is particularly suitable for scenarios where the source station has set whitelist restrictions. Users can obtain the IP whitelist of the CDN node provided by our company for back-to-source by calling this interface, so as to correctly configure and ensure that the data request can smoothly reach the source station through the CDN. | GET | /api/si/report/whiteip-list |
| Queryspecificipbelong | Query whether a given IP address belongs to our CDN IP. The user needs to provide a list of IP addresses. The returned result includes whether each IP address belongs to our CDN IP. | POST | /api/si/tools/ipCheck |
| Checkipowner | This API is used to query the affiliation information of specific IP addresses. Users can provide one or more IP addresses to check if they belong to the company CDN nodes, as well as their affiliated country, province, city, and carrier information. The returned results include an identifier indicating whether the IP is a company CDN node; if not, it will return 'unknown'. This API is applicable for users who wish to verify if an IP address belongs to our company. | POST | /api/tools/ip-info |