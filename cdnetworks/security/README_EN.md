# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/security
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/security"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &security.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := security.{ActionName}Response{}
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
| Vmpcreatesshkey | When creating an instance, it can specify the use of key pairs. Before that, it need to create a key pair. | POST | /vmp/osKeypairs |
| Vmpquerysshkey | Query all key pairs. | GET | /vmp/osKeypairs |
| Vmpremovesshkey | Delete the specified key pair by name. The instances created using this key pair are not affected and cannot be found on the query interface after deletion, nor can new instances be created using this key pair. | DELETE | /vmp/osKeypairs/* |
| Vmpquerysecuritygroup | Used to query the basic information of security groups created by customers. | GET | /vmp/security-group |
| Vmpbindsecuritygroup | Use to change the binding relationship between an instance and a security group. Only virtual instances support binding security groups, not bare metal instances. | PUT | /vmp/security-group/server |
| Deletionsecuritygrouprules | Used to delete security group rule information. | DELETE | /vmp/security-group/rule/* |
| Editsecuritygrouprules | Used to modify security group rule information. | PUT | /vmp/security-group/rule |
| Addsecuritygrouprules | Used to add security group rule information.<br>1) Within the same security group, the content of different rules cannot be completely the same;<br>2) When adding multiple rules in bulk, the result is either successful or failed addition. | POST | /vmp/security-group/rule |
| Deletesecuritygroup | Used to delete security group information. If the security group has defined rules, they will be deleted together with the rules.<br>Explanation:<br>1) If there is currently an instance bound to a security group, the deletion of the security group fails;<br>2) When deleting multiple security groups in bulk, if one security group fails to be deleted, it does not affect the normal deletion of other security groups. | DELETE | /vmp/security-group/* |
| Editsecuritygroup | Used to modify the security group name or comment information. | PUT | /vmp/security-group |
| Addsecuritygroup | Used to add a security group. | POST | /vmp/security-group |