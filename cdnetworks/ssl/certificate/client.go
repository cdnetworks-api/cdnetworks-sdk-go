package certificate

import (
	"errors"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
	"strconv"
)

type Client struct {
	common.Client
}

func NewClient(credential common.CredentialIface, httpProfile common.HttpProfileIface) (client *Client, err error) {
	client = &Client{}
	client.WithCredential(credential)
	client.WithHttpProfile(httpProfile)
	return
}

func (c *Client) AddCertificate(req *AddCertificateForTerraformRequest) (requestId string, response *AddCertificateForTerraformResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp AddCertificateForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificate", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return requestId, nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryCertificate(certificateId int64) (requestId string, response *QueryCertificateForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp QueryCertificateForTerraformResponse
	var certificateIdStr = strconv.FormatInt(certificateId, 10)
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificate/"+certificateIdStr, "GET")

	requestId, err = auth.Invoke(config, nil, &resp)
	if err != nil {
		return requestId, nil, err
	}
	return requestId, &resp, nil
}

func (c *Client) DeleteCertificate(certificateId int64) (requestId string, response *DeleteCertificateForTerraformResponse, err error) {
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var certificateIdStr = strconv.FormatInt(certificateId, 10)
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificate/"+certificateIdStr, "DELETE")
	requestId, err = auth.Invoke(config, nil, &response)

	if err != nil {
		return requestId, nil, err
	}

	return requestId, response, nil
}

func (c *Client) UpdateCertificate(certificateId int64, request *UpdateCertificateForTerraformRequest) (requestId string, response *UpdateCertificateForTerraformResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp UpdateCertificateForTerraformResponse
	var certificateIdStr = strconv.FormatInt(certificateId, 10)
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificate/"+certificateIdStr, "PUT")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return requestId, nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryCertificateList(request *QueryCertificateListForTerraformRequest) (requestId string, response *QueryCertificateListForTerraformResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryCertificateListForTerraformResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/terraform/certificates", "POST")

	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return requestId, nil, err
	}

	return requestId, &resp, nil
}
