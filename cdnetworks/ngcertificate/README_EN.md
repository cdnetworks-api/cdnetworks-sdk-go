# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngcertificate
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngcertificate"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &ngcertificate.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ngcertificate.{ActionName}Response{}
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
| CreateACertificate | Create a certificate along with its first version. You can choose to upload the key/certificate pair or generate a self-signed one. | POST | /cdn/certificates |
| GetListOfCertificates | Get a list of available certificates including each certificate's ID, name, version deployed to production, version deployed to staging, the latest version number, when the certificate was last updated, and when the certificate expires. By default, certificates are returned in order of most recently updated first. Use query parameters to filter the list of returned certificates. | GET | /cdn/certificates |
| GetACertificate | Gets details about a certificate including the versions of the certificate and who is using it. | GET | /cdn/certificates/* |
| UpdateACertificate | UpdateACertificate using this API. The fields are the same as those in the <a href="#operation/createCertificate">Create a certificate API</a>. If the certificate is currently used by properties deployed to production or staging, we recommend that you follow this API call by <a href="#operation/createDeployment">creating a deployment task</a> to deploy the updated certificate. Otherwise, CDN Pro may still serve content using an older version of the certificate. | PATCH | /cdn/certificates/* |
| DeleteACertificate | This API lets you delete a certificate that is not being used in production or staging. | DELETE | /cdn/certificates/* |
| DownloadTheCsr | Obtain the certificate signing request (CSR) information. You can take it to a certificate authority to apply for a certificate. Once you have it, you should update our system using the 'Update a certificate' API. | GET | /cdn/certificates/*/csr |
| GetACertificateVersionsDetails | Obtain details about a certificate version including the expiration date, algorithm and key length, fingerprint, and whether it is deployed in production and staging. | GET | /cdn/certificates/*/versions/* |