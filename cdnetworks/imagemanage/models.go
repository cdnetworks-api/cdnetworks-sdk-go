package imagemanage

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeployVmpImagePreheatingRequest struct {
  // {"en":"Image ID", "zh_CN":"镜像id"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
  // {"en":"Name of preheating node", "zh_CN":"预热节点"}
  NodeNames []*string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" require:"true" type:"Repeated"`
}

func (s DeployVmpImagePreheatingRequest) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingRequest) GoString() string {
  return s.String()
}

func (s *DeployVmpImagePreheatingRequest) SetImageId(v string) *DeployVmpImagePreheatingRequest {
  s.ImageId = &v
  return s
}

func (s *DeployVmpImagePreheatingRequest) SetNodeNames(v []*string) *DeployVmpImagePreheatingRequest {
  s.NodeNames = v
  return s
}

type DeployVmpImagePreheatingResponse struct {
}

func (s DeployVmpImagePreheatingResponse) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingResponse) GoString() string {
  return s.String()
}

type DeployVmpImagePreheatingPaths struct {
}

func (s DeployVmpImagePreheatingPaths) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingPaths) GoString() string {
  return s.String()
}

type DeployVmpImagePreheatingParameters struct {
}

func (s DeployVmpImagePreheatingParameters) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingParameters) GoString() string {
  return s.String()
}

type DeployVmpImagePreheatingRequestHeader struct {
}

func (s DeployVmpImagePreheatingRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingRequestHeader) GoString() string {
  return s.String()
}

type DeployVmpImagePreheatingResponseHeader struct {
}

func (s DeployVmpImagePreheatingResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeployVmpImagePreheatingResponseHeader) GoString() string {
  return s.String()
}




type VMPQueryImageRequest struct {
}

func (s VMPQueryImageRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageRequest) GoString() string {
  return s.String()
}

type VMPQueryImageResponse struct {
  // {"en":"Image information array", "zh_CN":"镜像信息数组"}
  Images []*string `json:"images,omitempty" xml:"images,omitempty" require:"true" type:"Repeated"`
  // {"en":"Image unique ID, global unique", "zh_CN":"镜像唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"Mirror name", "zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Mirror belongs to the Lord.", "zh_CN":"镜像属主"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en":"Image creation time, such as: 2017-11-04 14:17:41", "zh_CN":"镜像创建时间，如：2017-11-04 14:17:41"}
  CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty" require:"true"`
  // {"en":"Image size in GB", "zh_CN":"镜像大小，单位是GB"}
  Size *string `json:"size,omitempty" xml:"size,omitempty" require:"true"`
  // {"en":"Mirror Status", "zh_CN":"镜像状态"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
}

func (s VMPQueryImageResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageResponse) GoString() string {
  return s.String()
}

func (s *VMPQueryImageResponse) SetImages(v []*string) *VMPQueryImageResponse {
  s.Images = v
  return s
}

func (s *VMPQueryImageResponse) SetId(v string) *VMPQueryImageResponse {
  s.Id = &v
  return s
}

func (s *VMPQueryImageResponse) SetName(v string) *VMPQueryImageResponse {
  s.Name = &v
  return s
}

func (s *VMPQueryImageResponse) SetType(v string) *VMPQueryImageResponse {
  s.Type = &v
  return s
}

func (s *VMPQueryImageResponse) SetCreatedAt(v string) *VMPQueryImageResponse {
  s.CreatedAt = &v
  return s
}

func (s *VMPQueryImageResponse) SetSize(v string) *VMPQueryImageResponse {
  s.Size = &v
  return s
}

func (s *VMPQueryImageResponse) SetState(v string) *VMPQueryImageResponse {
  s.State = &v
  return s
}

type VMPQueryImagePaths struct {
}

func (s VMPQueryImagePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImagePaths) GoString() string {
  return s.String()
}

