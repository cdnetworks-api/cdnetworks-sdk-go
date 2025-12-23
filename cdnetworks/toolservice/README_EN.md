# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/toolservice
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/toolservice"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &toolservice.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := toolservice.{ActionName}Response{}
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
| Bandwidthlimitservice | Set/cancel bandwidth limits to certain specified domain. It's recommended that the call frequency is no higher than 1/5min.<br><br>Limit: no more than 3 domains can be operated on at a time. To set the limit higher, please contact technological customer service to evaluate. (The interface must be under strict control because it directly affects bandwidth) | POST | /api/tools/setBandwidthLimit |
| Querybandwidthlimittasklistservice | This interface is used to query the bandwidth limit task list under the account and return detailed information of all tasks with bandwidth. The returned content includes the domain name, task name, and the set maximum bandwidth value. This interface is suitable for scenarios where traffic control policies need to be evaluated and managed, helping users quickly identify and manage the currently set bandwidth limit tasks. | POST | /api/tools/queryBandwidthLimitTaskList |
| Icpqueryservice | Query whether specified domain has been registered in MIIT of Mainland China.<br><br>Limit: call frequency of the interface cannot exceed 50 times/day. | GET | /api/icp |
| Ipdomainservice | This interface is used to query the domain names that are using the IP address. The user enters the IP address to obtain the list of domain names associated with the IP. The information returned by the interface includes the current usage status of the IP and the list of domain names that use the IP. In actual applications, this interface can help users detect the domain name usage of a specific IP, which is suitable for network monitoring and management. | POST | /api/tools/ip/domain-list |
| Queryallbandwidthlimittasklistservice | This interface is used to query all bandwidth restriction tasks configured under the user account. When calling, the user can choose whether to include all customer domain names involved in the task and decide whether to return information on all task status. The returned data displays detailed information on each bandwidth restriction task in a list format, including task name, category, status, and related control policies and parameters. This interface helps users manage bandwidth settings and can effectively control and process specific traffic and requests in a timely manner. | POST | /api/tools/queryAllBandwidthLimitTaskList |