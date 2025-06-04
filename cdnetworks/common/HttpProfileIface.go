package common

type HttpProfileIface interface {
	GetProtocol() string
	GetEndpoint() string
	GetServiceType() string
}

type HttpProfile struct {
	Protocol    string
	Endpoint    string
	ServiceType string
}

func NewHttpProfile(endpoint, protocol, serviceType string) *HttpProfile {
	return &HttpProfile{
		Protocol:    protocol,
		Endpoint:    endpoint,
		ServiceType: serviceType,
	}
}

func (c *HttpProfile) GetProtocol() string {
	return c.Protocol
}

func (c *HttpProfile) GetEndpoint() string {
	return c.Endpoint
}

func (c *HttpProfile) GetServiceType() string {
	return c.ServiceType
}
