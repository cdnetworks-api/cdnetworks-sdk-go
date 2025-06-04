package cgmanage

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

func (c *Client) CreateControlGroup(req *CreateControlGroupRequest) (requestId string, response *CreateControlGroupResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp CreateControlGroupResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/user/control-groups", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryCustomizedControlGroupByName(req *QueryCustomizedControlGroupByNameRequest) (requestId string, response *QueryCustomizedControlGroupByNameResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryCustomizedControlGroupByNameResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/user/control-group/detail", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) EditControlGroupByCover(req *EditControlGroupByCoverRequest) (requestId string, response *EditControlGroupByCoverResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp EditControlGroupByCoverResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/user/control-group/full-update", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) DeleteControlGroup(paths *DeleteControlGroupPaths) (requestId string, response *DeleteControlGroupResponse, err error) {
	if paths == nil {
		return "", nil, errors.New("paths is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp DeleteControlGroupResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/user/control-groups/"+*paths.ControlGroupCode, "DELETE")
	requestId, err = auth.Invoke(config, nil, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryControlGroupList() (requestId string, response *QueryControlGroupListResponse, err error) {

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryControlGroupListResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/user/control-groups", "GET")
	requestId, err = auth.Invoke(config, nil, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
