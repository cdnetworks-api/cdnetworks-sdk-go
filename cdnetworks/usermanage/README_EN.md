# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/usermanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/usermanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &usermanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := usermanage.{ActionName}Response{}
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
| Addsubaccount | Create a sub account | POST | /sub-account |
| Querysubaccountinfo | Get sub account | GET | /sub-account/* |
| Getsubaccountlist | Get the list of sub-users by the main account | POST | /sub-account/list |
| Updatesubaccount | Update sub account | PUT | /sub-account |
| Deletesubaccount | Delete Specified sub-users | DELETE | /sub-account/* |
| Querypolicyattachedmainaccountorsubaccount | This interface is used to query the list of permission policies associated with a specified user. By entering their login name, users can obtain information about the policies linked to them, including the policy ID, name, description, type (system policy or custom policy), and support for multilingual descriptions. The results returned by the interface include status codes and related information to help users understand the specific attributes and classifications of the policies. This is very useful for managing user permissions and creating or modifying permission policies, enabling system administrators to configure and adjust permissions more efficiently. | GET | /user/policy-attached/* |
| Batchaddorrevokepolicytosubaccount | This interface is used to batch add or revoke permission policies for a specified sub-user. By inputting the sub-user's login name and a list of permission policy identifiers, you can choose to perform add or revoke operations. When adding permissions, the system will add the corresponding policies to the sub-account to expand its permission scope; when revoking permissions, it will remove the specified policies, reducing the sub-account's access permissions. The return value includes a request status code and operation information prompts, allowing users to confirm the success of the batch operation. This interface is suitable for scenarios where centralized management of sub-account permissions is needed, simplifying the process of batch permission adjustments.<br><br><br> | POST | /user/policies |