package rule

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

func (c *Client) CreateRealTimeRule(req *CreateCloudMonitorRealTimeAlarmRuleRequest) (requestId string, response *CreateCloudMonitorRealTimeAlarmRuleResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp CreateCloudMonitorRealTimeAlarmRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/cloudmonitor/alarm/real-time/add", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) EditRealTimeRule(req *EditCloudMonitorRealTimeAlarmRuleRequest) (requestId string, response *EditCloudMonitorRealTimeAlarmRuleResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp EditCloudMonitorRealTimeAlarmRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/cloudmonitor/alarm/real-time/edit", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) QueryRealTimeRule(req *QueryCloudMonitorRealTimeAlarmRuleRequest) (requestId string, response *QueryCloudMonitorRealTimeAlarmRuleResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp QueryCloudMonitorRealTimeAlarmRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/cloudmonitor/alarm/real-time/query", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}

func (c *Client) DeleteRealTimeRule(req *DeleteCloudMonitorRealTimeAlarmRuleRequest) (requestId string, response *DeleteCloudMonitorRealTimeAlarmRuleResponse, err error) {
	if req == nil {
		return "", nil, errors.New("request is required")
	}

	if c.GetCredential() == nil {
		return "", nil, errors.New("credential is required")
	}

	var resp DeleteCloudMonitorRealTimeAlarmRuleResponse
	config := auth.NewAkskConfig(c.GetCredential(), c.GetHttpProfile(), "/api/cloudmonitor/alarm/real-time/delete", "POST")
	requestId, err = auth.Invoke(config, req, &resp)

	if err != nil {
		return "", nil, err
	}

	return requestId, &resp, nil
}
