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
