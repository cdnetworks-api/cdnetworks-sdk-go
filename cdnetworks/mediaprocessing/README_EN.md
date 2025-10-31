# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/mediaprocessing
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/mediaprocessing"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &mediaprocessing.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := mediaprocessing.{ActionName}Response{}
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
| Videoclip | Call videoClip to cut the uploaded video interval and convert key parameters. | POST | /vod/videoManage/videoClip |
| Videoclipquery | Call videoClip to query the clipping result of the clipping task. | POST | /vod/videoManage/videoClipQuery |
| Videoconcat | Call videoConcat to splice multiple videos according to the specified sequence to generate a new video file. | POST | /vod/videoManage/videoConcat |
| Videoconcatquery | The videoConcatQuery command is used to query the result of the video stitching task | POST | /vod/videoManage/videoConcatQuery |
| Createclearadtask | Call createClearAdTask to create AI de-advertising task for the specified video, and the system will automatically conduct AI de-advertising operation for the video (value-added service function, if necessary, please contact customer service to open the value-added service of "AI de-advertising"). | POST | /vod/ai/createClearAdTask |
| Clearadtaskquery | Query AI clear AD task status and results | POST | /vod/ai/clearAdTaskQuery |
| Transcode | Use the transCode method to convert the specified video, mainly including transcoding the resolution and adding watermarks. | POST | /vod/videoManage/transCode |