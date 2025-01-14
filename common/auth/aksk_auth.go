package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common/constant"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common/model"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common/util"
	"io"
	"log"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type AkskConfig struct {
	AccessKey     string
	SecretKey     string
	Protocol      string
	EndPoint      string
	Uri           string
	Method        string
	SignedHeaders string
	CustomHeaders map[string]string //custom headers
}

func NewAkskConfig(credential common.CredentialIface, httpProfile common.HttpProfileIface, uri string, method string) (config AkskConfig) {
	config.AccessKey = credential.GetAccessKey()
	config.SecretKey = credential.GetSecretKey()

	config.Protocol = httpProfile.GetProtocol()
	if config.Protocol == "" {
		config.Protocol = constant.Https
	}
	config.Uri = uri
	config.EndPoint = httpProfile.GetEndpoint()
	if config.EndPoint == "" {
		config.EndPoint = constant.HttpRequestDomain
	}
	config.Method = method
	config.SignedHeaders = "content-type;host"
	config.CustomHeaders = map[string]string{}
	config.CustomHeaders[constant.ApiType] = constant.TypeTerraform

	if httpProfile.GetServiceType() != "" {
		config.CustomHeaders[constant.ServiceType] = httpProfile.GetServiceType()
	}

	return
}

func TransferRequestMsg(config AkskConfig) model.HttpRequestMsg {
	var requestMsg = model.HttpRequestMsg{Params: map[string]string{}, Headers: map[string]string{}}
	requestMsg.Uri = config.Uri
	requestMsg.Method = config.Method
	requestMsg.Url = constant.HttpRequestPrefix + config.Uri
	if len(config.EndPoint) == 0 || "{endPoint}" == config.EndPoint {
		requestMsg.Host = constant.HttpRequestDomain
		requestMsg.Url = constant.Https + "://" + constant.HttpRequestDomain + config.Uri
	} else {
		requestMsg.Host = config.EndPoint
		requestMsg.Url = config.Protocol + "://" + config.EndPoint + config.Uri
	}
	requestMsg.SignedHeaders = getSignedHeaders(config.SignedHeaders)
	return requestMsg
}

func Invoke(config AkskConfig, request interface{}, response interface{}) (reqeustId string, err error) {
	var requestMsg = TransferRequestMsg(config)

	if config.Method == "POST" || config.Method == "PUT" || config.Method == "PATCH" || config.Method == "DELETE" {
		requestMsg.Body = tea.Prettify(request)
	}

	timeStamp := getCurrentTimeSeconds()
	requestMsg.Headers[constant.HeadSignTimeStamp] = timeStamp
	requestMsg.Headers["Host"] = requestMsg.Host
	requestMsg.Headers[constant.ContentType] = constant.ApplicationJson
	requestMsg.Headers[constant.HeadSignAccessKey] = config.AccessKey
	requestMsg.Headers[constant.XCncAuthMethod] = constant.AKSK
	requestMsg.Headers[constant.ACCEPT] = constant.ApplicationJson
	for k, v := range config.CustomHeaders {
		requestMsg.Headers[k] = v
	}
	signature := getSignature(requestMsg, config.SecretKey, timeStamp)
	requestMsg.Headers[constant.Authorization] = genAuthorization(config.AccessKey, requestMsg.SignedHeaders, signature)

	resp, call_error := util.Call(requestMsg)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}(resp.Body)

	body, bodyReadError := io.ReadAll(resp.Body)

	if call_error != nil {
		log.Printf("Error: %v, response body: %s", err, body)
		return "", fmt.Errorf("unexpected HTTP status code: %d, response:%s", resp.StatusCode, body)
	}

	bodyReadError = json.Unmarshal(body, &response)
	if bodyReadError != nil {
		log.Printf("Unable to parse the response body : %v", err)
		return "", err
	}

	return resp.Header.Get("x-cnc-request-id"), nil
}

func getCurrentTimeSeconds() string {
	timeStamp := time.Now().UTC().Unix()
	return strconv.FormatInt(timeStamp, 10)
}

/*
*
拼接最后签名
*/
func genAuthorization(accessKey string, signedHeaders string, signature string) string {
	var build strings.Builder
	build.WriteString(constant.HeadSignAlgorithm)
	build.WriteString(" ")
	build.WriteString("Credential=")
	build.WriteString(accessKey)
	build.WriteString(", ")
	build.WriteString("SignedHeaders=")
	build.WriteString(signedHeaders)
	build.WriteString(", ")
	build.WriteString("Signature=")
	build.WriteString(signature)
	return build.String()
}

func getSignature(requestMsg model.HttpRequestMsg, secretKey string, timeStamp string) string {
	var bodyStr = requestMsg.Body
	if len(requestMsg.Body) == 0 || "GET" == requestMsg.Method {
		bodyStr = ""
	}
	hashedRequestPayload := hmacSha256(bodyStr)
	canonicalRequest := requestMsg.Method + "\n" +
		getRequestUri(requestMsg) + "\n" +
		getQueryString(requestMsg) + "\n" +
		getCanonicalHeaders(requestMsg.Headers, requestMsg.SignedHeaders) + "\n" +
		getSignedHeaders(requestMsg.SignedHeaders) + "\n" +
		hashedRequestPayload
	stringToSign := constant.HeadSignAlgorithm + "\n" + timeStamp + "\n" + hmacSha256(canonicalRequest)
	return hmac256(secretKey, stringToSign)
}

/**  获取uri  */
func getRequestUri(requestMsg model.HttpRequestMsg) string {
	indexOfQueryStringSeparator := strings.Index(requestMsg.Uri, "?")
	if indexOfQueryStringSeparator == -1 {
		return requestMsg.Uri
	}
	return string([]rune(requestMsg.Uri)[:indexOfQueryStringSeparator])
}

/*
*
获取uri参数
*/
func getQueryString(requestMsg model.HttpRequestMsg) string {
	indexOfQueryStringSeparator := strings.Index(requestMsg.Uri, "?")
	if "POST" == requestMsg.Method || indexOfQueryStringSeparator == -1 {
		return ""
	}
	s, err := url.QueryUnescape(requestMsg.Uri[indexOfQueryStringSeparator+1 : len(requestMsg.Uri)])
	if err != nil {
		fmt.Println("decode请求参数失败.")
	}
	return s
}

/*
*
获取并排序参与签名计算的头部
*/
func getSignedHeaders(signedHeaders string) string {
	if len(signedHeaders) == 0 {
		return "content-type;host"
	}
	headers := strings.Split(strings.ToLower(signedHeaders), ";")
	sort.Strings(headers)
	return strings.Join(headers, ";")
}

/*
*
获取k-v字符串
*/
func getCanonicalHeaders(headerMap map[string]string, signedHeaders string) string {
	keys := strings.Split(signedHeaders, ";")
	var headers = make(map[string]string)
	for k, v := range headerMap {
		headers[strings.ToLower(k)] = v
	}
	var build strings.Builder
	for i := 0; i < len(keys); i++ {
		build.WriteString(keys[i])
		build.WriteString(":")
		build.WriteString(strings.ToLower(headers[keys[i]]))
		build.WriteString("\n")
	}
	return build.String()
}

/*
*
加密算法
*/
func hmacSha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashCode := hash.Sum(nil)
	result := hex.EncodeToString(hashCode)
	return strings.ToLower(result)
}

func hmac256(secretKey string, stringToSign string) string {
	value := []byte(secretKey)
	key := hmac.New(sha256.New, value)
	key.Write([]byte(stringToSign))
	result := hex.EncodeToString(key.Sum(nil))
	return strings.ToLower(result)
}
