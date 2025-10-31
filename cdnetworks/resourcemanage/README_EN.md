# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/resourcemanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/resourcemanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &resourcemanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := resourcemanage.{ActionName}Response{}
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
| Vmpqueryflavor | You can use this interface to query the specification list of instances that a certain region can provide, and then create instances using this specification. | GET | /vmp/flavors |
| Vmpquerynode | Query which nodes of the edge computing platform can provide services. | GET | /vmp/nodes |
| Vmpquerybandwidth | Query the real-time redundant bandwidth of all the nodes you are using. Nodes with redundant bandwidth will return True, while nodes without redundant bandwidth will return False. If the bandwidth data collection results of the node are insufficient to support the timeliness requirements of the interface, that is, if the bandwidth statistics data is old and does not meet the timeliness requirements, undefined will be returned.<br>The judgment method for whether a node has redundant bandwidth: the upper limit bandwidth of the node is greater than the real-time usage bandwidth of the node. Data granularity: 5 minutes. | POST | /vmp/redundant-bandwidth |
| Getservergrouplist | Get servergroup list | GET | /api/clouddns/servergroups |
| Getservergroupbyid | Get specified servergroup by ID | GET | /api/clouddns/servergroups/* |
| Createservergroup | Create new servergroup | POST | /api/clouddns/servergroups |
| Editservergroupbyid | Update an existing servergroup | PUT | /api/clouddns/servergroups/* |
| Deleteservergoupbyid | Delete a specified servergroup | DELETE | /api/clouddns/servergroups/* |
| Getserverlistbyservergroup | Get specified servergroup's server list | GET | /api/clouddns/servergroups/*/servers |
| Createserver | Create server for specified servergroup | POST | /api/clouddns/servergroups/*/servers |
| Updateserver | Update specified server by servergroupId & serverId | PUT | /api/clouddns/servergroups/*/servers/* |
| Deleteserver | Delete Specified server | DELETE | /api/clouddns/servergroups/*/servers/* |
| Gethealthcheckerlist | Get the list of healthchecker methods | GET | /api/clouddns/infra/healthcheckers |
| Getcvlist | Get CV(Custom View) list | GET | /api/clouddns/cvs |
| Getregionlist | Get region list | GET | /api/clouddns/regions |
| Getisplist | Get ISP list | GET | /api/clouddns/isps |
| Gethealthcheckagentlist | Get health check agent list | GET | /api/clouddns/infra/healthcheckers/iplist |
| Enablehealthcheckwhitelist | Whether the health check detector can be configured with a whitelist. When configured as auto mode, you need to implement the function on your detected machine, and add the health check agent IP that has expired to the whitelist. | POST | /api/clouddns/infra/healthcheckers/whitelist |
| Gethealthcheckwhitelist | Get health check whitelist | GET | /api/clouddns/infra/healthcheckers/whitelist |