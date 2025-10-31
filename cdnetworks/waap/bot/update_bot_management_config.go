// This file is auto-generated, don't edit it. Thanks.
package bot

import (
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateBotManagementConfigRequest struct {
	// {"en":"Domain.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"General policies.","zh_CN":"通用策略。"}
	GeneralStrategy *UpdateBotManagementConfigRequestGeneralStrategy `json:"generalStrategy,omitempty" xml:"generalStrategy,omitempty" type:"Struct"`
	// {"en":"Web bot detection.","zh_CN":"Web风险检测。"}
	WebConfig *UpdateBotManagementConfigRequestWebConfig `json:"webConfig,omitempty" xml:"webConfig,omitempty" type:"Struct"`
	// {"en":"Traffic fraud detection.","zh_CN":"异常流量检测。"}
	TrafficDetection *UpdateBotManagementConfigRequestTrafficDetection `json:"trafficDetection,omitempty" xml:"trafficDetection,omitempty" type:"Struct"`
}

func (s UpdateBotManagementConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigRequest) SetDomain(v string) *UpdateBotManagementConfigRequest {
	s.Domain = &v
	return s
}

func (s *UpdateBotManagementConfigRequest) SetGeneralStrategy(v *UpdateBotManagementConfigRequestGeneralStrategy) *UpdateBotManagementConfigRequest {
	s.GeneralStrategy = v
	return s
}

func (s *UpdateBotManagementConfigRequest) SetWebConfig(v *UpdateBotManagementConfigRequestWebConfig) *UpdateBotManagementConfigRequest {
	s.WebConfig = v
	return s
}

func (s *UpdateBotManagementConfigRequest) SetTrafficDetection(v *UpdateBotManagementConfigRequestTrafficDetection) *UpdateBotManagementConfigRequest {
	s.TrafficDetection = v
	return s
}

type UpdateBotManagementConfigRequestGeneralStrategy struct {
	// {"en":"AI Bots configuration.","zh_CN":"AI Bots配置。"}
	AiBots []*UpdateBotManagementConfigRequestGeneralStrategyAiBots `json:"aiBots,omitempty" xml:"aiBots,omitempty" type:"Repeated"`
	// {"en":"Public Bots configuration.","zh_CN":"公开Bots配置。"}
	PublicBots []*UpdateBotManagementConfigRequestGeneralStrategyPublicBots `json:"publicBots,omitempty" xml:"publicBots,omitempty" type:"Repeated"`
	// {"en":"Definite Bots actions. NO_USE: Not Used.LOG: Log.BLOCK: Deny.","zh_CN":"绝对Bots处理动作。 \nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。","exampleValue":"NO_USE,LOG,BLOCK"}
	AbsoluteBotsAct *string `json:"absoluteBotsAct,omitempty" xml:"absoluteBotsAct,omitempty"`
	// {"en":"Bot Tagging.","zh_CN":"Bot标记。"}
	BotTagging []*UpdateBotManagementConfigRequestGeneralStrategyBotTagging `json:"botTagging,omitempty" xml:"botTagging,omitempty" type:"Repeated"`
}

func (s UpdateBotManagementConfigRequestGeneralStrategy) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigRequestGeneralStrategy) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigRequestGeneralStrategy) SetAiBots(v []*UpdateBotManagementConfigRequestGeneralStrategyAiBots) *UpdateBotManagementConfigRequestGeneralStrategy {
	s.AiBots = v
	return s
}

func (s *UpdateBotManagementConfigRequestGeneralStrategy) SetPublicBots(v []*UpdateBotManagementConfigRequestGeneralStrategyPublicBots) *UpdateBotManagementConfigRequestGeneralStrategy {
	s.PublicBots = v
	return s
}

func (s *UpdateBotManagementConfigRequestGeneralStrategy) SetAbsoluteBotsAct(v string) *UpdateBotManagementConfigRequestGeneralStrategy {
	s.AbsoluteBotsAct = &v
	return s
}

func (s *UpdateBotManagementConfigRequestGeneralStrategy) SetBotTagging(v []*UpdateBotManagementConfigRequestGeneralStrategyBotTagging) *UpdateBotManagementConfigRequestGeneralStrategy {
	s.BotTagging = v
	return s
}

type UpdateBotManagementConfigRequestGeneralStrategyAiBots struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_aiBotsCategory","en":"AI Bots categories.","zh_CN":"AI Bots类别。"}
	BotCategory *string `json:"botCategory,omitempty" xml:"botCategory,omitempty" require:"true"`
	// {"en":"Actions. \nNO_USE: Not Used.\nLOG: Log.\nBLOCK: Deny.","zh_CN":"处理动作。 \nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。","exampleValue":"NO_USE,LOG,BLOCK"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigRequestGeneralStrategyAiBots) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigRequestGeneralStrategyAiBots) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigRequestGeneralStrategyAiBots) SetBotCategory(v string) *UpdateBotManagementConfigRequestGeneralStrategyAiBots {
	s.BotCategory = &v
	return s
}

