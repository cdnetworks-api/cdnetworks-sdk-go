# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/certificationmanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/certificationmanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &certificationmanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := certificationmanage.{ActionName}Response{}
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
| Querycertificate | Check the details of SSL certificate, including certificate name, whether it shares or uses the current domain name information, etc. | GET | /api/ssl/certificate/* |
| Querycertificatelist | Check the SSL certificate lists and information, including certificate ID, certificate name, whether it shares or uses the current certificate domain information, etc. | GET | /api/ssl/certificate |
| Getcertificatecontent | get certificate content | GET | /api/ssl/content/*/download |
| Createcertificatev2 | Add a new certificate interface, including certificate name, certificate public key (CRT and Ca content merge), certificate key, csrid and comment. | POST | /api/certificate |
| Deletecertificate | Delete a certificate. A certificate cannot be deleted if it is in use. | DELETE | /api/certificate/* |
| Getcertificatev2 | Returns certificate detail by certificate ID. | GET | /api/certificate/* |
| Updatecertificatev2 | Re-upload a certificate. | PUT | /api/certificate/* |
| Getcertificatecontent | Query certificate content | GET | /api/certificate/*/content |
| Getcertificateassociatedhostnames | Query certificate related domains | GET | /api/certificate/*/domain |
| Querydomainmulticertconfig | Query Domain MultiCert Config | GET | /api/config/certificate/v2/* |
| Querycertificatepagination | This API is used to query paginated SSL certificate lists and their details. Users can specify the page number and page size to retrieve the certificate list. The results include certificate ID, certificate name, whether it is shared, and domain information using the current certificate. | POST | /api/ssl/certificate/pagination |
| Getcertificatecontentinfos | This API retrieves detailed information for multiple certificates based on the provided list of certificate IDs, including encrypted key file content, certificate expiration date, certificate ID, encrypted crt file content, and certificate name. | POST | /api/ssl/content |