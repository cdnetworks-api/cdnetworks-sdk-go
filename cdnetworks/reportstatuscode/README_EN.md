# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportstatuscode
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/reportstatuscode"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &reportstatuscode.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := reportstatuscode.{ActionName}Response{}
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
| Querystatuscodedistributionofeachispandprovince | Query the status code distribution for multiple domains across various ISPs and provinces. Users provide a time range and a list of domains, with optional selection of ISP and province. Results display the status code distribution by domain, ISP, and province, supporting queries at 1-minute/5-minute/1-hour granularity. The API supports the `Accept-Language` request header, with `zh-CN` and `en-US` as the only options; `zh-CN` is the default. When `en-US` is selected, provinces and ISPs are shown in code, otherwise in Chinese. | POST | /api/report/status-code/isp-province |
| Querystatuscodedistribution | This interface is used to count the edge status code data of multiple domain names. Users can select the query time range and domain name list to obtain data. The returned content includes the status code distribution of each domain name and its corresponding number of requests. It helps users analyze the domain name access status and optimize content distribution strategies. | POST | /api/report/status-code |
| Queryoriginstatuscodedistribution | Statistics of the distribution of origin status codes of multiple domains, and the statistics contain the origin data of all nodes. Around 5-15 minutes of data delay. It's recommended that the call frequency is no higher than 30/5min. | POST | /api/report/status-code/origin |
| Queryipv6statusofeachispandprovince | Query IPV6 status codes for multiple domains across provinces and ISPs. Users can specify domains, provinces, ISPs, and IP protocol types, with time granularity of 5 minutes or 1 hour. It returns status codes and request counts for each domain by province and ISP, aiding in monitoring website access conditions by IP protocol.<br>Supports the Accept-Language header with "zh-CN" and "en-US" options; default is "zh-CN". When set to "en-US", provinces and operators use codes for input and output; otherwise, they are in Chinese. | POST | /api/report/statusCode/isp-province/ipv6 |
| Queryrealtimeoriginstatuscode | This interface is used to query the back-to-origin request status code every minute within a specified time range. Users need to provide a time interval and a domain name list, and can select query dimensions, such as specific status codes or status code types (success, redirection, server error, etc.). The returned data includes detailed information on the request time and total number of requests. This interface is used to monitor the health status of the website, optimize the response of the origin station, and help analyze and adjust the network architecture to improve service stability and visitor experience. | POST | /api/report/status-code/real-time/origin/total |
| Queryispprovincestatuscode | Querying the minutes-level status code summary of the provincial ISP.<br>Support language request header Accept Language, only support zh-CN and en-US, default to zh-CN. Accept Language: en-US, both the province and isp input and return are in code, otherwise the return is in Chinese. | POST | /api/report/status-code/isp-province/total |
| Reportstatuscoderealtimeedgeservice | Query minute-level data of edge status codes. Users can specify the start time, end time, and domain name to query the request count data corresponding to the status codes. The data granularity can be selected as either 1 minute or 5 minutes in the request. The return result includes the number of requests per minute for each status code. | POST | /api/report/status-code/real-time/edge |
| Querymultidomainsipv6oripv4statuscode | ReportStatusCodeIpTypeService | POST | /api/report/status-code/ipv6 |
| Reportstatuscodelogservice | Query the hourly granular status code distribution of multiple domain names | POST | /api/report/status-code/log |
| Querystatuscodedistributionincountries | Query status code distribution by country granularity (Statistics by CDN IP) | POST | /api/report/status-code/country |