func (s *UpdateBotManagementConfigRequestGeneralStrategyAiBots) SetAction(v string) *UpdateBotManagementConfigRequestGeneralStrategyAiBots {
	s.Action = &v
	return s
}

type UpdateBotManagementConfigRequestGeneralStrategyPublicBots struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_publicBotsCategory","en":"Public Bots categories.","zh_CN":"公开Bots类别。"}
	BotCategory *string `json:"botCategory,omitempty" xml:"botCategory,omitempty" require:"true"`
	// {"en":"Actions. \nNO_USE: Not Used.\nLOG: Log.\nBLOCK: Deny.\nACCEPT: Release.","zh_CN":"处理动作。 \nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。\nACCEPT：放行。","exampleValue":"NO_USE,LOG,BLOCK,ACCEPT"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigRequestGeneralStrategyPublicBots) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigRequestGeneralStrategyPublicBots) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigRequestGeneralStrategyPublicBots) SetBotCategory(v string) *UpdateBotManagementConfigRequestGeneralStrategyPublicBots {
	s.BotCategory = &v
	return s
}

func (s *UpdateBotManagementConfigRequestGeneralStrategyPublicBots) SetAction(v string) *UpdateBotManagementConfigRequestGeneralStrategyPublicBots {
	s.Action = &v
	return s
}

