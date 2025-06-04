package livefetch

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type LiveFetchTaskRequest struct {
  // {'en':'url,Multiple values are separated by ","', 'zh_CN':'指要直播预拉取的频道集合，即url集合，URL必须是标准的http(s)://开头，且为具体的URL，不支持正则与目录，每次调用允许多个url，单次最多提交400条url。多个url时用引号和逗号隔开。'}
  Urls []*string `json:"urls,omitempty" xml:"urls,omitempty" require:"true" type:"Repeated"`
  // {"en":"view", "zh_CN":"要预热的地区；多个值用'，'分隔，单次调用最多同时传入10个view；允许为空，当view不传值时，表示预热所有区域；输入格式为：&bull; 假设要预热福建省，只需输入'福建'，如果想同时预热浙江省和福建省，只需输入'浙江，福建'；&bull; 如果要预热海外区域，则不需要传入运营商，只需要传入国家名，例如'奥地利''美国'等。注：由于数据的特殊性，如果想预热广东全省，需将广州和深圳同时传入，即需输入'广东，广州，深圳'。"}
  View *string `json:"view,omitempty" xml:"view,omitempty"`
  // {'en':'Operators,Such as China Telecom, China Unicom, etc.Multiple values are separated by ",",Allow to be empty.The default value is all operators', 'zh_CN':'预热的运营商，如电信、联通等；多个值用"，"分隔，允许为空，当carrier不传值时，表示预热所有运营商。'}
  Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
  // {'en':'0 start preftch 1 stop prefetch', 'zh_CN':'表示本次任务的操作类型，即传递给内容服务器一个操作的命令，可设置的数值有：0：开始预热，创建新的预热任务1：停止预热，停止已经在预热的任务'}
  Action *int32 `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s LiveFetchTaskRequest) String() string {
  return tea.Prettify(s)
}

func (s LiveFetchTaskRequest) GoString() string {
  return s.String()
}

func (s *LiveFetchTaskRequest) SetUrls(v []*string) *LiveFetchTaskRequest {
  s.Urls = v
  return s
}

func (s *LiveFetchTaskRequest) SetView(v string) *LiveFetchTaskRequest {
  s.View = &v
  return s
}

func (s *LiveFetchTaskRequest) SetCarrier(v string) *LiveFetchTaskRequest {
  s.Carrier = &v
  return s
}

func (s *LiveFetchTaskRequest) SetAction(v int32) *LiveFetchTaskRequest {
  s.Action = &v
  return s
}

type LiveFetchTaskResponse struct {
  // {'en':'result code', 'zh_CN':'表示任务创建的状态码，200表示成功，非200表示失败'}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {'en':'message', 'zh_CN':'表示任务提交后，系统的响应消息'}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s LiveFetchTaskResponse) String() string {
  return tea.Prettify(s)
}

func (s LiveFetchTaskResponse) GoString() string {
  return s.String()
}

func (s *LiveFetchTaskResponse) SetCode(v int32) *LiveFetchTaskResponse {
  s.Code = &v
  return s
}

func (s *LiveFetchTaskResponse) SetMessage(v string) *LiveFetchTaskResponse {
  s.Message = &v
  return s
}

type LiveFetchTaskPaths struct {
}

func (s LiveFetchTaskPaths) String() string {
  return tea.Prettify(s)
}

func (s LiveFetchTaskPaths) GoString() string {
  return s.String()
}

type LiveFetchTaskParameters struct {
}

func (s LiveFetchTaskParameters) String() string {
  return tea.Prettify(s)
}

func (s LiveFetchTaskParameters) GoString() string {
  return s.String()
}

type LiveFetchTaskRequestHeader struct {
}

func (s LiveFetchTaskRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveFetchTaskRequestHeader) GoString() string {
  return s.String()
}

type LiveFetchTaskResponseHeader struct {
}

func (s LiveFetchTaskResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s LiveFetchTaskResponseHeader) GoString() string {
  return s.String()
}




