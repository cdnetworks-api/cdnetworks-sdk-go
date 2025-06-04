package share_customizerule

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

func (c *Client) Add(request *CreateSharedCustomRuleRequest) (requestId string, response *CreateSharedCustomRuleResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp CreateSharedCustomRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/tf/share-customize-rule/add", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) GetList(request *ListSharedCustomRulesRequest) (requestId string, response *ListSharedCustomRulesResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp ListSharedCustomRulesResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/share-customize-rule/get-list", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) Delete(request *DeleteSharedCustomRulesRequest) (requestId string, response *DeleteSharedCustomRulesResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp DeleteSharedCustomRulesResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/share-customize-rule/delete", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) Update(request *UpdateSharedCustomRulesRequest) (requestId string, response *UpdateSharedCustomRulesResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateSharedCustomRulesResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/tf/share-customize-rule/update", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
