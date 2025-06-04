# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/nglogconfiguration
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/nglogconfiguration"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &nglogconfiguration.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := nglogconfiguration.{ActionName}Response{}
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
| GetAListOfLogConfigurations | Obtain a list of log configurations you have created. | GET | /cdn/report/logConfigs |
| CreateALogConfiguration | CreateALogConfiguration used to format the access logs for one or more hostnames. An applicable log configuration must exist for you to obtain logs for a hostname. It takes around 40 minutes for a newly created log configuration to take effect. | POST | /cdn/report/logConfigs |
| GetALogConfiguration | Retrieve a log configuration. | GET | /cdn/report/logConfigs/* |
| UpdateALogConfiguration | UpdateALogConfiguration's settings. It takes around 40 minutes for an updated configuration to take effect. | PATCH | /cdn/report/logConfigs/* |
| DeleteALogConfiguration | DeleteALogConfiguration. | DELETE | /cdn/report/logConfigs/* |
| GetAccessLogsForHostnames | Get access logs representing requests to one or more hostnames in your deployed properties. These logs contain data from requests made 2 or more hours earlier. By default, only logs within the past 14 days are available, and they are separated into files each covering part of a day.<br><br>To obtain logs for a hostname, you must first specify the format of the logs by creating a log configuration that applies to the hostname. Requests to your content made after the applicable log configuration has been created will be logged. Otherwise, no logs will be available for download.<br><br>For example, suppose you define a log configuration format as follows:<br>%cltip %rmtuser [%utctime] "%method %url %protocol" %statuscode %rspsize "%referer" "%ua" %rsptime<br><br>This API will return a list of available logs such as:<br>{<br>  "logs":[<br>          {<br>          "dateFrom":"2021-10-31T00:00:00Z",<br>          "dateTo":"2021-10-31T00:29:59Z",<br>          "fileSize":105878,<br>          "logUrl":"https://abc.example.com/logd/v2/download/0621c8fc885089805kea5f610797ff8ba92bc98c049c2bb308cbdb?traceId=ac6d696c657765625f74657374cf0000018dd01d89e8cd06d3",<br>          "hostname":"mydomain.domain.info"<br>          }<br>        ]<br>}<br><br>The logUrl field contains a URL to a gzip-compressed log file. The url will expire in 24 hours. Please call the API to fetch a new url if it expires. The log file consists of rows such as the following:<br>9.8.7.6 - [31/Oct/2021:19:59:57 +0000] "GET http://mydomain.domain.info/i/js/tab.js HTTP/1.1" 304 529 "http://mydomain.domain.info/?q=downloads" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36" 76<br><br> | POST | /cdn/report/logDownload |