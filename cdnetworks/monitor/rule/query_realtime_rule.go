// This file is auto-generated, don't edit it. Thanks.
package rule

import (
	"github.com/alibabacloud-go/tea/tea"
)

// 查询请求体
type QueryCloudMonitorRealTimeAlarmRuleRequest struct {
	// {"en":"Alert rule ID", "zh_CN":"报警规则id，需要鉴权是否该调用账号下或其下子账号的规则"}
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// {"en":"Alarm rule name, rule ID and rule name, either one, rule ID takes priority", "zh_CN":"报警规则名称，规则ID和规则名称二选一，规则ID优先"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s QueryCloudMonitorRealTimeAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudMonitorRealTimeAlarmRuleRequest) SetRuleId(v string) *QueryCloudMonitorRealTimeAlarmRuleRequest {
	s.RuleId = &v
	return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleRequest) SetRuleName(v string) *QueryCloudMonitorRealTimeAlarmRuleRequest {
	s.RuleName = &v
	return s
}

// 监控规则条件
type QueryConditionRule struct {
	// {"en":"Monitor item (BANDWIDTH, FLOW, REQUEST, etc.)", "zh_CN":"监控项（BANDWIDTH-带宽、FLOW-流量、REQUEST-请求数等）"}
	MonitorItem *string `json:"monitorItem,omitempty" xml:"monitorItem,omitempty" require:"true"`
	// {"en":"Condition type (MAX, MIN, UPRUSH, SLUMPED)", "zh_CN":"条件类型（MAX-大于、MIN-小于、UPRUSH-突增、SLUMPED-突降）"}
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
	// {"en":"Threshold value", "zh_CN":"监控项阈值"}
	Threshold *string `json:"threshold,omitempty" xml:"threshold,omitempty" require:"true"`
}

func (s QueryConditionRule) String() string {
	return tea.Prettify(s)
}

func (s QueryConditionRule) GoString() string {
	return s.String()
}

func (s *QueryConditionRule) SetMonitorItem(v string) *QueryConditionRule {
	s.MonitorItem = &v
	return s
}

func (s *QueryConditionRule) SetCondition(v string) *QueryConditionRule {
	s.Condition = &v
	return s
}

func (s *QueryConditionRule) SetThreshold(v string) *QueryConditionRule {
	s.Threshold = &v
	return s
}

// 规则项
type QueryRuleItem struct {
	// {"en":"Rule item ID", "zh_CN":"分时段子规则"}
	RuleItemId *string `json:"ruleItemId,omitempty" xml:"ruleItemId,omitempty" require:"true"`
	// {"en":"Start time (format: HH:00)", "zh_CN":"监控时段开始时间，格式：HH:00，示例：00:00"}
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
	// {"en":"End time (format: HH:59)", "zh_CN":"监控时段结束时间，格式：HH:59，示例：01:59"}
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
	// {"en":"Condition type (ANY/ALL)", "zh_CN":"告警需满足任意或全部条件，可选值：ANY-任意、ALL-全部"}
	ConditionType *string `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
	// {"en":"Monitoring period in minutes (1/5/10)", "zh_CN":"监控周期，周期单位为分钟（1-1分钟、5-5分钟、10-10分钟）"}
	Period *int32 `json:"period,omitempty" xml:"period,omitempty" require:"true"`
	// {"en":"Number of cycles (1/2/3/5/15/30)", "zh_CN":"持续几个周期满足条件告警（1/2/3/5/15/30）"}
	PeriodTimes *int32 `json:"periodTimes,omitempty" xml:"periodTimes,omitempty" require:"true"`
	// {"en":"Condition rules list", "zh_CN":"条件规则列表"}
	ConditionRules []*QueryConditionRule `json:"conditionRules,omitempty" xml:"conditionRules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRuleItem) String() string {
	return tea.Prettify(s)
}

func (s QueryRuleItem) GoString() string {
	return s.String()
}

func (s *QueryRuleItem) SetRuleItemId(v string) *QueryRuleItem {
	s.RuleItemId = &v
	return s
}

func (s *QueryRuleItem) SetStartTime(v string) *QueryRuleItem {
	s.StartTime = &v
	return s
}

func (s *QueryRuleItem) SetEndTime(v string) *QueryRuleItem {
	s.EndTime = &v
	return s
}

func (s *QueryRuleItem) SetConditionType(v string) *QueryRuleItem {
	s.ConditionType = &v
	return s
}

func (s *QueryRuleItem) SetPeriod(v int32) *QueryRuleItem {
	s.Period = &v
	return s
}

func (s *QueryRuleItem) SetPeriodTimes(v int32) *QueryRuleItem {
	s.PeriodTimes = &v
	return s
}

func (s *QueryRuleItem) SetConditionRules(v []*QueryConditionRule) *QueryRuleItem {
	s.ConditionRules = v
	return s
}

// 通知方式
type QueryNotice struct {
	// {"en":"Notice method (sms/email/robot/webhook)", "zh_CN":"通知方式（sms-短信、email-邮件、robot-机器人、webhook-webhook回调）"}
	NoticeMethod *string `json:"noticeMethod,omitempty" xml:"noticeMethod,omitempty" require:"true"`
	// {"en":"Notice object", "zh_CN":"告警通知对象，短信/邮件传联系人id，机器人传机器人id，webhook传地址"}
	NoticeObject *string `json:"noticeObject,omitempty" xml:"noticeObject,omitempty" require:"true"`
}

