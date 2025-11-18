# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/sharedconfiguration
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/sharedconfiguration"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &sharedconfiguration.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := sharedconfiguration.{ActionName}Response{}
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
| Listsharedwafruleexceptions | Return a list of WAF Rule Exceptions for the shared configuration. | POST | /api/v1/waf/share/exception/get-list |
| Createsharedwafruleexception | Create a WAF rule exception to shared configurations. | POST | /api/v1/waf/share/exception/create |
| Updatesharedwafruleexception | Modify the configuration of WAF Rule Exception in the shared configuration. | POST | /api/v1/waf/share/exception/update |
| Deletesharedwafruleexception | Delete the WAF rule exceptions in the shared configuration. | POST | /api/v1/waf/share/exception/delete |
| Createcustomaction | Create a custom action,up to 5 custom response actions can be configured. | POST | /api/v1/share-action/add-customize-act |
| Listcustomactions | Return a list of custom actions. | POST | /api/v1/share-action/get-customize-act-list |
| Updatecustomaction | Update the configuration of a custom action. | POST | /api/v1/share-action/update-customize-act |
| Deletecustomactions | Delete custom actions. Note: the referenced custom action cannot be deleted. | POST | /api/v1/share-action/delete-customize-act-batch |
| Createsharedwhitelistrule | Create a Whitelist rule for shared configurations. | POST | /api/v1/common/share-whitelist/add |
| Updatesharedwhitelistrule | Modify a Whitelist rule's configuration for the shared configuration. | POST | /api/v1/common/share-whitelist/update |
| Deletesharedwhitelistrule | Delete the Whitelist rules for the shared configuration. | POST | /api/v1/common/share-whitelist/delete |
| Listsharedwhitelistrules | Return a list of Whitelist rules for shared configurations. | POST | /api/v1/common/share-whitelist/get-list |
| Createsharedratelimitingrule | Create a Rate Limiting rule for shared configurations. | POST | /api/v1/share-rate-limit/add-rule |
| Updatesharedratelimitingrule | Modify a Rate Limiting rule's configuration for the shared configuration. | POST | /api/v1/share-rate-limit/update-rule |
| Deletesharedratelimitingrule | Delete the Rate Limiting rules for the shared configuration. | POST | /api/v1/share-rate-limit/delete-by-ids |
| Listsharedratelimitingrules | Return a list of Rate Limiting rules for shared configurations. | POST | /api/v1/share-rate-limit/get-rule-list |
| Createappapiexceptionfeature | Create APP/API exception feature. | POST | /api/v1/dms/service-feature/add |
| Queryappapiexceptionlist | Query APP/API exception list. | POST | /api/v1/dms/service-feature/get-list |
| Deleteappapiexceptionfeature | Delete APP/API exception feature. | POST | /api/v1/dms/service-feature/delete |
| Queryappapiexceptionfeaturedetail | Query APP/API exception feature details(Share Configurations). | POST | /api/v1/dms/service-feature/get-detail |
| Queryappapiexceptionfeaturereferenceddomains | Query APP/APIException Feature Referenced Domains(Share Configurations). | POST | /api/v1/dms/service-feature/get-relate-domain-list |
| Updateshareconfigurationsappapiexceptionfeature | Modify App/API exceptions(ShareConfigurations). | POST | /api/v1/dms/service-feature/update |
| Listsharecustomizebots | Query share custom Bots list. | POST | /api/v1/share-customize-bots/get-list |
| Deletesharecustomizebots | Delete share custom Bots. | POST | /api/v1/share-customize-bots/delete |
| Createsharedcustomrule | Create a Custom rule for shared configurations. | POST | /api/v1/share-customize-rule/add |
| Updatesharedcustomrules | Update the custom rule of the sharing configuration. | POST | /api/v1/share-customize-rule/update |
| Listsharedcustomrules | Get a list of  custom rules for shared configurations. | POST | /api/v1/share-customize-rule/get-list |
| Deletesharedcustomrules | Delete custom rule for shared configuration. | POST | /api/v1/share-customize-rule/delete |
| Disassociateshareratelimit | Disassociate shared configuration rate limiting rule from domain. | POST | /api/v1/common/share-rate-limit/disassociate |
| Associateshareratelimit | Associate shared configuration Rate limiting rule from Domain. | POST | /api/v1/common/share-rate-limit/associate |
| Associatesharecustomizerule | Associate the custom rules of shared configuration with the domain. | POST | /api/v1/common/share-customize-rule/associate |
| Disassociatesharecustomizerule | Dissociate custom rules from shared configuration with domain. | POST | /api/v1/common/share-customize-rule/disassociate |
| Associatesharecustomizebots | Associate shared configuration custom Bots from Hostname. | POST | /api/v1/common/share-customize-bots/associate |
| Disassociatesharecustomizebots | Disassociate shared configuration custom Bots from Hostname. | POST | /api/v1/common/share-customize-bots/disassociate |
| Associatesharedwhitelistrule | Associate shared configuration Whitelist rule from Domain. | POST | /api/v1/common/share-whitelist/associate |
| Disassociatesharedwhitelistrule | Disssociate shared configuration Whitelist rule from Domain. | POST | /api/v1/common/share-whitelist/disassociate |
| Associatedmsshareservicefeature | Associate the shared configuration APP/API exception with the domain. | POST | /api/v1/dms/service-feature/relateDomains |
| Disassociatedmsshareservicefeature | Disassociate the shared configuration app/API exception from the domain. | POST | /api/v1/dms/service-feature/disRelateDomains |