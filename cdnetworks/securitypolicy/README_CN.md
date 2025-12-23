# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 单独安装

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/securitypolicy
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/securitypolicy"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &securitypolicy.ActionNameRequest{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := securitypolicy.{ActionName}Response{}
    _, err := auth.Invoke(config, request, response)

    // Handle response
    if err != nil {
        log.Printf("error: %s\n", err)
        return
    }

    log.Printf("response body: %s\n", response.String())
}
```

## 错误处理

始终检查 API 调用返回的错误：

```go
_, err := auth.Invoke(config, request, response)
if err != nil {
    log.Printf("error: %s\n", err)
    // Handle the error appropriately
    return
}
```

## API列表
有关详细的 API 文档和可用方法，请参阅[官方 Cdnetworks API 文档](https://docs.cdnetworks.com/en/cdn/apidocs)。

| ActionName | description | client_methods | uri |
| --- | --- | --- | --- |
| Getblockpagesetting | 查询拦截页面定义。 | POST | /api/waf/domain/query-block-page-conf |
| Updateblockpagesetting | 修改拦截页面定义。 | POST | /api/waf/domain/update-block-page-conf |
| Getbuiltinrulelist | 查询内置规则列表。 | POST | /api/waf/template/rule/list-page |
| Getbuiltinruleinfo | 查询内置规则详情。 | POST | /api/waf/template/rule/detail |
| Createbuiltinrule | 新增内置规则。 | POST | /api/waf/template/rule/add |
| Updatebuiltinrule | 修改内置规则。 | POST | /api/waf/template/rule/update |
| Associatedomainsforbuiltinrule | 添加内置规则关联域名。 | POST | /api/waf/template/rule/add-rela |
| Removedomainsforbuiltinrule | 从指定的内置规则中移除对应的域名。 | POST | /api/waf/template/rule/remove-rela |
| Getcustomrulelist | 查询自定义规则列表。 | POST | /api/waf/custom/rule/list-page |
| Createcustomrule | 新增自定义规则。 | POST | /api/waf/custom/rule/add |
| Updatewafcustomrule | 修改自定义规则。 | POST | /api/waf/custom/rule/update |
| Deletecustomrule | 删除自定义规则。 | POST | /api/waf/custom/rule/delete |
| Associatedomainsforcustomrule | 添加自定义规则关联域名。 | POST | /api/waf/custom/rule/add-rela |
| Removedomainsforcustomrule | 从指定的自定义规则中移除对应的域名。 | POST | /api/waf/custom/rule/remove-rela |
| Getruleexceptionlist | 查询规则例外的规则列表。 | POST | /api/waf/template/rule/white/list-page |
| Createruleexception | 新增规则例外。 | POST | /api/waf/template/rule/white/add |
| Updateruleexception | 修改规则例外。 | POST | /api/waf/template/rule/white/update |
| Deleteruleexception | 删除规则例外。 | POST | /api/waf/template/rule/white/delete |
| Associatedomainsforruleexception | 添加规则例外关联域名。 | POST | /api/waf/template/rule/white/add-rela |
| Removedomainsforruleexception | 从指定的规则例外中移除对应的域名。 | POST | /api/waf/template/rule/white/remove-rela |
| Querycustomruleconfiguration | 查询自定义规则配置（安全策略-频率限制-自定义规则） | POST | /dms/api/getDomainRateLimitConfig |
| ModifyBuiltInProtection | 修改域名的内置防护开关和内置防护模式 | POST | /dms/api/UpdateL7InnerTmode |
| Updateaslconfig | 修改自定义规则配置（仅支持单个域名） | POST | /dms/api/UpdateL7DomainAslConfig |
| Getcrawlergood | 查询Bot情报 | POST | /api/bot/domainConfig/get-crawler-good |
| Querybotfeatureverification | 查询Bot特性检测 | POST | /api/bot/domainConfig/get-bot-challenge |
| Queryfingerprintanalysis | 查询指纹分析 | POST | /api/bot/domainConfig/get-fingerprint-challenge |
| Querycaptchaverification | 查询验证码校验 | POST | /api/bot/domainConfig/get-captcha-challenge |
| Getbehavioranalyse | 查询业务流分析 | POST | /api/bot/domainConfig/get-behavior-analyse |
| Getblockpage | 查询拦截页面 | POST | /api/bot/domainConfig/get-block-page |
| Deploycrawlergood | 部署Bot情报。注:该接口需要全量提交数据 | POST | /api/bot/domainConfig/deploy-crawler-good |
| Deploybotfeatureverification | 部署Bot特性检测 | POST | /api/bot/domainConfig/deploy-bot-challenge |
| Deployfingerprintanalysis | 部署指纹分析 | POST | /api/bot/domainConfig/deploy-fingerprint-challenge |
| Deploycaptchaverification | 部署验证码校验 | POST | /api/bot/domainConfig/deploy-captcha-challenge |
| Deploybehavioranalyse | 部署业务流分析。注:该接口需要全量提交数据 | POST | /api/bot/domainConfig/deploy-behavior-analyse |
| Deployblockpage | 部署拦截页面 | POST | /api/bot/domainConfig/deploy-block-page |
| Createconsumer | 新增消费方。 | POST | /api/sam/consumers/add |
| Editconsumer | 编辑消费方。 | POST | /api/sam/consumers/edit |
| Deleteconsumer | 删除消费方。 | POST | /api/sam/consumers/delete |
| Getconsumerlist | 获取消费方列表。 | POST | /api/sam/consumers/list |
| Getconsumerinfo | 获取消费方信息。 | GET | /api/sam/consumers |
| Getapiassetlist | 获取API资产列表 | POST | /api/sam/api-discovery/list |
| Getapiassetdetail | 获取API资产详情。 | POST | /api/sam/api-discovery/detail |
| Reportincorrectapiasset | 上报错误的API资产发现数据。 | POST | /api/sam/api-discovery/false-marking |
| Createandactiveapi | 新增并激活API。 | POST | /api/sam/api/add-active |
| Createquotarule | 新增配额规则。 | POST | /api/sam/policy-quota/add |
| Editquotarule | 修改配额规则。 | POST | /api/sam/policy-quota/edit |
| Deletequotarule | 删除配额规则。 | POST | /api/sam/policy-quota/delete |
| Getapilist | 获取api列表。 | POST | /api/sam/api/list |
| Getquotarulelist | 获取配额规则列表。 | GET | /api/sam/policy-quota/list |
| Deleteapi | 删除API。 | POST | /api/sam/api/delete |
| Editandactiveapi | 编辑API并部署激活。 | POST | /api/sam/api/edit |
| Createconcurrencylimitrule | 新增限流规则。 | POST | /api/sam/policy-flow/add |
| Editconcurrencylimitrule | 修改限流规则。 | POST | /api/sam/policy-flow/edit |
| Deleteconcurrencylimitrule | 删除限流规则。 | POST | /api/sam/policy-flow/delete |
| Getconcurrencylimitrulelist | 获取限流规则列表。 | GET | /api/sam/policy-flow/list |
| Queryaccesscontrolrulelist | 查询Bot规则 | POST | /api/bot/customAccessControl/get-rule-listcmm |
| Wssmpnetworkserviceadd | 非网站业务规则新增 | POST | /soc/WssMPTcpService/add |
| Wssmpnetworkservicedelete | 非网站业务规则删除 | POST | /soc/WssMPTcpService/delete |
| Wssmpnetworkserviceupdate | 非网站业务规则修改 | POST | /soc/WssMPTcpService/update |
| Wssmpnetworkservicequery | 非网站业务规则查询 | POST | /soc/WssMPTcpService/query |
| Deletebuiltinrule | 删除内置规则。 | POST | /api/waf/policy/delete-builtInRule |
| Setdefaultbuiltinrule | 设置默认内置规则。<br> | POST | /api/waf/policy/set-default-builtInRule |
| Getintelligentanalysislist | 查询规则托管列表。 | POST | /api/waf/policy/get-intelligentAnalysis-list |
| Rejectintelligentanalysis | 拒绝规则托管建议。 | POST | /api/waf/policy/reject-intelligentAnalysis |
| Deleteintelligentanalysis | 删除规则托管建议。 | POST | /api/waf/policy/delete-intelligentAnalysis |
| Queryexactrulelist(accessControl) | 查询规则列表（精准访问控制） | POST | /api/bot/exact/query-exact-rule-list |
| Querydomainlistbyrulename(accessControl) | 查询规则关联域名列表（精准访问控制） | POST | /api/bot/exact/query-domainlist-by-rulename |
| Addexactrule(accessControl) | 添加规则（精准访问控制） | POST | /api/bot/exact/add-exact-rule |
| Editexactrule(accessControl) | 编辑规则（精准访问控制） | POST | /api/bot/exact/edit-exact-rule |
| Deleteexactrule(accessControl) | 删除规则（精准访问控制） | POST | /api/bot/exact/delete-exact-rule |
| Batchassiociatedexactrule(accessControl) | 批量添加域名（精准访问控制） | POST | /api/bot/exact/batch-assiociated-exact-rule |
| Batchremoveexactrule(accessControl) | 批量移除域名（精准访问控制） | POST | /api/bot/exact/batch-remove-exact-rule |
| Queryexactrulelistbydomain(accessControl) | 根据域名查询规则列表（精准访问控制） | POST | /api/bot/exact/query-exact-rulelist-by-domain |
| Updatedomainbaseinfo | 修改域名基础开关/模式信息。 | POST | /api/v1/dms/ddosProtect/updateDomainBaseInfo |
| Querydomainbuiltinrules | 查询域名内置规则 | POST | /api/v1/dms/ddosProtect/queryDomainBuiltInRules |
| Updatedomainbuiltinrules | 修改域名内置规则 | POST | /api/v1/dms/ddosProtect/updateDomainBuiltInRules |
| Sharedomainbaseprotectconfig | 跨域名部署基础配置 | POST | /api/v1/dms/ddosProtect/shareDomainBaseProtectConfig |
| Sharedmsbuiltinrulesconfig | 跨域名部署全局内置规则 | POST | /api/v1/dms/ddosProtect/shareDmsBuiltInRulesConfig |
| Sharedmsnativeappapiconfig | 跨域名部署例外api | POST | /api/v1/dms/ddosProtect/shareDmsNativeAppApiConfig |
| Querydomainddosbaseinfo | 查询域名DDoS防护基础策略。 | POST | /api/v1/dms/ddosProtect/queryDomainBaseInfo |
| Queryadaptiveprotectionrules | 查询域名AI智能防护规则。 | POST | /api/v1/dms/ddosProtect/queryDomainAIRules |
| Updateadaptiveprotectionrules | 修改域名AI智能防护规则 | POST | /api/v1/dms/ddosProtect/updateDomainAIRules |
| Querydomainappapiexceptions | 查询域名关联的共享配置-App/API例外 | POST | /api/v1/dms/ddosProtect/queryDomainAppApiExceptionRules |
| Removedomainappapiexceptions | 移除域名的APP/API例外 | POST | /api/v1/dms/ddosProtect/removeDomainAppApiExceptionRules |
| Relatedappapiexceptionsfordomain | 域名关联App/API例外规则 | POST | /api/v1/dms/ddosProtect/relateDomainAppApiExceptionRules |
| Listnonsharedwafruleexceptionsforwafrules | 获取WAF规则的例外配置。 | POST | /api/v1/waf/rule-exception/list-normal |
| Domainsofnonlatestrulesetversion | 获取规则集版本不是最新版本的域名。 | GET | /api/v1/waf/conf/base/get-upgrade-domain-list |
| Listupgradedetails | 返回指定域名待更新的WAF规则信息，包括规则的名称、规则类型、描述等信息。 | POST | /api/v1/waf/rule/get-upgrade-rule-list |
| Upgradewafruleset | 将指定域名的WAF规则集升级到最新版本。 | POST | /api/v1/waf/rule/batch/upgrade |
| Listwafrules | 根据规则ID/规则名称/规则描述/漏洞编号查询特定的WAF规则。 | POST | /api/v1/waf/rule/get-rule-list |
| Updateactionforwafmanagedrules | 修改WAF内置规则的动作。 | POST | /api/v1/waf/rule/batch/update |
| Listwafbasicconfigofdomains | 返回指定域名的WAF防护模式、规则集模式和规则集版本号。 | POST | /api/v1/waf/conf/base/get-basic-conf-list |
| Updatemodeofwaf | 修改域名的防护模式和规则集模式。 | POST | /api/v1/waf/conf/base/update-basic-conf |
| Createexceptiontowafmanagedrules | 新增WAF内置规则的例外配置。 | POST | /api/v1/waf/rule-exception/add-normal |
| Associatesharewafruleexception | 将共享配置中的WAF例外关联给指定域名。 | POST | /api/v1/waf/rule-exception/add-share |
| Deleteexceptionforwafmanagedrules | 删除WAF内置规则的WAF规则例外。 | POST | /api/v1/waf/rule-exception/delete-normal |
| Updateexceptionforwafmanagedrules | 修改WAF内置规则的WAF规则例外。 | POST | /api/v1/waf/rule-exception/update |
| Deletesharedwafexceptionsassociatedwithwafmanagedrules | 删除WAF规则例外（来源：共享配置）。 | POST | /api/v1/waf/rule-exception/delete-share |
| Applyrecommendations | 将建议添加到例外。 | POST | /api/v1/waf/ai-rule-result/apply |
| Rejectrecommendations | 拒绝WAF规则例外建议。 | POST | /api/v1/waf/ai-rule-result/reject |
| Listrecommendations | 返回指定WAF规则的建议，包括待处理和已拒绝的建议。 | POST | /api/v1/waf/ai-rule-result/list |
| Listsharedwafruleexceptionsforwafrules | 返回指定WAF规则ID下的的例外配置（来源：共享配置）。 | POST | /api/v1/waf/rule-exception/list-share |
| Listdomaininfos | 获取域名信息列表。 | POST | /api/v1/common/sys-domain-info/get-list |
| Usingexistinghostnametoaddnewhostname | 参照防护域名的配置接入新的域名。 | POST | /api/v1/common/sys-domain-info/add-refer-to-other-domain |
| Copypoliciestootherhostnames | 将指定域名的策略复制给其他域名。<br>注意：<br>1. 不允许复制给未开启防护的策略；<br>2. 自定义规则和频率限制中有引用已定义API的规则不支持复制，策略复制过程中将跳过此类规则。 | POST | /api/v1/common/sys-domain-info/copy-domain |
| Modifypolicystatus | 修改策略开关。 | POST | /api/v1/common/sys-domain-info/update-switch |
| Removeprotecteddomain | 移除防护域名。 | GET | /api/v1/common/sys-domain-info/remove |
| Usingsystemrecommendedaccessdomain | 使用系统建议接入域名。 | POST | /api/v1/common/sys-domain-info/add-sys-domain |
| Getbotfunctionswitch | 查询Bot管理功能开关。 | POST | /api/v1/bot-manage/get-function-switch |
| Updatebotfunctionswitch | 修改Bot管理功能开关。 | POST | /api/v1/bot-manage/update-function-switch |
| Addratelimitingrule | 新增频率限制规则。 | POST | /api/v1/rate-limit/add-rule |
| Createwhitelistrule | 新增白名单规则。 | POST | /api/v1/common/whitelist/add |
| Updateratelimitingrule | 修改频率限制规则的配置。 | POST | /api/v1/rate-limit/update-rule |
| Updatewhitelistrule | 修改白名单规则的配置。 | POST | /api/v1/common/whitelist/update |
| Deletewhitelistrules | 删除白名单规则。 | POST | /api/v1/common/whitelist/delete |
| Addworkflowrule | 新增业业务流检测规则。 | POST | /api/v1/bot-manage/behavior/add-rule |
| Listratelimitingrules | 获取频率限制的规则列表。 | POST | /api/v1/rate-limit/get-rule-list |
| Updateworkflowrule | 修改业务流检测规则。 | POST | /api/v1/bot-manage/behavior/update-rule |
| Deleteratelimitingrules | 删除频率限制规则。 | POST | /api/v1/rate-limit/delete-rule-by-ids |
| Listworkflowrules | 查询业务流检测规则列表。 | POST | /api/v1/bot-manage/behavior/get-rule |
| Getratelimitingrulesforthesharedconfigurationasociatedwithdomain | 获取域名关联的共享配置的频率限制规则。 | POST | /api/v1/rate-limit/get-relation-by-domain |
| Deleteworkflowrule | 删除业务流检测规则。 | POST | /api/v1/bot-manage/behavior/delete-rule |
| Getthreatintelligencedomainconfig | 获取域名的威胁情报配置。 | POST | /api/v1/common/intelligence/query-list |
| Updatethreatintelligencedomainconfig | 修改威胁情报域名配置。 | POST | /api/v1/common/intelligence/update |
| Getdomainapisecurityconfiguration | 获取域名API安全配置。 | POST | /api/v1/sam/api-defend/get-list-by-domains |
| Updatedomainapisecurityconfiguration | 修改域名API安全配置。 | POST | /api/v1/sam/api-defend/update-defend-action |
| Addcustomizerule | 新增自定义规则。 | POST | /api/v1/customize-rule/add |
| Listwhitelistrules | 获取白名单规则列表。 | POST | /api/v1/common/whitelist/query-list |
| Getwhiterulesforthesharedconfigurationasociatedwithdomain | 获取域名关联的共享配置的白名单规则。 | GET | /api/v1/common/whitelist/get-imported-share-config |
| Updateipblocksettings | 更新IP封禁的配置，拦截特定IP/IP段发起的请求。注意：将覆盖原来的配置。 | POST | /api/v1/policy-block/save-deploy-ip-block |
| Updatecustomrule | 修改自定义规则的配置。 | POST | /api/v1/customize-rule/update |
| Deletecustomizerule | 删除自定义规则。 | POST | /api/v1/customize-rule/delete |
| Listcustomrules | 获取自定义规则列表。 | POST | /api/v1/customize-rule/get-list |
| Listipblocksettings | 获取域名的IP封禁配置。 | POST | /api/v1/policy-block/get-ip-block-list |
| Listgeoblocksettings | 返回指定域名已设置拦截的国家和地区。 | POST | /api/v1/policy-block/get-geo-block-list |
| Listcustomrulesforsharedconfigurationassociatedwithdomain | 获取域名关联的共享配置的自定义规则。 | POST | /api/v1/customize-rule/get-import-share-rule |
| Updategeoblocksettings | 拦截来自特定国家/地区的请求。注意：将覆盖原来的配置。 | POST | /api/v1/policy-block/save-deploy-geo-block |
| Listdomaincustomizebots | 查询域名的自定义Bot列表。 | POST | /api/v1/bot-manage/get-domain-customize-bots-list |
| Listspecificclienttrafficbypass | 查询应用请求白名单列表。 | POST | /api/v1/bot-manage/scene/whitelist/list |
| Addspecificclienttrafficbypass | 新增应用白名单。 | POST | /api/v1/bot-manage/scene/whitelist/add |
| Updatespecificclienttrafficbypass | 更新应用白名单。 | POST | /api/v1/bot-manage/scene/whitelist/update |
| Deletespecificclienttrafficbypass | 删除应用请求白名单。 | POST | /api/v1/bot-manage/scene/whitelist/delete |
| Getwebbotdetectionlist | 查询Web风险检测配置。 | POST | /api/v1/bot-manage/web/list |
| Updatewebbotdetection | 更新Web风险检测配置。<br>浏览器特性功能开关browserAnalyseSwitch默认为ON，并且这个字段不允许更新。 | POST | /api/v1/bot-manage/web/update |
| Addhtmlpageswithoutembeddingjs | 针对Web风险检测，新增不需要嵌入JS的html页面。 | POST | /api/v1/bot-manage/scene/web/js-exception/add |
| Updatehtmlpageswithoutembeddingjs | 针对Web风险检测，更新不需要嵌入JS的html页面。 | POST | /api/v1/bot-manage/scene/web/js-exception/update |
| Deletehtmlpageswithoutembeddingjs | 针对Web风险检测，删除不需要嵌入JS的html页面。 | POST | /api/v1/bot-manage/scene/web/js-exception/delete |
| Disableallpolicies | 禁用指定域名的全部安全策略，被禁用的域名将不再进行防护。 | POST | /api/v1/common/sys-domain-info/close-all-policy-switch |
| Updateknownbotsselectbotnames | 批量修改生效的已知Bot小类。 | POST | /api/v1/bot-manage/update-known-bots-select-bot-names |
| Listknownbots | 查询已知Bot列表。 | POST | /api/v1/bot-manage/get-known-bots |
| Updateknownbotsact | 批量修改已知Bot处理动作。 | POST | /api/v1/bot-manage/update-known-bots-act |
| Listuabots | 查询UA特征检测列表。 | POST | /api/v1/bot-manage/get-ua-bots |
| Updateuabotsselectbotnames | 批量修改生效的UA特征检测小类。 | POST | /api/v1/bot-manage/update-ua-bots-select-bot-names |
| Updateuabotsact | 批量修改UA特征检测处理动作。 | POST | /api/v1/bot-manage/update-ua-bots-act |
| Getresponsepageofdenyactiondetail | 获取基础配置-自定义拦截响应详情 | POST | /api/v1/other/sys-domain-basic/response/detail |
| Updateresponsepageofdenyactiondetail | 修改基础配置-自定义拦截响应详情 | POST | /api/v1/other/sys-domain-basic/response/update |
| Listhistoricalhostnames | 查询历史防护的域名。 | POST | /api/v1/other/domain/get-permanent-domain-list |
| Getwafscanprotectionconfig | 该接口用于查询指定域名的WAF扫描防护配置。用户可以通过提供域名来获取该域名的WAF扫描防护配置，包括扫描工具检测配置、反复违规检测配置、目录嗅探检测配置。 | POST | /api/v1/security-policy/waf/get-scan-protection-configuration |
| Updatewafscanprotectionconfig | 该接口用于批量修改指定域名的WAF扫描防护配置。用户可以通过该接口批量修改域名的扫描工具检测配置、反复违规检测配置、目录嗅探检测配置。 | POST | /api/v1/security-policy/waf/update-scan-protection-configuration |
| Getpredeployresult | 该接口用于查询预部署的结果。用户可以通过提供预部署id来获取预部署的结果，包括部署结果和Host信息。 | GET | /api/v1/security-policy/get-pre-deploy-result |
| Updatedetectionlengthlimitconfiguration | 该接口用于批量修改指定域名的检测长度限制配置。用户可以通过该接口批量修改域名的请求正文最大检测长度配置。 | POST | /api/v1/security-policy/basic/update-detection-length-limit-config |
| Getdetectionlengthlimitconfiguration | 该接口用于查询指定域名的检测长度限制配置。用户可以通过提供域名来获取该域名检测长度校验配置，包括请求正文最大检测长度。 | POST | /api/v1/security-policy/basic/get-detection-length-limit-config |
| Getbotgeneralconfig | 该接口用于查询Bot通用策略配置，包括公开Bots、AI Bots、绝对Bots和Bot标记的详细配置信息 | POST | /api/v1/security-policy/bot/get-general-config |
| Updatebotgeneralconfig | 该接口用于修改Bot通用策略配置，包括公开Bots、AI Bots、绝对Bots和Bot标记的详细配置信息 | POST | /api/v1/security-policy/bot/update-general-config |