package alertrules

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type QueryCloudMonitorRealTimeAlarmRuleRequest struct {
  // {"en":"Alert rule ID","zh_CN":"报警规则id，需要鉴权是否该调用账号下或其下子账号的规则"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
  // {"en":"Alarm rule name, rule ID and rule name, either one, rule ID takes priority","zh_CN":"报警规则名称，规则ID和规则名称二选一，规则ID优先"}
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

type QueryCloudMonitorRealTimeAlarmRuleRequestHeader struct {
}

func (s QueryCloudMonitorRealTimeAlarmRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleRequestHeader) GoString() string {
  return s.String()
}

type QueryCloudMonitorRealTimeAlarmRulePaths struct {
}

func (s QueryCloudMonitorRealTimeAlarmRulePaths) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRulePaths) GoString() string {
  return s.String()
}

type QueryCloudMonitorRealTimeAlarmRuleParameters struct {
}

func (s QueryCloudMonitorRealTimeAlarmRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleParameters) GoString() string {
  return s.String()
}

type QueryCloudMonitorRealTimeAlarmRuleResponse struct {
  // {"en":"Response code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"功能码说明"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *QueryCloudMonitorRealTimeAlarmRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
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

func (s *QueryCloudMonitorRealTimeAlarmRuleResponse) SetData(v *QueryCloudMonitorRealTimeAlarmRuleResponseData) *QueryCloudMonitorRealTimeAlarmRuleResponse {
  s.Data = v
  return s
}

type QueryCloudMonitorRealTimeAlarmRuleResponseData struct {
  // {"en":"Alarm rule ID","zh_CN":"报警规则id"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Alert rule name","zh_CN":"报警规则名称"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Monitor product or dimension","zh_CN":"按指定维度或指定已开通商品监控（userDimension-主账号维度、DG-域名组、CG-服务组、domainDimension-域名维度）"}
  MonitorProduct *string `json:"monitorProduct,omitempty" xml:"monitorProduct,omitempty" require:"true"`
  // {"en":"Resource type","zh_CN":"监控资源类型"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty" require:"true"`
  // {"en":"Monitor resources list or ALL","zh_CN":"监控的资源名称列表或ALL"}
  MonitorResources []*string `json:"monitorResources,omitempty" xml:"monitorResources,omitempty" require:"true" type:"Repeated"`
  // {"en":"Statistical method (CONSOLIDATED/SEPARATE)","zh_CN":"统计方式（CONSOLIDATED-合并统计、SEPARATE-逐一统计）"}
  StatisticalMethod *string `json:"statisticalMethod,omitempty" xml:"statisticalMethod,omitempty" require:"true"`
  // {"en":"Alert frequency (0/2/5/10/15/20)","zh_CN":"告警频率，每X分钟内产生的告警仅通知一次（0-仅首次告警、2/5/10/15/20分钟）"}
  AlertFrequency *int `json:"alertFrequency,omitempty" xml:"alertFrequency,omitempty" require:"true"`
  // {"en":"Restore notice (true/false)","zh_CN":"告警消警是否通知（true-是、false-否）"}
  RestoreNotice *string `json:"restoreNotice,omitempty" xml:"restoreNotice,omitempty" require:"true"`
  // {"en":"Rule items list","zh_CN":"规则项列表"}
  RuleItems []*QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems `json:"ruleItems,omitempty" xml:"ruleItems,omitempty" require:"true" type:"Repeated"`
  // {"en":"Notices list","zh_CN":"通知方式列表"}
  Notices []*QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices `json:"notices,omitempty" xml:"notices,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseData) GoString() string {
  return s.String()
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetRuleId(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.RuleId = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetRuleName(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.RuleName = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetMonitorProduct(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.MonitorProduct = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetResourceType(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.ResourceType = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetMonitorResources(v []*string) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.MonitorResources = v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetStatisticalMethod(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.StatisticalMethod = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetAlertFrequency(v int) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.AlertFrequency = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetRestoreNotice(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.RestoreNotice = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetRuleItems(v []*QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.RuleItems = v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseData) SetNotices(v []*QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices) *QueryCloudMonitorRealTimeAlarmRuleResponseData {
  s.Notices = v
  return s
}

type QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems struct     {
  // {"en":"Rule item ID","zh_CN":"分时段子规则"}
  RuleItemId *string `json:"ruleItemId,omitempty" xml:"ruleItemId,omitempty" require:"true"`
  // {"en":"Start time (format: HH:00)","zh_CN":"监控时段开始时间，格式：HH:00，示例：00:00"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"End time (format: HH:59)","zh_CN":"监控时段结束时间，格式：HH:59，示例：01:59"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Condition type (ANY/ALL)","zh_CN":"告警需满足任意或全部条件，可选值：ANY-任意、ALL-全部"}
  ConditionType *string `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
  // {"en":"Monitoring period in minutes (1/5/10)","zh_CN":"监控周期，周期单位为分钟（1-1分钟、5-5分钟、10-10分钟）"}
  Period *int `json:"period,omitempty" xml:"period,omitempty" require:"true"`
  // {"en":"Number of cycles (1/2/3/5/15/30)","zh_CN":"持续几个周期满足条件告警（1/2/3/5/15/30）"}
  PeriodTimes *int `json:"periodTimes,omitempty" xml:"periodTimes,omitempty" require:"true"`
  // {"en":"Condition rules list","zh_CN":"条件规则列表"}
  ConditionRules []*QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules `json:"conditionRules,omitempty" xml:"conditionRules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) GoString() string {
  return s.String()
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) SetRuleItemId(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems {
  s.RuleItemId = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) SetStartTime(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems {
  s.StartTime = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) SetEndTime(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems {
  s.EndTime = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) SetConditionType(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems {
  s.ConditionType = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) SetPeriod(v int) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems {
  s.Period = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) SetPeriodTimes(v int) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems {
  s.PeriodTimes = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems) SetConditionRules(v []*QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItems {
  s.ConditionRules = v
  return s
}

type QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules struct     {
  // {"en":"Monitor item (BANDWIDTH, FLOW, REQUEST, etc.)","zh_CN":"监控项（BANDWIDTH-带宽、FLOW-流量、REQUEST-请求数等）"}
  MonitorItem *string `json:"monitorItem,omitempty" xml:"monitorItem,omitempty" require:"true"`
  // {"en":"Condition type (MAX, MIN, UPRUSH, SLUMPED)","zh_CN":"条件类型（MAX-大于、MIN-小于、UPRUSH-突增、SLUMPED-突降）"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"Threshold value","zh_CN":"监控项阈值"}
  Threshold *string `json:"threshold,omitempty" xml:"threshold,omitempty" require:"true"`
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules) GoString() string {
  return s.String()
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules) SetMonitorItem(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules {
  s.MonitorItem = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules) SetCondition(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules {
  s.Condition = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules) SetThreshold(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataRuleItemsConditionRules {
  s.Threshold = &v
  return s
}

type QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices struct     {
  // {"en":"Notice method (sms/email/robot/webhook)","zh_CN":"通知方式（sms-短信、email-邮件、robot-机器人、webhook-webhook回调）"}
  NoticeMethod *string `json:"noticeMethod,omitempty" xml:"noticeMethod,omitempty" require:"true"`
  // {"en":"Notice object","zh_CN":"告警通知对象，短信/邮件传联系人id，机器人传机器人id，webhook传地址"}
  NoticeObject *string `json:"noticeObject,omitempty" xml:"noticeObject,omitempty" require:"true"`
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices) GoString() string {
  return s.String()
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices) SetNoticeMethod(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices {
  s.NoticeMethod = &v
  return s
}

func (s *QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices) SetNoticeObject(v string) *QueryCloudMonitorRealTimeAlarmRuleResponseDataNotices {
  s.NoticeObject = &v
  return s
}

type QueryCloudMonitorRealTimeAlarmRuleResponseHeader struct {
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryCloudMonitorRealTimeAlarmRuleResponseHeader) GoString() string {
  return s.String()
}




type EditCloudMonitorRealTimeAlarmRuleRequest struct {
  // {"en":"Alarm rule ID","zh_CN":"报警规则id"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
  // {"en":"Alert rule name, only supports Chinese, English, numbers, underscore, hyphen, max 100 characters","zh_CN":"报警规则名称，仅支持中英文、数字、下划线、中划线，最多100个字符"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Monitor by dimension or product. Dimensions: userDimension, DG, CG, domainDimension. Or input product code","zh_CN":"可选按指定维度或指定已开通商品监控。指定维度可选项：userDimension-主账号维度、DG-域名组、CG-服务组、domainDimension-域名维度。若按商品监控，请传入商品code"}
  MonitorProduct *string `json:"monitorProduct,omitempty" xml:"monitorProduct,omitempty" require:"true"`
  // {"en":"Resource type. Required when monitorProduct is specific product. Options: domain","zh_CN":"监控资源类型。当monitorPorduct传入指定商品时，需要必传resourceType。可选值：domain"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
  // {"en":"List of resource names to monitor or ALL","zh_CN":"传入具体要监控的资源名称列表或ALL"}
  MonitorResources []*string `json:"monitorResources,omitempty" xml:"monitorResources,omitempty" type:"Repeated"`
  // {"en":"Statistical method. Default: CONSOLIDATED. Options: CONSOLIDATED-consolidated statistics, SEPARATE-separate statistics","zh_CN":"统计方式。未传默认按合并统计。可选值：CONSOLIDATED-合并统计、SEPARATE-逐一统计"}
  StatisticalMethod *string `json:"statisticalMethod,omitempty" xml:"statisticalMethod,omitempty"`
  // {"en":"Alert frequency in minutes. Default: 5. Options: 0-first alert only, 2, 5, 10, 15, 20","zh_CN":"告警频率。避免告警风暴，每X分钟内产生的告警仅通知一次。未传默认值是5。可选值：0-仅首次告警、2、5、10、15、20"}
  AlertFrequency *int `json:"alertFrequency,omitempty" xml:"alertFrequency,omitempty"`
  // {"en":"Whether to notify on alert recovery. Default: true","zh_CN":"告警消警是否通知，未传默认为是。可选值：true-是、false-否"}
  RestoreNotice *string `json:"restoreNotice,omitempty" xml:"restoreNotice,omitempty"`
  // {"en":"Rule items list","zh_CN":"规则项列表"}
  RuleItems []*EditCloudMonitorRealTimeAlarmRuleRequestRuleItems `json:"ruleItems,omitempty" xml:"ruleItems,omitempty" require:"true" type:"Repeated"`
  // {"en":"Notification methods list","zh_CN":"通知方式列表"}
  Notices []*EditCloudMonitorRealTimeAlarmRuleRequestNotices `json:"notices,omitempty" xml:"notices,omitempty" require:"true" type:"Repeated"`
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

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetAlertFrequency(v int) *EditCloudMonitorRealTimeAlarmRuleRequest {
  s.AlertFrequency = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetRestoreNotice(v string) *EditCloudMonitorRealTimeAlarmRuleRequest {
  s.RestoreNotice = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetRuleItems(v []*EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) *EditCloudMonitorRealTimeAlarmRuleRequest {
  s.RuleItems = v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequest) SetNotices(v []*EditCloudMonitorRealTimeAlarmRuleRequestNotices) *EditCloudMonitorRealTimeAlarmRuleRequest {
  s.Notices = v
  return s
}

type EditCloudMonitorRealTimeAlarmRuleRequestRuleItems struct     {
  // {"en":"Time-divided sub-rules. If an id is passed in, it is considered to specify a modification of an item rule.","zh_CN":"分时段子规则id"}
  RuleItemId *string `json:"ruleItemId,omitempty" xml:"ruleItemId,omitempty"`
  // {"en":"Monitoring period start time, format: HH:00, example: 00:00","zh_CN":"监控时段开始时间，格式：HH:00，示例：00:00"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Monitoring period end time, format: HH:59, example: 01:59","zh_CN":"监控时段结束时间，格式：HH:59，示例：01:59"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Alert condition type. Options: ANY-any condition, ALL-all conditions","zh_CN":"告警需满足任意或全部条件。可选值：ANY-任意、ALL-全部"}
  ConditionType *string `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
  // {"en":"Monitoring cycle in minutes. Options: 1, 5, 10","zh_CN":"监控周期，周期单位为分钟。可选值：1-1分钟、5-5分钟、10-10分钟"}
  Period *int `json:"period,omitempty" xml:"period,omitempty" require:"true"`
  // {"en":"Number of cycles to meet conditions before alerting. Options: 1, 2, 3, 5, 15, 30","zh_CN":"持续几个周期满足条件告警。可选值：1、2、3、5、15、30"}
  PeriodTimes *int `json:"periodTimes,omitempty" xml:"periodTimes,omitempty" require:"true"`
  // {"en":"Condition rules list","zh_CN":"条件规则列表"}
  ConditionRules []*EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules `json:"conditionRules,omitempty" xml:"conditionRules,omitempty" require:"true" type:"Repeated"`
}

func (s EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) String() string {
  return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) GoString() string {
  return s.String()
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetRuleItemId(v string) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.RuleItemId = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetStartTime(v string) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.StartTime = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetEndTime(v string) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.EndTime = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetConditionType(v string) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.ConditionType = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetPeriod(v int) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.Period = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetPeriodTimes(v int) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.PeriodTimes = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetConditionRules(v []*EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.ConditionRules = v
  return s
}

type EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules struct     {
  // {"en":"Monitoring items currently only support domain-type monitoring items. Options: BANDWIDTH, FLOW, REQUEST, BTOB, FLOW_HIT_RATE, etc.","zh_CN":"当前仅支持监控对象类型为域名的监控项。可选监控项：BANDWIDTH-带宽、FLOW-流量、REQUEST-请求数、BTOB-回源带宽、FLOW_HIT_RATE-流量命中率等"}
  MonitorItem *string `json:"monitorItem,omitempty" xml:"monitorItem,omitempty" require:"true"`
  // {"en":"Condition type. Options: MAX-greater than, MIN-less than, UPRUSH-surge, SLUMPED-plunge","zh_CN":"条件类型。可选值：MAX-大于、MIN-小于、UPRUSH-突增、SLUMPED-突降"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"Threshold value, please provide a positive integer","zh_CN":"监控项阈值，请传递正整数"}
  Threshold *string `json:"threshold,omitempty" xml:"threshold,omitempty" require:"true"`
}

func (s EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) String() string {
  return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) GoString() string {
  return s.String()
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) SetMonitorItem(v string) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules {
  s.MonitorItem = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) SetCondition(v string) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules {
  s.Condition = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) SetThreshold(v string) *EditCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules {
  s.Threshold = &v
  return s
}

type EditCloudMonitorRealTimeAlarmRuleRequestNotices struct     {
  // {"en":"Notification method. Options: MOBILE, EMAIL, robot, webhook","zh_CN":"通知方式。可选值：MOBILE-短信、EMAIL-邮件、ROBOT-机器人、WEBHOOK-webhook回调"}
  NoticeMethod *string `json:"noticeMethod,omitempty" xml:"noticeMethod,omitempty" require:"true"`
  // {"en":"Notification object. For MOBILE/EMAIL: contact IDs separated by ;. For ROBOT: robot IDs separated by ;. For WEBHOOK: webhook URL","zh_CN":"告警通知对象。若noticeMethod为MOBILE、EMAIL：请传递联系人id，多个用;分隔。若为ROBOT：请传递机器人id，多个用;分隔。若为WEBHOOK：请直接传递webhook地址"}
  NoticeObject *string `json:"noticeObject,omitempty" xml:"noticeObject,omitempty" require:"true"`
}

func (s EditCloudMonitorRealTimeAlarmRuleRequestNotices) String() string {
  return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleRequestNotices) GoString() string {
  return s.String()
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestNotices) SetNoticeMethod(v string) *EditCloudMonitorRealTimeAlarmRuleRequestNotices {
  s.NoticeMethod = &v
  return s
}

func (s *EditCloudMonitorRealTimeAlarmRuleRequestNotices) SetNoticeObject(v string) *EditCloudMonitorRealTimeAlarmRuleRequestNotices {
  s.NoticeObject = &v
  return s
}

type EditCloudMonitorRealTimeAlarmRuleRequestHeader struct {
}

func (s EditCloudMonitorRealTimeAlarmRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleRequestHeader) GoString() string {
  return s.String()
}

type EditCloudMonitorRealTimeAlarmRulePaths struct {
}

func (s EditCloudMonitorRealTimeAlarmRulePaths) String() string {
  return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRulePaths) GoString() string {
  return s.String()
}

type EditCloudMonitorRealTimeAlarmRuleParameters struct {
}

func (s EditCloudMonitorRealTimeAlarmRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleParameters) GoString() string {
  return s.String()
}

type EditCloudMonitorRealTimeAlarmRuleResponse struct {
  // {"en":"Response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *EditCloudMonitorRealTimeAlarmRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
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

func (s *EditCloudMonitorRealTimeAlarmRuleResponse) SetData(v *EditCloudMonitorRealTimeAlarmRuleResponseData) *EditCloudMonitorRealTimeAlarmRuleResponse {
  s.Data = v
  return s
}

type EditCloudMonitorRealTimeAlarmRuleResponseData struct {
  // {"en":"Rule ID","zh_CN":"规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s EditCloudMonitorRealTimeAlarmRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleResponseData) GoString() string {
  return s.String()
}

func (s *EditCloudMonitorRealTimeAlarmRuleResponseData) SetRuleId(v string) *EditCloudMonitorRealTimeAlarmRuleResponseData {
  s.RuleId = &v
  return s
}

type EditCloudMonitorRealTimeAlarmRuleResponseHeader struct {
}

func (s EditCloudMonitorRealTimeAlarmRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s EditCloudMonitorRealTimeAlarmRuleResponseHeader) GoString() string {
  return s.String()
}




type DeleteCloudMonitorRealTimeAlarmRuleRequest struct {
  // {"en":"Alert rule ID","zh_CN":"报警规则id，需要鉴权是否该调用账号下或其下子账号的规则"}
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

type DeleteCloudMonitorRealTimeAlarmRuleRequestHeader struct {
}

func (s DeleteCloudMonitorRealTimeAlarmRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCloudMonitorRealTimeAlarmRuleRequestHeader) GoString() string {
  return s.String()
}

type DeleteCloudMonitorRealTimeAlarmRulePaths struct {
}

func (s DeleteCloudMonitorRealTimeAlarmRulePaths) String() string {
  return tea.Prettify(s)
}

func (s DeleteCloudMonitorRealTimeAlarmRulePaths) GoString() string {
  return s.String()
}

type DeleteCloudMonitorRealTimeAlarmRuleParameters struct {
}

func (s DeleteCloudMonitorRealTimeAlarmRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s DeleteCloudMonitorRealTimeAlarmRuleParameters) GoString() string {
  return s.String()
}

type DeleteCloudMonitorRealTimeAlarmRuleResponse struct {
  // {"en":"Response code","zh_CN":"功能码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"功能码说明"}
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

type DeleteCloudMonitorRealTimeAlarmRuleResponseHeader struct {
}

func (s DeleteCloudMonitorRealTimeAlarmRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeleteCloudMonitorRealTimeAlarmRuleResponseHeader) GoString() string {
  return s.String()
}




type CreateCloudMonitorRealTimeAlarmRuleRequest struct {
  // {"en":"Alert rule name, only supports Chinese, English, numbers, underscore, hyphen, max 100 characters","zh_CN":"报警规则名称，仅支持中英文、数字、下划线、中划线，最多100个字符"}
  RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
  // {"en":"Monitor by dimension or product. Dimensions: userDimension, DG, CG, domainDimension. Or input product code","zh_CN":"可选按指定维度或指定已开通商品监控。指定维度可选项：userDimension-主账号维度、DG-域名组、CG-服务组、domainDimension-域名维度。若按商品监控，请传入商品code"}
  MonitorProduct *string `json:"monitorProduct,omitempty" xml:"monitorProduct,omitempty" require:"true"`
  // {"en":"Resource type. Required when monitorProduct is specific product. Options: domain","zh_CN":"监控资源类型。当monitorPorduct传入指定商品时，需要必传resourceType。可选值：domain"}
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
  // {"en":"List of resource names to monitor or ALL","zh_CN":"传入具体要监控的资源名称列表或ALL"}
  MonitorResources []*string `json:"monitorResources,omitempty" xml:"monitorResources,omitempty" type:"Repeated"`
  // {"en":"Statistical method. Default: CONSOLIDATED. Options: CONSOLIDATED-consolidated statistics, SEPARATE-separate statistics","zh_CN":"统计方式。未传默认按合并统计。可选值：CONSOLIDATED-合并统计、SEPARATE-逐一统计"}
  StatisticalMethod *string `json:"statisticalMethod,omitempty" xml:"statisticalMethod,omitempty"`
  // {"en":"Alert frequency in minutes. Default: 5. Options: 0-first alert only, 2, 5, 10, 15, 20","zh_CN":"告警频率。避免告警风暴，每X分钟内产生的告警仅通知一次。未传默认值是5。可选值：0-仅首次告警、2、5、10、15、20"}
  AlertFrequency *int `json:"alertFrequency,omitempty" xml:"alertFrequency,omitempty"`
  // {"en":"Whether to notify on alert recovery. Default: true","zh_CN":"告警消警是否通知，未传默认为是。可选值：true-是、false-否"}
  RestoreNotice *string `json:"restoreNotice,omitempty" xml:"restoreNotice,omitempty"`
  // {"en":"Rule items list","zh_CN":"规则项列表"}
  RuleItems []*CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems `json:"ruleItems,omitempty" xml:"ruleItems,omitempty" require:"true" type:"Repeated"`
  // {"en":"Notification methods list","zh_CN":"通知方式列表"}
  Notices []*CreateCloudMonitorRealTimeAlarmRuleRequestNotices `json:"notices,omitempty" xml:"notices,omitempty" require:"true" type:"Repeated"`
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequest) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequest) GoString() string {
  return s.String()
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetRuleName(v string) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.RuleName = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetMonitorProduct(v string) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.MonitorProduct = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetResourceType(v string) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.ResourceType = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetMonitorResources(v []*string) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.MonitorResources = v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetStatisticalMethod(v string) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.StatisticalMethod = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetAlertFrequency(v int) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.AlertFrequency = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetRestoreNotice(v string) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.RestoreNotice = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetRuleItems(v []*CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.RuleItems = v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequest) SetNotices(v []*CreateCloudMonitorRealTimeAlarmRuleRequestNotices) *CreateCloudMonitorRealTimeAlarmRuleRequest {
  s.Notices = v
  return s
}

type CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems struct     {
  // {"en":"Monitoring period start time, format: HH:00, example: 00:00","zh_CN":"监控时段开始时间，格式：HH:00，示例：00:00"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"Monitoring period end time, format: HH:59, example: 01:59","zh_CN":"监控时段结束时间，格式：HH:59，示例：01:59"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en":"Alert condition type. Options: ANY-any condition, ALL-all conditions","zh_CN":"告警需满足任意或全部条件。可选值：ANY-任意、ALL-全部"}
  ConditionType *string `json:"conditionType,omitempty" xml:"conditionType,omitempty" require:"true"`
  // {"en":"Monitoring cycle in minutes. Options: 1, 5, 10","zh_CN":"监控周期，周期单位为分钟。可选值：1-1分钟、5-5分钟、10-10分钟"}
  Period *int `json:"period,omitempty" xml:"period,omitempty" require:"true"`
  // {"en":"Number of cycles to meet conditions before alerting. Options: 1, 2, 3, 5, 15, 30","zh_CN":"持续几个周期满足条件告警。可选值：1、2、3、5、15、30"}
  PeriodTimes *int `json:"periodTimes,omitempty" xml:"periodTimes,omitempty" require:"true"`
  // {"en":"Condition rules list","zh_CN":"条件规则列表"}
  ConditionRules []*CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules `json:"conditionRules,omitempty" xml:"conditionRules,omitempty" require:"true" type:"Repeated"`
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) GoString() string {
  return s.String()
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetStartTime(v string) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.StartTime = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetEndTime(v string) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.EndTime = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetConditionType(v string) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.ConditionType = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetPeriod(v int) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.Period = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetPeriodTimes(v int) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.PeriodTimes = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems) SetConditionRules(v []*CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItems {
  s.ConditionRules = v
  return s
}

type CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules struct     {
  // {"en":"Monitoring items currently only support domain-type monitoring items. Options: BANDWIDTH, FLOW, REQUEST, BTOB, FLOW_HIT_RATE, etc.","zh_CN":"当前仅支持监控对象类型为域名的监控项。可选监控项：BANDWIDTH-带宽、FLOW-流量、REQUEST-请求数、BTOB-回源带宽、FLOW_HIT_RATE-流量命中率等"}
  MonitorItem *string `json:"monitorItem,omitempty" xml:"monitorItem,omitempty" require:"true"`
  // {"en":"Condition type. Options: MAX-greater than, MIN-less than, UPRUSH-surge, SLUMPED-plunge","zh_CN":"条件类型。可选值：MAX-大于、MIN-小于、UPRUSH-突增、SLUMPED-突降"}
  Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
  // {"en":"Threshold value, please provide a positive integer","zh_CN":"监控项阈值，请传递正整数"}
  Threshold *string `json:"threshold,omitempty" xml:"threshold,omitempty" require:"true"`
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) GoString() string {
  return s.String()
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) SetMonitorItem(v string) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules {
  s.MonitorItem = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) SetCondition(v string) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules {
  s.Condition = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules) SetThreshold(v string) *CreateCloudMonitorRealTimeAlarmRuleRequestRuleItemsConditionRules {
  s.Threshold = &v
  return s
}

type CreateCloudMonitorRealTimeAlarmRuleRequestNotices struct     {
  // {"en":"Notification method. Options: MOBILE, EMAIL, ROBOT, WEBHOOK","zh_CN":"通知方式。可选值：MOBILE-短信、EMAIL-邮件、ROBOT-机器人、WEBHOOK-webhook回调","exampleValue":"MOBILE, EMAIL, ROBOT, WEBHOOK"}
  NoticeMethod *string `json:"noticeMethod,omitempty" xml:"noticeMethod,omitempty" require:"true"`
  // {"en":"Notification object. For MOBILE/EMAIL: contact IDs separated by ;. For ROBOT: robot IDs separated by ;. For WEBHOOK: webhook URL","zh_CN":"告警通知对象。若noticeMethod为MOBILE、EMAIL：请传递联系人id，多个用;分隔。若为ROBOT：请传递机器人id，多个用;分隔。若为WEBHOOK：请直接传递webhook地址"}
  NoticeObject *string `json:"noticeObject,omitempty" xml:"noticeObject,omitempty" require:"true"`
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequestNotices) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequestNotices) GoString() string {
  return s.String()
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestNotices) SetNoticeMethod(v string) *CreateCloudMonitorRealTimeAlarmRuleRequestNotices {
  s.NoticeMethod = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleRequestNotices) SetNoticeObject(v string) *CreateCloudMonitorRealTimeAlarmRuleRequestNotices {
  s.NoticeObject = &v
  return s
}

type CreateCloudMonitorRealTimeAlarmRuleRequestHeader struct {
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleRequestHeader) GoString() string {
  return s.String()
}

type CreateCloudMonitorRealTimeAlarmRulePaths struct {
}

func (s CreateCloudMonitorRealTimeAlarmRulePaths) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRulePaths) GoString() string {
  return s.String()
}

type CreateCloudMonitorRealTimeAlarmRuleParameters struct {
}

func (s CreateCloudMonitorRealTimeAlarmRuleParameters) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleParameters) GoString() string {
  return s.String()
}

type CreateCloudMonitorRealTimeAlarmRuleResponse struct {
  // {"en":"Response code","zh_CN":"响应码"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response message","zh_CN":"响应消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Response data","zh_CN":"响应数据"}
  Data *CreateCloudMonitorRealTimeAlarmRuleResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateCloudMonitorRealTimeAlarmRuleResponse) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleResponse) GoString() string {
  return s.String()
}

func (s *CreateCloudMonitorRealTimeAlarmRuleResponse) SetCode(v string) *CreateCloudMonitorRealTimeAlarmRuleResponse {
  s.Code = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleResponse) SetMessage(v string) *CreateCloudMonitorRealTimeAlarmRuleResponse {
  s.Message = &v
  return s
}

func (s *CreateCloudMonitorRealTimeAlarmRuleResponse) SetData(v *CreateCloudMonitorRealTimeAlarmRuleResponseData) *CreateCloudMonitorRealTimeAlarmRuleResponse {
  s.Data = v
  return s
}

type CreateCloudMonitorRealTimeAlarmRuleResponseData struct {
  // {"en":"Rule ID","zh_CN":"规则ID"}
  RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty" require:"true"`
}

func (s CreateCloudMonitorRealTimeAlarmRuleResponseData) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleResponseData) GoString() string {
  return s.String()
}

func (s *CreateCloudMonitorRealTimeAlarmRuleResponseData) SetRuleId(v string) *CreateCloudMonitorRealTimeAlarmRuleResponseData {
  s.RuleId = &v
  return s
}

type CreateCloudMonitorRealTimeAlarmRuleResponseHeader struct {
}

func (s CreateCloudMonitorRealTimeAlarmRuleResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreateCloudMonitorRealTimeAlarmRuleResponseHeader) GoString() string {
  return s.String()
}




