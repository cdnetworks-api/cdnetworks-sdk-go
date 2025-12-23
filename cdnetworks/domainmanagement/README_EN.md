# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/domainmanagement
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/domainmanagement"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &domainmanagement.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := domainmanagement.{ActionName}Response{}
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
| Enablesingledomainservice | Enable an accelerated domain with a state of "disabled" and provide accelerated service with an existing configuration. | PUT | /api/domain/*/enable |
| Disablesingledomainservice | Disables the specified accelerated domain name, and requests for the accelerated domain name after it is disabled will be rejected directly and will not be back to origin. | PUT | /api/domain/*/disable |
| Deleteapidomainservice | Delete one of the added accelerator domains. Cannot be enabled after deletion, only the accelerated domain name can be re-created. | DELETE | /api/domain/* |
| Getfuzzypagingdomainlist | Queries all, or specified, acceleration domains and states under the user's account. Each acceleration domain name contains profile information. The returned accelerated domain name list is sorted in alphabetical order of the initials. | POST | /api/domain/domainList |
| Querydomainbyoriginip | Query the list of all domain name names corresponding to the origin station IP under the user account. | GET | /api/originaldomainlist |
| Disablemultidomainservice | Disable multiple accelerated domain names. Once disabled, all requests for accelerated domain names will be directly rejected and will not be forwarded to the source server. | POST | /api/domain/disable |
| Enablemultidomainservice | Enable multiple acceleration domain names that are currently disabled, and provide acceleration services using the existing configuration. | POST | /api/domain/enable |
| Addcdndomain | When you apply the non-live-streaming acceleration service for specified domain. Content Acceleration, Dynamic Web Acceleration, Download Acceleration are supported.<br><br>The newly added domains can be deployed in a fast mode after the following requirements are met service form for relevant acceleration type; and application that no manual review needed are finished. By default, manual review is required.<br><br>When creating Accelerated domain inside china or outside china, the requirements are as below. <br><br>1. To create the accelerated domain inside china <br>- The accelerated domain name MUST obtain an ICP license<br>- Keep the core requirement of obtaining ICP certification <br> Accelerated domain names should not contain the following information: the sale of controlled drugs and tools, unlawful remarks and information.<br> Video-related accelerated domain names must have Audio Visual Services License; blog forums or community-related accelerated domain names must have BBS special approval qualifications.<br><br>2. To create the accelerated domain outside china <br>- The accelerated domain name MUST use accelerate-no-china = true that indicates that the customer domain name is not accelerated in China.<br><br>Note:<br>1. Data format: Both request and response support XML.<br>2. Recommended frequency: 1 times per15min. | POST | /cdnw/api/domain |
| Updatecdndomain | Modify the configurations of the specified accelerated domain name. | PUT | /cdnw/api/domain/* |
| Querycdndomainservice | View the configuration of the specified accelerated domain name. | GET | /cdnw/api/domain/* |
| Queryapidomainlistservice | Query all the accelerated domain names and states indicated by the user account or the specific cname-label. Each accelerated domain name contains summary information, and the returned accelerated domain name list is sorted in alphabetical order.<br>Note: Disabled domains("enabled" value equals false) could not be modified. | GET | /api/domain |
| Adddomainalias | The accelerated domain name is created by alias, which is configured the same as the associated domain name. The contract information, service type, acceleration area, and certificate for the alias are consistent with the domain name associated. | POST | /api/domain/alias/ |
| Queryapidomainattribution | Get domain list. | GET | /api/domainlist |
| Queryrelationshipbetweendomainandaliasbasedonalias | Query the master domain and alias:<br>1)Query master domain and the aliases under master domain according to the aliases.<br>2)Query for aliases based on the master domain. | GET | /api/domain/domain-and-alias/* |
| Deletedomainalias | the interface of delete the domain alias | DELETE | /api/domain-alias |
| Queryrelationshipbetweendomainandaliasbasedondomain | Query the master domain and alias relation. | POST | /api/domain-alias |
| Getpagingdomainlist | Queries all, or specified, acceleration domains and states under the user's account. Each acceleration domain name contains profile information. The returned accelerated domain name list is sorted by version. | POST | /api/domain/list |
| Querycustomeranycastipforwplus | QueryCustomerAnycastIP | GET | /api/anycast-ips |
| Predeploychangeserverconfig | Predeployed change server. The interface * can be domain name or domain id. | PUT | /api/predeploy/changeserver/* |
| Updatechangeserver | Update change server. The interface * can be domain name or domain id. | PUT | /api/config/changeserver/* |
| Querychangeserver | Query change server. The interface * can be domain name or domain id. | GET | /api/config/changeserver/* |
| Updatecustomeranycastiprecordstatus | Update the record status of Anycast IP | POST | /api/anycast-ips/record-status |