package bot_scene_whitelist

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

func (c *Client) Add(request *AddSpecificClientTrafficBypassRequest) (requestId string, response *AddSpecificClientTrafficBypassResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp AddSpecificClientTrafficBypassResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/bot-manage/scene/whitelist/add", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) GetList(request *ListSpecificClientTrafficBypassRequest) (requestId string, response *ListSpecificClientTrafficBypassResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp ListSpecificClientTrafficBypassResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/bot-manage/scene/whitelist/list", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) Delete(request *DeleteSpecificClientTrafficBypassRequest) (requestId string, response *DeleteSpecificClientTrafficBypassResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp DeleteSpecificClientTrafficBypassResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/bot-manage/scene/whitelist/delete", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) Update(request *UpdateSpecificClientTrafficBypassRequest) (requestId string, response *UpdateSpecificClientTrafficBypassResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateSpecificClientTrafficBypassResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/bot-manage/scene/whitelist/update", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
