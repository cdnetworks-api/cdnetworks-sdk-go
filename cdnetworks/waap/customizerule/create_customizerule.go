// This file is auto-generated, don't edit it. Thanks.
package customizerule

import (
	"github.com/alibabacloud-go/tea/tea"
)

type AddCustomizeRuleRequest struct {
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
	// {'en':'API ID under API business, multiple separated by ; sign.
	// When the protected target is APIThis field is required.', 'zh_CN':'API业务下的API ID，多个用 ; 隔开。
	// 当业务场景为API业务时此字段必填。'}
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// {'en':'Action.
	// NO_USE:Not Used
	// LOG:Log
	// DELAY:Delay
	// BLOCK:Deny
	// RESET:Reset Connection', 'zh_CN':'处理动作。
	// NO_USE：不使用
	// LOG：监控
	// DELAY：延迟响应
	// BLOCK：拦截
	// RESET:断开连接'}
	Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
	// {'en':'Match Conditions.', 'zh_CN':'匹配条件。'}
	Condition *CommonCustomizeRuleConditionDTO `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
}

func (s AddCustomizeRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomizeRuleRequest) GoString() string {
	return s.String()
}

func (s *AddCustomizeRuleRequest) SetDomain(v string) *AddCustomizeRuleRequest {
	s.Domain = &v
	return s
}

func (s *AddCustomizeRuleRequest) SetRuleName(v string) *AddCustomizeRuleRequest {
	s.RuleName = &v
	return s
}

func (s *AddCustomizeRuleRequest) SetDescription(v string) *AddCustomizeRuleRequest {
	s.Description = &v
	return s
}

func (s *AddCustomizeRuleRequest) SetScene(v string) *AddCustomizeRuleRequest {
	s.Scene = &v
	return s
}

func (s *AddCustomizeRuleRequest) SetApiId(v string) *AddCustomizeRuleRequest {
	s.ApiId = &v
	return s
}

func (s *AddCustomizeRuleRequest) SetAct(v string) *AddCustomizeRuleRequest {
	s.Act = &v
	return s
}

func (s *AddCustomizeRuleRequest) SetCondition(v *CommonCustomizeRuleConditionDTO) *AddCustomizeRuleRequest {
	s.Condition = v
	return s
}

type AddCustomizeRuleResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述信息。'}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {'en':'Data.', 'zh_CN':'出参数据。'}
	Data *string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s AddCustomizeRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomizeRuleResponse) GoString() string {
	return s.String()
}

func (s *AddCustomizeRuleResponse) SetCode(v string) *AddCustomizeRuleResponse {
	s.Code = &v
	return s
}

func (s *AddCustomizeRuleResponse) SetMessage(v string) *AddCustomizeRuleResponse {
	s.Message = &v
	return s
}

func (s *AddCustomizeRuleResponse) SetData(v string) *AddCustomizeRuleResponse {
	s.Data = &v
	return s
}