type UpdateBotManagementConfigRequestGeneralStrategyBotTagging struct {
	// {"en":"Customizing HTTP request headers.","zh_CN":"自定义HTTP请求头。"}
	RequestHeaderKey *string `json:"requestHeaderKey,omitempty" xml:"requestHeaderKey,omitempty" require:"true"`
	// {"en":"Tagging Bot traffic characteristics.\nAI_BOT: AI Bots.\nPUBLIC_BOT: Public Bots.\nCUSTOMIZE_BOT: Custom Bots.","zh_CN":"标记Bot流量特征。\nAI_BOT：AI Bots。\nPUBLIC_BOT：公开Bots。\nCUSTOMIZE_BOT：自定义Bots。","exampleValue":"AI_BOT,PUBLIC_BOT,CUSTOMIZE_BOT"}
	TrafficCharacteristics *string `json:"trafficCharacteristics,omitempty" xml:"trafficCharacteristics,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigRequestGeneralStrategyBotTagging) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigRequestGeneralStrategyBotTagging) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigRequestGeneralStrategyBotTagging) SetRequestHeaderKey(v string) *UpdateBotManagementConfigRequestGeneralStrategyBotTagging {
	s.RequestHeaderKey = &v
	return s
}

func (s *UpdateBotManagementConfigRequestGeneralStrategyBotTagging) SetTrafficCharacteristics(v string) *UpdateBotManagementConfigRequestGeneralStrategyBotTagging {
	s.TrafficCharacteristics = &v
	return s
}

type UpdateBotManagementConfigRequestWebConfig struct {
	// {"en":"Action.\nNO_USE: Do not use\nBLOCK: Intercept\nLOG: Monitor","zh_CN":"处理动作。\nNO_USE:不使用\nBLOCK:拦截\nLOG:监控"}
	Act *string `json:"act,omitempty" xml:"act,omitempty"`
	// {"en":"Browser feature function switch.\nON: Enable\nOFF: Disable","zh_CN":"浏览器特性功能开关。\nON：开启\nOFF：关闭"}
	BrowserAnalyseSwitch *string `json:"browserAnalyseSwitch,omitempty" xml:"browserAnalyseSwitch,omitempty"`
	// {"en":"Automated tool detection function switch.\nON: Enable\nOFF: Disable","zh_CN":"自动化工具检测功能开关。\nON：开启\nOFF：关闭"}
	AutoToolSwitch *string `json:"autoToolSwitch,omitempty" xml:"autoToolSwitch,omitempty"`
	// {"en":"Cracking the behavior detection function switch.\nON: Enable\nOFF: Disable","zh_CN":"破解行为检测功能开关。\nON：开启\nOFF：关闭"}
	CrackAnalyseSwitch *string `json:"crackAnalyseSwitch,omitempty" xml:"crackAnalyseSwitch,omitempty"`
	// {"en":"Page anti-debugging function switch.\nON: Enable\nOFF: Disable","zh_CN":"页面反调试功能开关。\nON：开启\nOFF：关闭"}
	PageDebugSwitch *string `json:"pageDebugSwitch,omitempty" xml:"pageDebugSwitch,omitempty"`
	// {"en":"Interactive behavior verification function switch.\nON: Enable\nOFF: Disable","zh_CN":"交互行为验证开关。\nON：开启\nOFF：关闭"}
	InteractionAnalyseSwitch *string `json:"interactionAnalyseSwitch,omitempty" xml:"interactionAnalyseSwitch,omitempty"`
	// {"en":"Ajax request exception function switch.ON: EnableOFF: Disable","zh_CN":"Ajax请求例外开关。ON：开启OFF：关闭"}
	AjaxExceptionSwitch *string `json:"ajaxExceptionSwitch,omitempty" xml:"ajaxExceptionSwitch,omitempty"`
}

func (s UpdateBotManagementConfigRequestWebConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigRequestWebConfig) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigRequestWebConfig) SetAct(v string) *UpdateBotManagementConfigRequestWebConfig {
	s.Act = &v
	return s
}

func (s *UpdateBotManagementConfigRequestWebConfig) SetBrowserAnalyseSwitch(v string) *UpdateBotManagementConfigRequestWebConfig {
	s.BrowserAnalyseSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigRequestWebConfig) SetAutoToolSwitch(v string) *UpdateBotManagementConfigRequestWebConfig {
	s.AutoToolSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigRequestWebConfig) SetCrackAnalyseSwitch(v string) *UpdateBotManagementConfigRequestWebConfig {
	s.CrackAnalyseSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigRequestWebConfig) SetPageDebugSwitch(v string) *UpdateBotManagementConfigRequestWebConfig {
	s.PageDebugSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigRequestWebConfig) SetInteractionAnalyseSwitch(v string) *UpdateBotManagementConfigRequestWebConfig {
	s.InteractionAnalyseSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigRequestWebConfig) SetAjaxExceptionSwitch(v string) *UpdateBotManagementConfigRequestWebConfig {
	s.AjaxExceptionSwitch = &v
	return s
}

type UpdateBotManagementConfigRequestTrafficDetection struct {
	// {"en":"Start time, format: HH:mm.","zh_CN":"开始时间，格式：HH:mm。"}
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// {"en":"End time, format: HH:mm.","zh_CN":"结束时间，格式：HH:mm。"}
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// {"en":"Action.\nRESET: Reset connection.\nLOG: Log.\nBLOCK: Block.\nNO_USE: Not Used.","zh_CN":"处理动作。\nRESET：断开连接。\nLOG：监控。\nBLOCK： 拦截。\nNO_USE：不使用。"}
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// {"en":"Whitelist.","zh_CN":"例外。"}
	Whitelist []*string `json:"whitelist" xml:"whitelist" type:"Repeated"`
}

func (s UpdateBotManagementConfigRequestTrafficDetection) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigRequestTrafficDetection) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigRequestTrafficDetection) SetStartTime(v string) *UpdateBotManagementConfigRequestTrafficDetection {
	s.StartTime = &v
	return s
}

func (s *UpdateBotManagementConfigRequestTrafficDetection) SetEndTime(v string) *UpdateBotManagementConfigRequestTrafficDetection {
	s.EndTime = &v
	return s
}

func (s *UpdateBotManagementConfigRequestTrafficDetection) SetAction(v string) *UpdateBotManagementConfigRequestTrafficDetection {
	s.Action = &v
	return s
}

func (s *UpdateBotManagementConfigRequestTrafficDetection) SetWhitelist(v []*string) *UpdateBotManagementConfigRequestTrafficDetection {
	s.Whitelist = v
	return s
}

type UpdateBotManagementConfigResponse struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_retCodeEnum","en":"Please refer to the error code for exceptions.","zh_CN":"请参照错误码。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述信息。"}
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	// {"en":"Data.","zh_CN":"出参数据。"}
	Data *UpdateBotManagementConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s UpdateBotManagementConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponse) SetCode(v string) *UpdateBotManagementConfigResponse {
	s.Code = &v
	return s
}

func (s *UpdateBotManagementConfigResponse) SetMsg(v string) *UpdateBotManagementConfigResponse {
	s.Msg = &v
	return s
}

func (s *UpdateBotManagementConfigResponse) SetData(v *UpdateBotManagementConfigResponseData) *UpdateBotManagementConfigResponse {
	s.Data = v
	return s
}

type UpdateBotManagementConfigResponseData struct {
	// {"en":"Configuration list.","zh_CN":"配置列表。"}
	ConfigList []*UpdateBotManagementConfigResponseDataConfigList `json:"configList,omitempty" xml:"configList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateBotManagementConfigResponseData) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseData) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseData) SetConfigList(v []*UpdateBotManagementConfigResponseDataConfigList) *UpdateBotManagementConfigResponseData {
	s.ConfigList = v
	return s
}

