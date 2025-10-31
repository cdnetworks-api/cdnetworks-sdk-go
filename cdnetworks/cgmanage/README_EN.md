# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cgmanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cgmanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &cgmanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := cgmanage.{ActionName}Response{}
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
| Editcontrolgroup | This interface is used to modify the information of a specified ControlGroup, including the name, domain list, and account list of the ControlGroup. Users can choose to update by either overwriting or appending as needed. For custom-type ControlGroups, users can modify their name and domain list, while client-type and contract-type ControlGroups do not allow such modifications. Additionally, users can specify a list of accounts with access permissions for the ControlGroup. This interface facilitates flexible management of ControlGroup configuration, allowing users to control access permissions and applicability by adjusting these settings, thereby better serving business needs. | PUT | /user/control-groups/* |
| Deletecontrolgroup | The function of the Delete ControlGroup interface is to remove a specified control group. Users need to provide the identifier information of the ControlGroup to be deleted as input to execute the delete operation. The return value of this interface includes three key pieces of data: a success indicator, a status code, and a message to help users understand the result of the operation and any error reasons. This interface is particularly useful in scenarios where there is a need to manage and clean up specified control groups in the system, ensuring dynamic adjustment and optimization of system resources. | DELETE | /user/control-groups/* |
| Createcontrolgroup | The Add ControlGroup interface is used to create a new control group. Users need to provide key information such as the control group code and name. The response parameters include a list of specified domains and a list of accounts with access permissions. The domain string array specifies the domains included in the control group, while the account object array specifies the accounts authorized to access the control group, ensuring that the appropriate personnel can access and manage this group. This interface is suitable for scenarios requiring organization and permission management, such as managing access control for multiple subsystems in a large system. By using this interface, users can efficiently allocate and control the permission structure. | POST | /user/control-groups |
| Querycontrolgrouplist | This interface is used to query a list of control groups with access permissions, helping users obtain detailed information about control groups associated with their account. By calling this interface, users can retrieve basic attributes of the control groups, including the name, code, ID, and type. Additionally, the interface provides success or failure status indicators and status messages to help users determine the outcome of the request. This interface is suitable for scenarios where managing and retrieving user permission control group information is needed, assisting users in efficiently maintaining and managing permissions. | GET | /user/control-groups |
| Querycustomizedcontrolgroup | This interface is used to query information about a specified Control Group (CG) and only supports querying custom CGs. Users can obtain detailed information about a specific Control Group, including its code, name, the list of domains it contains, and the list of accounts with access permissions. The response structure of the interface provides status codes and message prompts to indicate whether the request was successful. This is primarily used for managing and monitoring custom Control Groups to ensure correct access control for specific domains and facilitate permission management. By analyzing the returned information, users can understand the current settings and status of the Control Group. | GET | /user/control-groups/* |
| Querysubaccountrelatedcontrolgroup | This interface is used to query the list of control groups associated with a specific sub-user. When a user provides the login name of the sub-user as a parameter, the system returns information about the control groups linked to that user, including the control group ID and name. This allows users to easily view and manage the permissions and control group affiliations of different sub-users. The response structure of the interface also includes status codes and messages, helping users determine whether the request was successful and providing details.  | GET | /user/sub-account/control-groups/* |
| Batchassociateordetachcontrolgroupwithsubaccount | This interface is used to batch add or revoke associated control groups for a specified user. Users need to provide the login name of the sub-user, a list of specified Control Group Codes, and select the operation type as either adding or revoking the association. The interface returns the status of the request, including success or failure status codes and detailed result messages. Application scenarios include the quick management of association relationships between multiple sub-users and control groups, enabling administrators to efficiently perform batch operations to meet complex and dynamic control group management needs within an organization. | POST | /user/sub-account/control-groups |
| Getdomainlistofcontrolgroup | The function of this interface is to obtain a list of control groups (CGs) that the current user account has permission to access, along with their associated domain information. Users do not need to provide specific parameters when calling this interface; the system will return a complete list containing the CG names and corresponding domains. The results include the ID of each control group and the associated domains. This interface is particularly useful for scenarios where there is a need to quickly understand the full scope of domains that an account can manage or monitor, helping users effectively manage resources and permissions.<br><br><br> | POST | /user/cgdomainlist |