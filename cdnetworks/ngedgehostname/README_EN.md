# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngedgehostname
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ngedgehostname"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &ngedgehostname.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := ngedgehostname.{ActionName}Response{}
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
| GetAnEdgeHostname | This API returns information about an edge hostname including the associated client zone rules and operation history. <br> | GET | /cdn/edgeHostnames/* |
| DeleteAnEdgeHostname | Use this API to delete an edge hostname. If you have configured your DNS server with this edge hostname, you should update your DNS records before deleting it. Otherwise, visitors to the associated hostnames will see an error. | DELETE | /cdn/edgeHostnames/* |
| UpdateAnEdgeHostnameAll | Use this API to update an edge hostname. Be sure to enter all fields describing your edge hostname including ones that are not changing. After a request is submitted, it takes a few minutes to deploy the edge hostname. You can call 'Get an edge hostname' and look at the history array in the response to check the deployment status. | PUT | /cdn/edgeHostnames/* |
| UpdateAnEdgeHostnamePart | This API allows you to update an edge hostname by sending only the fields that need to be changed. After a request is submitted, it takes a few minutes to deploy the edge hostname. You can call 'Get an edge hostname' and look at the history array in the response to check the deployment status. | PATCH | /cdn/edgeHostnames/* |
| CreateAnEdgeHostname | This API enables you to create edge hostnames to control how requests from different client zones are handled. You must create CNAME DNS records to point your properties' hostnames to an edge hostname created using this API in order for the CDN to handle your content. <br><br>After a request is submitted, it takes a few minutes to deploy the edge hostname. Please make sure the deployment is finished and the configurations take effect before pointing your hostnames to the edge hostname. You can call 'Get an edge hostname' and look at the history array in the response to check the deployment status. | POST | /cdn/edgeHostnames |
| GetAListOfEdgeHostnames | GetAListOfEdgeHostnames. Specify search, offset, and limit parameters to filter your results. The default is to return the most recently updated edge hostname first. | GET | /cdn/edgeHostnames |
| GetAListOfIsps | Obtain a list of ISPs that can be used in edge hostname client zone rules. | GET | /cdn/edgeHostnames/isps |
| Getclientregions | Get a list of client regions that can be used in edge hostname client zone rules. Query parameters allow you to filter and sort the results. | GET | /cdn/edgeHostnames/clientRegions |