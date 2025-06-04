# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/alertrules
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/alertrules"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &alertrules.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := alertrules.{ActionName}Response{}
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
| Createcloudmonitorrealtimealarmrule | Create real-time monitoring alarm rules for cloud monitoring. Users can provide the monitoring rule name, monitored resource object, alarm trigger threshold, and alarm contact information to generate real-time monitoring rules. This is suitable for scenarios where timely attention to monitoring data changes of cloud products on the console is required. The interface for creating monitoring rules currently only supports some CDN monitoring items by domain name.  | POST | /api/cloudmonitor/alarm/real-time/add |
| Editcloudmonitorrealtimealarmrule | Edit the real-time monitoring alarm rules of Cloud monitoring. This interface modifies the rules in an overwrite manner. Users can pass in the monitoring rule ID, the name of the monitoring rule to be modified, the monitoring resource object, the alarm trigger condition threshold, the alarm contact person and other information to overwrite and modify the monitoring rule.<br><br>Interface management monitoring rules currently only support some monitoring items by domain. | POST | /api/cloudmonitor/alarm/real-time/edit |
| Querycloudmonitorrealtimealarmrule | Query the real-time monitoring alarm rules of Cloud monitoring. Users can enter the monitoring rule ID, query the monitoring rule name, monitoring resource object, alarm trigger condition threshold, alarm contact and other monitoring rule details.<br><br>Interface management monitoring rules currently only support some monitoring items by domain. | POST | /api/cloudmonitor/alarm/real-time/query |
| Deletecloudmonitorrealtimealarmrule | Delete the real-time monitoring alarm rule of Cloud monitoring. Users can pass in the monitoring rule ID to delete the specified monitoring rule. | POST | /api/cloudmonitor/alarm/real-time/delete |