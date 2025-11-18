# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/zonemanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/zonemanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &zonemanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := zonemanage.{ActionName}Response{}
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
| Getzonelist | Get zone list | GET | /api/clouddns/zones |
| Getzonebyid | Get zone by zoneId | GET | /api/clouddns/zones/* |
| Updatazonebyid | Update a zone's configuration | PUT | /api/clouddns/zones/* |
| Deployzoneconfig | Deploy zone's configuration to staging or production | POST | /api/clouddns/zones/*/deployment |
| Getrecordlistbyzone | Get all record list of a zone | GET | /api/clouddns/zones/*/records |
| Getrecordbyid | Get record info by zoneId & recordId | GET | /api/clouddns/zones/*/records/* |
| Bulkcreaterecordsbyzone | Bulk create records for a specified zone. This method requires the request body to be in the form of a list. OpenApi does not currently support configuration calls. Please use another tool to test the calls. | POST | /api/clouddns/zones/*/records |
| Updaterecordbyid | Update a specified record | PUT | /api/clouddns/zones/*/records/* |
| Bulkupdaterecordsbyzoneid | Bulk update specified zone's specified records. This method requires the request body to be in the form of a list. OpenApi does not currently support configuration calls. Please use another tool to test the calls. | PUT | /api/clouddns/zones/*/records |
| Deleterecordbyid | Delete specified record by zoneId & recordId | DELETE | /api/clouddns/zones/*/records/* |
| Createzone | Create new Zone | POST | /api/clouddns/zones |
| Deletezone | delete Zone | DELETE | /api/clouddns/zones/* |
| Batchcreatezone | Batch Create new Zone | POST | /api/clouddns/zones/bulk |
| Batchdeletezone | Batch delete zones, maximum 20 zones can be deleted at the same time | DELETE | /api/clouddns/zones/bulk/* |
| Updateztsbulk | Batch create or update ZTS configuration information. Users can create or update ZTS configuration information autonomously through the interface to achieve Zone data synchronization | PUT | /api/clouddns/zts/bulk |
| Deleteztsbulk | The ZTS configuration deletion interface removes the ZTS information configured on the Zone. After deletion, configuration changes to the Zone will no longer be automatically synchronized. | DELETE | /api/clouddns/zones/zts/bulk |