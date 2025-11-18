package webhooks

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type DeleteAWebhookPaths struct {
  // {"en" : "ID of a webhook", "zh_CN": "webhook接口id。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s DeleteAWebhookPaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteAWebhookPaths) GoString() string {
  return s.String()
}

func (s *DeleteAWebhookPaths) SetId(v string) *DeleteAWebhookPaths {
  s.Id = &v
  return s
}

type DeleteAWebhookParameters struct {
}

func (s DeleteAWebhookParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteAWebhookParameters) GoString() string {
  return s.String()
}

type DeleteAWebhookRequestHeader struct {
}

func (s DeleteAWebhookRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAWebhookRequestHeader) GoString() string {
  return s.String()
}

type DeleteAWebhookRequest struct {
}

func (s DeleteAWebhookRequest) String() string {
  return tea.Prettify(s)
}

func (s DeleteAWebhookRequest) GoString() string {
  return s.String()
}

type DeleteAWebhookResponseHeader struct {
}

func (s DeleteAWebhookResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteAWebhookResponseHeader) GoString() string {
  return s.String()
}

type DeleteAWebhookResponse struct {
}

func (s DeleteAWebhookResponse) String() string {
  return tea.Prettify(s)
}

func (s DeleteAWebhookResponse) GoString() string {
  return s.String()
}




type UpdateAWebhookRequest struct {
  // {"en":"Range: <= 250 characters\nName of the webhook.","zh_CN":"取值范围: <= 250 字符\nwebhook接口名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"Range: <= 500 characters\nAn optional description of the webhook.","zh_CN":"取值范围: <= 500 字符\nwebhook接口描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Range: <= 250 characters\nThat URL that will be called when the asynchronous task has completed. The HTTP POST method will be used. The body will in JSON format.","zh_CN":"取值范围: <= 250 字符\n当关联的异步任务执行完成时，需触发的webhook接口的地址。CDN Pro将使用HTTP POST方法调webhook接口，请求体为JSON格式。"}
  Url *string `json:"url,omitempty" xml:"url,omitempty"`
  // {"en":"If your webhook endpoint requires authentication, the authentication method should be HTTP Basic authentication. You should provide credentials which we will use to generate the Authorization header.","zh_CN":"用于鉴权的账号信息。当您的服务器有鉴权要求时，需支持HTTP Basic鉴权方式。CDN Pro将用当前日期对secretKey进行加密，生成密码(password)。"}
  Credentials *UpdateAWebhookRequestCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Struct"`
}

func (s UpdateAWebhookRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdateAWebhookRequest) GoString() string {
  return s.String()
}

func (s *UpdateAWebhookRequest) SetName(v string) *UpdateAWebhookRequest {
  s.Name = &v
  return s
}

func (s *UpdateAWebhookRequest) SetDescription(v string) *UpdateAWebhookRequest {
  s.Description = &v
  return s
}

func (s *UpdateAWebhookRequest) SetUrl(v string) *UpdateAWebhookRequest {
  s.Url = &v
  return s
}

func (s *UpdateAWebhookRequest) SetCredentials(v *UpdateAWebhookRequestCredentials) *UpdateAWebhookRequest {
  s.Credentials = v
  return s
}

type UpdateAWebhookRequestCredentials struct {
  // {"en":"The username passed to the URL on your server.","zh_CN":"用于鉴权的用户名。"}
  User *string `json:"user,omitempty" xml:"user,omitempty"`
  // {"en":"A string that is encoded with the date and passed in the Authorization header to your server.","zh_CN":"用于鉴权的密钥。CDN Pro将用当期日期对密钥进行加密生成密码(password)，然后通过Authorization请求头传给你方服务器。"}
  SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
}

func (s UpdateAWebhookRequestCredentials) String() string {
  return tea.Prettify(s)
}

func (s UpdateAWebhookRequestCredentials) GoString() string {
  return s.String()
}

func (s *UpdateAWebhookRequestCredentials) SetUser(v string) *UpdateAWebhookRequestCredentials {
  s.User = &v
  return s
}

