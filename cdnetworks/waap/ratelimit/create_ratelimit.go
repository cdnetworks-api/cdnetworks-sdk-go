// This file is auto-generated, don't edit it. Thanks.
package ratelimit

import (
	"github.com/alibabacloud-go/tea/tea"
)

type CreatRateLimitingRuleRequest struct {
	// {'en':'Hostname.', 'zh_CN':'域名。'}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {'en':'Rule Name, maximum 50 characters.
	// Does not support special characters and spaces.', 'zh_CN':'规则名称，最多50个字符。
	// 不支持特殊字符和空格。'}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty" require:"true"`
	// {'en':'Description, maximum 200 characters.', 'zh_CN':'规则描述，最多200个字符。'}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {'en':'Protected target.
	// WEB:Website
	// API:API', 'zh_CN':'业务场景。
	// WEB：网站业务
	// API：API业务'}
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
	// {'en':'Client identifier.
	// IP:Client IP
	// IP_UA:Client IP and User-Agent
	// COOKIE:Cookie
	// IP_COOKIE:Client IP and Cookie
	// HEADER:Request Header
	// When there is a status code in the matching condition,this client identifier is not supported.
	// IP_HEADER:Client IP and Request Header
	// When there is a status code in the matching condition,this client identifier is not supported .', 'zh_CN':'统计粒度。
	// IP：客户端IP
	// IP_UA：客户端IP和User-Agent
	// COOKIE：Cookie
	// IP_COOKIE：客户端IP和Cookie
	// HEADER：请求头，当匹配条件中存在状态码时不支持此统计粒度
	// IP_HEADER：客户端IP和请求头，当匹配条件中存在状态码时不支持此统计粒度'}
	StatisticalItem *string `json:"statisticalItem,omitempty" xml:"statisticalItem,omitempty" require:"true"`
	// {'en':'Statistical key value.
	// When the client identifier is cookie/header value, the corresponding key value needs to be entered.', 'zh_CN':'统计key值。
	// 当统计粒度cookie/header值，需要输入对应的key值。'}
	StatisticsKey *string `json:"statisticsKey,omitempty" xml:"statisticsKey,omitempty"`
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
	// {'en':'Effective time period.When the effective status is effective within the cycle or not effective within the cycle, this field must have a value.', 'zh_CN':'规则生效周期。
	// 生效状态为周期内生效或周期内不生效时，此字段必须有值。'}
	RateLimitEffective *RateLimitEffective `json:"rateLimitEffective,omitempty" xml:"rateLimitEffective,omitempty"`
	// {'en':'API ID under API business, multiple separated by ; sign.
	// When the protected target is APIThis field is required.', 'zh_CN':'API业务下的API ID，多个用 ; 隔开。
	// 当业务场景为API业务时此字段必填。'}
	AssetApiId *string `json:"assetApiId,omitempty" xml:"assetApiId,omitempty"`
	// {'en':'Action.
	// NO_USE:Not Used
	// ACCEPT:Skip
	// LOG:Log
	// COOKIE:Cookie verification
	// JS_CHECK:Javascript verification
	// DELAY:Delay
	// BLOCK:Deny
	// RESET:Reset Connection
	// Custom response ID:Custom response ID
	// When there is a status code in the matching condition, the supported actions are Log, Deny,Not Used, and Reset Connection.', 'zh_CN':'处理动作。
	// NO_USE：不使用
	// ACCEPT：放行
	// LOG：监控
	// COOKIE：Cookie校验
	// JS_CHECK：JavaScript校验
	// DELAY：延迟响应
	// BLOCK：拦截
	// RESET：断开连接
	// 自定义响应ID：自定义响应ID
	// 当匹配条件中存在状态码时，支持处理动作为监控、拦截、不使用、断开连接。'}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {'en':'Matching conditions.', 'zh_CN':'匹配条件。'}
	RateLimitRuleCondition *RateLimitRuleCondition `json:"rateLimitRuleCondition,omitempty" xml:"rateLimitRuleCondition,omitempty" require:"true"`
}

func (s CreatRateLimitingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatRateLimitingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreatRateLimitingRuleRequest) SetDomain(v string) *CreatRateLimitingRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetRuleName(v string) *CreatRateLimitingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetDescription(v string) *CreatRateLimitingRuleRequest {
	s.Description = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetScene(v string) *CreatRateLimitingRuleRequest {
	s.Scene = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetStatisticalItem(v string) *CreatRateLimitingRuleRequest {
	s.StatisticalItem = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetStatisticsKey(v string) *CreatRateLimitingRuleRequest {
	s.StatisticsKey = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetStatisticalPeriod(v int) *CreatRateLimitingRuleRequest {
	s.StatisticalPeriod = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetTriggerThreshold(v int) *CreatRateLimitingRuleRequest {
	s.TriggerThreshold = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetInterceptTime(v int) *CreatRateLimitingRuleRequest {
	s.InterceptTime = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetEffectiveStatus(v string) *CreatRateLimitingRuleRequest {
	s.EffectiveStatus = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetRateLimitEffective(v *RateLimitEffective) *CreatRateLimitingRuleRequest {
	s.RateLimitEffective = v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetAssetApiId(v string) *CreatRateLimitingRuleRequest {
	s.AssetApiId = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetAction(v string) *CreatRateLimitingRuleRequest {
	s.Action = &v
	return s
}

func (s *CreatRateLimitingRuleRequest) SetRateLimitRuleCondition(v *RateLimitRuleCondition) *CreatRateLimitingRuleRequest {
	s.RateLimitRuleCondition = v
	return s
}

type CreatRateLimitingRuleResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述信息。'}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {'en':'Rule ID.', 'zh_CN':'规则ID。'}
	Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreatRateLimitingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatRateLimitingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreatRateLimitingRuleResponse) SetCode(v string) *CreatRateLimitingRuleResponse {
	s.Code = &v
	return s
}

func (s *CreatRateLimitingRuleResponse) SetMessage(v string) *CreatRateLimitingRuleResponse {
	s.Message = &v
	return s
}

func (s *CreatRateLimitingRuleResponse) SetData(v string) *CreatRateLimitingRuleResponse {
	s.Data = &v
	return s
}