type UpdateBotManagementConfigResponseDataConfigList struct {
	// {"en":"Domain.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Policy switch.ON: Enable.OFF: Disable.","zh_CN":"策略开关。\nON：开启。\nOFF：关闭。","exampleValue":"ON,OFF"}
	ConfigSwitch *string `json:"configSwitch,omitempty" xml:"configSwitch,omitempty" require:"true"`
	// {"en":"General policies.","zh_CN":"通用策略。"}
	GeneralStrategy *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy `json:"generalStrategy,omitempty" xml:"generalStrategy,omitempty" require:"true" type:"Struct"`
	// {"en":"Web risk detection.","zh_CN":"Web风险检测。"}
	WebConfig *UpdateBotManagementConfigResponseDataConfigListWebConfig `json:"webConfig,omitempty" xml:"webConfig,omitempty" require:"true" type:"Struct"`
	// {"en":"Application request whitelist.","zh_CN":"应用请求白名单。"}
	SceneWhitelist []*UpdateBotManagementConfigResponseDataConfigListSceneWhitelist `json:"sceneWhitelist,omitempty" xml:"sceneWhitelist,omitempty" require:"true" type:"Repeated"`
	// {"en":"Abnormal traffic detection.","zh_CN":"异常流量检测。"}
	TrafficDetection *UpdateBotManagementConfigResponseDataConfigListTrafficDetection `json:"trafficDetection,omitempty" xml:"trafficDetection,omitempty" require:"true" type:"Struct"`
	// {"en":"Custom bots.","zh_CN":"自定义Bots。"}
	CustomBots []*UpdateBotManagementConfigResponseDataConfigListCustomBots `json:"customBots,omitempty" xml:"customBots,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateBotManagementConfigResponseDataConfigList) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigList) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigList) SetDomain(v string) *UpdateBotManagementConfigResponseDataConfigList {
	s.Domain = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigList) SetConfigSwitch(v string) *UpdateBotManagementConfigResponseDataConfigList {
	s.ConfigSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigList) SetGeneralStrategy(v *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy) *UpdateBotManagementConfigResponseDataConfigList {
	s.GeneralStrategy = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigList) SetWebConfig(v *UpdateBotManagementConfigResponseDataConfigListWebConfig) *UpdateBotManagementConfigResponseDataConfigList {
	s.WebConfig = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigList) SetSceneWhitelist(v []*UpdateBotManagementConfigResponseDataConfigListSceneWhitelist) *UpdateBotManagementConfigResponseDataConfigList {
	s.SceneWhitelist = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigList) SetTrafficDetection(v *UpdateBotManagementConfigResponseDataConfigListTrafficDetection) *UpdateBotManagementConfigResponseDataConfigList {
	s.TrafficDetection = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigList) SetCustomBots(v []*UpdateBotManagementConfigResponseDataConfigListCustomBots) *UpdateBotManagementConfigResponseDataConfigList {
	s.CustomBots = v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListGeneralStrategy struct {
	// {"en":"AI Bots configuration。","zh_CN":"AI Bots配置。"}
	AiBots []*UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots `json:"aiBots,omitempty" xml:"aiBots,omitempty" require:"true" type:"Repeated"`
	// {"en":"Public Bots configuration.","zh_CN":"公开Bots配置。"}
	PublicBots []*UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots `json:"publicBots,omitempty" xml:"publicBots,omitempty" require:"true" type:"Repeated"`
	// {"en":"Definite Bots actions.\nNO_USE: Not Used.\nLOG: Log.\nBLOCK: Deny.","zh_CN":"绝对Bots处理动作。 \nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。","exampleValue":"NO_USE,LOG,BLOCK"}
	AbsoluteBotsAct *string `json:"absoluteBotsAct,omitempty" xml:"absoluteBotsAct,omitempty" require:"true"`
	// {"en":"Bot Tagging.","zh_CN":"Bot标记。"}
	BotTagging []*UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging `json:"botTagging,omitempty" xml:"botTagging,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateBotManagementConfigResponseDataConfigListGeneralStrategy) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListGeneralStrategy) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy) SetAiBots(v []*UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy {
	s.AiBots = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy) SetPublicBots(v []*UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy {
	s.PublicBots = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy) SetAbsoluteBotsAct(v string) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy {
	s.AbsoluteBotsAct = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy) SetBotTagging(v []*UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategy {
	s.BotTagging = v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_aiBotsCategory","en":"AI Bots categories.","zh_CN":"AI Bots类别。"}
	BotCategory *string `json:"botCategory,omitempty" xml:"botCategory,omitempty" require:"true"`
	// {"en":"Actions. \nNO_USE: Not Used.\nLOG: Log.\nBLOCK: Deny.","zh_CN":"处理动作。 \nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。","exampleValue":"NO_USE,LOG,BLOCK"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots) SetBotCategory(v string) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots {
	s.BotCategory = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots) SetAction(v string) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyAiBots {
	s.Action = &v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots struct {
	// {"dictionary":"belong=WAAP-MS-Ext|dict=waap_publicBotsCategory","en":"Public Bots categories.","zh_CN":"公开Bots类别。"}
	BotCategory *string `json:"botCategory,omitempty" xml:"botCategory,omitempty" require:"true"`
	// {"en":"Actions. \nNO_USE: Not Used.\nLOG: Log.\nBLOCK: Deny.\nACCEPT: Release.","zh_CN":"处理动作。 \nNO_USE：不使用。\nLOG：监控。\nBLOCK：拦截。\nACCEPT：放行。","exampleValue":"NO_USE,LOG,BLOCK,ACCEPT"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots) SetBotCategory(v string) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots {
	s.BotCategory = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots) SetAction(v string) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyPublicBots {
	s.Action = &v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging struct {
	// {"en":"Customizing HTTP request headers.","zh_CN":"自定义HTTP请求头。"}
	RequestHeaderKey *string `json:"requestHeaderKey,omitempty" xml:"requestHeaderKey,omitempty" require:"true"`
	// {"en":"Tagging Bot traffic characteristics.\nAI_BOT: AI Bots.\nPUBLIC_BOT: Public Bots.\nCUSTOMIZE_BOT: Custom Bots.","zh_CN":"标记Bot流量特征。\nAI_BOT：AI Bots。\nPUBLIC_BOT：公开Bots。\nCUSTOMIZE_BOT：自定义Bots。","exampleValue":"AI_BOT,PUBLIC_BOT,CUSTOMIZE_BOT"}
	TrafficCharacteristics *string `json:"trafficCharacteristics,omitempty" xml:"trafficCharacteristics,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging) SetRequestHeaderKey(v string) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging {
	s.RequestHeaderKey = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging) SetTrafficCharacteristics(v string) *UpdateBotManagementConfigResponseDataConfigListGeneralStrategyBotTagging {
	s.TrafficCharacteristics = &v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListWebConfig struct {
	// {"en":"Action.\nNO_USE: Not used.\nLOG: Log.\nBLOCK: Deny.","zh_CN":"处理动作。\nNO_USE：不使用。\nBLOCK：拦截。\nLOG：监控。"}
	Act *string `json:"act,omitempty" xml:"act,omitempty" require:"true"`
	// {"en":"Browser feature switch.\nON: Enable\nOFF: Disable","zh_CN":"浏览器特性功能开关。\nON：开启\nOFF：关闭"}
	BrowserAnalyseSwitch *string `json:"browserAnalyseSwitch,omitempty" xml:"browserAnalyseSwitch,omitempty" require:"true"`
	// {"en":"Automated tool detection function switch.\nON: Enable\nOFF: Disable","zh_CN":"自动化工具检测功能开关。\nON：开启\nOFF：关闭"}
	AutoToolSwitch *string `json:"autoToolSwitch,omitempty" xml:"autoToolSwitch,omitempty" require:"true"`
	// {"en":"Cracking the behavior detection switch.\nON: Enable\nOFF: Disable","zh_CN":"破解行为检测功能开关。\nON：开启\nOFF：关闭"}
	CrackAnalyseSwitch *string `json:"crackAnalyseSwitch,omitempty" xml:"crackAnalyseSwitch,omitempty" require:"true"`
	// {"en":"Page anti-debugging switch.\nON: Enable\nOFF: Disable","zh_CN":"页面反调试功能开关。\nON：开启\nOFF：关闭"}
	PageDebugSwitch *string `json:"pageDebugSwitch,omitempty" xml:"pageDebugSwitch,omitempty" require:"true"`
	// {"en":"Interactive behavior verification switch.\nON: Enable\nOFF: Disable","zh_CN":"交互行为验证开关。\nON：开启\nOFF：关闭"}
	InteractionAnalyseSwitch *string `json:"interactionAnalyseSwitch,omitempty" xml:"interactionAnalyseSwitch,omitempty" require:"true"`
	// {"en":"Rule list of html pages without embedding JS.","zh_CN":"不需要嵌入JS的html页面规则列表。"}
	JsExceptionList []*UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList `json:"jsExceptionList,omitempty" xml:"jsExceptionList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Interactive behavior validation rule list.","zh_CN":"交互行为验证规则列表。"}
	InteractionRuleList []*UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList `json:"interactionRuleList,omitempty" xml:"interactionRuleList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Ajax request exception function switch.ON: EnableOFF: Disable","zh_CN":"Ajax请求例外开关。ON：开启OFF：关闭"}
	AjaxExceptionSwitch *string `json:"ajaxExceptionSwitch,omitempty" xml:"ajaxExceptionSwitch,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigResponseDataConfigListWebConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListWebConfig) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetAct(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.Act = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetBrowserAnalyseSwitch(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.BrowserAnalyseSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetAutoToolSwitch(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.AutoToolSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetCrackAnalyseSwitch(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.CrackAnalyseSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetPageDebugSwitch(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.PageDebugSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetInteractionAnalyseSwitch(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.InteractionAnalyseSwitch = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetJsExceptionList(v []*UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.JsExceptionList = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetInteractionRuleList(v []*UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.InteractionRuleList = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfig) SetAjaxExceptionSwitch(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfig {
	s.AjaxExceptionSwitch = &v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList struct {
	// {"en":"Rule ID of html pages without embedding JS.","zh_CN":"不需要嵌入JS的html页面规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Hostname.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Path.","zh_CN":"验证路径。"}
	Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"Whether it is REGEX.\nfalse: No\ntrue: Yes","zh_CN":"是否正则匹配。\nfalse：否\ntrue：是"}
	Regex *bool `json:"regex,omitempty" xml:"regex,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList) SetId(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList {
	s.Id = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList) SetDomain(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList {
	s.Domain = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList) SetPath(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList {
	s.Path = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList) SetDescription(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList {
	s.Description = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList) SetRegex(v bool) *UpdateBotManagementConfigResponseDataConfigListWebConfigJsExceptionList {
	s.Regex = &v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList struct {
	// {"en":"Interactive behavior validation rule ID.","zh_CN":"交互行为验证规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Hostname.","zh_CN":"域名。"}
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
	// {"en":"Path.","zh_CN":"验证路径。"}
	Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
	// {"en":"Minimum number of triggers.","zh_CN":"最少交互次数。"}
	TriggerTimes *int `json:"triggerTimes,omitempty" xml:"triggerTimes,omitempty" require:"true"`
	// {"en":"Whether it is REGEX.\nfalse: No\ntrue: Yes","zh_CN":"是否正则匹配。\nfalse：否\ntrue：是"}
	Regex *bool `json:"regex,omitempty" xml:"regex,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList) SetId(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList {
	s.Id = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList) SetDomain(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList {
	s.Domain = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList) SetPath(v string) *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList {
	s.Path = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList) SetTriggerTimes(v int) *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList {
	s.TriggerTimes = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList) SetRegex(v bool) *UpdateBotManagementConfigResponseDataConfigListWebConfigInteractionRuleList {
	s.Regex = &v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListSceneWhitelist struct {
	// {"en":"ID.","zh_CN":"ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Whitelist name.","zh_CN":"白名单名称。"}
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"描述。"}
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	// {"en":"List of matching conditions.","zh_CN":"匹配条件列表。"}
	Conditions []*UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateBotManagementConfigResponseDataConfigListSceneWhitelist) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListSceneWhitelist) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListSceneWhitelist) SetId(v string) *UpdateBotManagementConfigResponseDataConfigListSceneWhitelist {
	s.Id = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListSceneWhitelist) SetName(v string) *UpdateBotManagementConfigResponseDataConfigListSceneWhitelist {
	s.Name = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListSceneWhitelist) SetDescription(v string) *UpdateBotManagementConfigResponseDataConfigListSceneWhitelist {
	s.Description = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListSceneWhitelist) SetConditions(v []*UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions) *UpdateBotManagementConfigResponseDataConfigListSceneWhitelist {
	s.Conditions = v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions struct {
	// {"en":"Matching condition name.\nIP_IPS: IP/CIDR\nPATH: Path\nURI: Path with parameters\nHEADER: Request Header\nUA: User Agent\nREQUEST_METHOD: Request Method\nREFERER: Referer","zh_CN":"匹配条件名称。\nIP_IPS：IP/IP段\nPATH：路径\nURI：带参数路径\nHEADER：请求头\nUA: User-Agent\nREQUEST_METHOD: 请求方法\nREFERER: Referer"}
	MatchName *string `json:"matchName,omitempty" xml:"matchName,omitempty" require:"true"`
	// {"en":"When matchName is IP_IPS, maximum 300 IP/CIDR in match value list, the optional value of matchType is:\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nWhen matchName is a URI, the optional value of matchType is:\nEQUAL: Equals, the matching value is case-sensitive and needs to start with \"/\" and include parameters\nNOT_EQUAL: Does not equal, the matching value is case-sensitive, needs to start with \"/\", and contains parameters\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, only one value is allowed for the match value\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is PATH, the optional value of matchType is:\nEQUAL: Equals, the matching value is case-sensitive and needs to start with \"/\" , and does not contain parameters\nNOT_EQUAL: Does not equal, the matching value is case-sensitive, needs to start with \"/\", and does not contain parameters\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is HEADER, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is UA, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is REFERER, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive\nCONTAIN: Contains, match values are case insensitive\nNOT_CONTAIN: Does not contains, match values are case insensitive\nREGEX: Regex match, match values are case insensitive and only one value is allowed\nNONE: Empty or does not exist\nNOT_REGEX: regular does not match\nSTART_WITH: starts with\nEND_WITH: ends with\nWILDCARD: wildcard matches\nNOT_WILDCARD: wildcard does not match\nWhen matchName is REQUEST_METHOD, the optional value of matchType is:\nEQUAL: Equals, match values are case sensitive\nNOT_EQUAL: Does not equal, the matching value is case-sensitive","zh_CN":"当matchName是IP_IPS时，匹配值列表中最多300个IP或IP段，matchType可选值为：\nEQUAL：等于\nNOT_EQUAL：不等于\n当matchName是URI时，matchType可选值为：\nEQUAL：等于匹配值大小写敏感，需要以\"/\"开头，含参数\nNOT_EQUAL：不等于，匹配值大小写敏感，需要以\"/\"开头，含参数\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值只允许有一个值\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是PATH时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感，需要以\"/\"开头，不含参数\nNOT_EQUAL：不等于，匹配值大小写敏感，需要以\"/\"开头，不含参数\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值大小写不敏感，只允许有一个值\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是HEADER时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是UA时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是REFERER时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感\nCONTAIN：包含，匹配值大小写不敏感\nNOT_CONTAIN：不包含，匹配值大小写不敏感\nREGEX：匹配正则，匹配值大小写不敏感，只允许有一个值\nNONE：为空或不存在\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配\nNOT_WILDCARD：通配符不匹配\n当matchName是REQUEST_METHOD时，matchType可选值为：\nEQUAL：等于，匹配值大小写敏感\nNOT_EQUAL：不等于，匹配值大小写敏感"}
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty" require:"true"`
	// {"en":"Mathing key, this value is not empty and valid only when matchName=HEADER.\nMaximum 100 characters, case insensitive.","zh_CN":"条件key，只有当matchName=HEADER时，该值不为空且有效。\n最多100个字符，大小写不敏感。"}
	MatchKey *string `json:"matchKey,omitempty" xml:"matchKey,omitempty" require:"true"`
	// {"en":"List of matching values.","zh_CN":"匹配值列表。"}
	MatchValueList []*string `json:"matchValueList,omitempty" xml:"matchValueList,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions) SetMatchName(v string) *UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions {
	s.MatchName = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions) SetMatchType(v string) *UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions {
	s.MatchType = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions) SetMatchKey(v string) *UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions {
	s.MatchKey = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions) SetMatchValueList(v []*string) *UpdateBotManagementConfigResponseDataConfigListSceneWhitelistConditions {
	s.MatchValueList = v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListTrafficDetection struct {
	// {"en":"Start time, format: HH:mm.","zh_CN":"开始时间，格式：HH:mm。"}
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
	// {"en":"End time, format: HH:mm.","zh_CN":"结束时间，格式：HH:mm。"}
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
	// {"en":"Action.\nRESET: Reset connection.\nLOG: Log.\nBLOCK: Block.\nNO_USE: Not Used.","zh_CN":"处理动作。\nRESET：断开连接。\nLOG：监控。\nBLOCK： 拦截。\nNO_USE：不使用。"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {"en":"Whitelist.","zh_CN":"例外。"}
	Whitelist []*string `json:"whitelist,omitempty" xml:"whitelist,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateBotManagementConfigResponseDataConfigListTrafficDetection) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListTrafficDetection) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListTrafficDetection) SetStartTime(v string) *UpdateBotManagementConfigResponseDataConfigListTrafficDetection {
	s.StartTime = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListTrafficDetection) SetEndTime(v string) *UpdateBotManagementConfigResponseDataConfigListTrafficDetection {
	s.EndTime = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListTrafficDetection) SetAction(v string) *UpdateBotManagementConfigResponseDataConfigListTrafficDetection {
	s.Action = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListTrafficDetection) SetWhitelist(v []*string) *UpdateBotManagementConfigResponseDataConfigListTrafficDetection {
	s.Whitelist = v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListCustomBots struct {
	// {"en":"Rule ID.","zh_CN":"规则ID。"}
	Id *string `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// {"en":"Rule Name.","zh_CN":"规则名称。"}
	BotName *string `json:"botName,omitempty" xml:"botName,omitempty" require:"true"`
	// {"en":"Description.","zh_CN":"规则描述。"}
	BotDescription *string `json:"botDescription,omitempty" xml:"botDescription,omitempty" require:"true"`
	// {"en":"Actions:\nBLOCK: block\nLOG: log\nACCEPT: release","zh_CN":"处理动作：\nBLOCK：拦截\nLOG：监控\nACCEPT：放行"}
	BotAct *string `json:"botAct,omitempty" xml:"botAct,omitempty" require:"true"`
	// {"en":"Matching conditions.","zh_CN":"匹配条件。"}
	ConditionList []*UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList `json:"conditionList,omitempty" xml:"conditionList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Creator.","zh_CN":"创建者。"}
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigResponseDataConfigListCustomBots) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListCustomBots) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBots) SetId(v string) *UpdateBotManagementConfigResponseDataConfigListCustomBots {
	s.Id = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBots) SetBotName(v string) *UpdateBotManagementConfigResponseDataConfigListCustomBots {
	s.BotName = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBots) SetBotDescription(v string) *UpdateBotManagementConfigResponseDataConfigListCustomBots {
	s.BotDescription = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBots) SetBotAct(v string) *UpdateBotManagementConfigResponseDataConfigListCustomBots {
	s.BotAct = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBots) SetConditionList(v []*UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList) *UpdateBotManagementConfigResponseDataConfigListCustomBots {
	s.ConditionList = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBots) SetCreator(v string) *UpdateBotManagementConfigResponseDataConfigListCustomBots {
	s.Creator = &v
	return s
}

type UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList struct {
	// {"en":"Matching condition name. \nIP_IPS: IP/IP segment \nJA3: JA3 Fingerprint\nJA4: JA4 Fingerprint\nUA: User-agent \nHEADER: Request Header \nASN: AS Number \nCLIENT_GROUP: Client Group \nPUBLIC_BOT: Public Bots","zh_CN":"匹配条件名称。 \nIP_IPS：IP/IP段 \nJA3：JA3指纹\nJA4：JA4指纹\nUA：User-agent \nHEADER：请求头 \nASN：AS号 \nCLIENT_GROUP：客户端分组 \nPUBLIC_BOT：公开Bots"}
	ConditionName *string `json:"conditionName,omitempty" xml:"conditionName,omitempty" require:"true"`
	// {"en":"Condition value list.","zh_CN":"条件值列表。"}
	ConditionValueList []*string `json:"conditionValueList,omitempty" xml:"conditionValueList,omitempty" require:"true" type:"Repeated"`
	// {"en":"Matching condition function.\nEQUAL: Equals\nNOT_EQUAL: Does not equal\nCONTAIN: Contains\nNOT_CONTAIN: Does not Contains\nNONE:Empty or non-existent\nREGEX: Regex match\nNOT_REGEX: Regular does not match\nSTART_WITH: Starts with\nEND_WITH: Ends with\nWILDCARD: Wildcard matches, * represents zero or more arbitrary characters, ? represents any single character\nNOT_WILDCARD: Wildcard does not match, * represents zero or more arbitrary characters, ? represents any single character\n\nWhen conditionName is IP_IPS, the value can be:EQUAL,NOT_EQUAL \nWhen conditionName is JA3, the value can be:EQUAL,NOT_EQUAL\nWhen conditionName is JA4, the value can be:EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD\nWhen conditionName is UA, the value can be:EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD \nWhen conditionName is HEADER, the value can be:EQUAL,NOT_EQUAL,CONTAIN,NOT_CONTAIN,NONE,REGEX,NOT_REGEX,START_WITH,END_WITH,WILDCARD,NOT_WILDCARD \nWhen conditionName is ASN, the value can be:EQUAL,NOT_EQUAL \nWhen conditionName is CLIENT_GROUP, the value can be:EQUAL,NOT_EQUAL \nWhen conditionName is PUBLIC_BOT, the value can be:EQUAL,NOT_EQUAL","zh_CN":"匹配条件函数。\nEQUAL：等于\nNOT_EQUAL：不等于\nCONTAIN：包含\nNOT_CONTAIN：不包含\nNONE：为空或不存在\nREGEX：匹配正则\nNOT_REGEX：正则不匹配\nSTART_WITH：开头是\nEND_WITH：结尾是\nWILDCARD：通配符匹配，*代表零个或多个任意字符，?代表任意单个字符\nNOT_WILDCARD：通配符不匹配，*代表零个或多个任意字符，?代表任意单个字符\n\n当conditionName为IP_IPS时，取值范围为：EQUAL、NOT_EQUAL \n当conditionName为JA3时，取值范围为：EQUAL、NOT_EQUAL \n当conditionName为JA4时，取值范围为：EQUAL、NOT_EQUAL、CONTAIN、NOT_CONTAIN、START_WITH、END_WITH、WILDCARD、NOT_WILDCARD\n当conditionName为UA时，取值范围为：EQUAL、NOT_EQUAL、CONTAIN、NOT_CONTAIN、NONE、REGEX、NOT_REGEX、START_WITH、END_WITH、WILDCARD、NOT_WILDCARD\n当conditionName为HEADER时，取值范围为：EQUAL、NOT_EQUAL、CONTAIN、NOT_CONTAIN、NONE、REGEX、NOT_REGEX、START_WITH、END_WITH、WILDCARD、NOT_WILDCARD\n当conditionName为ASN时，取值范围为：EQUAL、NOT_EQUAL \n当conditionName为CLIENT_GROUP时，取值范围为：EQUAL、NOT_EQUAL \n当conditionName为PUBLIC_BOT时，取值范围为：EQUAL"}
	ConditionFunc *string `json:"conditionFunc,omitempty" xml:"conditionFunc,omitempty" require:"true"`
	// {"en":"Request header name.case insensitive.","zh_CN":"头部名称，大小写不敏感。"}
	ConditionKey *string `json:"conditionKey,omitempty" xml:"conditionKey,omitempty" require:"true"`
}

func (s UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList) String() string {
	return tea.Prettify(s)
}

func (s UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList) GoString() string {
	return s.String()
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList) SetConditionName(v string) *UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList {
	s.ConditionName = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList) SetConditionValueList(v []*string) *UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList {
	s.ConditionValueList = v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList) SetConditionFunc(v string) *UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList {
	s.ConditionFunc = &v
	return s
}

func (s *UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList) SetConditionKey(v string) *UpdateBotManagementConfigResponseDataConfigListCustomBotsConditionList {
	s.ConditionKey = &v
	return s
}
