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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
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
| Querycertificatelist | Check the SSL certificate lists and information, including certificate ID, certificate name, whether it shares or uses the current certificate domain information, etc. | GET | /api/ssl/certificate |
| Getcertificatecontent | get certificate content | GET | /api/ssl/content/*/download |
| Addcertificateservicev2 | Add a new certificate interface, including certificate name, certificate public key (CRT and Ca content merge), certificate key, csrid and commont. | POST | /api/certificate |
| Deletecertificate | Delete a certificate. A certificate cannot be deleted if it is in use. | DELETE | /api/certificate/* |
| Querycertificateinfo | Returns certificate detail by certificate ID. | GET | /api/certificate/* |
| Editcertificatev2 | Re-upload a certificate. | PUT | /api/certificate/* |
| Querycertificatecontent | Query certificate content | GET | /api/certificate/*/content |
| Querycertificaterelateddomains | Query certificate related domains | GET | /api/certificate/*/domain |
| Querydomainmulticertconfig | Query Domain MultiCert Config | GET | /api/config/certificate/v2/* |
| Reissuecertificateforwplus | This interface is used for reissuing certificates. You can reissue the certificate by providing the certificate ID, certificate description, certificate algorithm, verification method, whether it is automatically verified, whether it is automatically deployed, common name, and subject alternate name. When the call is successful, the interface will return the sales order ID. | POST | /api/certificate/reissue |