func (s *UpdateAWebhookRequestCredentials) SetSecretKey(v string) *UpdateAWebhookRequestCredentials {
  s.SecretKey = &v
  return s
}

type UpdateAWebhookRequestHeader struct {
}

func (s UpdateAWebhookRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAWebhookRequestHeader) GoString() string {
  return s.String()
}

type UpdateAWebhookPaths struct {
  // {"en":"ID of a webhook","zh_CN":"webhook接口id。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s UpdateAWebhookPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdateAWebhookPaths) GoString() string {
  return s.String()
}

func (s *UpdateAWebhookPaths) SetId(v string) *UpdateAWebhookPaths {
  s.Id = &v
  return s
}

type UpdateAWebhookParameters struct {
}

func (s UpdateAWebhookParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdateAWebhookParameters) GoString() string {
  return s.String()
}

type UpdateAWebhookResponse struct {
}

func (s UpdateAWebhookResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdateAWebhookResponse) GoString() string {
  return s.String()
}

type UpdateAWebhookResponseHeader struct {
}

func (s UpdateAWebhookResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdateAWebhookResponseHeader) GoString() string {
  return s.String()
}




type CreateAWebhookRequest struct {
  // {"en":"Range: <= 250 characters\nName of the webhook.","zh_CN":"取值范围: <= 250 字符\nwebhook接口名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"Range: <= 500 characters\nOptional description of the webhook.","zh_CN":"取值范围: <= 500 字符\nwebhook接口描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en":"Range: <= 250 characters\nThat URL that will be called when the asynchronous task has completed. The HTTP POST method will be used. The body will in JSON format:\n<code>{\"subject\":\"{subject}\",\"taskType\":\"{task type}\",\"taskList\":[{\"taskId\":\"{task id 1}\",...},{ \"taskId\":\"{task id 2}\",...}]}<code>","zh_CN":"取值范围: <= 250 字符\n当关联的异步任务执行完成时，需触发的webhook接口的地址。CDN Pro将使用HTTP POST方法调用webhook接口，请求体为JSON格式。请求体示例：\n<code>{\"subject\":\"{subject}\",\"taskType\":\"{task type}\",\"taskList\":[{\"taskId\":\"{task id 1}\",...},{ \"taskId\":\"{task id 2}\",...}]}<code>"}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en":"If your webhook endpoint requires authentication, the authentication method should be HTTP Basic authentication. You should provide credentials which we will use to generate the Authorization header. The example below show how the credentials will be used and how your webhook endpoint will be requested. \n\n<pre>\n#!/bin/bash\n# credentials you provide\nUSER=YourUser\nSECRET_KEY=YourSecretKey\n\nDATE=`date '+%a, %d %b %Y %H:%M:%S %Z'`\necho $DATE\n\n# generate password\npassword=$(echo -n '$DATE' | openssl dgst -sha1 -hmac '$SECRET_KEY' -binary | base64)\necho ' '\n\n# invoke webhook request\napiCall='curl -i --url 'https://a.webhook.com'\n-X POST \n-u $USER:$password \n-H 'Date: $DATE' \n-H 'Content-Type: application/json' \n-d '{\n...\n}' \n'\neval $apiCall\n</pre>","zh_CN":"用于鉴权的账号信息。当您的服务器有鉴权要求时，需支持HTTP Basic鉴权方式。CDN Pro将用当前日期对secretKey进行加密，生成密码(password)。以下是一个secretKey加密和webhook接口调用的示例：\n\n<pre>\n#!/bin/bash\n# 鉴权账号\nUSER=YourUser\nSECRET_KEY=YourSecretKey\n\nDATE=`date '+%a, %d %b %Y %H:%M:%S %Z'`\necho $DATE\n\n# 生成密码\npassword=$(echo -n '$DATE' | openssl dgst -sha1 -hmac '$SECRET_KEY' -binary | base64)\necho ' '\n\n# 调用webhook接口\napiCall='curl -i --url 'https://a.webhook.com'\n-X POST \n-u $USER:$password \n-H 'Date: $DATE' \n-H 'Content-Type: application/json' \n-d '{\n...\n}' \n'\neval $apiCall\n</pre>"}
  Credentials *CreateAWebhookRequestCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Struct"`
}

