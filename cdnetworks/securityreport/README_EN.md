# Cdnetworks SDK for Go

This README provides documentation for using the Cdnetworks SDK for Go.

## SDK Installation(Recommended)

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go
```

## Product Single Installation

```bash
go get github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/securityreport
```

## Example Usage

The SDK uses AKSK (Access Key/Secret Key) authentication. Configure your credentials as follows:

```go
package main

import (
    "github.com/cdnetworks-api/cdnetworks-sdk-go/common/auth"
    "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/securityreport"
    "log"
)

func main() {
	// Refer to the API list at the end of this document and modify the corresponding {ActionName}, Method, and Uri
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
| Queryccattackqps | Query CC Attack QPS | POST | /soc/api/report/Queryccattackqps |
| Getwafrequestandattackevent | Query the requests statistics and key incidents statistics. | POST | /api/waf/report/query-total-request |
| Getwafattacktrend | Query the trend of waf attack. | POST | /api/waf/report/query-attack-hit-trend |
| Getwafattackevent | Query the key incidents of domain. | POST | /api/waf/report/query-attack-event-list |
| Getwafattackeddomain | Query the attacked domain. | POST | /api/waf/report/query-attacked-domain-list |
| Getwafattackedurl | Query the attacked url. | POST | /api/waf/report/query-attacked-url-list |
| Getdistributionofwafattacktype | Query the distribution of waf attack type. | POST | /api/waf/report/query-top-by-attack-type |
| Getwaftriggerrulelist | Query the triggered rules and policies. | POST | /api/waf/report/query-trrigger-rule-list |
| Getwafpolicydetails | Query the configuration of waf policy. | POST | /api/waf/report/query-last-attack-info |
| Getwafattackip | Query the top attack ip of waf. | POST | /api/waf/report/query-attack-ip-list |
| Getdistributionofwafattacksource | Query the distribution of waf attack source. | POST | /api/waf/report/query-attack-country-list |
| Getwafdomainhistory | Query waf historical domains. | POST | /api/waf/report/query-waf-history-domain-list |
| Queryccattackdetails | CC Attack Details | POST | /soc/wss_report/shieldOverview/ccAttackDetails |
| Queryddosattackdetails | DDoS Attack Details | POST | /soc/wss_report/shieldOverview/ddosAttackDetails |
| Getdomainbotvisitdetails | Get the total number of domain requests, bot requests, bot relief attack requests, and Top5 bot type requests | POST | /api/bot/report/bot-visit-domain |
| Getbotrequestoverviewdata | Get the overview data, including: total requests, known bot type attacks, unknown bot type attacks, and bot relief attacks | POST | /api/bot/report/request-visit |
| Getbotrequeststatisticperdomain | Get the top number of  bot  requests  by domain | POST | /api/bot/report/bot-visit-domain-top |
| Getbotrequesttypedistributedata | Get the bot request type distribution data of the specified domain | POST | /api/bot/report/attack-type-distribute |
| Getbotaccessurltopdata | Get the TOP data of the request URL of the specified domain | POST | /api/bot/report/top-url |
| Getbotrequestsourceiptopdata | get the bot request IP TOP data of the specified domain | POST | /api/bot/report/top-ip |
| Getbotrequestreferertopdata | Get the request  Referer TOP data | POST | /api/bot/report/top-referer |
| GetbotrequestuseragentTopdata | Get the  request top data for user agent | POST | /api/bot/report/top-ua |
| Getbotrequestsourcedistributiondata | Getbotrequestsourcedistributiondata | POST | /api/bot/report/top-area |
| Getbotrequesttrendsandtriggerrulesdata | Get Bot Request Trends And Rules Top Data | POST | /api/bot/report/trend-rule |
| Getacttypedistributiondata | Get the bot protection action type distribution data | POST | /api/bot/report/act-type-distribute |
| Getbotruletypetopdata | Get the bot rule data | POST | /api/bot/report/top-rule-type |
| Getapieventlogs | Get API event logs. | POST | /api/sam/attack-log/query-attack-log |
| Getapinumber | The number of defined api. | POST | /api/sam/view-api-count |
| Getriskeventnumber | The number of risk events. | POST | /api/sam/view-attack-count |
| Getriskeventtop5 | Get risk event Top5 | POST | /api/sam/view-attack-type-top5 |
| Getactiveapitop5 | Get active APITop5. | POST | /api/sam/view-call-api-top5 |
| Getprivacystatusdistribution | Get privacy status distribution. | POST | /api/sam/view-conceal-count |
| Getconsumernumber | The number of consumers. | POST | /api/sam/view-consumer-count |
| Getrequesttrend | The trend of requests | POST | /api/sam/view-request-trend |
| Gettotalrequestnumber | The total number of requests | POST | /api/sam/view-total-call-count |
| Getrulestatusstatistics | Query the rule status statistics.<br> | POST | /api/waf/report/get-ruleStatus-statistics |
| Queryeventtrend | Query traffic trend (by request count). | POST | /api/v1/overview/trend-info |
| Listattacklogs | Query attack logs based on policy type or specified domain.Only the first 10,000 log entries are returned. To retrieve all logs, please adjust the filter criteria or contact technical support to use a non-API method. | POST | /api/v1/attack-log/attack-log-list |
| Gettrendsbyrps | Get traffic trends (by rps). | POST | /api/v1/overview/trend-info-qps |
| Queryattacklogdetailinfo | Query attack log details. | POST | /api/v1/attack-log/all-detail-info |
| Getoriginalrequestinformation | Get the original request information of the attack log. | POST | /api/v1/attack-log/original-request-info |
| Gettriggeredwafmanagedrules | Return a list of WAF managed rules that have been triggered and how many times they have been triggered. | POST | /api/v1/situation/waf-rule-hit |
| Gettriggeredddosmanagedrules | Return a list of DDoS managed rules that have been triggered and how many times they have been triggered. | POST | /api/v1/situation/ddos-rule-hit |
| Gettriggeredratelimitingrules | Return a list of Rate limiting rules that have been triggered and how many times they have been triggered. | POST | /api/v1/situation/rate-limit-rule-hit |
| Gettriggeredcustomrules | Return a list of custom rules that have been triggered and how many times they have been triggered. | POST | /api/v1/situation/customize-rule-hit |
| Gettopattacktargetsbydomain | Get the most attacked domains. | POST | /api/v1/situation/target-domain |
| Gettopattacktargetsbypath | Get the most attacked path. | POST | /api/v1/situation/target-url |
| Gettopattacksourcesforclientips | Get Top Attack Sources (Client IP). | POST | /api/v1/situation/source-ip-top |
| Gettopattacksourcesforglobal | Obtain the Top data of global attack sources, you can see which countries/regions the attack requests come from. | POST | /api/v1/situation/source-country-top |
| Gettopattacksourcesforchinamainland | Obtain the Top data of attack sources in Mainland China, you can see which cities the attack requests come from. | POST | /api/v1/situation/source-province-top |
| Getsummaryrequests | Get the number of requests, including total requests, mitigated requests, monitored requests, and the whitelisted requests. | POST | /api/v1/overview/event-summary |
| Getinfrastructureloglist | Get a list of infrastructure protection logs. | POST | /api/v1/csec-attack-log/get-infrastructure-log-list |
| L4ddosevent | Obtain L3/4 DDoS events. | POST | /api/v1/dms-overview/ddos-event |
| L4ddostrend | Obtain data related to the L3/4 DDoS attack trend. | POST | /api/v1/dms-overview/ddos-trend |
| Gettoppoliciestriggered | Gets the number of triggers for each policy type. | POST | /api/v1/overview/event-distribution |
| Listserviceusage | This interface is used to query service usage. Users can obtain service usage data by specifying a time period, including the main domain, list of subdomains, defense start time, and defense end time. The current fs2.0 billing logic charges based on the number of main domains, helping users control costs and perform billing audits. | POST | /api/v1/service-usage/service-usage-list |