type VMPQueryImageParameters struct {
  // {"en":"There can be multiple field names for sorting. The values are:
  // Name, createdat, type, size, state", "zh_CN":"排序的字段名称，可以有多个，取值：
  // name、createdAt、type、size、state"}
  SortKey *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
  // {"en":"Sorting direction must follow sortkey. Value:
  // Desc: descending, default
  // ASC: ascending order", "zh_CN":"排序方向，必须跟在sortKey后面出现，取值：
  // desc：降序，默认值
  // asc：升序"}
  SortDir *string `json:"sortDir,omitempty" xml:"sortDir,omitempty"`
  // {"en":"The number of items displayed on each page is 20 by default", "zh_CN":"每个页面显示条数，默认是20"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
  // {"en":"Query from the image ID specified by the marker", "zh_CN":"从marker指定的镜像id开始查询"}
  Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
  // {"en":"Mirror ID. A maximum of 100 IDS can be queried at a time. The IDs are separated by a half angle comma character ','.", "zh_CN":"镜像 ID。单次最多查询 100 条 ID，ID 之间用半角逗号字符','隔开。"}
  Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
  // {"en":"Mirror name", "zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"The image belongs to the master. Is it a public image or a user-defined image? Value:
  // 
  // Common: official image
  // 
  // Snapshot: user snapshot image
  // 
  // Custom: user defined image'", "zh_CN":"镜像属主，是公共镜像还是用户自定义镜像，取值：
  // COMMON：官方镜像
  // SNAPSHOT：用户快照镜像
  // CUSTOM：用户自定义镜像"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"Image status, value:
  // 
  // Active: available status
  // 
  // Building: Creating
  // 
  // Inactive: not available (such as creation failure, etc.)'", "zh_CN":"镜像状态，取值：
  // ACTIVE：可用状态
  // BUILDING：创建中
  // INACTIVE：不可用（如创建失败等）"}
  State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s VMPQueryImageParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageParameters) GoString() string {
  return s.String()
}

func (s *VMPQueryImageParameters) SetSortKey(v string) *VMPQueryImageParameters {
  s.SortKey = &v
  return s
}

func (s *VMPQueryImageParameters) SetSortDir(v string) *VMPQueryImageParameters {
  s.SortDir = &v
  return s
}

func (s *VMPQueryImageParameters) SetLimit(v int) *VMPQueryImageParameters {
  s.Limit = &v
  return s
}

func (s *VMPQueryImageParameters) SetMarker(v string) *VMPQueryImageParameters {
  s.Marker = &v
  return s
}

func (s *VMPQueryImageParameters) SetIds(v string) *VMPQueryImageParameters {
  s.Ids = &v
  return s
}

func (s *VMPQueryImageParameters) SetName(v string) *VMPQueryImageParameters {
  s.Name = &v
  return s
}

func (s *VMPQueryImageParameters) SetType(v string) *VMPQueryImageParameters {
  s.Type = &v
  return s
}

func (s *VMPQueryImageParameters) SetState(v string) *VMPQueryImageParameters {
  s.State = &v
  return s
}

type VMPQueryImageRequestHeader struct {
}

func (s VMPQueryImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageRequestHeader) GoString() string {
  return s.String()
}

type VMPQueryImageResponseHeader struct {
}

func (s VMPQueryImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPQueryImageResponseHeader) GoString() string {
  return s.String()
}




type QueryVmpImagePreheatingStateRequest struct {
}

func (s QueryVmpImagePreheatingStateRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateRequest) GoString() string {
  return s.String()
}

type QueryVmpImagePreheatingStateResponse struct {
  QueryVmpImagePreheatingStatePreHeatingInfo []*QueryVmpImagePreheatingStatePreHeatingInfo `json:"preHeatingInfo,omitempty" xml:"preHeatingInfo,omitempty" require:"true" type:"Repeated"`
}

func (s QueryVmpImagePreheatingStateResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateResponse) GoString() string {
  return s.String()
}

