// This file is auto-generated, don't edit it. Thanks.
package rule

import (
	"github.com/alibabacloud-go/tea/tea"
)

// 监控规则条件
type EditRuleCondition struct {
	// {"en":"Monitoring items currently only support domain-type monitoring items. Options: BANDWIDTH, FLOW, REQUEST, BTOB, FLOW_HIT_RATE, etc.", "zh_CN":"当前仅支持监控对象类型为域名的监控项。可选监控项：BANDWIDTH-带宽、FLOW-流量、REQUEST-请求数、BTOB-回源带宽、FLOW_HIT_RATE-流量命中率等"}
	MonitorItem *string `json:"monitorItem,omitempty" xml:"monitorItem,omitempty" require:"true"`
	// {"en":"Condition type. Options: MAX-greater than, MIN-less than, UPRUSH-surge, SLUMPED-plunge", "zh_CN":"条件类型。可选值：MAX-大于、MIN-小于、UPRUSH-突增、SLUMPED-突降"}
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
	// {"en":"Threshold value, please provide a positive integer", "zh_CN":"监控项阈值，请传递正整数"}
	Threshold *string `json:"threshold,omitempty" xml:"threshold,omitempty" require:"true"`
}

func (s EditRuleCondition) String() string {
	return tea.Prettify(s)
}

func (s EditRuleCondition) GoString() string {
	return s.String()
}

func (s *EditRuleCondition) SetMonitorItem(v string) *EditRuleCondition {
	s.MonitorItem = &v
	return s
}

func (s *EditRuleCondition) SetCondition(v string) *EditRuleCondition {
	s.Condition = &v
	return s
}

func (s *EditRuleCondition) SetThreshold(v string) *EditRuleCondition {
	s.Threshold = &v
	return s
}

