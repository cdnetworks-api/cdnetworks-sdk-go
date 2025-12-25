# Cdnetworks SDK for Go

文档提供了使用 Cdnetworks SDK for Go 的文档。

## SDK 安装（推荐）

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## 示例用法

该 SDK 使用 AKSK（访问密钥/秘密密钥）认证。按如下方式配置您的凭据：

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/securityreport"
    "log"
)

func main() {
    // 参考本文档最后的API列表，修改一下对应的{ActionName}、Method、Uri
    request := &securityreport.{ActionName}Request{}

    // Configure authentication
    var config auth.AkskConfig
    config.AccessKey = "YOUR_ACCESS_KEY"
    config.SecretKey = "YOUR_SECRET_KEY"
    config.Uri = "/your/api/path"
    config.Method = "HTTP_METHOD"  // GET, POST, PUT, DELETE, etc.

    // Create response object and make API call
    response := securityreport.{ActionName}Response{}
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
| Queryccattackqps | 查询CC攻击QPS | POST | /soc/api/report/Queryccattackqps |
| Getwafrequestandattackevent | 查询WAF请求数及攻击事件汇总。 | POST | /api/waf/report/query-total-request |
| Getwafattacktrend | 查询WAF的攻击请求数趋势。 | POST | /api/waf/report/query-attack-hit-trend |
| Getwafattackevent | 查询域名的重点攻击事件。 | POST | /api/waf/report/query-attack-event-list |
| Getwafattackeddomain | 查询域名的攻击次数。 | POST | /api/waf/report/query-attacked-domain-list |
| Getwafattackedurl | 查询域名的受攻击URL攻击情况。 | POST | /api/waf/report/query-attacked-url-list |
| Getdistributionofwafattacktype | 查询WAF攻击类型分布。 | POST | /api/waf/report/query-top-by-attack-type |
| Getwaftriggerrulelist | 查询WAF触发规则/策略。 | POST | /api/waf/report/query-trrigger-rule-list |
| Getwafpolicydetails | 查询WAF防护策略详情。 | POST | /api/waf/report/query-last-attack-info |
| Getwafattackip | 查询WAF攻击来源Top IP。 | POST | /api/waf/report/query-attack-ip-list |
| Getdistributionofwafattacksource | 查询WAF攻击来源地区分布。 | POST | /api/waf/report/query-attack-country-list |
| Getwafdomainhistory | 查询WAF历史域名。 | POST | /api/waf/report/query-waf-history-domain-list |
| Queryccattackdetails | CC 攻击详情查询接口<br>接口请求和返回结果只接受json格式 | POST | /soc/wss_report/shieldOverview/ccAttackDetails |
| Queryddosattackdetails | DDoS攻击详情接口<br>接口请求和返回结果只接受json数据 | POST | /soc/wss_report/shieldOverview/ddosAttackDetails |
| Getdomainbotvisitdetails | 统计域名的总请求次数、Bot请求数、缓解Bot攻击数及Top5 Bot类型的请求次数 | POST | /api/bot/report/bot-visit-domain |
| Getbotrequestoverviewdata | 统计指定域名的请求攻击数据，包括：总请求数、已知Bot类型攻击数、未知Bot类型攻击数、缓解bot攻击数 | POST | /api/bot/report/request-visit |
| Getbotrequeststatisticperdomain | 统计Bot访问域名请求数量TOP数据。 | POST | /api/bot/report/bot-visit-domain-top |
| Getbotrequesttypedistributedata | 统计指定域名的Bot请求类型分布数据 | POST | /api/bot/report/attack-type-distribute |
| Getbotaccessurltopdata | 统计指定域名的请求访问URL TOP数据 | POST | /api/bot/report/top-url |
| Getbotrequestsourceiptopdata | 统计指定域名的请求访问IP TOP数据 | POST | /api/bot/report/top-ip |
| Getbotrequestreferertopdata | 获取请求访问Referer TOP数据 | POST | /api/bot/report/top-referer |
| GetbotrequestuseragentTopdata | 获取请求访问Ua TOP数据 | POST | /api/bot/report/top-ua |
| Getbotrequestsourcedistributiondata | 获取请求来源区域分布数据。 | POST | /api/bot/report/top-area |
| Getbotrequesttrendsandtriggerrulesdata | 获取域名的Bot请求趋势 | POST | /api/bot/report/trend-rule |
| Getacttypedistributiondata | 获取处置状态分布数据 | POST | /api/bot/report/act-type-distribute |
| Getbotruletypetopdata | 获取域名的Bot触发规则 | POST | /api/bot/report/top-rule-type |
| Getapieventlogs | 查询API事件日志。 | POST | /api/sam/attack-log/query-attack-log |
| Getapinumber | 获取已定义的API数。 | POST | /api/sam/view-api-count |
| Getriskeventnumber | 获取风险事件数。 | POST | /api/sam/view-attack-count |
| Getriskeventtop5 | 获取风险事件top5 | POST | /api/sam/view-attack-type-top5 |
| Getactiveapitop5 | 获取活跃APITop5。 | POST | /api/sam/view-call-api-top5 |
| Getprivacystatusdistribution | 获取隐私状态分布。 | POST | /api/sam/view-conceal-count |
| Getconsumernumber | 获取消费方个数。 | POST | /api/sam/view-consumer-count |
| Getrequesttrend | 获取请求趋势 | POST | /api/sam/view-request-trend |
| Gettotalrequestnumber | 获取总调用次数 | POST | /api/sam/view-total-call-count |
| Getrulestatusstatistics | 查询域名的规则状态统计数据。<br> | POST | /api/waf/report/get-ruleStatus-statistics |
| Queryeventtrend | 查询流量趋势（按请求数）。 | POST | /api/v1/overview/trend-info |
| Listattacklogs | 按策略类型或指定域名查询攻击日志。只返回前10000条日志，要获取全部日志请调整过滤条件或者联系技术支持使用非API方式。 | POST | /api/v1/attack-log/attack-log-list |
| Gettrendsbyrps | 获取流量趋势（按rps）。 | POST | /api/v1/overview/trend-info-qps |
| Queryattacklogdetailinfo | 查询攻击日志详情信息。 | POST | /api/v1/attack-log/all-detail-info |
| Getoriginalrequestinformation | 查询攻击日志原始请求信息。 | POST | /api/v1/attack-log/original-request-info |
| Gettriggeredwafmanagedrules | 获取已触发的WAF内置规则及其触发次数。 | POST | /api/v1/situation/waf-rule-hit |
| Gettriggeredddosmanagedrules | 获取已触发的DDoS内置规则及其触发次数。 | POST | /api/v1/situation/ddos-rule-hit |
| Gettriggeredratelimitingrules | 获取已触发的频率限制规则及其触发次数。 | POST | /api/v1/situation/rate-limit-rule-hit |
| Gettriggeredcustomrules | 获取已触发的自定义规则及其触发次数。 | POST | /api/v1/situation/customize-rule-hit |
| Gettopattacktargetsbydomain | 获取遭受攻击最多的域名。 | POST | /api/v1/situation/target-domain |
| Gettopattacktargetsbypath | 获取遭受攻击最多的路径。 | POST | /api/v1/situation/target-url |
| Gettopattacksourcesforclientips | 获取攻击来源Top数据。 | POST | /api/v1/situation/source-ip-top |
| Gettopattacksourcesforglobal | 获取全球的攻击来源Top数据，可查看攻击请求来自于哪些国家/地区。 | POST | /api/v1/situation/source-country-top |
| Gettopattacksourcesforchinamainland | 获取中国大陆的攻击来源Top数据，可查看攻击请求来自于哪些城市。 | POST | /api/v1/situation/source-province-top |
| Getsummaryrequests | 获取汇总请求数，包括总请求数、已缓解的请求数、观察的请求数和被白名单放行的请求数。 | POST | /api/v1/overview/event-summary |
| Getinfrastructureloglist | 获取基础设施保护日志。 | POST | /api/v1/csec-attack-log/get-infrastructure-log-list |
| L4ddosevent | 查询DDoS攻击事件。 | POST | /api/v1/dms-overview/ddos-event |
| L4ddostrend | 网络层DDoS攻击趋势。 | POST | /api/v1/dms-overview/ddos-trend |
| Gettoppoliciestriggered | 获取各策略类型的触发次数。 | POST | /api/v1/overview/event-distribution |
| Listserviceusage | 该接口用于查询服务用量。 用户可以通过时间来获取服务用量，包括主域名、子域名列表、开始防护时间、结束防护时间。 当前fs2.0计费逻辑是按照主域名数量来进行计费，可以帮助用户控制成本和对账审计。 | POST | /api/v1/service-usage/service-usage-list |