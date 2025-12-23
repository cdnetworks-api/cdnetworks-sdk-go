# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/uploadassets
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/uploadassets"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &uploadassets.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := uploadassets.{ActionName}Response{}
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
| Getuploadtoken | Invoke getUploadToken to obtain the video file upload address and credentials | POST | /vod/videoManage/getUploadToken |
| Getaudiouploadtoken | You can invoke getAudioUploadToken to obtain audio file upload addresses and credentials. You can obtain a maximum of 50 audio file upload addresses and credentials in batches. | GET | /vod/audioManage/getAudioUploadToken |
| Getmaterialuploadtoken | Call getMaterialUploadToken to obtain the upload address and credentials of materials. You can batch obtain a maximum of 50 material addresses and credentials. | POST | /vod/material/getMaterialUploadToken |
| Pullvideo | Call pullVideo to set the url of the video to be pulled to the background. The background automatically pulls and saves url videos from third-party platforms. You can set pull tasks in batches. | POST | /vod/videoManage/pullVideo |
| Pullvideoquery | You can query the completion of pullVideoQuery. | POST | /vod/videoManage/pullVideoQuery |
| Getworkflowlist | Get WorkflowList | POST | /workflow/getWorkflowList |
| Gettranscodecombinelist | Get TranscodeCombine List | POST | /vod/transcodeManage/getTranscodeCombineList |
| Getsubtitlelist | Get Subtitle List | POST | /vod/material/getSubtitleList |
| Getdomainlistforcloudvuploadtool | Get Domain List | POST | /vod/domainManage/getDomainList |
| Getwatermarklist | get Watermark List | POST | /vod/watermarkManage/getWatermarkList |
| Getcategorylist | Get Category List | POST | /vod/categoryManage/getCategoryList |