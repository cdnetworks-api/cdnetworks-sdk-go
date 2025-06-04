package common

type CredentialIface interface {
	GetAccessKey() string
	GetSecretKey() string
	GetCredential() (string, string)
}

type Credential struct {
	SecretId  string
	SecretKey string
}

func (c *Credential) needRefresh() bool {
	return false
}

func (c *Credential) refresh() {
}

func NewCredential(secretId, secretKey string) *Credential {
	return &Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
	}
}

func (c *Credential) GetAccessKey() string {
	return c.SecretId
}

func (c *Credential) GetSecretKey() string {
	return c.SecretKey
}

func (c *Credential) GetCredential() (string, string) {
	return c.SecretId, c.SecretKey
}
