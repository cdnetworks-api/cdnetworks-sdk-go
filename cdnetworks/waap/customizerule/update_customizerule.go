// This file is auto-generated, don't edit it. Thanks.
package customizerule

import (
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateCustomRuleRequest struct {
	// {'en':'Rule ID.', 'zh_CN':'规则ID。'}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {'en':'Rule Name, maximum 50 characters.
	// Does not support special characters and spaces.', 'zh_CN':'规则名称，最多50个字符。
	// 不支持特殊字符和空格。'}
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// {'en':'Description, maximum 200 characters.', 'zh_CN':'规则描述，最多200个字符。'}
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// {'en':'Protected target.
	// WEB:Website
	// API:API', 'zh_CN':'业务场景。
	// WEB：WEB业务
	// API：API业务'}
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
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
	Act *string `json:"act,omitempty" xml:"act,omitempty"`
	// {'en':'Matching conditions. Except for header conditions, there can only be at most one record per match type under each type of condition.', 'zh_CN':'匹配条件。除了请求头条件，其它类型的条件下一种匹配类型最多只能有一条记录。'}
	Condition *CommonCustomizeRuleConditionDTO `json:"condition,omitempty" xml:"condition,omitempty"`
}

func (s UpdateCustomRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRuleRequest) SetId(v string) *UpdateCustomRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateCustomRuleRequest) SetRuleName(v string) *UpdateCustomRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateCustomRuleRequest) SetDescription(v string) *UpdateCustomRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateCustomRuleRequest) SetScene(v string) *UpdateCustomRuleRequest {
	s.Scene = &v
	return s
}

func (s *UpdateCustomRuleRequest) SetApiId(v string) *UpdateCustomRuleRequest {
	s.ApiId = &v
	return s
}

func (s *UpdateCustomRuleRequest) SetAct(v string) *UpdateCustomRuleRequest {
	s.Act = &v
	return s
}

func (s *UpdateCustomRuleRequest) SetCondition(v *CommonCustomizeRuleConditionDTO) *UpdateCustomRuleRequest {
	s.Condition = v
	return s
}

type UpdateCustomRuleResponse struct {
	// {'en':'Please refer to the error code for exceptions.', 'zh_CN':'请参照错误码。','dictionary':'belong=WAAP-MS-Ext|dict=waap_retCodeEnum'}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {'en':'Description.', 'zh_CN':'描述信息。'}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {'en':'Data.', 'zh_CN':'出参数据。'}
	Data *bool `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdateCustomRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRuleResponse) SetCode(v string) *UpdateCustomRuleResponse {
	s.Code = &v
	return s
}

func (s *UpdateCustomRuleResponse) SetMessage(v string) *UpdateCustomRuleResponse {
	s.Message = &v
	return s
}

func (s *UpdateCustomRuleResponse) SetData(v bool) *UpdateCustomRuleResponse {
	s.Data = &v
	return s
}
