package usermanage

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

func (c *Client) BatchAddOrRevokePolicyToSubAccount(req *BatchAddOrRevokePolicyToSubAccountRequest) (requestId string, response *BatchAddOrRevokePolicyToSubAccountResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp BatchAddOrRevokePolicyToSubAccountResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/user/policies", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryPolicyAttachedMainAccountOrSubAccount(req *QueryPolicyAttachedMainAccountOrSubAccountPaths) (requestId string, response *QueryPolicyAttachedMainAccountOrSubAccountResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryPolicyAttachedMainAccountOrSubAccountResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/user/policy-attached/"+*req.LoginName, "GET")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) CreateUser(req *AddSubAccountRequest) (requestId string, response *AddSubAccountResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp AddSubAccountResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/sub-account", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) EditUser(req *UpdateSubAccountRequest) (requestId string, response *UpdateSubAccountResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}
	if *req.LoginName == "" {
		return "", nil, errors.New("login name is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateSubAccountResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/sub-account", "PUT")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryUser(req *QuerySubAccountInfoRequest, path *QuerySubAccountInfoPaths) (requestId string, response *QuerySubAccountInfoResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}
	if path == nil || *path.LoginName == "" {
		return "", nil, errors.New("login name is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QuerySubAccountInfoResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/sub-account/"+*path.LoginName, "GET")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) DeleteUser(req *DeleteSubAccountRequest, path *DeleteSubAccountPaths) (requestId string, response *DeleteSubAccountResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}
	if path == nil && *path.LoginName == "" {
		return "", nil, errors.New("login name is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp DeleteSubAccountResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/sub-account/"+*path.LoginName, "DELETE")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) ListUsers(req *GetSubAccountListRequest) (requestId string, response *GetSubAccountListResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}
	if req.PageIndex == nil || req.PageSize == nil {
		return "", nil, errors.New("page number and page size are required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp GetSubAccountListResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/sub-account/list", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
