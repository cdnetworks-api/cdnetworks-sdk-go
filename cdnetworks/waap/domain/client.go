package domain

import (
	"errors"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
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

func (c *Client) AddDomain(request *AccessDomainRequest) (requestId string, response *AccessDomainResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp AccessDomainResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/tf/sys-domain-info/add", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}

func (c *Client) AddDomainByCopy(request *UsingExistingHostnameToAddNewHostnameRequest) (requestId string, response *UsingExistingHostnameToAddNewHostnameResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp UsingExistingHostnameToAddNewHostnameResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/sys-domain-info/add-refer-to-other-domain", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}

/*func (c *Client) GetDomainList(request *ListDomainCommonConfRequest) (requestId string, response *ListDomainCommonConfResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp ListDomainCommonConfResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/tf/sys-domain-info/get-common-conf-list", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}*/

func (c *Client) GetDomainList(request *ListDomainInfoRequest) (requestId string, response *ListDomainInfoResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp ListDomainInfoResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/sys-domain-info/get-list", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil

}

/*func (c *Client) UpdateDomain(request *UpdateDomainCommonConfRequest) (requestId string, response *UpdateDomainCommonConfResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp UpdateDomainCommonConfResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/tf/sys-domain-info/update-common-conf", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil

}*/

func (c *Client) UpdateDomainPolicy(request *ModifyPolicyStatusRequest) (requestId string, response *ModifyPolicyStatusResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	var resp ModifyPolicyStatusResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/sys-domain-info/update-switch", "POST")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}

func (c *Client) DeleteDomain(request *RemoveProtectedHostnameParameters) (requestId string, response *RemoveProtectedHostnameResponse, err error) {
	if request == nil {
		return "", nil, errors.New("request is required")
	}
	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}
	domain := *request.Domain
	var resp RemoveProtectedHostnameResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/v1/common/sys-domain-info/remove?domain="+domain, "GET")
	requestId, err = auth.Invoke(config, request, &resp)
	if err != nil {
		return "", nil, err
	}
	return requestId, &resp, nil
}
