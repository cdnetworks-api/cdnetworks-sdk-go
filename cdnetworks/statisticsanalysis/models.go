package statisticsanalysis

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetRealTimeChannelOnlineNumberRequest struct {
  // {"en":"Pull id, multiple values separated by \",\"", "zh_CN":"拉流 id，多个值通过\",\"隔开"}
  PullIds *string `json:"pullIds,omitempty" xml:"pullIds,omitempty"`
}

func (s GetRealTimeChannelOnlineNumberRequest) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberRequest) GoString() string {
  return s.String()
}

func (s *GetRealTimeChannelOnlineNumberRequest) SetPullIds(v string) *GetRealTimeChannelOnlineNumberRequest {
  s.PullIds = &v
  return s
}

type GetRealTimeChannelOnlineNumberResponse struct {
  // {"en":"200 success", "zh_CN":"200，操作成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Successful operation", "zh_CN":"操作成功"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  GetRealTimeChannelOnlineNumberData []*GetRealTimeChannelOnlineNumberData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s GetRealTimeChannelOnlineNumberResponse) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberResponse) GoString() string {
  return s.String()
}

func (s *GetRealTimeChannelOnlineNumberResponse) SetCode(v int32) *GetRealTimeChannelOnlineNumberResponse {
  s.Code = &v
  return s
}

func (s *GetRealTimeChannelOnlineNumberResponse) SetMessage(v string) *GetRealTimeChannelOnlineNumberResponse {
  s.Message = &v
  return s
}

func (s *GetRealTimeChannelOnlineNumberResponse) SetData(v []*GetRealTimeChannelOnlineNumberData) *GetRealTimeChannelOnlineNumberResponse {
  s.GetRealTimeChannelOnlineNumberData = v
  return s
}

type GetRealTimeChannelOnlineNumberData struct {
  // {"en":"Online list of people", "zh_CN":"在线人数列表"}
  OnlineNumberList []*GetRealTimeChannelOnlineNumberOnlineNumberItem `json:"onlineNumberList,omitempty" xml:"onlineNumberList,omitempty" require:"true" type:"Repeated"`
}

func (s GetRealTimeChannelOnlineNumberData) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberData) GoString() string {
  return s.String()
}

func (s *GetRealTimeChannelOnlineNumberData) SetOnlineNumberList(v []*GetRealTimeChannelOnlineNumberOnlineNumberItem) *GetRealTimeChannelOnlineNumberData {
  s.OnlineNumberList = v
  return s
}

type GetRealTimeChannelOnlineNumberOnlineNumberItem struct {
  // {"en":"Online population", "zh_CN":"在线人数"}
  OnlineNumber *int32 `json:"onlineNumber,omitempty" xml:"onlineNumber,omitempty" require:"true"`
  // {"en":"pullId", "zh_CN":"拉流 id"}
  PullId *string `json:"pullId,omitempty" xml:"pullId,omitempty" require:"true"`
}

func (s GetRealTimeChannelOnlineNumberOnlineNumberItem) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberOnlineNumberItem) GoString() string {
  return s.String()
}

func (s *GetRealTimeChannelOnlineNumberOnlineNumberItem) SetOnlineNumber(v int32) *GetRealTimeChannelOnlineNumberOnlineNumberItem {
  s.OnlineNumber = &v
  return s
}

func (s *GetRealTimeChannelOnlineNumberOnlineNumberItem) SetPullId(v string) *GetRealTimeChannelOnlineNumberOnlineNumberItem {
  s.PullId = &v
  return s
}

type GetRealTimeChannelOnlineNumberPaths struct {
}

func (s GetRealTimeChannelOnlineNumberPaths) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberPaths) GoString() string {
  return s.String()
}

type GetRealTimeChannelOnlineNumberParameters struct {
}

func (s GetRealTimeChannelOnlineNumberParameters) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberParameters) GoString() string {
  return s.String()
}

type GetRealTimeChannelOnlineNumberRequestHeader struct {
}

func (s GetRealTimeChannelOnlineNumberRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberRequestHeader) GoString() string {
  return s.String()
}

type GetRealTimeChannelOnlineNumberResponseHeader struct {
}

func (s GetRealTimeChannelOnlineNumberResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetRealTimeChannelOnlineNumberResponseHeader) GoString() string {
  return s.String()
}




