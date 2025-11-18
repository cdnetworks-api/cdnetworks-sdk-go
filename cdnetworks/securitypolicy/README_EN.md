# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/securitypolicy
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/securitypolicy"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
    request := &securitypolicy.{ActionName}Request{}

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

## Error Handling

Always check for errors returned by the API calls:

```go
_, err := auth.Invoke(config, request, response)
if err != nil {
    log.Printf("error: %s\n", err)
    // Handle the error appropriately
    return
}
```

## API List
For detailed API documentation and available methods, please refer to the [official Cdnetworks API documentation](https://docs.cdnetworks.com/en/cdn/apidocs).

| ActionName | enDescription | client_methods | uri |
| --- | --- | --- | --- |
| Getblockpagesetting | Query the block page setting. | POST | /api/waf/domain/query-block-page-conf |
| Updateblockpagesetting | Update the block page setting. | POST | /api/waf/domain/update-block-page-conf |
| Getbuiltinrulelist | Query the list of built-in rule. | POST | /api/waf/template/rule/list-page |
| Getbuiltinruleinfo | Query the built-in rule. | POST | /api/waf/template/rule/detail |
| Createbuiltinrule | Create the built-in rule. | POST | /api/waf/template/rule/add |
| Updatebuiltinrule | Update the built-in rule. | POST | /api/waf/template/rule/update |
| Associatedomainsforbuiltinrule | Associate the domains with built-in rule. | POST | /api/waf/template/rule/add-rela |
| Removedomainsforbuiltinrule | Remove the domains from the built-in rule. | POST | /api/waf/template/rule/remove-rela |
| Getcustomrulelist | Query the rule list of custom rule. | POST | /api/waf/custom/rule/list-page |
| Createcustomrule | Create custom rule. | POST | /api/waf/custom/rule/add |
| Updatewafcustomrule | Update custom rule. | POST | /api/waf/custom/rule/update |
| Deletecustomrule | Delete custom rule. | POST | /api/waf/custom/rule/delete |
| Associatedomainsforcustomrule | Associate the domains with custom rule. | POST | /api/waf/custom/rule/add-rela |
| Removedomainsforcustomrule | Remove the domains from the specified custom rule. | POST | /api/waf/custom/rule/remove-rela |
| Getruleexceptionlist | Query the rule list of rule exception. | POST | /api/waf/template/rule/white/list-page |
| Createruleexception | Create the rule exception. | POST | /api/waf/template/rule/white/add |
| Updateruleexception | Update the rule exception. | POST | /api/waf/template/rule/white/update |
| Deleteruleexception | Delete the rule exception. | POST | /api/waf/template/rule/white/delete |
| Associatedomainsforruleexception | Associate the domains with rule exception. | POST | /api/waf/template/rule/white/add-rela |
| Removedomainsforruleexception | Remove the domains from the specified rule exception. | POST | /api/waf/template/rule/white/remove-rela |
| Querycustomruleconfiguration | Query custom rule configuration (Security Policy - L7 DDoS Policy - Custom Rule). | POST | /dms/api/getDomainRateLimitConfig |
| ModifyBuiltInProtection | 修改域名的内置防护开关和内置防护模式 | POST | /dms/api/UpdateL7InnerTmode |
| Updateaslconfig | Update custom rule configuration(Only a single domain is supported) | POST | /dms/api/UpdateL7DomainAslConfig |
| Getcrawlergood | Get bot intelligence | POST | /api/bot/domainConfig/get-crawler-good |
| Querybotfeatureverification | Get bot challenge rules | POST | /api/bot/domainConfig/get-bot-challenge |
| Queryfingerprintanalysis | Query fingerprint analysis rules | POST | /api/bot/domainConfig/get-fingerprint-challenge |
| Querycaptchaverification | Query captcha verification rules | POST | /api/bot/domainConfig/get-captcha-challenge |
| Getbehavioranalyse | Get behavior analyse rules | POST | /api/bot/domainConfig/get-behavior-analyse |
| Getblockpage | Get block page content | POST | /api/bot/domainConfig/get-block-page |
| Deploycrawlergood | Deploycrawlergood | POST | /api/bot/domainConfig/deploy-crawler-good |
| Deploybotfeatureverification | Deploy Bot challenge rules | POST | /api/bot/domainConfig/deploy-bot-challenge |
| Deployfingerprintanalysis | Deploy fingerprint analysis rules | POST | /api/bot/domainConfig/deploy-fingerprint-challenge |
| Deploycaptchaverification | Deploy captcha verification rules | POST | /api/bot/domainConfig/deploy-captcha-challenge |
| Deploybehavioranalyse | Deploy behavior analyse rules | POST | /api/bot/domainConfig/deploy-behavior-analyse |
| Deployblockpage | Deploy block page | POST | /api/bot/domainConfig/deploy-block-page |
| Createconsumer | Createconsumer. | POST | /api/sam/consumers/add |
| Editconsumer | Edit consumer. | POST | /api/sam/consumers/edit |
| Deleteconsumer | Delete consumer. | POST | /api/sam/consumers/delete |
| Getconsumerlist | Get consumer list. | POST | /api/sam/consumers/list |
| Getconsumerinfo | Get consumer info. | GET | /api/sam/consumers |
| Getapiassetlist | Get API asset list | POST | /api/sam/api-discovery/list |
| Getapiassetdetail | Get API asset detail. | POST | /api/sam/api-discovery/detail |
| Reportincorrectapiasset | Report incorrect API asset discovery data. | POST | /api/sam/api-discovery/false-marking |
| Createandactiveapi | Create and deployment activation API. | POST | /api/sam/api/add-active |
| Createquotarule | Create quota rule. | POST | /api/sam/policy-quota/add |
| Editquotarule | Edit quota rule. | POST | /api/sam/policy-quota/edit |
| Deletequotarule | Delete quota rule. | POST | /api/sam/policy-quota/delete |
| Getapilist | Get API list. | POST | /api/sam/api/list |
| Getquotarulelist | Get quota rule list. | GET | /api/sam/policy-quota/list |
| Deleteapi | Delete API. | POST | /api/sam/api/delete |
| Editandactiveapi | Edit and active API. | POST | /api/sam/api/edit |
| Createconcurrencylimitrule | Create high concurrency limit rule. | POST | /api/sam/policy-flow/add |
| Editconcurrencylimitrule | Modify high concurrency limit rule. | POST | /api/sam/policy-flow/edit |
| Deleteconcurrencylimitrule | Delete high concurrency limit rule. | POST | /api/sam/policy-flow/delete |
| Getconcurrencylimitrulelist | Get high concurrency limit rule list. | GET | /api/sam/policy-flow/list |
| Queryaccesscontrolrulelist | Get  bot  rules | POST | /api/bot/customAccessControl/get-rule-listcmm |
| Wssmpnetworkserviceadd | Non-website business rules addition | POST | /soc/WssMPTcpService/add |
| Wssmpnetworkservicedelete | Non-website business rules deletion | POST | /soc/WssMPTcpService/delete |
| Wssmpnetworkserviceupdate | Non-website business rule modification | POST | /soc/WssMPTcpService/update |
| Wssmpnetworkservicequery | Query rules for non-website business | POST | /soc/WssMPTcpService/query |
| Deletebuiltinrule | Delete the built-in rule. | POST | /api/waf/policy/delete-builtInRule |
| Setdefaultbuiltinrule | Setting the default built-in rule.<br> | POST | /api/waf/policy/set-default-builtInRule |
| Getintelligentanalysislist | Query the list of intelligent analysis List. | POST | /api/waf/policy/get-intelligentAnalysis-list |
| Rejectintelligentanalysis | Reject the recommendation of intelligent analysis. | POST | /api/waf/policy/reject-intelligentAnalysis |
| Deleteintelligentanalysis | Delete the recommendation of intelligent analysis. | POST | /api/waf/policy/delete-intelligentAnalysis |
| Queryexactrulelist(accessControl) | Query rule list (Advanced Access Control) | POST | /api/bot/exact/query-exact-rule-list |
| Querydomainlistbyrulename(accessControl) | Query domainList by ruleName (Advanced Access Control) | POST | /api/bot/exact/query-domainlist-by-rulename |
| Addexactrule(accessControl) | Add rules (Advanced Access Control) | POST | /api/bot/exact/add-exact-rule |
| Editexactrule(accessControl) | Editing rules (Advanced Access Control)<br> | POST | /api/bot/exact/edit-exact-rule |
| Deleteexactrule(accessControl) | Delete Rule (Advanced Access Control) | POST | /api/bot/exact/delete-exact-rule |
| Batchassiociatedexactrule(accessControl) | Batch add domain names(Advanced Access Control)<br> | POST | /api/bot/exact/batch-assiociated-exact-rule |
| Batchremoveexactrule(accessControl) | Batch removal of domain names(Advanced Access Control) | POST | /api/bot/exact/batch-remove-exact-rule |
| Queryexactrulelistbydomain(accessControl) | Query rule list based on domain name(Advanced Access Control) | POST | /api/bot/exact/query-exact-rulelist-by-domain |
| Updatedomainbaseinfo | Modify domain name basic switch/mode information. | POST | /api/v1/dms/ddosProtect/updateDomainBaseInfo |
| Querydomainbuiltinrules | Query domain built-in rules | POST | /api/v1/dms/ddosProtect/queryDomainBuiltInRules |
| Updatedomainbuiltinrules | Update domain built-in rules | POST | /api/v1/dms/ddosProtect/updateDomainBuiltInRules |
| Sharedomainbaseprotectconfig | Cross domain deployment basic configuration | POST | /api/v1/dms/ddosProtect/shareDomainBaseProtectConfig |
| Sharedmsbuiltinrulesconfig | Cross domain deployment common buildin config | POST | /api/v1/dms/ddosProtect/shareDmsBuiltInRulesConfig |
| Sharedmsnativeappapiconfig | Cross domain deployment exception api | POST | /api/v1/dms/ddosProtect/shareDmsNativeAppApiConfig |
| Querydomainddosbaseinfo | Query domain basic DDoS protection strategy. | POST | /api/v1/dms/ddosProtect/queryDomainBaseInfo |
| Queryadaptiveprotectionrules | Querying adaptive protection rules for domains. | POST | /api/v1/dms/ddosProtect/queryDomainAIRules |
| Updateadaptiveprotectionrules | Update the adaptive protection rules of domain. | POST | /api/v1/dms/ddosProtect/updateDomainAIRules |
| Querydomainappapiexceptions | Query the shared configuration associated with domains - App/API exceptions | POST | /api/v1/dms/ddosProtect/queryDomainAppApiExceptionRules |
| Removedomainappapiexceptions | Remove APP/API exception rules for a domain | POST | /api/v1/dms/ddosProtect/removeDomainAppApiExceptionRules |
| Relatedappapiexceptionsfordomain | Domain related App/API exception rules | POST | /api/v1/dms/ddosProtect/relateDomainAppApiExceptionRules |
| Listnonsharedwafruleexceptionsforwafrules | Get a list of exceptions for the WAF rules. | POST | /api/v1/waf/rule-exception/list-normal |
| Domainsofnonlatestrulesetversion | Get domains with ruleset versions that are not the latest. | GET | /api/v1/waf/conf/base/get-upgrade-domain-list |
| Listupgradedetails | Return a list of the upgrade ruleset details for the specified domains. | POST | /api/v1/waf/rule/get-upgrade-rule-list |
| Upgradewafruleset | Upgrade the WAF ruleset for specified domains to the latest version. | POST | /api/v1/waf/rule/batch/upgrade |
| Listwafrules | Return a list of specific WAF rules that are filtered based on rule ID, rule name, rule description or vulnerability number. | POST | /api/v1/waf/rule/get-rule-list |
| Updateactionforwafmanagedrules | Modify the action for WAF managed rules. | POST | /api/v1/waf/rule/batch/update |
| Listwafbasicconfigofdomains | Return a list of protection mode, ruleset mode and the version of ruleset for specified domains. | POST | /api/v1/waf/conf/base/get-basic-conf-list |
| Updatemodeofwaf | Update the protection mode and ruleset mode of the domain. | POST | /api/v1/waf/conf/base/update-basic-conf |
| Createexceptiontowafmanagedrules | Create an exception rule for WAF managed rules. | POST | /api/v1/waf/rule-exception/add-normal |
| Associatesharewafruleexception | Associate the WAF exceptions in the shared configuration with the specified domains. | POST | /api/v1/waf/rule-exception/add-share |
| Deleteexceptionforwafmanagedrules | Remove exception for WAF managed rules. | POST | /api/v1/waf/rule-exception/delete-normal |
| Updateexceptionforwafmanagedrules | Modify exception for WAF managed rules. | POST | /api/v1/waf/rule-exception/update |
| Deletesharedwafexceptionsassociatedwithwafmanagedrules | Delete the exceptions(source: Shared configurations) of the WAF managed rules. | POST | /api/v1/waf/rule-exception/delete-share |
| Applyrecommendations | Apply the recommendations to the WAF exception. | POST | /api/v1/waf/ai-rule-result/apply |
| Rejectrecommendations | Reject the recommendations for WAF exception. | POST | /api/v1/waf/ai-rule-result/reject |
| Listrecommendations | Return a list of recommendations for WAF rules, including pending and rejected recommendations. | POST | /api/v1/waf/ai-rule-result/list |
| Listsharedwafruleexceptionsforwafrules | Return a list of exceptions(resource: shared configuration) for the specified WAF managed rules. | POST | /api/v1/waf/rule-exception/list-share |
| Listdomaininfos | Get Domain Information List. | POST | /api/v1/common/sys-domain-info/get-list |
| Usingexistinghostnametoaddnewhostname | Access new domain by referring to the configuration of protective domain. | POST | /api/v1/common/sys-domain-info/add-refer-to-other-domain |
| Copypoliciestootherhostnames | Copy security policy to other domains. <br>Note: <br>1. Copying to domains without protection enabled is not allowed; <br>2. Custom rules and Rate limiting that reference predefined API rules are not supported for copying, and such rules will be skipped during the policy copying process. | POST | /api/v1/common/sys-domain-info/copy-domain |
| Modifypolicystatus | Modify Policy Switch. | POST | /api/v1/common/sys-domain-info/update-switch |
| Removeprotecteddomain | Remove the protected domain. | GET | /api/v1/common/sys-domain-info/remove |
| Usingsystemrecommendedaccessdomain | Use the system-suggested access domain. | POST | /api/v1/common/sys-domain-info/add-sys-domain |
| Getbotfunctionswitch | Query the Bot management function switch. | POST | /api/v1/bot-manage/get-function-switch |
| Updatebotfunctionswitch | Modify the Bot management function switch. | POST | /api/v1/bot-manage/update-function-switch |
| Addratelimitingrule | Add a Rate Limiting rule. | POST | /api/v1/rate-limit/add-rule |
| Createwhitelistrule | Create a Whitelist rule. | POST | /api/v1/common/whitelist/add |
| Updateratelimitingrule | Update the configuration of a Rate Limiting rule. | POST | /api/v1/rate-limit/update-rule |
| Updatewhitelistrule | Update the configuration of a Whitelist rule. | POST | /api/v1/common/whitelist/update |
| Deletewhitelistrules | Delete Whitelist rules. | POST | /api/v1/common/whitelist/delete |
| Addworkflowrule | Add workflow detection rule. | POST | /api/v1/bot-manage/behavior/add-rule |
| Listratelimitingrules | Get a list of rules  for Rate Limiting. | POST | /api/v1/rate-limit/get-rule-list |
| Updateworkflowrule | Modify workflow detection rule. | POST | /api/v1/bot-manage/behavior/update-rule |
| Deleteratelimitingrules | Delete Rate Limiting rules. | POST | /api/v1/rate-limit/delete-rule-by-ids |
| Listworkflowrules | Query the list of workflow detection rule. | POST | /api/v1/bot-manage/behavior/get-rule |
| Getratelimitingrulesforthesharedconfigurationasociatedwithdomain | Return a list of Rate Limiting rules for the shared configuration associated with the domain. | POST | /api/v1/rate-limit/get-relation-by-domain |
| Deleteworkflowrule | Delete workflow detection rule. | POST | /api/v1/bot-manage/behavior/delete-rule |
| Getthreatintelligencedomainconfig | Get Threat Intelligence configurations of domain. | POST | /api/v1/common/intelligence/query-list |
| Updatethreatintelligencedomainconfig | Update Threat Intelligence Configuration of Domain. | POST | /api/v1/common/intelligence/update |
| Getdomainapisecurityconfiguration | Get Domain API security configuration. | POST | /api/v1/sam/api-defend/get-list-by-domains |
| Updatedomainapisecurityconfiguration | Update Domain API security configuration. | POST | /api/v1/sam/api-defend/update-defend-action |
| Addcustomizerule | Add custom rules. | POST | /api/v1/customize-rule/add |
| Listwhitelistrules | Get a list of rules for Whitelist. | POST | /api/v1/common/whitelist/query-list |
| Getwhiterulesforthesharedconfigurationasociatedwithdomain | Get the whitelist rules of the shared configuration associated with the domain. | GET | /api/v1/common/whitelist/get-imported-share-config |
| Updateipblocksettings | Update the settings of IP Block. Block requests from specific IP/CIDR. Tips: It will overwrite the original configuration. | POST | /api/v1/policy-block/save-deploy-ip-block |
| Updatecustomrule | Update the configuration of a custom rule. | POST | /api/v1/customize-rule/update |
| Deletecustomizerule | Delete custom rules. | POST | /api/v1/customize-rule/delete |
| Listcustomrules | Get a list of Custom Rules. | POST | /api/v1/customize-rule/get-list |
| Listipblocksettings | Get IP block configuration for a domain. | POST | /api/v1/policy-block/get-ip-block-list |
| Listgeoblocksettings | Return the countries and regions that are blocked for the specified domains. | POST | /api/v1/policy-block/get-geo-block-list |
| Listcustomrulesforsharedconfigurationassociatedwithdomain | Get the custom rules of the shared configuration associated with the domain. | POST | /api/v1/customize-rule/get-import-share-rule |
| Updategeoblocksettings | Block requests from specific countries and areas. Tips: It will overwrite the original configuration. | POST | /api/v1/policy-block/save-deploy-geo-block |
| Listdomaincustomizebots | Query the custom Bots list of the hostname. | POST | /api/v1/bot-manage/get-domain-customize-bots-list |
| Listspecificclienttrafficbypass | Query the list of Bypass Traffic from Specific Clients. | POST | /api/v1/bot-manage/scene/whitelist/list |
| Addspecificclienttrafficbypass | Add the rules of Bypass Traffic from Specific Clients. | POST | /api/v1/bot-manage/scene/whitelist/add |
| Updatespecificclienttrafficbypass | Update the rules of Bypass Traffic from Specific Clients. | POST | /api/v1/bot-manage/scene/whitelist/update |
| Deletespecificclienttrafficbypass | Delete the rules of Bypass Traffic from Specific Clients. | POST | /api/v1/bot-manage/scene/whitelist/delete |
| Addhtmlpageswithoutembeddingjs | For web Bot detection, add html pages without embedding JS. | POST | /api/v1/bot-manage/scene/web/js-exception/add |
| Updatehtmlpageswithoutembeddingjs | For web Bot detection, update html pages without embedding JS. | POST | /api/v1/bot-manage/scene/web/js-exception/update |
| Deletehtmlpageswithoutembeddingjs | For web Bot detection, delete html pages without embedding JS. | POST | /api/v1/bot-manage/scene/web/js-exception/delete |
| Disableallpolicies | Disable all security policies for the specified domains. The disabled domains will no longer be protected. | POST | /api/v1/common/sys-domain-info/close-all-policy-switch |
| Updateknownbotsselectbotnames | Batch modification of known Bot subcategories that take effect. | POST | /api/v1/bot-manage/update-known-bots-select-bot-names |
| Listknownbots | Query the list of known Bot list. | POST | /api/v1/bot-manage/get-known-bots |
| Updateknownbotsact | Batch modification known Bot actions. | POST | /api/v1/bot-manage/update-known-bots-act |
| Listuabots | Query the list of User-Agent based detection. | POST | /api/v1/bot-manage/get-ua-bots |
| Updateuabotsselectbotnames | Batch modification of User-Agent based detection subcategories that take effect. | POST | /api/v1/bot-manage/update-ua-bots-select-bot-names |
| Updateuabotsact | Batch modification User-Agent based detection actions. | POST | /api/v1/bot-manage/update-ua-bots-act |
| Getresponsepageofdenyactiondetail | Get General Settings-Response Page of Deny Action details | POST | /api/v1/other/sys-domain-basic/response/detail |
| Updateresponsepageofdenyactiondetail | Update General Settings-Response Page of Deny Action details | POST | /api/v1/other/sys-domain-basic/response/update |
| Listhistoricalhostnames | Query the list of historically protected domains. | POST | /api/v1/other/domain/get-permanent-domain-list |
| Getwafscanprotectionconfig | This API is used to retrieve the WAF scan protection configuration for specified domains. Users can use this API to obtain the configurations for scanning tool detection, repeated violation detection, and directory probing detection for domains. | POST | /api/v1/security-policy/waf/get-scan-protection-configuration |
| Updatewafscanprotectionconfig | This API is used to batch modify the WAF scan protection configuration for specified domains. Users can use this API to batch modify the configurations for scanning tool detection, repeated violation detection, and directory probing detection for domains. | POST | /api/v1/security-policy/waf/update-scan-protection-configuration |
| Getpredeployresult | This API is used to query the results of a pre-deployment. Users can obtain the pre-deployment results, including deployment results and Host information, by providing the pre-deployment id. | GET | /api/v1/security-policy/get-pre-deploy-result |
| Updatedetectionlengthlimitconfiguration | This API is used to batch modify the detection length limit configuration for specified domains. Users can use this API to batch modify the maximum detection length configuration of the request body for domains. | POST | /api/v1/security-policy/basic/update-detection-length-limit-config |
| Getdetectionlengthlimitconfiguration | This API is used to query the detection length limit configuration for a specified domain. Users can provide a domain to obtain the detection length verification configuration for that domain, including the maximum detection length of the request body. | POST | /api/v1/security-policy/basic/get-detection-length-limit-config |
| Getbotgeneralconfig | This API is used to query the general configuration of Bot strategies, including detailed configuration information for Public Bots, AI Bots, Definite Bots, and Bot tagging config. | POST | /api/v1/security-policy/bot/get-general-config |
| Updatebotgeneralconfig | This API is used to modify the general configuration of Bot strategies, including detailed configuration information for Public Bots, AI Bots, Definite Bots, and Bot tagging. | POST | /api/v1/security-policy/bot/update-general-config |