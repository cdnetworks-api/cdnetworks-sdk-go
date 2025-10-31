# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/gamingshield
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/gamingshield"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &gamingshield.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := gamingshield.{ActionName}Response{}
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
| Listgroupdetails | Used to query the website group details list, users can provide website group ID, name, domain name to query the corresponding website group details information, including protocol, port, domain name, etc. | POST | /api/v1/hijack/group-info/get-detail-list |
| Createwebsitegroup | Used for creating new website groups. Add websites with identical content to a single group, and user access will be automatically distributed to a healthy domain within the group to ensure friendly access experience. Each website group includes domain, protocol, and port information. | POST | /api/v1/hijack/group-info/add |
| Updatewebsitegroup | Used to modify website grouping information, users can modify the protocol, port, associated domain name, and other information of website grouping through this interface. | POST | /api/v1/hijack/group-info/update |
| Deletewebsitegroup | Used to delete website groups, users can delete unrelated website groups by providing the website group ID. | POST | /api/v1/hijack/group-info/delete |
| Listentrydomains | Used to query the detailed list of entrance hostnames, users can obtain detailed information by providing the entrance host name ID and other information, including the entrance hostname protocol and port, as well as associated website grouping information. | POST | /api/v1/hijack/entry-domain/get-detail-list |
| Addentrydomain | Used to add an entrance hostname, users can provide the protocol and port of the entrance hostname, as well as the associated website group ID, to add the entrance hostname. | POST | /api/v1/hijack/entry-domain/add |
| Updateentrydomain | Used to modify the entrance hostname, when there are changes in the entrance domain name information, users can provide the entrance hostname ID to modify related information, including associated website group ID, port, protocol, etc. | POST | /api/v1/hijack/entry-domain/update |
| Deleteentrydomain | Used to delete the entrance hostname, when the entrance hostname is not needed, users can provide it to delete. | POST | /api/v1/hijack/entry-domain/delete |
| Restartentrydomain | Used to restart the portal hostname, when the portal hostname is disabled, users can provide the portal hostname to restart the portal hostname and restore the service. | POST | /api/v1/hijack/entry-domain/restart |
| Listnavigationrules | Used to query the navigation rule list, the complete information of the specified navigation rule can be obtained by providing a list of website group IDs or website group names, including rule switches, matching conditions, and other information. | POST | /api/v1/hijack/group-rule/get-list |
| Updatedistributionrule | Used for modifying distribution rules, this interface allows updates when there are changes to the user requests or targets that require distribution. | POST | /api/v1/hijack/group-rule/save |