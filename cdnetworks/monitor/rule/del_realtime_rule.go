// This file is auto-generated, don't edit it. Thanks.
package rule

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteCloudMonitorRealTimeAlarmRuleRequest struct {
	// {"en":"Alert rule ID", "zh_CN":"报警规则id，需要鉴权是否该调用账号下或其下子账号的规则"}
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s DeleteCloudMonitorRealTimeAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCloudMonitorRealTimeAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudMonitorRealTimeAlarmRuleRequest) SetRuleId(v string) *DeleteCloudMonitorRealTimeAlarmRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteCloudMonitorRealTimeAlarmRuleResponse struct {
	// {"en":"Response code", "zh_CN":"功能码"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response message", "zh_CN":"功能码说明"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeleteCloudMonitorRealTimeAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCloudMonitorRealTimeAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudMonitorRealTimeAlarmRuleResponse) SetCode(v string) *DeleteCloudMonitorRealTimeAlarmRuleResponse {
	s.Code = &v
	return s
}

func (s *DeleteCloudMonitorRealTimeAlarmRuleResponse) SetMessage(v string) *DeleteCloudMonitorRealTimeAlarmRuleResponse {
	s.Message = &v
	return s
}
