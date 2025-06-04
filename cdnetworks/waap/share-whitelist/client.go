package share_whitelist

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

func (c *Client) AddWaapShareWhitelistRule(request *CreateShareWhitelistRuleRequest) (requestId string, response *CreateShareWhitelistRuleResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp CreateShareWhitelistRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/tf/share-whitelist/add", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) GetWaapShareWhitelistList(request *ListShareWhitelistRulesRequest) (requestId string, response *ListShareWhitelistRulesResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp ListShareWhitelistRulesResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/share-whitelist/get-list", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil

}

func (c *Client) DeleteWaapShareWhitelist(request *DeleteShareWhitelistRuleRequest) (requestId string, response *DeleteShareWhitelistRuleResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp DeleteShareWhitelistRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/share-whitelist/delete", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) UpdateWaapShareWhitelist(request *UpdateShareWhitelistRuleRequest) (requestId string, response *UpdateShareWhitelistRuleResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateShareWhitelistRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/tf/share-whitelist/update", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