func (s CreateAWebhookRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateAWebhookRequest) GoString() string {
  return s.String()
}

func (s *CreateAWebhookRequest) SetName(v string) *CreateAWebhookRequest {
  s.Name = &v
  return s
}

func (s *CreateAWebhookRequest) SetDescription(v string) *CreateAWebhookRequest {
  s.Description = &v
  return s
}

func (s *CreateAWebhookRequest) SetUrl(v string) *CreateAWebhookRequest {
  s.Url = &v
  return s
}

func (s *CreateAWebhookRequest) SetCredentials(v *CreateAWebhookRequestCredentials) *CreateAWebhookRequest {
  s.Credentials = v
  return s
}

type CreateAWebhookRequestCredentials struct {
  // {"en":"The username passed to the URL on your server.","zh_CN":"用于鉴权的用户名。"}
  User *string `json:"user,omitempty" xml:"user,omitempty" require:"true"`
  // {"en":"A string that is encoded with the date and passed in the Authorization header to your server.","zh_CN":"用于鉴权的密钥。CDN Pro将用当期日期对密钥进行加密生成密码(password)，然后通过Authorization请求头传给你方服务器。"}
  SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty" require:"true"`
}

func (s CreateAWebhookRequestCredentials) String() string {
  return tea.Prettify(s)
}

func (s CreateAWebhookRequestCredentials) GoString() string {
  return s.String()
}

func (s *CreateAWebhookRequestCredentials) SetUser(v string) *CreateAWebhookRequestCredentials {
  s.User = &v
  return s
}

func (s *CreateAWebhookRequestCredentials) SetSecretKey(v string) *CreateAWebhookRequestCredentials {
  s.SecretKey = &v
  return s
}

type CreateAWebhookRequestHeader struct {
}

func (s CreateAWebhookRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAWebhookRequestHeader) GoString() string {
  return s.String()
}

type CreateAWebhookPaths struct {
}

func (s CreateAWebhookPaths) String() string {
  return tea.Prettify(s)
}

func (s CreateAWebhookPaths) GoString() string {
  return s.String()
}

type CreateAWebhookParameters struct {
}

func (s CreateAWebhookParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateAWebhookParameters) GoString() string {
  return s.String()
}

type CreateAWebhookResponse struct {
}

func (s CreateAWebhookResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateAWebhookResponse) GoString() string {
  return s.String()
}

type CreateAWebhookResponseHeader struct {
  // {"en":"Returns a URL pointing to the new webhook created, if the request is accepted. The URL contains the ID of the new webhook. </br> URL format: <code>{scheme}://{domain}/cdn/webhooks/{webhookId}</code> Example URL: <code>https://api.example.com/cdn/webhooks/60a2e69806d5583a81ad9500</code>","zh_CN":"当接口调用成功时，通过Location响应头返回新建的wehbook的地址。地址中包含webhook的ID，可使用该ID调用获取‘webhook详细信息’接口来查看webhook的详细信息。</br> URL格式：<code>{协议}://{域名}/cdn/webhooks/{webhookId}</code> URL示例： <code>https://open.chinanetcenter.com/cdn/webhooks/60a2e69806d5583a81ad9500</code>"}
  Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
}

func (s CreateAWebhookResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateAWebhookResponseHeader) GoString() string {
  return s.String()
}

func (s *CreateAWebhookResponseHeader) SetLocation(v string) *CreateAWebhookResponseHeader {
  s.Location = &v
  return s
}




type GetAWebhookPaths struct {
  // {"en" : "ID of a webhook", "zh_CN": "webhook接口id。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s GetAWebhookPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAWebhookPaths) GoString() string {
  return s.String()
}

func (s *GetAWebhookPaths) SetId(v string) *GetAWebhookPaths {
  s.Id = &v
  return s
}

type GetAWebhookParameters struct {
}

