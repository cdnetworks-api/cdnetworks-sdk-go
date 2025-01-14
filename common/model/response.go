package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
}

type ErrorResponse struct {
	Response struct {
		Error struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

func ParseFromHttpResponse(hr *http.Response, resp Response) error {
	switch hr.Header.Get("Content-Type") {
	default:
		return parseFromJson(hr, resp)
	}
}

func parseFromJson(hr *http.Response, resp Response) error {
	defer hr.Body.Close()
	body, err := io.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return NewCdnetworksSDKError("ClientError.IOError", msg, "")
	}
	if hr.StatusCode != 200 {
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return NewCdnetworksSDKError("ClientError.HttpStatusCodeError", msg, "")
	}
	err = resp.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return NewCdnetworksSDKError("ClientError.ParseJsonError", msg, "")
	}
	return nil
}
