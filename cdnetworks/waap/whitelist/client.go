package whitelist

import (
	"errors"
	common2 "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
)

type Client struct {
	common2.Client
}

func NewClient(credential common2.CredentialIface, httpProfile common2.HttpProfileIface) (client *Client, err error) {
	client = &Client{}
	client.WithCredential(credential)
	client.WithHttpProfile(httpProfile)
	return
}

func (c *Client) AddWaapWhitelistRule(request *CreateWhitelistRuleRequest) (requestId string, response *CreateWhitelistRuleResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp CreateWhitelistRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/whitelist/add", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) GetWaapWhitelistList(request *ListWhitelistRulesRequest) (requestId string, response *ListWhitelistRulesResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp ListWhitelistRulesResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/whitelist/query-list", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil

}

func (c *Client) DeleteWaapWhitelist(request *DeleteWhitelistRulesRequest) (requestId string, response *DeleteWhitelistRulesResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp DeleteWhitelistRulesResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/whitelist/delete", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) UpdateWaapWhitelist(request *UpdateWhitelistRuleRequest) (requestId string, response *UpdateWhitelistRuleResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateWhitelistRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/whitelist/update", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
