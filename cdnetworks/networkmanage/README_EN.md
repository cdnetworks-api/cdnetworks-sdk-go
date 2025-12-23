# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/networkmanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/networkmanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &networkmanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := networkmanage.{ActionName}Response{}
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
| Vmpreleaseedgeip | Used to release exclusive and drifting additional public IP addresses.<br>Explanation:<br>1) Drift type additional IP needs to be unbound with the cloud host before it can be released. If it is not unbound, the release will fail;<br>2) If multiple drifting IPs are released in bulk, some IPs have been unbound and some IPs have not been unbound, the unbound ones will be released normally, while the unbound ones will fail to be released.<br>3) Exclusive IP is released without unbinding. | PUT | /vmp/edgeIp/release |
| Vmpallocateedgeip | Used to apply for drift type additional public IP. Drift mode supports the simultaneous use of multiple instances of the same IP, and is commonly used in master-slave switching scenarios, such as LVS. | POST | /vmp/edgeIp/allocate |
| Vmpassignedgeip | Bind the requested drifting additional IP to the specified instance.<br>Drift type additional IP supports the simultaneous use of multiple instances of the same IP, commonly used in master-slave switching scenarios, such as LVS. | PUT | /vmp/edgeIp/assign |
| Vmpunassignedgeip | Can be used to unbind drifting additional public network IPs and instances. Drift type additional public IP supports the simultaneous use of multiple instances of the same IP, and is commonly used in master-slave switching scenarios, such as LVS.<br>Explanation:<br>1) The drifting additional public IP to unbind must be an IP bound to the specified instance, and cannot be an IP bound to other virtual machines;<br>2) If multiple drifting additional public network IPs are unbound in bulk, and some IPs are bound to other instances, all IPs fail to be unbound, and the interface returns an error message. | PUT | /vmp/edgeIp/unassign |
| Vmpqueryedgeip | Used to query edge public IP addresses that have been applied for. | GET | /vmp/edgeIp |