func (s GetAWebhookParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAWebhookParameters) GoString() string {
  return s.String()
}

type GetAWebhookRequestHeader struct {
}

func (s GetAWebhookRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAWebhookRequestHeader) GoString() string {
  return s.String()
}

type GetAWebhookRequest struct {
}

func (s GetAWebhookRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAWebhookRequest) GoString() string {
  return s.String()
}

type GetAWebhookResponseHeader struct {
}

func (s GetAWebhookResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAWebhookResponseHeader) GoString() string {
  return s.String()
}

type GetAWebhookResponse struct {
  // {"en" : "Range: <= 250 characters 
  // Name of the webhook.", "zh_CN": "取值范围: <= 250 字符 
  // webhook接口名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en" : "Range: <= 500 characters 
  // An optional description of the webhook.", "zh_CN": "取值范围: <= 500 字符 
  // webhook接口描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en" : "Range: <= 250 characters 
  // That URL that will be called when the asynchronous task has completed. The HTTP POST method will be used. The body will in JSON format:
  // 
  // { 'subject': '{some text}',
  //   'taskType': '{task type}',
  //   'taskList': [
  //     { 'taskId': '{task id 1}',
  //       ...
  //     },
  //     { 'taskId': '{task id 2}',
  //       ...
  //     }
  //   ]
  // }
  // 
  // ", "zh_CN": "取值范围: <= 250 字符 
  // 当关联的异步任务执行完成时，需触发的webhook接口的地址。CDN Pro将使用HTTP POST方法调webhook接口，请求体为JSON格式。请求体示例：
  // 
  // { 'subject': '{some text}',
  //   'taskType': '{task type}',
  //   'taskList': [
  //     { 'taskId': '{task id 1}',
  //       ...
  //     },
  //     { 'taskId': '{task id 2}',
  //       ...
  //     }
  //   ]
  // }
  // 
  // "}
  Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
  // {"en" : "Optional credentials passed to the URL. If requiring credentials your server should support HTTP Basic authentication the same way we do. Refer to Authentication summary. In particular, the password will be the secretKey encoded with the current date.", "zh_CN": "用于鉴权的账号信息。当您的服务器有鉴权要求时，需支持HTTP Basic鉴权方式。CDN Pro将用当前日期对secretKey进行加密，生成密码(password)。"}
  Credentials *GetAWebhookResponseCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" require:"true" type:"Struct"`
  // {"en" : "", "zh_CN": ""}
  MetaData *GetAWebhookResponseMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" require:"true" type:"Struct"`
}

func (s GetAWebhookResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAWebhookResponse) GoString() string {
  return s.String()
}

func (s *GetAWebhookResponse) SetName(v string) *GetAWebhookResponse {
  s.Name = &v
  return s
}

func (s *GetAWebhookResponse) SetDescription(v string) *GetAWebhookResponse {
  s.Description = &v
  return s
}

func (s *GetAWebhookResponse) SetUrl(v string) *GetAWebhookResponse {
  s.Url = &v
  return s
}

func (s *GetAWebhookResponse) SetCredentials(v *GetAWebhookResponseCredentials) *GetAWebhookResponse {
  s.Credentials = v
  return s
}

func (s *GetAWebhookResponse) SetMetaData(v *GetAWebhookResponseMetaData) *GetAWebhookResponse {
  s.MetaData = v
  return s
}

type GetAWebhookResponseCredentials struct {
  // {"en" : "The username passed to the URL on your server.", "zh_CN": "用于鉴权的用户名。"}
  User *string `json:"user,omitempty" xml:"user,omitempty" require:"true"`
  // {"en" : "A string that is encoded with the date and passed in the Authorization header to your server.", "zh_CN": "用于鉴权的密钥。CDN Pro将用当期日期对密钥进行加密生成密码(password)，然后通过Authorization请求头传给你方服务器。"}
  SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty" require:"true"`
}

func (s GetAWebhookResponseCredentials) String() string {
  return tea.Prettify(s)
}

func (s GetAWebhookResponseCredentials) GoString() string {
  return s.String()
}

