# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/productmanagement
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/productmanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &productmanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := productmanagement.{ActionName}Response{}
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
| GetSystemConfiguration | This API returns system configuration settings including supported <a href="/cdn/docs/edge-logic/supported-directives">Edge Logic directives</a>, advanced Edge Logic directives, the default pricing for traffic, and nonstandard ports that can handle HTTP and HTTPS requests. Please note that your account may not have access to the advanced Edge Logic directives or features. Please contact our support team if you require access to a particular directive, feature, or port that is not listed. | GET | /cdn/systemConfigs |
| GetIpAddressesOfTheCdn | Get a list of IP addresses used by CDN Pro. If your property's origin configuration uses the directConnection setting 'noDirect', and its Edge Logic does not use the origin_fast_route directive, you can restrict access to your origin servers by updating your allow lists with these IP addresses. Since the CDN's IP addresses can periodically change, we recommend that you check once a day and update your allow lists. | GET | /cdn/publicIpList |
| GetAListOfOriginShields | Get a list of shields that can be used as an extra layer between the CDN Pro edge servers and your origin servers. A shield can further reduce the number of requests to your origin and limit the requests to a smaller set of servers. <br>This is an advanced feature. If you require this feature, please contact our support team. | GET | /cdn/shields |
| GetAShieldByItsId | Obtain information about a shield which can be used as an extra layer between the CDN Pro edge servers and your origin servers. This is an advanced feature. If you require this feature, please contact our support team. | GET | /cdn/shields/* |
| CheckIfIpAddressesBelongToTheCdnProPlatform | This API allows you to check whether one or more IP addresses belong to the CDN Pro platform. | POST | /ngadmin/ipDetails |
| GetOriginFastRouteIpList | This API returns a list of IP addresses used by our origin fast route servers. If you use our origin_fast_route feature to make the connections faster and more reliable and wish to restrict access to your origin servers, add these IP addresses to your allow lists.<br> | GET | /cdn/originFastRouteIpList |
| GetStagingServersList | This API returns a list of IP addresses of servers that can be used to test deployments of properties and certificates to 'staging'. We recommend deploying to 'staging' and testing before deploying to 'production'. | GET | /cdn/stagingServers |