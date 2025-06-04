package model

import (
	"fmt"
)

type CdnetworksSDKError struct {
	Code      string
	Message   string
	RequestId string
}

func (e *CdnetworksSDKError) Error() string {
	if e.RequestId == "" {
		return fmt.Sprintf("[CdnetworksSDKError] Code=%s, Message=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("[CdnetworksSDKError] Code=%s, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewCdnetworksSDKError(code, message, requestId string) error {
	return &CdnetworksSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *CdnetworksSDKError) GetCode() string {
	return e.Code
}

func (e *CdnetworksSDKError) GetMessage() string {
	return e.Message
}

func (e *CdnetworksSDKError) GetRequestId() string {
	return e.RequestId
}
