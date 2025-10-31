package share_customizebot

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

func (c *Client) Add(request *AddShareCustomizeBotTFRequest) (requestId string, response *AddShareCustomizeBotTFResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp AddShareCustomizeBotTFResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/tf/share-customize-bots/add", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) GetList(request *ListShareCustomizeBotsRequest) (requestId string, response *ListShareCustomizeBotsResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp ListShareCustomizeBotsResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/share-customize-bots/get-list", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) Delete(request *DeleteShareCustomizeBotsRequest) (requestId string, response *DeleteShareCustomizeBotsResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp DeleteShareCustomizeBotsResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/share-customize-bots/delete", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) Update(request *UpdateShareCustomizeBotTFRequest) (requestId string, response *UpdateShareCustomizeBotTFResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateShareCustomizeBotTFResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/tf/share-customize-bots/update", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
