# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/manageassets
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/manageassets"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &manageassets.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := manageassets.{ActionName}Response{}
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
| Getvideolist | Call the getVideoList client to get a list of uploaded videos (including: upload time, video name, video ID, video status, etc.) | POST | /vod/videoManage/getVideoList |
| Videoedit | The basic information about a single video can be edited by calling videoEdit. | POST | /vod/videoManage/videoEdit |
| Videoblock | Invoke videoBlock to mask and disable uploaded videos. When the video is disabled, you cannot watch it any more. | POST | /vod/videoManage/videoBlock |
| Videoenable | Call videoEnable to re-enable the disabled video. | POST | /vod/videoManage/videoEnable |
| Deletevideo | Delete the uploaded video by calling deleteVideo. | POST | /vod/videoManage/deleteVideo |
| Getmateriallist | Call getMaterialList to get the information list of uploaded material | POST | /vod/material/getMaterialList |
| Materialedit | Calling materialEdit allows you to edit the underlying information of a single material file. | POST | /vod/material/materialEdit |
| Deletematerial | Delete the uploaded material file by calling deleteVideo. | POST | /vod/material/deleteMaterial |
| Getaudiolist | Query the list of audio files and sort them in reverse order by upload time. Support paging to get a list | POST | /vod/audioManage/getAudioList |
| Editaudio | Edits the basic information that specifies a single tone frequency | POST | /vod/audioManage/editAudio |
| Deleteaudio | Deletes the specified audio through this interface. | POST | /vod/audioManage/deleteAudio |
| Getcategorylist | Get category list  | POST | /vod/videoCategory/getCategoryList |
| Createcategory | Create category | POST | /vod/videoCategory/createCategory |
| Deletecategory | Delete category | POST | /vod/videoCategory/deleteCategory |
| Deletevideobycategory | Delete video by category  | POST | /vod/videoCategory/deleteVideoByCategory |