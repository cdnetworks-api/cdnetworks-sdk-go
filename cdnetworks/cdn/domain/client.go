package domain

import (
	"errors"
	common2 "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
	"log"
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

func (c *Client) AddCdnDomain(req *AddDomainForTerraformRequest) (requestId string, response *AddDomainForTerraformResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp AddDomainForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/domain", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryCdnDomain(domainName string) (response *QueryDomainForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}
	//invoke auth.Invoke
	var resp QueryDomainForTerraformResponse
	var requestId string
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/domain/"+domainName, "GET")

	requestId, err = auth.Invoke(config, nil, &resp)
	if err != nil {
		return nil, err
	}
	log.Printf("requestId: %s", requestId)
	return &resp, nil
}

func (c *Client) DeleteCdnDomain(domainName string) (requestId string, response *DeleteDomainForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/domain/"+domainName, "DELETE")
	requestId, err = auth.Invoke(config, nil, &response)

	if err != nil {
		return "", nil, err
	}

	return requestId, response, nil
}

func (c *Client) QueryDomainDeployStatus(xCncRequestId string) (response *QueryDeployResultForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return nil, errors.New("credential is required")
	}
	var resp QueryDeployResultForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/request/"+xCncRequestId, "GET")

	_, err = auth.Invoke(config, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) UpdateCdnDomain(request *UpdateDomainForTerraformRequest, domainName string) (requestId string, response *UpdateDomainForTerraformResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateDomainForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/domain/"+domainName, "PUT")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryCdnDomainList(request *QueryPagingDomainListForTerraformRequest) (requestId string, response *QueryPagingDomainListForTerraformResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryPagingDomainListForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/domain/list", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
