// This file is auto-generated, don't edit it. Thanks.
package ratelimit

import (
	"github.com/alibabacloud-go/tea/tea"
)

type ListRateLimitingRulesRequest struct {
	// {'en':'Hostname list.', 'zh_CN':'域名列表。'}
	DomainList []*string `json:"domainList,omitempty" xml:"domainList,omitempty" require:"true" type:"Repeated"`
	// {'en':'Rule name, fuzzy query.', 'zh_CN':'规则名称，模糊查询。'}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s ListRateLimitingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRateLimitingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRateLimitingRulesRequest) SetDomainList(v []*string) *ListRateLimitingRulesRequest {
	s.DomainList = v
	return s
}

func (s *ListRateLimitingRulesRequest) SetRuleName(v string) *ListRateLimitingRulesRequest {
	s.RuleName = &v
	return s
}

type ListRateLimitingRulesResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述信息。'}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {'en':'Data.', 'zh_CN':'出参数据。'}
	Data []*CommonRateLimitVO `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListRateLimitingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRateLimitingRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRateLimitingRulesResponse) SetCode(v string) *ListRateLimitingRulesResponse {
	s.Code = &v
	return s
}

func (s *ListRateLimitingRulesResponse) SetMessage(v string) *ListRateLimitingRulesResponse {
	s.Message = &v
	return s
}

func (s *ListRateLimitingRulesResponse) SetData(v []*CommonRateLimitVO) *ListRateLimitingRulesResponse {
	s.Data = v
	return s
}

type CommonRateLimitVO struct {
	// {'en':'Rule ID.', 'zh_CN':'规则ID。'}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {'en':'Hostname.', 'zh_CN':'域名。'}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {'en':'Rule Name.', 'zh_CN':'规则名称。'}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'规则描述。'}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {'en':'Protected target.
	// WEB:Website
	// API:API', 'zh_CN':'业务场景。
	// WEB：网站业务
	// API：API业务'}
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
	// {'en':'Count on.
	// REQUEST:Request
	// RESPONSE:Response', 'zh_CN':'统计阶段。
	// REQUEST：请求
	// RESPONSE：响应'}
	StatisticalStage *string `json:"statisticalStage,omitempty" xml:"statisticalStage,omitempty" require:"true"`
	// {'en':'Client identifier.
	// IP:Client IP
	// IP_UA:Client IP and User-Agent
	// COOKIE:Cookie
	// IP_COOKIE:Client IP and Cookie
	// HEADER:Request Header
	// IP_HEADER:Client IP and Request Header', 'zh_CN':'统计粒度。
	// IP：客户端IP
	// IP_UA：客户端IP和User-Agent
	// COOKIE：Cookie
	// IP_COOKIE：客户端IP和Cookie
	// HEADER：请求头
	// IP_HEADER：客户端IP和请求头'}
	StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty" require:"true"`
	// {'en':'Statistical key value .', 'zh_CN':'统计key值。'}
	StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty" require:"true"`
	// {'en':'Statistics period, unit: seconds.', 'zh_CN':'统计周期，单位：秒。'}
	StatisticalPeriod *int `json:"statisticalPeriod,omitempty" xml:"statisticalPeriod,omitempty" require:"true"`
	// {'en':'Trigger threshold, unit: times.', 'zh_CN':'触发阈值，单位：次。'}
	TriggerThreshold *int `json:"triggerThreshold,omitempty" xml:"triggerThreshold,omitempty" require:"true"`
	// {'en':'Action duration, unit: seconds.', 'zh_CN':'处理动作持续时间，单位：秒。'}
	InterceptTime *int `json:"interceptTime,omitempty" xml:"interceptTime,omitempty" require:"true"`
	// {'en':'Cycle effective status.
	// PERMANENT:All time
	// WITHOUT:Excluded time
	// WITHIN:Selected time', 'zh_CN':'周期生效状态。
	// PERMANENT：永久生效
	// WITHOUT：周期内不生效
	// WITHIN：周期内生效'}
	EffectiveStatus *string `json:"effectiveStatus,omitempty" xml:"effectiveStatus,omitempty" require:"true"`
	// {'en':'Effective time period.', 'zh_CN':'规则生效周期。'}
	RateLimitEffective *RateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty" require:"true"`
	// {'en':'API ID under API business, multiple separated by ; sign.', 'zh_CN':'API业务下的API ID，多个用 ; 隔开。'}
	AssetApiId *string `json:"assetApiId,omitempty" xml:"assetApiId,omitempty" require:"true"`
	// {'en':'Action.
	// NO_USE:Not Used
	// ACCEPT:Skip
	// LOG:Log
	// COOKIE:Cookie verification
	// JS_CHECK:Javascript verification
	// DELAY:Delay
	// BLOCK:Deny
	// RESET:Reset Connection
	// Custom response ID:Custom response ID', 'zh_CN':'处理动作。
	// NO_USE：不使用
	// ACCEPT：放行
	// LOG：监控
	// COOKIE：Cookie校验
	// JS_CHECK：JavaScript校验
	// DELAY：延迟响应
	// BLOCK：拦截
	// RESET：断开连接
	// 自定义响应ID：自定义响应ID'}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {'en':'Matching conditions.', 'zh_CN':'匹配条件。'}
	RateLimitRuleCondition *RateLimitRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty" require:"true"`
	// {'en':'Update time.', 'zh_CN':'更新时间。'}
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s CommonRateLimitVO) String() string {
	return tea.Prettify(s)
}