func (s *GetAWebhookResponseCredentials) SetUser(v string) *GetAWebhookResponseCredentials {
  s.User = &v
  return s
}

func (s *GetAWebhookResponseCredentials) SetSecretKey(v string) *GetAWebhookResponseCredentials {
  s.SecretKey = &v
  return s
}

type GetAWebhookResponseMetaData struct {
  // {"en" : "RFC 3339 date indicating when the webhook was created.", "zh_CN": "webhook接口创建时间，以RFC 3339日期格式展示。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating when the webhook was last updated.", "zh_CN": "webhook接口最近一次更新的时间，以RFC 3339日期格式展示。"}
  LastUpdatetime *string `json:"lastUpdatetime,omitempty" xml:"lastUpdatetime,omitempty"`
  // {"en" : "RFC 3339 date indicating when the webhook was last called", "zh_CN": "webhook接口最近一次调用的时间，以RFC 3339日期格式展示。"}
  LastCallTime *string `json:"lastCallTime,omitempty" xml:"lastCallTime,omitempty"`
  // {"en" : "Range: >= 0 
  // Number of times the webhook has been called.", "zh_CN": "取值范围: >= 0 
  // webhook接口累计被调用的次数。"}
  TotalCalls *int `json:"totalCalls,omitempty" xml:"totalCalls,omitempty"`
}

func (s GetAWebhookResponseMetaData) String() string {
  return tea.Prettify(s)
}

func (s GetAWebhookResponseMetaData) GoString() string {
  return s.String()
}

func (s *GetAWebhookResponseMetaData) SetCreationTime(v string) *GetAWebhookResponseMetaData {
  s.CreationTime = &v
  return s
}

func (s *GetAWebhookResponseMetaData) SetLastUpdatetime(v string) *GetAWebhookResponseMetaData {
  s.LastUpdatetime = &v
  return s
}

func (s *GetAWebhookResponseMetaData) SetLastCallTime(v string) *GetAWebhookResponseMetaData {
  s.LastCallTime = &v
  return s
}

func (s *GetAWebhookResponseMetaData) SetTotalCalls(v int) *GetAWebhookResponseMetaData {
  s.TotalCalls = &v
  return s
}




type GetAListOfWebhooksPaths struct {
}

func (s GetAListOfWebhooksPaths) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfWebhooksPaths) GoString() string {
  return s.String()
}

type GetAListOfWebhooksParameters struct {
  // {"en" : "Range: >= 0 
  // Index of the first webhook to return.", "zh_CN": "默认值: 0 取值范围: >= 0 
  // 查询起始位置。"}
  Offset *int `json:"offset,omitempty" xml:"offset,omitempty"`
  // {"en" : "Default: 100 Range: [ 1 .. 200 ] 
  // Maximum number of webhooks to return.", "zh_CN": "默认值: 100 取值范围: <= 200 
  // 每次查询的最大条数。"}
  Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s GetAListOfWebhooksParameters) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfWebhooksParameters) GoString() string {
  return s.String()
}

func (s *GetAListOfWebhooksParameters) SetOffset(v int) *GetAListOfWebhooksParameters {
  s.Offset = &v
  return s
}

func (s *GetAListOfWebhooksParameters) SetLimit(v int) *GetAListOfWebhooksParameters {
  s.Limit = &v
  return s
}

type GetAListOfWebhooksRequestHeader struct {
}

func (s GetAListOfWebhooksRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfWebhooksRequestHeader) GoString() string {
  return s.String()
}

type GetAListOfWebhooksRequest struct {
}

func (s GetAListOfWebhooksRequest) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfWebhooksRequest) GoString() string {
  return s.String()
}

type GetAListOfWebhooksResponseHeader struct {
}

func (s GetAListOfWebhooksResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfWebhooksResponseHeader) GoString() string {
  return s.String()
}