func (s QueryNotice) String() string {
	return tea.Prettify(s)
}

func (s QueryNotice) GoString() string {
	return s.String()
}

func (s *QueryNotice) SetNoticeMethod(v string) *QueryNotice {
	s.NoticeMethod = &v
	return s
}

func (s *QueryNotice) SetNoticeObject(v string) *QueryNotice {
	s.NoticeObject = &v
	return s
}

// 响应数据
type QueryMonitorRuleData struct {
	// {"en":"Alarm rule ID", "zh_CN":"报警规则id"}
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
	// {"en":"Alert rule name", "zh_CN":"报警规则名称"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Monitor product or dimension", "zh_CN":"按指定维度或指定已开通商品监控（userDimension-主账号维度、DG-域名组、CG-服务组、domainDimension-域名维度）"}
	MonitorProduct *string `json:"monitorProduct,omitempty" xml:"monitorProduct,omitempty" require:"true"`
	// {"en":"Resource type", "zh_CN":"监控资源类型"}
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty" require:"true"`
	// {"en":"Monitor resources list or ALL", "zh_CN":"监控的资源名称列表或ALL"}
	MonitorResources []*string `json:"monitorResources,omitempty" xml:"monitorResources,omitempty" require:"true" type:"Repeated"`
	// {"en":"Statistical method (CONSOLIDATED/SEPARATE)", "zh_CN":"统计方式（CONSOLIDATED-合并统计、SEPARATE-逐一统计）"}
	StatisticalMethod *string `json:"statisticalMethod,omitempty" xml:"statisticalMethod,omitempty" require:"true"`
	// {"en":"Alert frequency (0/2/5/10/15/20)", "zh_CN":"告警频率，每X分钟内产生的告警仅通知一次（0-仅首次告警、2/5/10/15/20分钟）"}
	AlertFrequency *int32 `json:"alertFrequency,omitempty" xml:"alertFrequency,omitempty" require:"true"`
	// {"en":"Restore notice (true/false)", "zh_CN":"告警消警是否通知（true-是、false-否）"}
	RestoreNotice *string `json:"restoreNotice,omitempty" xml:"restoreNotice,omitempty" require:"true"`
	// {"en":"Rule items list", "zh_CN":"规则项列表"}
	RuleItems []*QueryRuleItem `json:"ruleItems,omitempty" xml:"ruleItems,omitempty" require:"true" type:"Repeated"`
	// {"en":"Notices list", "zh_CN":"通知方式列表"}
	Notices []*QueryNotice `json:"notices,omitempty" xml:"notices,omitempty" require:"true" type:"Repeated"`
}

func (s QueryMonitorRuleData) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorRuleData) GoString() string {
	return s.String()
}

func (s *QueryMonitorRuleData) SetRuleId(v string) *QueryMonitorRuleData {
	s.RuleId = &v
	return s
}

func (s *QueryMonitorRuleData) SetRuleName(v string) *QueryMonitorRuleData {
	s.RuleName = &v
	return s
}

func (s *QueryMonitorRuleData) SetMonitorProduct(v string) *QueryMonitorRuleData {
	s.MonitorProduct = &v
	return s
}

func (s *QueryMonitorRuleData) SetResourceType(v string) *QueryMonitorRuleData {
	s.ResourceType = &v
	return s
}

func (s *QueryMonitorRuleData) SetMonitorResources(v []*string) *QueryMonitorRuleData {
	s.MonitorResources = v
	return s
}

func (s *QueryMonitorRuleData) SetStatisticalMethod(v string) *QueryMonitorRuleData {
	s.StatisticalMethod = &v
	return s
}

func (s *QueryMonitorRuleData) SetAlertFrequency(v int32) *QueryMonitorRuleData {
	s.AlertFrequency = &v
	return s
}

func (s *QueryMonitorRuleData) SetRestoreNotice(v string) *QueryMonitorRuleData {
	s.RestoreNotice = &v
	return s
}

func (s *QueryMonitorRuleData) SetRuleItems(v []*QueryRuleItem) *QueryMonitorRuleData {
	s.RuleItems = v
	return s
}

func (s *QueryMonitorRuleData) SetNotices(v []*QueryNotice) *QueryMonitorRuleData {
	s.Notices = v
	return s
}

// 响应体
type QueryCloudMonitorRealTimeAlarmRuleResponse struct {
	// {"en":"Response code", "zh_CN":"功能码"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response message", "zh_CN":"功能码说明"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data", "zh_CN":"响应数据"}
	Data *QueryMonitorRuleData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponse) SetCode(v string) *QueryCloudMonitorRealTimeAlarmRuleResponse {
	s.Code = &v
	return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponse) SetMessage(v string) *QueryCloudMonitorRealTimeAlarmRuleResponse {
	s.Message = &v
	return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponse) SetData(v *QueryMonitorRuleData) *QueryCloudMonitorRealTimeAlarmRuleResponse {
	s.Data = v
	return s
}
