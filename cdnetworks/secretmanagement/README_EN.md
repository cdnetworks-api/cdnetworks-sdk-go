# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/secretmanagement
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/secretmanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &secretmanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := secretmanagement.{ActionName}Response{}
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
| CreateASecret | CreateASecret to protect sensitive text from being exposed in a property's Edge Logic. Refer to the secret in your property's Edge Logic using the $SECRET(secretName) syntax. When you deploy the property, references to $SECRET(secretName) will be replaced by the text. | POST | /cdn/secrets |
| GetAListOfSecrets | GetAListOfSecrets that have been defined. Filter the results using the query parameters. | GET | /cdn/secrets |
| GetASecret | Get details about a secret including its usage in properties deployed to staging and production. | GET | /cdn/secrets/* |
| UpdateASecret | UpdateASecret using this API. Deployed properties using the secret will continue to use its original value at the time of deployment. If you change the secret's name, property versions referring to the old name will not pass future validations until you edit the properties to remove the references or update them to use a valid secret's name. If you change the secret's value, properties using the secret must be revalidated before they can be redeployed.  | PATCH | /cdn/secrets/* |
| DeleteASecret | DeleteASecret by its ID. A secret used in a deployed property cannot be deleted. If a secret is obsolete, you must remove its references in properties, revalidate, and redeploy before you can delete the secret. | DELETE | /cdn/secrets/* |