func (s *QueryVmpImagePreheatingStateResponse) SetPreHeatingInfo(v []*QueryVmpImagePreheatingStatePreHeatingInfo) *QueryVmpImagePreheatingStateResponse {
  s.QueryVmpImagePreheatingStatePreHeatingInfo = v
  return s
}

type QueryVmpImagePreheatingStatePreHeatingInfo struct {
  // {"en":"Node name", "zh_CN":"节点名称"}
  NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty" require:"true"`
  // {"en":"Preheating status: SUCCESS - preheated; FAIL - preheating failed; SENDING - preheating", "zh_CN":"预热状态：SUCCESS-已预热；FAIL-预热失败；SENDING-预热中"}
  State *string `json:"state,omitempty" xml:"state,omitempty" require:"true"`
}

func (s QueryVmpImagePreheatingStatePreHeatingInfo) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStatePreHeatingInfo) GoString() string {
  return s.String()
}

func (s *QueryVmpImagePreheatingStatePreHeatingInfo) SetNodeName(v string) *QueryVmpImagePreheatingStatePreHeatingInfo {
  s.NodeName = &v
  return s
}

func (s *QueryVmpImagePreheatingStatePreHeatingInfo) SetState(v string) *QueryVmpImagePreheatingStatePreHeatingInfo {
  s.State = &v
  return s
}

type QueryVmpImagePreheatingStatePaths struct {
  // {"en":"no", "zh_CN":"镜像id"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s QueryVmpImagePreheatingStatePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStatePaths) GoString() string {
  return s.String()
}

func (s *QueryVmpImagePreheatingStatePaths) SetId(v string) *QueryVmpImagePreheatingStatePaths {
  s.Id = &v
  return s
}

type QueryVmpImagePreheatingStateParameters struct {
}

func (s QueryVmpImagePreheatingStateParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateParameters) GoString() string {
  return s.String()
}

type QueryVmpImagePreheatingStateRequestHeader struct {
}

func (s QueryVmpImagePreheatingStateRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateRequestHeader) GoString() string {
  return s.String()
}

type QueryVmpImagePreheatingStateResponseHeader struct {
}

func (s QueryVmpImagePreheatingStateResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryVmpImagePreheatingStateResponseHeader) GoString() string {
  return s.String()
}




type VMPCreateImageRequest struct {
  // {"en":"Mirror name", "zh_CN":"镜像名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Virtual machine instance ID, optional parameter. This parameter and imagesrcurl must be one of two choices", "zh_CN":"虚拟机实例标识，可选参数，该参数与imageSrcUrl必须二选一"}
  InstanceUuid *string `json:"instanceUuid,omitempty" xml:"instanceUuid,omitempty"`
  // {"en":"Virtual machine image URL address, optional parameter. The parameter and instanceuuid must be one of two choices.", "zh_CN":"虚拟机镜像Url地址，可选参数，该参数与instanceUuid必须二选一"}
  ImageSrcUrl *string `json:"imageSrcUrl,omitempty" xml:"imageSrcUrl,omitempty"`
  // {"en":"MD5 value of virtual machine image, used with imagesrcurl", "zh_CN":"虚拟机镜像的md5值，与imageSrcUrl配合使用"}
  Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
  // {"en":"Operating system type, if the URL address is carried, this parameter is required; if the virtual machine is specified to be created, the ostype of the virtual machine shall prevail.
  // 
  // There are two values: windows and Linux'", "zh_CN":"操作系统类型，如果携带的是url地址时，该参数必填；如果是指定虚拟机创建，则以虚拟机的ostype为准。
  // 有2种取值：windows、linux"}
  Ostype *string `json:"ostype,omitempty" xml:"ostype,omitempty"`
  // {"en":"Minimum requirements for system disk, unit: GB. The system disk size of the selected template must be greater than or equal to this value, otherwise virtual machine creation fails", "zh_CN":"系统盘最小要求，单位是GB。选择的模板的系统盘大小必须大于等于该值，否则虚拟机创建失败"}
  MinDisk *int `json:"minDisk,omitempty" xml:"minDisk,omitempty"`
  // {"en":"Whether QEMU guest agent is enabled for the image, and password reset is supported for the opened image
  // 
  // There are two values: true and false'", "zh_CN":"镜像是否开启了qemu guest agent，有开启的镜像支持密码重置
  // 有2种取值：TRUE、FALSE"}
  QgaEnabled *string `json:"qgaEnabled,omitempty" xml:"qgaEnabled,omitempty"`
}

