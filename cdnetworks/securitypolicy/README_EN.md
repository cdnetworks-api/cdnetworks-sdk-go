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
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
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
| Updatecustomrule | Update custom rule. | POST | /api/waf/custom/rule/update |
| Deletecustomrule | Delete custom rule. | POST | /api/waf/custom/rule/delete |
| Associatedomainsforcustomrule | Associate the domains with custom rule. | POST | /api/waf/custom/rule/add-rela |
| Removedomainsforcustomrule | Remove the domains from the specified custom rule. | POST | /api/waf/custom/rule/remove-rela |
| Getruleexceptionlist | Query the rule list of rule exception. | POST | /api/waf/template/rule/white/list-page |
| Createruleexception | Create the rule exception. | POST | /api/waf/template/rule/white/add |
| Updateruleexception | Update the rule exception. | POST | /api/waf/template/rule/white/update |
| Deleteruleexception | Delete the rule exception. | POST | /api/waf/template/rule/white/delete |
| Associatedomainsforruleexception | Associate the domains with rule exception. | POST | /api/waf/template/rule/white/add-rela |
| Removedomainsforruleexception | Remove the domains from the specified rule exception. | POST | /api/waf/template/rule/white/remove-rela |
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
| Querydomainbuiltinrules | Query domain name built-in rules | POST | /api/v1/dms/ddosProtect/queryDomainBuiltInRules |
| Listnonsharedwafruleexceptionsforwafrules | Return a list of exceptions for the WAF rules, including the exception configuration for "Source: Manual" and "Source: Recommendations". | POST | /api/v1/waf/rule-exception/list-normal |
| Listhosnamesofnonlatestrulesetversion | Return the hostnames which ruleset is not latest version. | GET | /api/v1/waf/conf/base/get-upgrade-domain-list |
| Listupgradedetails | Return a list of the upgrade ruleset details for the specified hostnames. | POST | /api/v1/waf/rule/get-upgrade-rule-list |
| Upgradewafruleset | Upgrade the WAF ruleset for specified hostnames to the latest version. | POST | /api/v1/waf/rule/batch/upgrade |
| Listwafrules | Return a list of specific WAF rules that are filtered based on rule ID, rule name, rule description or vulnerability number. | POST | /api/v1/waf/rule/get-rule-list |
| Updateactionforwafmanagedrules | Modify the action for WAF managed rules. | POST | /api/v1/waf/rule/batch/update |
| Listwafbasicconfigofhostnames | Return a list of protection mode, ruleset mode and the version of ruleset for specified hostnames. | POST | /api/v1/waf/conf/base/get-basic-conf-list |
| Updatemodeofwaf | Modify the Protection Mode and Ruleset Mode for specified hostnames. | POST | /api/v1/waf/conf/base/update-basic-conf |
| Createexceptiontowafmanagedrules | Create an exception rule for WAF managed rules. | POST | /api/v1/waf/rule-exception/add-normal |
| Associatesharewafruleexception | Associate the WAF exceptions in the shared configuration with the specified hostnames. | POST | /api/v1/waf/rule-exception/add-share |
| Deleteexceptionforwafmanagedrules | Remove exception for WAF managed rules. | POST | /api/v1/waf/rule-exception/delete-normal |
| Updateexceptionforwafmanagedrules | Modify exception for WAF managed rules. | POST | /api/v1/waf/rule-exception/update |
| Deletesharedwafexceptionsassociatedwithwafmanagedrules | Delete the exceptions(source: Shared configurations) of the WAF managed rules. | POST | /api/v1/waf/rule-exception/delete-share |
| Applyrecommendations | Apply the recommendations to the WAF exception. | POST | /api/v1/waf/ai-rule-result/apply |
| Rejectrecommendations | Reject the recommendations for WAF exception. | POST | /api/v1/waf/ai-rule-result/reject |
| Listrecommendations | Return a list of recommendations for WAF rules, including pending and rejected recommendations. | POST | /api/v1/waf/ai-rule-result/list |
| Listsharedwafruleexceptionsforwafrules | Return a list of exceptions(resource: shared configuration) for the specified WAF managed rules. | POST | /api/v1/waf/rule-exception/list-share |
| Listdomaininfos | Query list of hostname information. | POST | /api/v1/common/sys-domain-info/get-list |
| Usingexistinghostnametoaddnewhostname | Add a New Hostname by Using Existing Hostname's configuration.  | POST | /api/v1/common/sys-domain-info/add-refer-to-other-domain |
| Copypoliciestootherhostnames | Copy specified hostname's policies to other hostnames.<br>Note: <br>1. Copying to policies without protection enabled is not allowed;<br>2. Rules that reference defined APIs in custom rules and frequency limits are not supported for copying, and such rules will be skipped during the policy copying process. | POST | /api/v1/common/sys-domain-info/copy-domain |
| Modifypolicystatus | Modify policy status. | POST | /api/v1/common/sys-domain-info/update-switch |
| Removeprotectedhostname | Remove the protected hostname. | GET | /api/v1/common/sys-domain-info/remove |
| Usingsystemrecommendedaccessdomain | Using the system recommended policy access hostname. | POST | /api/v1/common/sys-domain-info/add-sys-domain |
| Getbotfunctionswitch | Query the Bot management function switch. | POST | /api/v1/bot-manage/get-function-switch |
| Updatebotfunctionswitch | Modify the Bot management function switch. | POST | /api/v1/bot-manage/update-function-switch |
| Creatratelimitingrule | Creat a Rate Limiting rule. | POST | /api/v1/rate-limit/add-rule |
| Createwhitelistrule | Create a Whitelist rule. | POST | /api/v1/common/whitelist/add |
| Updateratelimitingrule | Update the configuration of a Rate Limiting rule. | POST | /api/v1/rate-limit/update-rule |
| Updatewhitelistrule | Update the configuration of a Whitelist rule. | POST | /api/v1/common/whitelist/update |
| Deletewhitelistrules | Delete Whitelist rules. | POST | /api/v1/common/whitelist/delete |
| Addworkflowrule | Add workflow detection rule. | POST | /api/v1/bot-manage/behavior/add-rule |
| Listratelimitingrules | Return a list of rules  for Rate Limiting. | POST | /api/v1/rate-limit/get-rule-list |
| Updateworkflowrule | Modify workflow detection rule. | POST | /api/v1/bot-manage/behavior/update-rule |
| Deleteratelimitingrules | Delete Rate Limiting rules. | POST | /api/v1/rate-limit/delete-rule-by-ids |
| Listworkflowrules | Query the list of workflow detection rule. | POST | /api/v1/bot-manage/behavior/get-rule |
| Getratelimitingrulesforthesharedconfigurationasociatedwithhostname | Return a list of Rate Limiting rules for the shared configuration associated with the hostname. | POST | /api/v1/rate-limit/get-relation-by-domain |
| Deleteworkflowrule | Delete workflow detection rule. | POST | /api/v1/bot-manage/behavior/delete-rule |
| Getthreatintelligencedomainconfig | Get threat intelligence hostname config. | POST | /api/v1/common/intelligence/query-list |
| Updatethreatintelligencedomainconfig | Update threat intelligence hostname config. | POST | /api/v1/common/intelligence/update |
| Getdomainapisecurityconfiguration | Get hostname API security configuration. | POST | /api/v1/sam/api-defend/get-list-by-domains |
| Updatedomainapisecurityconfiguration | Update hostname API security configuration. | POST | /api/v1/sam/api-defend/update-defend-action |
| Addcustomizerule | Add custom rules. | POST | /api/v1/customize-rule/add |
| Listwhitelistrules | Return a list of rules for Whitelist. | POST | /api/v1/common/whitelist/query-list |
| Getwhitelistrulesforthesharedconfigurationasociatedwithhostname | Return a list of Whitelist rules for the shared configuration associated with the hostname. | GET | /api/v1/common/whitelist/get-imported-share-config |
| Updateipblocksettings | Update the settings of IP Block. Block requests from specific IP/CIDR. Tips: It will overwrite the original configuration. | POST | /api/v1/policy-block/save-deploy-ip-block |
| Updatecustomrule | Update the configuration of a custom rule. | POST | /api/v1/customize-rule/update |
| Deletecustomizerule | Delete custom rules. | POST | /api/v1/customize-rule/delete |
| Listcustomrules | Return a list of rules for Custom Rules. | POST | /api/v1/customize-rule/get-list |
| Listipblocksettings | Return a list of the blocked IP for the specified hostnames. | POST | /api/v1/policy-block/get-ip-block-list |
| Listgeoblocksettings | Return a list of the blocked countries and areas for the specified hostnames. | POST | /api/v1/policy-block/get-geo-block-list |
| Listcustomrulesforsharedconfigurationassociatedwithhostname | Return a list of custom rules for the shared configuration associated with the specified hostnames. | POST | /api/v1/customize-rule/get-import-share-rule |
| Updategeoblocksettings | Block requests from specific countries and areas. Tips: It will overwrite the original configuration. | POST | /api/v1/policy-block/save-deploy-geo-block |
| Listdomaincustomizebots | Query the custom Bots list of the hostname. | POST | /api/v1/bot-manage/get-domain-customize-bots-list |
| Updateconfigurationofattackerippunishment | Modify the Configuration that specifies the Attacker IP Punishment rules. | POST | /api/v1/waf/waf/ip-punishment/update |
| Listconfigurationofattackerippunishment | Return the configuration of the "Attacker IP Punishment" for the specified hostnames. | POST | /api/v1/waf/waf/ip-punishment/list |
| Disableallpolicies | Disable all security policies for the specified hostnames. The disabled hostname will not be checked. | POST | /api/v1/common/sys-domain-info/close-all-policy-switch |
| Updateknownbotsselectbotnames | Batch modification of known Bot subcategories that take effect. | POST | /api/v1/bot-manage/update-known-bots-select-bot-names |
| Listknownbots | Query the list of known Bot list. | POST | /api/v1/bot-manage/get-known-bots |
| Updateknownbotsact | Batch modification known Bot actions. | POST | /api/v1/bot-manage/update-known-bots-act |
| Listuabots | Query the list of User-Agent based detection. | POST | /api/v1/bot-manage/get-ua-bots |
| Updateuabotsselectbotnames | Batch modification of User-Agent based detection subcategories that take effect. | POST | /api/v1/bot-manage/update-ua-bots-select-bot-names |
| Updateuabotsact | Batch modification User-Agent based detection actions. | POST | /api/v1/bot-manage/update-ua-bots-act |
| Getresponsepageofdenyactiondetail | Get General Settings-Response Page of Deny Action details | POST | /api/v1/other/sys-domain-basic/response/detail |
| Updateresponsepageofdenyactiondetail | Update General Settings-Response Page of Deny Action details | POST | /api/v1/other/sys-domain-basic/response/update |
| Listhistoricalhostnames | Query the domain name for historical protection. | POST | /api/v1/other/domain/get-permanent-domain-list |