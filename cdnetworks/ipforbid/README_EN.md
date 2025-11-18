# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ipforbid
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ipforbid"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &ipforbid.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ipforbid.{ActionName}Response{}
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
| Forbidorresumevisitoripsbydomainservice | Forbid and Resume guest IP addresses from accessing the domain. | POST | /api/spider/ip-forbid |
| Queryforbiddingvisitoripsbydomainservice | Query Forbidding IPs by domain, support paging query. | POST | /api/spider/ip-forbid/query |
| Queryforbiddingvisitoripsbylabelcodeservice | Query the forbidding IPs by LabelCode, and support pagination.<br><br> | POST | /api/spider/label-ip-forbid/query |
| Forbidorresumevisitoripsbylabelcodeservice | A customer can create a label that can be associated with multiple domains and used to forbid or resume visitor IPs through that label . This has the same effect as forbidding or resuming the designated visitor IPs for all domains associated with the label. | POST | /api/spider/label-ip-forbid/operate |