func (s VMPCreateImageRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageRequest) GoString() string {
  return s.String()
}

func (s *VMPCreateImageRequest) SetName(v string) *VMPCreateImageRequest {
  s.Name = &v
  return s
}

func (s *VMPCreateImageRequest) SetInstanceUuid(v string) *VMPCreateImageRequest {
  s.InstanceUuid = &v
  return s
}

func (s *VMPCreateImageRequest) SetImageSrcUrl(v string) *VMPCreateImageRequest {
  s.ImageSrcUrl = &v
  return s
}

func (s *VMPCreateImageRequest) SetMd5(v string) *VMPCreateImageRequest {
  s.Md5 = &v
  return s
}

func (s *VMPCreateImageRequest) SetOstype(v string) *VMPCreateImageRequest {
  s.Ostype = &v
  return s
}

func (s *VMPCreateImageRequest) SetMinDisk(v int) *VMPCreateImageRequest {
  s.MinDisk = &v
  return s
}

func (s *VMPCreateImageRequest) SetQgaEnabled(v string) *VMPCreateImageRequest {
  s.QgaEnabled = &v
  return s
}

type VMPCreateImageResponse struct {
  // {"en":"Image unique ID, global unique", "zh_CN":"镜像唯一标识，全局唯一"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s VMPCreateImageResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageResponse) GoString() string {
  return s.String()
}

func (s *VMPCreateImageResponse) SetId(v string) *VMPCreateImageResponse {
  s.Id = &v
  return s
}

type VMPCreateImagePaths struct {
}

func (s VMPCreateImagePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImagePaths) GoString() string {
  return s.String()
}

type VMPCreateImageParameters struct {
}

func (s VMPCreateImageParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageParameters) GoString() string {
  return s.String()
}

type VMPCreateImageRequestHeader struct {
}

func (s VMPCreateImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageRequestHeader) GoString() string {
  return s.String()
}

type VMPCreateImageResponseHeader struct {
}

func (s VMPCreateImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPCreateImageResponseHeader) GoString() string {
  return s.String()
}




type VMPRemoveImageRequest struct {
  // {"en":"Image unique identification", "zh_CN":"镜像唯一标识"}
  ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty" require:"true"`
}

func (s VMPRemoveImageRequest) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageRequest) GoString() string {
  return s.String()
}

func (s *VMPRemoveImageRequest) SetImageId(v string) *VMPRemoveImageRequest {
  s.ImageId = &v
  return s
}

type VMPRemoveImageResponse struct {
}

func (s VMPRemoveImageResponse) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageResponse) GoString() string {
  return s.String()
}

type VMPRemoveImagePaths struct {
}

func (s VMPRemoveImagePaths) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImagePaths) GoString() string {
  return s.String()
}

type VMPRemoveImageParameters struct {
}

func (s VMPRemoveImageParameters) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageParameters) GoString() string {
  return s.String()
}

type VMPRemoveImageRequestHeader struct {
}

func (s VMPRemoveImageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageRequestHeader) GoString() string {
  return s.String()
}

type VMPRemoveImageResponseHeader struct {
}

func (s VMPRemoveImageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s VMPRemoveImageResponseHeader) GoString() string {
  return s.String()
}




