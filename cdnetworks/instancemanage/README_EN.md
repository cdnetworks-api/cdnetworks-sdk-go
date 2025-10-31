# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/instancemanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/instancemanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &instancemanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := instancemanage.{ActionName}Response{}
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
| Vmpremoveinstance | Destroy the specified instance, supporting batch destruction. After destruction, all physical resources used by the instance will be reclaimed, including disks and snapshots. All related data will be lost and permanently irreparable. | POST | /vmp/servers/delete |
| Vmpqueryinstance | instance query | GET | /vmp/servers |
| Vmpinstanceoperation | If you find that an instance is not functioning properly (such as being able to ping but unable to log in), you can call the interface to attempt to forcibly restart the machine. Forcing a reboot is equivalent to a traditional server power outage reboot, which may result in the loss of data in the instance operating system that has not been written to the disk. A normal shutdown is a normal shutdown operation. After successfully calling the interface, it is necessary to call the instance query interface again to confirm the latest status of the instance and verify whether the restart was successful. | POST | /vmp/servers/*/action |
| Vmpcreateinstance | Through this interface, you can apply for a cloud host instance of a specified specification in a certain region. After the instance is created, you can obtain the latest status of the instance by using the instance query interface. | POST | /vmp/servers |
| Editinstance | Used to modify instance information. Currently, only instance name modification is supported. | PUT | /vmp/servers |
| Instanceipv6management | Assign/remove IPv6 addresses for existing cloud host instances, and bare metal instances do not support it.<br>Instructions for assigning IPv6:<br>1) Action=ALLOCATION;<br>2) The instance status must be running or down, and the node where the host is located must support IPv6;<br>3) If the instance already has IPv6, the interface will directly return the existing IPv6 address.<br>Instructions for removing IPv6:<br>1) Action=REMOVE;<br>2) The instance state must be running or down. | POST | /vmp/servers/ipv6 |
| Instancediskscaling | Supports attaching disks to instances online. | POST | /vmp/servers/attachDisk |
| Manageinstancetags | Manage virtual machine instance labels and support modifying or deleting instance labels. | PUT | /vmp/server-tags |
| Convertfreetypeinstancetochargetype | Convert free type instance to charge type.Transfer the designated free instances to regular status to prevent them from being deleted upon expiration, and support batch transfer to regular status. After being confirmed, the instance will be converted to a billing instance. | POST | /vmp/servers/charge |
| InstanceRebuild | Through this api, it is possible to reinstall the system for the already created instances. | POST | /vmp/servers/rebuild |