package certificateapplication

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

func (c *Client) CreateCertificateApplication(req *CreateCertificateApplicationOrderForTerraformRequest) (requestId string, response *CreateCertificateApplicationOrderForTerraformResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificate/order/create", "POST")
	requestId, err = auth.Invoke(config, req, &response)
	if err != nil {
		return requestId, nil, err
	}
	return requestId, response, nil
}

func (c *Client) ListCertificateApplication(req *ListCertificateApplicationOrdersForTerraformRequest) (requestId string, response *ListCertificateApplicationOrdersForTerraformResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificate/order/list", "POST")
	requestId, err = auth.Invoke(config, req, &response)
	if err != nil {
		return requestId, nil, err
	}
	return requestId, response, nil
}

func (c *Client) GetCertificateApplicationDetail(req *GetCertificateApplicationOrderForTerraformRequest) (requestId string, response *GetCertificateApplicationOrderForTerraformResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificate/order/detail", "POST")
	requestId, err = auth.Invoke(config, req, &response)
	if err != nil {
		return requestId, nil, err
	}
	return requestId, response, nil
}

func (c *Client) CancelCertificateApplication(req *CancelCertificateApplicationOrderForTerraformRequest) (requestId string, response *CancelCertificateApplicationOrderForTerraformResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificate/order/cancel", "POST")
	requestId, err = auth.Invoke(config, req, &response)
	if err != nil {
		return requestId, nil, err
	}
	return requestId, response, nil
}