type GetAListOfWebhooksResponse struct {
  // {"en" : "Range: >= 0 
  // Total number of webhooks. The number returned in the webhooks array may be fewer depending on your query parameters.", "zh_CN": "取值范围: >= 0 
  // webhook接口数量，该数量与查询参数直接相关。"}
  Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
  // {"en" : "", "zh_CN": ""}
  Webooks []*GetAListOfWebhooksResponseWebooks `json:"webooks,omitempty" xml:"webooks,omitempty" require:"true" type:"Repeated"`
}

func (s GetAListOfWebhooksResponse) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfWebhooksResponse) GoString() string {
  return s.String()
}

func (s *GetAListOfWebhooksResponse) SetCount(v int) *GetAListOfWebhooksResponse {
  s.Count = &v
  return s
}

func (s *GetAListOfWebhooksResponse) SetWebooks(v []*GetAListOfWebhooksResponseWebooks) *GetAListOfWebhooksResponse {
  s.Webooks = v
  return s
}

type GetAListOfWebhooksResponseWebooks struct     {
  // {"en" : "ID of the webhook.", "zh_CN": "webhook接口ID。"}
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  // {"en" : "Range: <= 250 characters 
  // Name of the webhook.", "zh_CN": "取值范围: <= 250 字符 
  // webhook接口名称。"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en" : "Range: <= 500 characters 
  // An optional description of the webhook.", "zh_CN": "取值范围: <= 500 字符 
  // webhook接口描述。"}
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // {"en" : "RFC 3339 date indicating when the webhook was created.", "zh_CN": "webhook接口创建时间，以RFC 3339日期格式展示。"}
  CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
  // {"en" : "RFC 3339 date indicating when the webhook was last updated.", "zh_CN": "webhook接口最近一次更新时间，以RFC 3339日期格式展示。"}
  LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
  // {"en" : "RFC 3339 date indicating when the webhook was last called.", "zh_CN": "webhook接口最近一次调用时间，以RFC 3339日期格式展示。"}
  LastCallTime *string `json:"lastCallTime,omitempty" xml:"lastCallTime,omitempty"`
  // {"en" : "Range: >= 0 
  // Total number of times the webhook has been called.", "zh_CN": "取值范围: >= 0 
  // webhook接口累计被调用的次数。"}
  TotalCalls *int `json:"totalCalls,omitempty" xml:"totalCalls,omitempty"`
  // {"en" : "Whether the webhook requires credentials to access.", "zh_CN": "是否要求鉴权。"}
  HasCredentials *bool `json:"hasCredentials,omitempty" xml:"hasCredentials,omitempty"`
}

func (s GetAListOfWebhooksResponseWebooks) String() string {
  return tea.Prettify(s)
}

func (s GetAListOfWebhooksResponseWebooks) GoString() string {
  return s.String()
}

func (s *GetAListOfWebhooksResponseWebooks) SetId(v string) *GetAListOfWebhooksResponseWebooks {
  s.Id = &v
  return s
}

func (s *GetAListOfWebhooksResponseWebooks) SetName(v string) *GetAListOfWebhooksResponseWebooks {
  s.Name = &v
  return s
}

func (s *GetAListOfWebhooksResponseWebooks) SetDescription(v string) *GetAListOfWebhooksResponseWebooks {
  s.Description = &v
  return s
}

func (s *GetAListOfWebhooksResponseWebooks) SetCreationTime(v string) *GetAListOfWebhooksResponseWebooks {
  s.CreationTime = &v
  return s
}

func (s *GetAListOfWebhooksResponseWebooks) SetLastUpdateTime(v string) *GetAListOfWebhooksResponseWebooks {
  s.LastUpdateTime = &v
  return s
}

func (s *GetAListOfWebhooksResponseWebooks) SetLastCallTime(v string) *GetAListOfWebhooksResponseWebooks {
  s.LastCallTime = &v
  return s
}

func (s *GetAListOfWebhooksResponseWebooks) SetTotalCalls(v int) *GetAListOfWebhooksResponseWebooks {
  s.TotalCalls = &v
  return s
}

func (s *GetAListOfWebhooksResponseWebooks) SetHasCredentials(v bool) *GetAListOfWebhooksResponseWebooks {
  s.HasCredentials = &v
  return s
}




