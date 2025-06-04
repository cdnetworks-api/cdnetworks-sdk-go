package usagestatistics

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type WcsQpsRequest struct {
  // {"en":"bucket", "zh_CN":"空间名"}
  BucketList []*string `json:"bucketList,omitempty" xml:"bucketList,omitempty" type:"Repeated"`
  // {"en":"regionname", "zh_CN":"节点名称"}
  RegionNameList []*string `json:"regionNameList,omitempty" xml:"regionNameList,omitempty" type:"Repeated"`
  // {"en":"dateFrom", "zh_CN":"查询开始时间，例如：2024-12-12T10:00+08:00，包含开始时间"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty" require:"true"`
  // {"en":"dateTo", "zh_CN":"查询j结束时间，例如：2024-12-12T10:01+08:00，包含结束时间"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty" require:"true"`
  // {"en":"type", "zh_CN":"write：写, read：读, 不传则查读和写"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"timezone", "zh_CN":"返回数据的时区。 格式：+0N:00，-12<= n <= 12。如:+08:00。  默认+08:00"}
  Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s WcsQpsRequest) String() string {
  return tea.Prettify(s)
}

func (s WcsQpsRequest) GoString() string {
  return s.String()
}

func (s *WcsQpsRequest) SetBucketList(v []*string) *WcsQpsRequest {
  s.BucketList = v
  return s
}

func (s *WcsQpsRequest) SetRegionNameList(v []*string) *WcsQpsRequest {
  s.RegionNameList = v
  return s
}

func (s *WcsQpsRequest) SetDateFrom(v string) *WcsQpsRequest {
  s.DateFrom = &v
  return s
}

func (s *WcsQpsRequest) SetDateTo(v string) *WcsQpsRequest {
  s.DateTo = &v
  return s
}

func (s *WcsQpsRequest) SetType(v string) *WcsQpsRequest {
  s.Type = &v
  return s
}

func (s *WcsQpsRequest) SetTimezone(v string) *WcsQpsRequest {
  s.Timezone = &v
  return s
}

type WcsQpsResponse struct {
  // {"en":"code", "zh_CN":"错误码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"string", "zh_CN":"错误信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"", "zh_CN":""}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
  // {"en":"bucket", "zh_CN":"空间名"}
  Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
  // {"en":"regionname", "zh_CN":"节点名称"}
  RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
  // {"en":"writeDetail", "zh_CN":"写"}
  Detail []*string `json:"detail,omitempty" xml:"detail,omitempty" require:"true" type:"Repeated"`
  // {"en":"time", "zh_CN":"时间"}
  Time *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
  // {"en":"value", "zh_CN":"每分钟的读qps。保留两位小数，每分钟的qps=请求数/60"}
  ReadValue *float64 `json:"readValue,omitempty" xml:"readValue,omitempty" require:"true"`
  // {"en":"value", "zh_CN":"每分钟的写qps。保留两位小数，每分钟的qps=请求数/60"}
  WriteValue *float64 `json:"writeValue,omitempty" xml:"writeValue,omitempty" require:"true"`
}

func (s WcsQpsResponse) String() string {
  return tea.Prettify(s)
}

func (s WcsQpsResponse) GoString() string {
  return s.String()
}

func (s *WcsQpsResponse) SetCode(v string) *WcsQpsResponse {
  s.Code = &v
  return s
}

func (s *WcsQpsResponse) SetMessage(v string) *WcsQpsResponse {
  s.Message = &v
  return s
}

func (s *WcsQpsResponse) SetData(v []*string) *WcsQpsResponse {
  s.Data = v
  return s
}

func (s *WcsQpsResponse) SetBucket(v string) *WcsQpsResponse {
  s.Bucket = &v
  return s
}

func (s *WcsQpsResponse) SetRegionName(v string) *WcsQpsResponse {
  s.RegionName = &v
  return s
}

func (s *WcsQpsResponse) SetDetail(v []*string) *WcsQpsResponse {
  s.Detail = v
  return s
}

func (s *WcsQpsResponse) SetTime(v string) *WcsQpsResponse {
  s.Time = &v
  return s
}

func (s *WcsQpsResponse) SetReadValue(v float64) *WcsQpsResponse {
  s.ReadValue = &v
  return s
}

func (s *WcsQpsResponse) SetWriteValue(v float64) *WcsQpsResponse {
  s.WriteValue = &v
  return s
}

type WcsQpsPaths struct {
}

func (s WcsQpsPaths) String() string {
  return tea.Prettify(s)
}

func (s WcsQpsPaths) GoString() string {
  return s.String()
}

type WcsQpsParameters struct {
}

func (s WcsQpsParameters) String() string {
  return tea.Prettify(s)
}

func (s WcsQpsParameters) GoString() string {
  return s.String()
}

type WcsQpsRequestHeader struct {
}

func (s WcsQpsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s WcsQpsRequestHeader) GoString() string {
  return s.String()
}

type WcsQpsResponseHeader struct {
}

func (s WcsQpsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s WcsQpsResponseHeader) GoString() string {
  return s.String()
}




