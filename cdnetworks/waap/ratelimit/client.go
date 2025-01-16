package ratelimit

import (
	"errors"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
)

type Client struct {
	common.Client
}

func NewClient(credential common.CredentialIface, httpProfile common.HttpProfileIface) (client *Client, err error) {
	client = &Client{}
	client.WithCredential(credential)
	client.WithHttpProfile(httpProfile)
	return
}

func (c *Client) AddRateLimit(request *CreatRateLimitingRuleRequest) (requestId string, response *CreatRateLimitingRuleResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp CreatRateLimitingRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/rate-limit/add-rule", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}

func (c *Client) GetRateLimitList(request *ListRateLimitingRulesRequest) (requestId string, response *ListRateLimitingRulesResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp ListRateLimitingRulesResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/rate-limit/get-rule-list", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}

func (c *Client) DeleteRateLimit(request *DeleteRateLimitingRulesRequest) (requestId string, response *DeleteRateLimitingRulesResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp DeleteRateLimitingRulesResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/rate-limit/delete-rule-by-ids", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}

func (c *Client) UpdateRateLimit(request *UpdateRateLimitingRuleRequest) (requestId string, response *UpdateRateLimitingRuleResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp UpdateRateLimitingRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/rate-limit/update-rule", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}
