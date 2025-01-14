package common

import "net/http"

type Client struct {
	httpClient  *http.Client
	credential  CredentialIface
	httpProfile HttpProfileIface
}

func (c *Client) GetCredential() CredentialIface {
	return c.credential
}

func (c *Client) GetHttpProfile() HttpProfileIface {
	return c.httpProfile
}

func (c *Client) WithHttpProfile(profile HttpProfileIface) *Client {
	c.httpProfile = profile
	return c
}

func (c *Client) WithCredential(cred CredentialIface) *Client {
	c.credential = cred
	return c
}