// 规则项
type EditRuleItem struct {
	// {"en":"Time-divided sub-rules. If an id is passed in, it is considered to specify a modification of an item rule.", "zh_CN":"分时段子规则id"}
	RuleItemId *string `json:"ruleItemId,omitempty" xml:"ruleItemId,omitempty" require:"true"`
	// {"en":"Monitoring period start time, format: HH:00, example: 00:00", "zh_CN":"监控时段开始时间，格式：HH:00，示例：00:00"}
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
	// {"en":"Monitoring period end time, format: HH:59, example: 01:59", "zh_CN":"监控时段结束时间，格式：HH:59，示例：01:59"}
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
	// {"en":"Alert condition type. Options: ANY-any condition, ALL-all conditions", "zh_CN":"告警需满足任意或全部条件。可选值：ANY-任意、ALL-全部"}
	ConditionType *string `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
	// {"en":"Monitoring cycle in minutes. Options: 1, 5, 10", "zh_CN":"监控周期，周期单位为分钟。可选值：1-1分钟、5-5分钟、10-10分钟"}
	Period *int32 `json:"period,omitempty" xml:"period,omitempty" require:"true"`
	// {"en":"Number of cycles to meet conditions before alerting. Options: 1, 2, 3, 5, 15, 30", "zh_CN":"持续几个周期满足条件告警。可选值：1、2、3、5、15、30"}
	PeriodTimes *int32 `json:"periodTimes,omitempty" xml:"periodTimes,omitempty" require:"true"`
	// {"en":"Condition rules list", "zh_CN":"条件规则列表"}
	ConditionRules []*EditRuleCondition `json:"conditionRules,omitempty" xml:"conditionRules,omitempty" require:"true" type:"Repeated"`
}

func (s EditRuleItem) String() string {
	return tea.Prettify(s)
}

func (s EditRuleItem) GoString() string {
	return s.String()
}

func (s *EditRuleItem) SetRuleItemId(v string) *EditRuleItem {
	s.RuleItemId = &v
	return s
}

func (s *EditRuleItem) SetStartTime(v string) *EditRuleItem {
	s.StartTime = &v
	return s
}

func (s *EditRuleItem) SetEndTime(v string) *EditRuleItem {
	s.EndTime = &v
	return s
}

func (s *EditRuleItem) SetConditionType(v string) *EditRuleItem {
	s.ConditionType = &v
	return s
}

func (s *EditRuleItem) SetPeriod(v int32) *EditRuleItem {
	s.Period = &v
	return s
}

func (s *EditRuleItem) SetPeriodTimes(v int32) *EditRuleItem {
	s.PeriodTimes = &v
	return s
}

func (s *EditRuleItem) SetConditionRules(v []*EditRuleCondition) *EditRuleItem {
	s.ConditionRules = v
	return s
}

// 通知方式
type EditRuleNotice struct {
	// {"en":"Notification method. Options: MOBILE, EMAIL, robot, webhook", "zh_CN":"通知方式。可选值：MOBILE-短信、EMAIL-邮件、ROBOT-机器人、WEBHOOK-webhook回调"}
	NoticeMethod *string `json:"noticeMethod,omitempty" xml:"noticeMethod,omitempty" require:"true"`
	// {"en":"Notification object. For MOBILE/EMAIL: contact IDs separated by ;. For ROBOT: robot IDs separated by ;. For WEBHOOK: webhook URL", "zh_CN":"告警通知对象。若noticeMethod为MOBILE、EMAIL：请传递联系人id，多个用;分隔。若为ROBOT：请传递机器人id，多个用;分隔。若为WEBHOOK：请直接传递webhook地址"}
	NoticeObject *string `json:"noticeObject,omitempty" xml:"noticeObject,omitempty" require:"true"`
}

func (s EditRuleNotice) String() string {
	return tea.Prettify(s)
}

func (s EditRuleNotice) GoString() string {
	return s.String()
}

func (s *EditRuleNotice) SetNoticeMethod(v string) *EditRuleNotice {
	s.NoticeMethod = &v
	return s
}

func (s *EditRuleNotice) SetNoticeObject(v string) *EditRuleNotice {
	s.NoticeObject = &v
	return s
}

// 请求体
type EditCloudMonitorRealTimeAlarmRuleRequest struct {
	// {"en":"Alarm rule ID", "zh_CN":"报警规则id"}
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
	// {"en":"Alert rule name, only supports Chinese, English, numbers, underscore, hyphen, max 20 characters", "zh_CN":"报警规则名称，仅支持中英文、数字、下划线、中划线，最多20个字符"}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {"en":"Monitor by dimension or product. Dimensions: userDimension, DG, CG, domainDimension. Or input product code", "zh_CN":"可选按指定维度或指定已开通商品监控。指定维度可选项：userDimension-主账号维度、DG-域名组、CG-服务组、domainDimension-域名维度。若按商品监控，请传入商品code"}
	MonitorProduct *string `json:"monitorProduct,omitempty" xml:"monitorProduct,omitempty" require:"true"`
	// {"en":"Resource type. Required when monitorProduct is specific product. Options: domain", "zh_CN":"监控资源类型。当monitorPorduct传入指定商品时，需要必传resourceType。可选值：domain"}
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// {"en":"List of resource names to monitor or ALL", "zh_CN":"传入具体要监控的资源名称列表或ALL"}
	MonitorResources []*string `json:"monitorResources,omitempty" xml:"monitorResources,omitempty" type:"Repeated"`
	// {"en":"Statistical method. Default: CONSOLIDATED. Options: CONSOLIDATED-consolidated statistics, SEPARATE-separate statistics", "zh_CN":"统计方式。未传默认按合并统计。可选值：CONSOLIDATED-合并统计、SEPARATE-逐一统计"}
	StatisticalMethod *string `json:"statisticalMethod,omitempty" xml:"statisticalMethod,omitempty" require:"true"`
	// {"en":"Alert frequency in minutes. Default: 5. Options: 0-first alert only, 2, 5, 10, 15, 20", "zh_CN":"告警频率。避免告警风暴，每X分钟内产生的告警仅通知一次。未传默认值是5。可选值：0-仅首次告警、2、5、10、15、20"}
	AlertFrequency *int32 `json:"alertFrequency,omitempty" xml:"alertFrequency,omitempty" require:"true"`
	// {"en":"Whether to notify on alert recovery. Default: true", "zh_CN":"告警消警是否通知，未传默认为是。可选值：true-是、false-否"}
	RestoreNotice *string `json:"restoreNotice,omitempty" xml:"restoreNotice,omitempty" require:"true"`
	// {"en":"Rule items list", "zh_CN":"规则项列表"}
	RuleItems []*EditRuleItem `json:"ruleItems,omitempty" xml:"ruleItems,omitempty" require:"true" type:"Repeated"`
	// {"en":"Notification methods list", "zh_CN":"通知方式列表"}
	Notices []*EditRuleNotice `json:"notices,omitempty" xml:"notices,omitempty" require:"true" type:"Repeated"`
}

func (s EditCloudMonitorRealTimeAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetRuleId(v string) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.RuleId = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetRuleName(v string) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.RuleName = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetMonitorProduct(v string) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.MonitorProduct = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetResourceType(v string) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.ResourceType = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetMonitorResources(v []*string) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.MonitorResources = v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetStatisticalMethod(v string) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.StatisticalMethod = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetAlertFrequency(v int32) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.AlertFrequency = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetRestoreNotice(v string) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.RestoreNotice = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetRuleItems(v []*EditRuleItem) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.RuleItems = v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetNotices(v []*EditRuleNotice) *EditCloudMonitorRealTimeAlarmRuleRequest {
	s.Notices = v
	return s
}

// 响应数据
type EditMonitorRuleData struct {
	// {"en":"Rule ID", "zh_CN":"规则ID"}
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s EditMonitorRuleData) String() string {
	return tea.Prettify(s)
}

func (s EditMonitorRuleData) GoString() string {
	return s.String()
}

func (s *EditMonitorRuleData) SetRuleId(v string) *EditMonitorRuleData {
	s.RuleId = &v
	return s
}

// 响应体
type EditCloudMonitorRealTimeAlarmRuleResponse struct {
	// {"en":"Response code", "zh_CN":"响应码"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response message", "zh_CN":"响应消息"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data", "zh_CN":"响应数据"}
	Data *EditMonitorRuleData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s EditCloudMonitorRealTimeAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *EditCloudMonitorRealTimeAlarmRuleResponse) SetCode(v string) *EditCloudMonitorRealTimeAlarmRuleResponse {
	s.Code = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleResponse) SetMessage(v string) *EditCloudMonitorRealTimeAlarmRuleResponse {
	s.Message = &v
	return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleResponse) SetData(v *EditMonitorRuleData) *EditCloudMonitorRealTimeAlarmRuleResponse {
	s.Data = v
	return s
}
