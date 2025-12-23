# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/imagemanage
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/imagemanage"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &imagemanage.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := imagemanage.{ActionName}Response{}
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
| Vmpqueryimage | Query the list of images that users can use. The list of image resources displayed includes user-defined images and public images provided by the edge computing platform. | GET | /vmp/images |
| Vmpcreateimage | You can create a virtual machine system disk as a mirror, and then use it to create a new virtual machine. It is recommended to shut down the virtual machine or stop applications or services on the virtual machine during the production of the image to avoid affecting the integrity of the image data. After the image production is completed, restart the virtual machine and its applications. The image created by this type of operation is returned as SNAPSHOT in the image query interface, representing the image created using a virtual machine snapshot. | POST | /vmp/images |
| Vmpremoveimage | Delete your custom image.<br>Deleting the image does not affect the already created virtual machine, but it cannot be used to create new virtual machines in the future. <br>You can only delete custom images created by yourself, and other customers' custom images and public images cannot be deleted. | DELETE | /vmp/images/* |
| Deployvmpimagepreheating | Used to extract private images of preheating customers. This interface is asynchronous, and the image preheating results need to be queried separately. | PUT | /vmp/images/preHeating |
| Queryvmpimagepreheatingstate | Used to query the image preheating status. | GET | /vmp/images/preHeatingInfo/* |