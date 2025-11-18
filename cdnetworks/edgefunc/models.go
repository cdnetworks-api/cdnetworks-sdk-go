package edgefunc

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UploadFuncCodeRequest struct {
  // {"en":"func name","zh_CN":"函数名称"}
  FuncName *string `json:"FuncName,omitempty" xml:"FuncName,omitempty"`
  // {"en":"commit message","zh_CN":"提交信息"}
  CommitMessage *string `json:"CommitMessage,omitempty" xml:"CommitMessage,omitempty"`
  // {"en":"func file list","zh_CN":"函数文件列表"}
  Files []*UploadFuncCodeRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
}

func (s UploadFuncCodeRequest) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequest) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeRequest) SetFuncName(v string) *UploadFuncCodeRequest {
  s.FuncName = &v
  return s
}

func (s *UploadFuncCodeRequest) SetCommitMessage(v string) *UploadFuncCodeRequest {
  s.CommitMessage = &v
  return s
}

func (s *UploadFuncCodeRequest) SetFiles(v []*UploadFuncCodeRequestFiles) *UploadFuncCodeRequest {
  s.Files = v
  return s
}

type UploadFuncCodeRequestFiles struct     {
  // {"en":"FileName","zh_CN":"文件名称"}
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // {"en":"File code (base64 encoding)","zh_CN":"文件所含代码（Base64 编码）"}
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UploadFuncCodeRequestFiles) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequestFiles) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeRequestFiles) SetFileName(v string) *UploadFuncCodeRequestFiles {
  s.FileName = &v
  return s
}

func (s *UploadFuncCodeRequestFiles) SetCode(v string) *UploadFuncCodeRequestFiles {
  s.Code = &v
  return s
}

type UploadFuncCodeRequestHeader struct {
}

func (s UploadFuncCodeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeRequestHeader) GoString() string {
  return s.String()
}

type UploadFuncCodePaths struct {
}

func (s UploadFuncCodePaths) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodePaths) GoString() string {
  return s.String()
}

type UploadFuncCodeParameters struct {
}

func (s UploadFuncCodeParameters) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeParameters) GoString() string {
  return s.String()
}

type UploadFuncCodeResponse struct {
  // {"en":"code","zh_CN":"状态码"}
  Code *string `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
  // {"en":"message","zh_CN":"状态码描述"}
  Message *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
}

func (s UploadFuncCodeResponse) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeResponse) GoString() string {
  return s.String()
}

func (s *UploadFuncCodeResponse) SetCode(v string) *UploadFuncCodeResponse {
  s.Code = &v
  return s
}

func (s *UploadFuncCodeResponse) SetMessage(v string) *UploadFuncCodeResponse {
  s.Message = &v
  return s
}

type UploadFuncCodeResponseHeader struct {
}

func (s UploadFuncCodeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UploadFuncCodeResponseHeader) GoString() string {
  return s.String()
}