func (s CommonRateLimitVO) GoString() string {
	return s.String()
}

func (s *CommonRateLimitVO) SetId(v string) *CommonRateLimitVO {
	s.Id = &v
	return s
}

func (s *CommonRateLimitVO) SetDomain(v string) *CommonRateLimitVO {
	s.Domain = &v
	return s
}

func (s *CommonRateLimitVO) SetRuleName(v string) *CommonRateLimitVO {
	s.RuleName = &v
	return s
}

func (s *CommonRateLimitVO) SetDescription(v string) *CommonRateLimitVO {
	s.Description = &v
	return s
}

func (s *CommonRateLimitVO) SetScene(v string) *CommonRateLimitVO {
	s.Scene = &v
	return s
}

func (s *CommonRateLimitVO) SetStatisticalStage(v string) *CommonRateLimitVO {
	s.StatisticalStage = &v
	return s
}

func (s *CommonRateLimitVO) SetStatisticalItem(v string) *CommonRateLimitVO {
	s.StatisticalItem = &v
	return s
}

func (s *CommonRateLimitVO) SetStatisticsKey(v string) *CommonRateLimitVO {
	s.StatisticsKey = &v
	return s
}

func (s *CommonRateLimitVO) SetStatisticalPeriod(v int) *CommonRateLimitVO {
	s.StatisticalPeriod = &v
	return s
}

func (s *CommonRateLimitVO) SetTriggerThreshold(v int) *CommonRateLimitVO {
	s.TriggerThreshold = &v
	return s
}

func (s *CommonRateLimitVO) SetInterceptTime(v int) *CommonRateLimitVO {
	s.InterceptTime = &v
	return s
}

func (s *CommonRateLimitVO) SetEffectiveStatus(v string) *CommonRateLimitVO {
	s.EffectiveStatus = &v
	return s
}

func (s *CommonRateLimitVO) SetRateLimitEffective(v *RateLimitEffective) *CommonRateLimitVO {
	s.RateLimitEffective = v
	return s
}

func (s *CommonRateLimitVO) SetAssetApiId(v string) *CommonRateLimitVO {
	s.AssetApiId = &v
	return s
}

func (s *CommonRateLimitVO) SetAction(v string) *CommonRateLimitVO {
	s.Action = &v
	return s
}

func (s *CommonRateLimitVO) SetRateLimitRuleCondition(v *RateLimitRuleCondition) *CommonRateLimitVO {
	s.RateLimitRuleCondition = v
	return s
}

func (s *CommonRateLimitVO) SetUpdateTime(v string) *CommonRateLimitVO {
	s.UpdateTime = &v
	return s
}
