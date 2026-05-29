package ngreports

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetASummaryOfTrafficBandwidthPaths struct {
}

func (s GetASummaryOfTrafficBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthPaths) GoString() string {
  return s.String()
}

type GetASummaryOfTrafficBandwidthParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z ", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z. Due to latency associated with new traffic data, enddate should be no later than five minutes before the current time. This ensures you get the most accurate results.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。由于数据处理存在延迟，所指定的结束时间必须至少比当前时间早5分钟，否则返回的数据可能不准确。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: http https all 
  // Default: all 
  // Limit the results to the specified scheme. By default, all traffic is returned.", "zh_CN": "[ 0 .. 5 ] 字符 
  // 取值范围: http, https, all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetASummaryOfTrafficBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthParameters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficBandwidthParameters) SetStartdate(v string) *GetASummaryOfTrafficBandwidthParameters {
  s.Startdate = &v
  return s
}

func (s *GetASummaryOfTrafficBandwidthParameters) SetEnddate(v string) *GetASummaryOfTrafficBandwidthParameters {
  s.Enddate = &v
  return s
}

func (s *GetASummaryOfTrafficBandwidthParameters) SetScheme(v string) *GetASummaryOfTrafficBandwidthParameters {
  s.Scheme = &v
  return s
}

type GetASummaryOfTrafficBandwidthRequestHeader struct {
}

func (s GetASummaryOfTrafficBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthRequestHeader) GoString() string {
  return s.String()
}

type GetASummaryOfTrafficBandwidthRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetASummaryOfTrafficBandwidthRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
  // {"en" : "<= 2 items 
  // You can group results using a combination of up to two of the following: 'hostnames', 'serverGroups', and 'customerIds'.", "zh_CN": "<= 2 条目 
  // 指定分组依据对数据进行分组汇总。支持按'hostnames'，'serverGroups'单独进行分组汇总，也支持同时指定这2个参数进行分组汇总。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetASummaryOfTrafficBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthRequest) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficBandwidthRequest) SetFilters(v *GetASummaryOfTrafficBandwidthRequestFilters) *GetASummaryOfTrafficBandwidthRequest {
  s.Filters = v
  return s
}

func (s *GetASummaryOfTrafficBandwidthRequest) SetGroupBy(v []*string) *GetASummaryOfTrafficBandwidthRequest {
  s.GroupBy = v
  return s
}

type GetASummaryOfTrafficBandwidthRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Indicates one or more server groups.", "zh_CN": "指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetASummaryOfTrafficBandwidthRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthRequestFilters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficBandwidthRequestFilters) SetHostnames(v []*string) *GetASummaryOfTrafficBandwidthRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetASummaryOfTrafficBandwidthRequestFilters) SetServerGroups(v []*string) *GetASummaryOfTrafficBandwidthRequestFilters {
  s.ServerGroups = v
  return s
}

type GetASummaryOfTrafficBandwidthResponseHeader struct {
}

func (s GetASummaryOfTrafficBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthResponseHeader) GoString() string {
  return s.String()
}

type GetASummaryOfTrafficBandwidthResponse struct {
  // {"en" : "Metadata that describes the data in the response.", "zh_CN": "对响应体中的报表数据的相关说明。"}
  MetaData *GetASummaryOfTrafficBandwidthResponseMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" require:"true" type:"Struct"`
  // {"en" : "This object contains the breakdown of traffic by group. Groups are determined by the request body.", "zh_CN": "每个分组及其带宽。"}
  Groups []*GetASummaryOfTrafficBandwidthResponseGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetASummaryOfTrafficBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthResponse) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficBandwidthResponse) SetMetaData(v *GetASummaryOfTrafficBandwidthResponseMetaData) *GetASummaryOfTrafficBandwidthResponse {
  s.MetaData = v
  return s
}

func (s *GetASummaryOfTrafficBandwidthResponse) SetGroups(v []*GetASummaryOfTrafficBandwidthResponseGroups) *GetASummaryOfTrafficBandwidthResponse {
  s.Groups = v
  return s
}

type GetASummaryOfTrafficBandwidthResponseMetaData struct {
  // {"en" : "RFC 3339 date indicating the beginning of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的起始时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en" : "RFC 3339 date indicating the end.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en" : "The response can contain up to 10000 groups. If there are more groups, isComplete will be false.", "zh_CN": "返回最多10000个分组的数据。如果实际组数超过10000，则isComplete将为false。"}
  IsComplete *bool `json:"isComplete,omitempty" xml:"isComplete,omitempty" require:"true"`
  // {"en" : "Indicates the type of data returned. 'edge bandwidth' represents edge bandwidth traffic. ", "zh_CN": "表示返回的数据类型。'edge bandwidth'指边缘带宽。"}
  DataNames []*string `json:"dataNames,omitempty" xml:"dataNames,omitempty" type:"Repeated"`
  // {"en" : "Indicates the unit of measurement of the returned values.", "zh_CN": "表示返回值的计量单位。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
}

func (s GetASummaryOfTrafficBandwidthResponseMetaData) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthResponseMetaData) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficBandwidthResponseMetaData) SetStartTime(v string) *GetASummaryOfTrafficBandwidthResponseMetaData {
  s.StartTime = &v
  return s
}

func (s *GetASummaryOfTrafficBandwidthResponseMetaData) SetEndTime(v string) *GetASummaryOfTrafficBandwidthResponseMetaData {
  s.EndTime = &v
  return s
}

func (s *GetASummaryOfTrafficBandwidthResponseMetaData) SetIsComplete(v bool) *GetASummaryOfTrafficBandwidthResponseMetaData {
  s.IsComplete = &v
  return s
}

func (s *GetASummaryOfTrafficBandwidthResponseMetaData) SetDataNames(v []*string) *GetASummaryOfTrafficBandwidthResponseMetaData {
  s.DataNames = v
  return s
}

func (s *GetASummaryOfTrafficBandwidthResponseMetaData) SetDataUnit(v string) *GetASummaryOfTrafficBandwidthResponseMetaData {
  s.DataUnit = &v
  return s
}

type GetASummaryOfTrafficBandwidthResponseGroups struct     {
  // {"en" : "Name of a group.  '__all__' is a special group encompassing all groups.", "zh_CN": "分组名称。'__all__' 是一个特殊分组，表示总带宽。"}
  Group *string `json:"group,omitempty" xml:"group,omitempty"`
  // {"en" : "Data values. The units of measurement are determined by the dataUnit field.", "zh_CN": "每个分组的带宽值。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetASummaryOfTrafficBandwidthResponseGroups) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficBandwidthResponseGroups) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficBandwidthResponseGroups) SetGroup(v string) *GetASummaryOfTrafficBandwidthResponseGroups {
  s.Group = &v
  return s
}

func (s *GetASummaryOfTrafficBandwidthResponseGroups) SetData(v []*string) *GetASummaryOfTrafficBandwidthResponseGroups {
  s.Data = v
  return s
}




type GetEdgeHostnameSummaryStatisticsPaths struct {
}

func (s GetEdgeHostnameSummaryStatisticsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsPaths) GoString() string {
  return s.String()
}

type GetEdgeHostnameSummaryStatisticsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-11-14T00:00:00Z", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time span. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z. Due to latency associated with new traffic data, enddate should be no later than five minutes before the current time. This ensures you get the most accurate results.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。由于数据处理存在延迟，所指定的结束时间必须至少比当前时间早5分钟，否则返回的数据可能不准确。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
}

func (s GetEdgeHostnameSummaryStatisticsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsParameters) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameSummaryStatisticsParameters) SetStartdate(v string) *GetEdgeHostnameSummaryStatisticsParameters {
  s.Startdate = &v
  return s
}

func (s *GetEdgeHostnameSummaryStatisticsParameters) SetEnddate(v string) *GetEdgeHostnameSummaryStatisticsParameters {
  s.Enddate = &v
  return s
}

type GetEdgeHostnameSummaryStatisticsRequestHeader struct {
}

func (s GetEdgeHostnameSummaryStatisticsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsRequestHeader) GoString() string {
  return s.String()
}

type GetEdgeHostnameSummaryStatisticsRequest struct {
  // {"en" : "Limit statistics to specific edge hostnames.
  // ", "zh_CN": "指定查询范围。"}
  Filters *GetEdgeHostnameSummaryStatisticsRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
  // {"en" : "Specify an array containing 'edgeHostnames' to get data for each edge hostname. Omit to get only a cumulative total.", "zh_CN": "指定edgeHostnames，按调度域名分组汇总返回数据。如未指定，则只返回所有调度域名的汇总数据。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetEdgeHostnameSummaryStatisticsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsRequest) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameSummaryStatisticsRequest) SetFilters(v *GetEdgeHostnameSummaryStatisticsRequestFilters) *GetEdgeHostnameSummaryStatisticsRequest {
  s.Filters = v
  return s
}

func (s *GetEdgeHostnameSummaryStatisticsRequest) SetGroupBy(v []*string) *GetEdgeHostnameSummaryStatisticsRequest {
  s.GroupBy = v
  return s
}

type GetEdgeHostnameSummaryStatisticsRequestFilters struct {
  // {"en" : "One or more edge hostnames.", "zh_CN": "指定一个或多个调度域名进行查询。"}
  EdgeHostnames []*string `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty" type:"Repeated"`
}

func (s GetEdgeHostnameSummaryStatisticsRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsRequestFilters) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameSummaryStatisticsRequestFilters) SetEdgeHostnames(v []*string) *GetEdgeHostnameSummaryStatisticsRequestFilters {
  s.EdgeHostnames = v
  return s
}

type GetEdgeHostnameSummaryStatisticsResponseHeader struct {
}

func (s GetEdgeHostnameSummaryStatisticsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsResponseHeader) GoString() string {
  return s.String()
}

type GetEdgeHostnameSummaryStatisticsResponse struct {
  // {"en" : "This object contains fields describing the data returned in the groups object.", "zh_CN": "此对象包含的字段是对groups对象中返回数据的描述。"}
  MetaData *GetEdgeHostnameSummaryStatisticsResponseMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" require:"true" type:"Struct"`
  // {"en" : "This object contains the breakdown of DNS requests by group. Groups are determined by the request body.", "zh_CN": "按调度域名分组汇总的数据。"}
  Groups []*GetEdgeHostnameSummaryStatisticsResponseGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetEdgeHostnameSummaryStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsResponse) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameSummaryStatisticsResponse) SetMetaData(v *GetEdgeHostnameSummaryStatisticsResponseMetaData) *GetEdgeHostnameSummaryStatisticsResponse {
  s.MetaData = v
  return s
}

func (s *GetEdgeHostnameSummaryStatisticsResponse) SetGroups(v []*GetEdgeHostnameSummaryStatisticsResponseGroups) *GetEdgeHostnameSummaryStatisticsResponse {
  s.Groups = v
  return s
}

type GetEdgeHostnameSummaryStatisticsResponseMetaData struct {
  // {"en" : "RFC 3339 date indicating the beginning of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的起始时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en" : "RFC 3339 date indicating the end of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en" : "The response can contain up to 10000 groups. If there are more groups, isComplete will be false.", "zh_CN": "该接口最多返回10000个分组的数据。如果实际分组数量大于10000，则isComplete的值将为false。"}
  IsComplete *bool `json:"isComplete,omitempty" xml:"isComplete,omitempty" require:"true"`
  // {"en" : "Indicates the type of data returned: 'edgeHostname dns request'", "zh_CN": "返回的数据类型，即调度域名DNS解析请求数。"}
  DataNames []*string `json:"dataNames,omitempty" xml:"dataNames,omitempty" type:"Repeated"`
  // {"en" : "Indicates the unit of measurement of the returned values.", "zh_CN": "返回值的计量单位。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
}

func (s GetEdgeHostnameSummaryStatisticsResponseMetaData) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsResponseMetaData) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameSummaryStatisticsResponseMetaData) SetStartTime(v string) *GetEdgeHostnameSummaryStatisticsResponseMetaData {
  s.StartTime = &v
  return s
}

func (s *GetEdgeHostnameSummaryStatisticsResponseMetaData) SetEndTime(v string) *GetEdgeHostnameSummaryStatisticsResponseMetaData {
  s.EndTime = &v
  return s
}

func (s *GetEdgeHostnameSummaryStatisticsResponseMetaData) SetIsComplete(v bool) *GetEdgeHostnameSummaryStatisticsResponseMetaData {
  s.IsComplete = &v
  return s
}

func (s *GetEdgeHostnameSummaryStatisticsResponseMetaData) SetDataNames(v []*string) *GetEdgeHostnameSummaryStatisticsResponseMetaData {
  s.DataNames = v
  return s
}

func (s *GetEdgeHostnameSummaryStatisticsResponseMetaData) SetDataUnit(v string) *GetEdgeHostnameSummaryStatisticsResponseMetaData {
  s.DataUnit = &v
  return s
}

type GetEdgeHostnameSummaryStatisticsResponseGroups struct     {
  // {"en" : "Name of a group.  '__all__' is a special group encompassing all groups.
  // ", "zh_CN": "分组名称。'__all__' 是一个特殊分组，包含其它所有分组的数据。"}
  Group *string `json:"group,omitempty" xml:"group,omitempty"`
  // {"en" : "Data value. The units of measurement are determined by the dataUnit field.", "zh_CN": "DNS解析请求数。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetEdgeHostnameSummaryStatisticsResponseGroups) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameSummaryStatisticsResponseGroups) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameSummaryStatisticsResponseGroups) SetGroup(v string) *GetEdgeHostnameSummaryStatisticsResponseGroups {
  s.Group = &v
  return s
}

func (s *GetEdgeHostnameSummaryStatisticsResponseGroups) SetData(v []*string) *GetEdgeHostnameSummaryStatisticsResponseGroups {
  s.Data = v
  return s
}




type GetOriginStatusCodeDetailsPaths struct {
}

func (s GetOriginStatusCodeDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsPaths) GoString() string {
  return s.String()
}

type GetOriginStatusCodeDetailsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes,hourly,daily,monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8  means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Range: [ 0 .. 5 ] characters 
  // Enum: http,https,all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetOriginStatusCodeDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsParameters) GoString() string {
  return s.String()
}

func (s *GetOriginStatusCodeDetailsParameters) SetStartdate(v string) *GetOriginStatusCodeDetailsParameters {
  s.Startdate = &v
  return s
}

func (s *GetOriginStatusCodeDetailsParameters) SetEnddate(v string) *GetOriginStatusCodeDetailsParameters {
  s.Enddate = &v
  return s
}

func (s *GetOriginStatusCodeDetailsParameters) SetType(v string) *GetOriginStatusCodeDetailsParameters {
  s.Type = &v
  return s
}

func (s *GetOriginStatusCodeDetailsParameters) SetScheme(v string) *GetOriginStatusCodeDetailsParameters {
  s.Scheme = &v
  return s
}

type GetOriginStatusCodeDetailsRequestHeader struct {
}

func (s GetOriginStatusCodeDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsRequestHeader) GoString() string {
  return s.String()
}

type GetOriginStatusCodeDetailsRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetOriginStatusCodeDetailsRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetOriginStatusCodeDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsRequest) GoString() string {
  return s.String()
}

func (s *GetOriginStatusCodeDetailsRequest) SetFilters(v *GetOriginStatusCodeDetailsRequestFilters) *GetOriginStatusCodeDetailsRequest {
  s.Filters = v
  return s
}

type GetOriginStatusCodeDetailsRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
}

func (s GetOriginStatusCodeDetailsRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsRequestFilters) GoString() string {
  return s.String()
}

func (s *GetOriginStatusCodeDetailsRequestFilters) SetHostnames(v []*string) *GetOriginStatusCodeDetailsRequestFilters {
  s.Hostnames = v
  return s
}

type GetOriginStatusCodeDetailsResponseHeader struct {
}

func (s GetOriginStatusCodeDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsResponseHeader) GoString() string {
  return s.String()
}

type GetOriginStatusCodeDetailsResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Enum: counts 
  // Unit of measurement.  For status code APIs, it will be 'counts', which indicates the number of times a status code appears.", "zh_CN": "取值范围: counts 
  // 计量单位。对于状态码报表接口，单位为'次数'，表示某个状态码出现的次数。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "Contains information about status codes at different points in time.", "zh_CN": "不同时间点的状态码统计信息。"}
  DataSeries []*GetOriginStatusCodeDetailsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginStatusCodeDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsResponse) GoString() string {
  return s.String()
}

func (s *GetOriginStatusCodeDetailsResponse) SetDataName(v string) *GetOriginStatusCodeDetailsResponse {
  s.DataName = &v
  return s
}

func (s *GetOriginStatusCodeDetailsResponse) SetDataUnit(v string) *GetOriginStatusCodeDetailsResponse {
  s.DataUnit = &v
  return s
}

func (s *GetOriginStatusCodeDetailsResponse) SetDataSeries(v []*GetOriginStatusCodeDetailsResponseDataSeries) *GetOriginStatusCodeDetailsResponse {
  s.DataSeries = v
  return s
}

type GetOriginStatusCodeDetailsResponseDataSeries struct     {
  // {"en" : "RFC 3339 format date indicate the beginning of an interval.", "zh_CN": "RFC 3339格式的日期，表示每个时间段的开始时间。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "", "zh_CN": ""}
  Details []*GetOriginStatusCodeDetailsResponseDataSeriesDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s GetOriginStatusCodeDetailsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetOriginStatusCodeDetailsResponseDataSeries) SetTimestamp(v string) *GetOriginStatusCodeDetailsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetOriginStatusCodeDetailsResponseDataSeries) SetDetails(v []*GetOriginStatusCodeDetailsResponseDataSeriesDetails) *GetOriginStatusCodeDetailsResponseDataSeries {
  s.Details = v
  return s
}

type GetOriginStatusCodeDetailsResponseDataSeriesDetails struct     {
  // {"en" : "Indicates an HTTP status code, for example, '200'.", "zh_CN": "HTTP状态码，例如'200'。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  // {"en" : "Range: >= 0 
  // Indicates the number of times the status code was returned.", "zh_CN": "取值范围: >= 0 
  // 状态码出现的次数。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetOriginStatusCodeDetailsResponseDataSeriesDetails) String() string {
  return tea.Prettify(s)
}

func (s GetOriginStatusCodeDetailsResponseDataSeriesDetails) GoString() string {
  return s.String()
}

func (s *GetOriginStatusCodeDetailsResponseDataSeriesDetails) SetStatusCode(v string) *GetOriginStatusCodeDetailsResponseDataSeriesDetails {
  s.StatusCode = &v
  return s
}

func (s *GetOriginStatusCodeDetailsResponseDataSeriesDetails) SetData(v int) *GetOriginStatusCodeDetailsResponseDataSeriesDetails {
  s.Data = &v
  return s
}




type GetEdgeBandwidthPaths struct {
}

func (s GetEdgeBandwidthPaths) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeBandwidthPaths) GoString() string {
  return s.String()
}

type GetEdgeBandwidthParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes hourly daily monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8  means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes, hourly, daily, monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "[ 0 .. 5 ] characters 
  // Enum: http https all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "[ 0 .. 5 ] 字符 
  // 取值范围: http, https, all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetEdgeBandwidthParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeBandwidthParameters) GoString() string {
  return s.String()
}

func (s *GetEdgeBandwidthParameters) SetStartdate(v string) *GetEdgeBandwidthParameters {
  s.Startdate = &v
  return s
}

func (s *GetEdgeBandwidthParameters) SetEnddate(v string) *GetEdgeBandwidthParameters {
  s.Enddate = &v
  return s
}

func (s *GetEdgeBandwidthParameters) SetType(v string) *GetEdgeBandwidthParameters {
  s.Type = &v
  return s
}

func (s *GetEdgeBandwidthParameters) SetScheme(v string) *GetEdgeBandwidthParameters {
  s.Scheme = &v
  return s
}

type GetEdgeBandwidthRequestHeader struct {
}

func (s GetEdgeBandwidthRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeBandwidthRequestHeader) GoString() string {
  return s.String()
}

type GetEdgeBandwidthRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetEdgeBandwidthRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetEdgeBandwidthRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeBandwidthRequest) GoString() string {
  return s.String()
}

func (s *GetEdgeBandwidthRequest) SetFilters(v *GetEdgeBandwidthRequestFilters) *GetEdgeBandwidthRequest {
  s.Filters = v
  return s
}

type GetEdgeBandwidthRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Enum: standard premium deluxe ultra 
  // Indicates one or more server groups.", "zh_CN": "取值范围: standard, premium, deluxe, ultra 
  // 指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetEdgeBandwidthRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeBandwidthRequestFilters) GoString() string {
  return s.String()
}

func (s *GetEdgeBandwidthRequestFilters) SetHostnames(v []*string) *GetEdgeBandwidthRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetEdgeBandwidthRequestFilters) SetServerGroups(v []*string) *GetEdgeBandwidthRequestFilters {
  s.ServerGroups = v
  return s
}

type GetEdgeBandwidthResponseHeader struct {
}

func (s GetEdgeBandwidthResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeBandwidthResponseHeader) GoString() string {
  return s.String()
}

type GetEdgeBandwidthResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetEdgeBandwidthResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetEdgeBandwidthResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeBandwidthResponse) GoString() string {
  return s.String()
}

func (s *GetEdgeBandwidthResponse) SetDataName(v string) *GetEdgeBandwidthResponse {
  s.DataName = &v
  return s
}

func (s *GetEdgeBandwidthResponse) SetDataUnit(v string) *GetEdgeBandwidthResponse {
  s.DataUnit = &v
  return s
}

func (s *GetEdgeBandwidthResponse) SetDataSeries(v []*GetEdgeBandwidthResponseDataSeries) *GetEdgeBandwidthResponse {
  s.DataSeries = v
  return s
}

type GetEdgeBandwidthResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetEdgeBandwidthResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeBandwidthResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetEdgeBandwidthResponseDataSeries) SetTimestamp(v string) *GetEdgeBandwidthResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetEdgeBandwidthResponseDataSeries) SetData(v int) *GetEdgeBandwidthResponseDataSeries {
  s.Data = &v
  return s
}




type GetEdgeStatusCodeDetailsPaths struct {
}

func (s GetEdgeStatusCodeDetailsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsPaths) GoString() string {
  return s.String()
}

type GetEdgeStatusCodeDetailsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes,hourly,daily,monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8  means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Range: [ 0 .. 5 ] characters 
  // Enum: http,https,all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetEdgeStatusCodeDetailsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsParameters) GoString() string {
  return s.String()
}

func (s *GetEdgeStatusCodeDetailsParameters) SetStartdate(v string) *GetEdgeStatusCodeDetailsParameters {
  s.Startdate = &v
  return s
}

func (s *GetEdgeStatusCodeDetailsParameters) SetEnddate(v string) *GetEdgeStatusCodeDetailsParameters {
  s.Enddate = &v
  return s
}

func (s *GetEdgeStatusCodeDetailsParameters) SetType(v string) *GetEdgeStatusCodeDetailsParameters {
  s.Type = &v
  return s
}

func (s *GetEdgeStatusCodeDetailsParameters) SetScheme(v string) *GetEdgeStatusCodeDetailsParameters {
  s.Scheme = &v
  return s
}

type GetEdgeStatusCodeDetailsRequestHeader struct {
}

func (s GetEdgeStatusCodeDetailsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsRequestHeader) GoString() string {
  return s.String()
}

type GetEdgeStatusCodeDetailsRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetEdgeStatusCodeDetailsRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetEdgeStatusCodeDetailsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsRequest) GoString() string {
  return s.String()
}

func (s *GetEdgeStatusCodeDetailsRequest) SetFilters(v *GetEdgeStatusCodeDetailsRequestFilters) *GetEdgeStatusCodeDetailsRequest {
  s.Filters = v
  return s
}

type GetEdgeStatusCodeDetailsRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
}

func (s GetEdgeStatusCodeDetailsRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsRequestFilters) GoString() string {
  return s.String()
}

func (s *GetEdgeStatusCodeDetailsRequestFilters) SetHostnames(v []*string) *GetEdgeStatusCodeDetailsRequestFilters {
  s.Hostnames = v
  return s
}

type GetEdgeStatusCodeDetailsResponseHeader struct {
}

func (s GetEdgeStatusCodeDetailsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsResponseHeader) GoString() string {
  return s.String()
}

type GetEdgeStatusCodeDetailsResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Enum: counts 
  // Unit of measurement.  For status code APIs, it will be 'counts', which indicates the number of times a status code appears.", "zh_CN": "取值范围: counts 
  // 计量单位。对于状态码报表接口，单位为'次数'，表示某个状态码出现的次数。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "Contains information about status codes at different points in time.", "zh_CN": "不同时间点的状态码统计信息。"}
  DataSeries []*GetEdgeStatusCodeDetailsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetEdgeStatusCodeDetailsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsResponse) GoString() string {
  return s.String()
}

func (s *GetEdgeStatusCodeDetailsResponse) SetDataName(v string) *GetEdgeStatusCodeDetailsResponse {
  s.DataName = &v
  return s
}

func (s *GetEdgeStatusCodeDetailsResponse) SetDataUnit(v string) *GetEdgeStatusCodeDetailsResponse {
  s.DataUnit = &v
  return s
}

func (s *GetEdgeStatusCodeDetailsResponse) SetDataSeries(v []*GetEdgeStatusCodeDetailsResponseDataSeries) *GetEdgeStatusCodeDetailsResponse {
  s.DataSeries = v
  return s
}

type GetEdgeStatusCodeDetailsResponseDataSeries struct     {
  // {"en" : "RFC 3339 format date indicate the beginning of an interval.", "zh_CN": "RFC 3339格式的日期，表示每个时间段的开始时间。"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "", "zh_CN": ""}
  Details []*GetEdgeStatusCodeDetailsResponseDataSeriesDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s GetEdgeStatusCodeDetailsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetEdgeStatusCodeDetailsResponseDataSeries) SetTimestamp(v string) *GetEdgeStatusCodeDetailsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetEdgeStatusCodeDetailsResponseDataSeries) SetDetails(v []*GetEdgeStatusCodeDetailsResponseDataSeriesDetails) *GetEdgeStatusCodeDetailsResponseDataSeries {
  s.Details = v
  return s
}

type GetEdgeStatusCodeDetailsResponseDataSeriesDetails struct     {
  // {"en" : "Indicates an HTTP status code, for example, '200'.", "zh_CN": "HTTP状态码，例如'200'。"}
  StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  // {"en" : "Range: >= 0 
  // Indicates the number of times the status code was returned.", "zh_CN": "取值范围: >= 0 
  // 状态码出现的次数。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetEdgeStatusCodeDetailsResponseDataSeriesDetails) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeStatusCodeDetailsResponseDataSeriesDetails) GoString() string {
  return s.String()
}

func (s *GetEdgeStatusCodeDetailsResponseDataSeriesDetails) SetStatusCode(v string) *GetEdgeStatusCodeDetailsResponseDataSeriesDetails {
  s.StatusCode = &v
  return s
}

func (s *GetEdgeStatusCodeDetailsResponseDataSeriesDetails) SetData(v int) *GetEdgeStatusCodeDetailsResponseDataSeriesDetails {
  s.Data = &v
  return s
}




type GetTheNumberOfRequestsToOriginPaths struct {
}

func (s GetTheNumberOfRequestsToOriginPaths) String() string {
  return tea.Prettify(s)
}

func (s GetTheNumberOfRequestsToOriginPaths) GoString() string {
  return s.String()
}

type GetTheNumberOfRequestsToOriginParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes,hourly,daily,monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8  means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Range: [ 0 .. 5 ] characters 
  // Enum: http,https,all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetTheNumberOfRequestsToOriginParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTheNumberOfRequestsToOriginParameters) GoString() string {
  return s.String()
}

func (s *GetTheNumberOfRequestsToOriginParameters) SetStartdate(v string) *GetTheNumberOfRequestsToOriginParameters {
  s.Startdate = &v
  return s
}

func (s *GetTheNumberOfRequestsToOriginParameters) SetEnddate(v string) *GetTheNumberOfRequestsToOriginParameters {
  s.Enddate = &v
  return s
}

func (s *GetTheNumberOfRequestsToOriginParameters) SetType(v string) *GetTheNumberOfRequestsToOriginParameters {
  s.Type = &v
  return s
}

func (s *GetTheNumberOfRequestsToOriginParameters) SetScheme(v string) *GetTheNumberOfRequestsToOriginParameters {
  s.Scheme = &v
  return s
}

type GetTheNumberOfRequestsToOriginRequestHeader struct {
}

func (s GetTheNumberOfRequestsToOriginRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTheNumberOfRequestsToOriginRequestHeader) GoString() string {
  return s.String()
}

type GetTheNumberOfRequestsToOriginRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetTheNumberOfRequestsToOriginRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetTheNumberOfRequestsToOriginRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTheNumberOfRequestsToOriginRequest) GoString() string {
  return s.String()
}

func (s *GetTheNumberOfRequestsToOriginRequest) SetFilters(v *GetTheNumberOfRequestsToOriginRequestFilters) *GetTheNumberOfRequestsToOriginRequest {
  s.Filters = v
  return s
}

type GetTheNumberOfRequestsToOriginRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
}

func (s GetTheNumberOfRequestsToOriginRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetTheNumberOfRequestsToOriginRequestFilters) GoString() string {
  return s.String()
}

func (s *GetTheNumberOfRequestsToOriginRequestFilters) SetHostnames(v []*string) *GetTheNumberOfRequestsToOriginRequestFilters {
  s.Hostnames = v
  return s
}

type GetTheNumberOfRequestsToOriginResponseHeader struct {
}

func (s GetTheNumberOfRequestsToOriginResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTheNumberOfRequestsToOriginResponseHeader) GoString() string {
  return s.String()
}

type GetTheNumberOfRequestsToOriginResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetTheNumberOfRequestsToOriginResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTheNumberOfRequestsToOriginResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTheNumberOfRequestsToOriginResponse) GoString() string {
  return s.String()
}

func (s *GetTheNumberOfRequestsToOriginResponse) SetDataName(v string) *GetTheNumberOfRequestsToOriginResponse {
  s.DataName = &v
  return s
}

func (s *GetTheNumberOfRequestsToOriginResponse) SetDataUnit(v string) *GetTheNumberOfRequestsToOriginResponse {
  s.DataUnit = &v
  return s
}

func (s *GetTheNumberOfRequestsToOriginResponse) SetDataSeries(v []*GetTheNumberOfRequestsToOriginResponseDataSeries) *GetTheNumberOfRequestsToOriginResponse {
  s.DataSeries = v
  return s
}

type GetTheNumberOfRequestsToOriginResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetTheNumberOfRequestsToOriginResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTheNumberOfRequestsToOriginResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetTheNumberOfRequestsToOriginResponseDataSeries) SetTimestamp(v string) *GetTheNumberOfRequestsToOriginResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTheNumberOfRequestsToOriginResponseDataSeries) SetData(v int) *GetTheNumberOfRequestsToOriginResponseDataSeries {
  s.Data = &v
  return s
}




type GetEdgeHostnameStatisticsPaths struct {
}

func (s GetEdgeHostnameStatisticsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameStatisticsPaths) GoString() string {
  return s.String()
}

type GetEdgeHostnameStatisticsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-11-14T00:00:00Z", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time span. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-11-14T00:00:00Z。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes,hourly,daily,monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8 means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s GetEdgeHostnameStatisticsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameStatisticsParameters) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameStatisticsParameters) SetStartdate(v string) *GetEdgeHostnameStatisticsParameters {
  s.Startdate = &v
  return s
}

func (s *GetEdgeHostnameStatisticsParameters) SetEnddate(v string) *GetEdgeHostnameStatisticsParameters {
  s.Enddate = &v
  return s
}

func (s *GetEdgeHostnameStatisticsParameters) SetType(v string) *GetEdgeHostnameStatisticsParameters {
  s.Type = &v
  return s
}

type GetEdgeHostnameStatisticsRequestHeader struct {
}

func (s GetEdgeHostnameStatisticsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameStatisticsRequestHeader) GoString() string {
  return s.String()
}

type GetEdgeHostnameStatisticsRequest struct {
  // {"en" : "Limit statistics to specific edge hostnames.", "zh_CN": "指定调度域名进行查询。"}
  Filters *GetEdgeHostnameStatisticsRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetEdgeHostnameStatisticsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameStatisticsRequest) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameStatisticsRequest) SetFilters(v *GetEdgeHostnameStatisticsRequestFilters) *GetEdgeHostnameStatisticsRequest {
  s.Filters = v
  return s
}

type GetEdgeHostnameStatisticsRequestFilters struct {
  // {"en" : "One or more edge hostnames.", "zh_CN": "一个或多个调度域名。"}
  EdgeHostnames []*string `json:"edgeHostnames,omitempty" xml:"edgeHostnames,omitempty" type:"Repeated"`
}

func (s GetEdgeHostnameStatisticsRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameStatisticsRequestFilters) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameStatisticsRequestFilters) SetEdgeHostnames(v []*string) *GetEdgeHostnameStatisticsRequestFilters {
  s.EdgeHostnames = v
  return s
}

type GetEdgeHostnameStatisticsResponseHeader struct {
}

func (s GetEdgeHostnameStatisticsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameStatisticsResponseHeader) GoString() string {
  return s.String()
}

type GetEdgeHostnameStatisticsResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetEdgeHostnameStatisticsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetEdgeHostnameStatisticsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameStatisticsResponse) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameStatisticsResponse) SetDataName(v string) *GetEdgeHostnameStatisticsResponse {
  s.DataName = &v
  return s
}

func (s *GetEdgeHostnameStatisticsResponse) SetDataUnit(v string) *GetEdgeHostnameStatisticsResponse {
  s.DataUnit = &v
  return s
}

func (s *GetEdgeHostnameStatisticsResponse) SetDataSeries(v []*GetEdgeHostnameStatisticsResponseDataSeries) *GetEdgeHostnameStatisticsResponse {
  s.DataSeries = v
  return s
}

type GetEdgeHostnameStatisticsResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetEdgeHostnameStatisticsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeHostnameStatisticsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetEdgeHostnameStatisticsResponseDataSeries) SetTimestamp(v string) *GetEdgeHostnameStatisticsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetEdgeHostnameStatisticsResponseDataSeries) SetData(v int) *GetEdgeHostnameStatisticsResponseDataSeries {
  s.Data = &v
  return s
}




type GetCpuTimeUsedPaths struct {
}

func (s GetCpuTimeUsedPaths) String() string {
  return tea.Prettify(s)
}

func (s GetCpuTimeUsedPaths) GoString() string {
  return s.String()
}

type GetCpuTimeUsedParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14. For example, type=daily+8 means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetCpuTimeUsedParameters) String() string {
  return tea.Prettify(s)
}

func (s GetCpuTimeUsedParameters) GoString() string {
  return s.String()
}

func (s *GetCpuTimeUsedParameters) SetStartdate(v string) *GetCpuTimeUsedParameters {
  s.Startdate = &v
  return s
}

func (s *GetCpuTimeUsedParameters) SetEnddate(v string) *GetCpuTimeUsedParameters {
  s.Enddate = &v
  return s
}

func (s *GetCpuTimeUsedParameters) SetType(v string) *GetCpuTimeUsedParameters {
  s.Type = &v
  return s
}

func (s *GetCpuTimeUsedParameters) SetScheme(v string) *GetCpuTimeUsedParameters {
  s.Scheme = &v
  return s
}

type GetCpuTimeUsedRequestHeader struct {
}

func (s GetCpuTimeUsedRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCpuTimeUsedRequestHeader) GoString() string {
  return s.String()
}

type GetCpuTimeUsedRequest struct {
  // {"en" : "Filter results by specifying hostnames or server groups.", "zh_CN": "Filter results by specifying hostnames or server groups."}
  Filters *GetCpuTimeUsedRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetCpuTimeUsedRequest) String() string {
  return tea.Prettify(s)
}

func (s GetCpuTimeUsedRequest) GoString() string {
  return s.String()
}

func (s *GetCpuTimeUsedRequest) SetFilters(v *GetCpuTimeUsedRequestFilters) *GetCpuTimeUsedRequest {
  s.Filters = v
  return s
}

type GetCpuTimeUsedRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "One or more server groups for which to return data. If unspecified, data for all server groups will be returned.", "zh_CN": "指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetCpuTimeUsedRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetCpuTimeUsedRequestFilters) GoString() string {
  return s.String()
}

func (s *GetCpuTimeUsedRequestFilters) SetHostnames(v []*string) *GetCpuTimeUsedRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetCpuTimeUsedRequestFilters) SetServerGroups(v []*string) *GetCpuTimeUsedRequestFilters {
  s.ServerGroups = v
  return s
}

type GetCpuTimeUsedResponseHeader struct {
}

func (s GetCpuTimeUsedResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetCpuTimeUsedResponseHeader) GoString() string {
  return s.String()
}

type GetCpuTimeUsedResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetCpuTimeUsedResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetCpuTimeUsedResponse) String() string {
  return tea.Prettify(s)
}

func (s GetCpuTimeUsedResponse) GoString() string {
  return s.String()
}

func (s *GetCpuTimeUsedResponse) SetDataName(v string) *GetCpuTimeUsedResponse {
  s.DataName = &v
  return s
}

func (s *GetCpuTimeUsedResponse) SetDataUnit(v string) *GetCpuTimeUsedResponse {
  s.DataUnit = &v
  return s
}

func (s *GetCpuTimeUsedResponse) SetDataSeries(v []*GetCpuTimeUsedResponseDataSeries) *GetCpuTimeUsedResponse {
  s.DataSeries = v
  return s
}

type GetCpuTimeUsedResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetCpuTimeUsedResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetCpuTimeUsedResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetCpuTimeUsedResponseDataSeries) SetTimestamp(v string) *GetCpuTimeUsedResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetCpuTimeUsedResponseDataSeries) SetData(v int) *GetCpuTimeUsedResponseDataSeries {
  s.Data = &v
  return s
}




type GetASummaryOfRequestsPaths struct {
}

func (s GetASummaryOfRequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsPaths) GoString() string {
  return s.String()
}

type GetASummaryOfRequestsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z ", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z. Due to latency associated with new traffic data, enddate should be no later than five minutes before the current time. This ensures you get the most accurate results.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。由于数据处理存在延迟，所指定的结束时间必须至少比当前时间早5分钟，否则返回的数据可能不准确。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: http https all 
  // Default: all 
  // Limit the results to the specified scheme. By default, data from HTTPS and HTTP requests is returned.", "zh_CN": "[ 0 .. 5 ] 字符 
  // 取值范围: http, https, all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetASummaryOfRequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsParameters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfRequestsParameters) SetStartdate(v string) *GetASummaryOfRequestsParameters {
  s.Startdate = &v
  return s
}

func (s *GetASummaryOfRequestsParameters) SetEnddate(v string) *GetASummaryOfRequestsParameters {
  s.Enddate = &v
  return s
}

func (s *GetASummaryOfRequestsParameters) SetScheme(v string) *GetASummaryOfRequestsParameters {
  s.Scheme = &v
  return s
}

type GetASummaryOfRequestsRequestHeader struct {
}

func (s GetASummaryOfRequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsRequestHeader) GoString() string {
  return s.String()
}

type GetASummaryOfRequestsRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetASummaryOfRequestsRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
  // {"en" : "<= 2 items 
  // items Enum: hostnames, serverGroups 
  // You can group results using a combination of up to two of the following: 'hostnames', and 'serverGroups'.", "zh_CN": "<= 2 条目 
  // 取值范围: hostnames, serverGroups 
  // 指定分组依据对数据进行分组汇总。支持按'hostnames'，'serverGroups'单独进行分组汇总，也支持同时指定这2个参数进行分组汇总。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetASummaryOfRequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsRequest) GoString() string {
  return s.String()
}

func (s *GetASummaryOfRequestsRequest) SetFilters(v *GetASummaryOfRequestsRequestFilters) *GetASummaryOfRequestsRequest {
  s.Filters = v
  return s
}

func (s *GetASummaryOfRequestsRequest) SetGroupBy(v []*string) *GetASummaryOfRequestsRequest {
  s.GroupBy = v
  return s
}

type GetASummaryOfRequestsRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Indicates one or more server groups.", "zh_CN": "指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetASummaryOfRequestsRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsRequestFilters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfRequestsRequestFilters) SetHostnames(v []*string) *GetASummaryOfRequestsRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetASummaryOfRequestsRequestFilters) SetServerGroups(v []*string) *GetASummaryOfRequestsRequestFilters {
  s.ServerGroups = v
  return s
}

type GetASummaryOfRequestsResponseHeader struct {
}

func (s GetASummaryOfRequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsResponseHeader) GoString() string {
  return s.String()
}

type GetASummaryOfRequestsResponse struct {
  // {"en" : "This object contains fields describing the data returned in the groups object.", "zh_CN": "此对象包含的字段是对groups对象中返回数据的描述。"}
  MetaData *GetASummaryOfRequestsResponseMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" require:"true" type:"Struct"`
  // {"en" : "This object contains the breakdown of requests by group.", "zh_CN": "每个分组及其请求数。"}
  Groups []*GetASummaryOfRequestsResponseGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetASummaryOfRequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsResponse) GoString() string {
  return s.String()
}

func (s *GetASummaryOfRequestsResponse) SetMetaData(v *GetASummaryOfRequestsResponseMetaData) *GetASummaryOfRequestsResponse {
  s.MetaData = v
  return s
}

func (s *GetASummaryOfRequestsResponse) SetGroups(v []*GetASummaryOfRequestsResponseGroups) *GetASummaryOfRequestsResponse {
  s.Groups = v
  return s
}

type GetASummaryOfRequestsResponseMetaData struct {
  // {"en" : "RFC 3339 date indicating the beginning of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的起始时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en" : "RFC 3339 date indicating the end of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en" : "The response can contain up to 10000 groups. If there are more groups, isComplete will be false.", "zh_CN": "该接口最多返回10000个分组的数据。如果实际分组数量大于10000，则isComplete将为false。"}
  IsComplete *bool `json:"isComplete,omitempty" xml:"isComplete,omitempty" require:"true"`
  // {"en" : "Indicates the type of data returned. 'edge request' represents edge traffic. 'fast route request' refers to traffic from your origin accelerated through the HDT product. 'origin request' represents origin traffic. 'intermediate request' represents intermediate traffic. The order of the entries in dataNames array corresponds to the order of values returned in the data data array in the response.", "zh_CN": "表示返回的数据类型。'edge request'表示边缘请求数，'fast route request'表示快速回源请求数，'origin request'表示回源请求数，'intermediate request'表示回父节点的请求数。dataNames数组中条目的顺序与groups[].data中返回值的顺序一一对应。"}
  DataNames []*string `json:"dataNames,omitempty" xml:"dataNames,omitempty" type:"Repeated"`
  // {"en" : "Indicates the unit of measurement of the returned values.", "zh_CN": "返回值的计量单位。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
}

func (s GetASummaryOfRequestsResponseMetaData) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsResponseMetaData) GoString() string {
  return s.String()
}

func (s *GetASummaryOfRequestsResponseMetaData) SetStartTime(v string) *GetASummaryOfRequestsResponseMetaData {
  s.StartTime = &v
  return s
}

func (s *GetASummaryOfRequestsResponseMetaData) SetEndTime(v string) *GetASummaryOfRequestsResponseMetaData {
  s.EndTime = &v
  return s
}

func (s *GetASummaryOfRequestsResponseMetaData) SetIsComplete(v bool) *GetASummaryOfRequestsResponseMetaData {
  s.IsComplete = &v
  return s
}

func (s *GetASummaryOfRequestsResponseMetaData) SetDataNames(v []*string) *GetASummaryOfRequestsResponseMetaData {
  s.DataNames = v
  return s
}

func (s *GetASummaryOfRequestsResponseMetaData) SetDataUnit(v string) *GetASummaryOfRequestsResponseMetaData {
  s.DataUnit = &v
  return s
}

type GetASummaryOfRequestsResponseGroups struct     {
  // {"en" : "Name of a group.  '__all__' is a special group encompassing all groups.", "zh_CN": "分组名称。'__all__' 是一个特殊分组，表示总请求数。"}
  Group *string `json:"group,omitempty" xml:"group,omitempty"`
  // {"en" : "Data values. The units of measurement are determined by the dataUnit field.", "zh_CN": "每个分组的请求数。注意：当分组条件包含serverGroups时，极个别情况下，可能会出现'__all__' 组的值明显大于其它组累加的和。这是由于存在未知原因导致某些请求无法映射到serverGroup所致。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetASummaryOfRequestsResponseGroups) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfRequestsResponseGroups) GoString() string {
  return s.String()
}

func (s *GetASummaryOfRequestsResponseGroups) SetGroup(v string) *GetASummaryOfRequestsResponseGroups {
  s.Group = &v
  return s
}

func (s *GetASummaryOfRequestsResponseGroups) SetData(v []*string) *GetASummaryOfRequestsResponseGroups {
  s.Data = v
  return s
}




type GetOriginFastRouteRequestsPaths struct {
}

func (s GetOriginFastRouteRequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteRequestsPaths) GoString() string {
  return s.String()
}

type GetOriginFastRouteRequestsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8 means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetOriginFastRouteRequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteRequestsParameters) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteRequestsParameters) SetStartdate(v string) *GetOriginFastRouteRequestsParameters {
  s.Startdate = &v
  return s
}

func (s *GetOriginFastRouteRequestsParameters) SetEnddate(v string) *GetOriginFastRouteRequestsParameters {
  s.Enddate = &v
  return s
}

func (s *GetOriginFastRouteRequestsParameters) SetType(v string) *GetOriginFastRouteRequestsParameters {
  s.Type = &v
  return s
}

func (s *GetOriginFastRouteRequestsParameters) SetScheme(v string) *GetOriginFastRouteRequestsParameters {
  s.Scheme = &v
  return s
}

type GetOriginFastRouteRequestsRequestHeader struct {
}

func (s GetOriginFastRouteRequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteRequestsRequestHeader) GoString() string {
  return s.String()
}

type GetOriginFastRouteRequestsRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetOriginFastRouteRequestsRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetOriginFastRouteRequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteRequestsRequest) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteRequestsRequest) SetFilters(v *GetOriginFastRouteRequestsRequestFilters) *GetOriginFastRouteRequestsRequest {
  s.Filters = v
  return s
}

type GetOriginFastRouteRequestsRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
}

func (s GetOriginFastRouteRequestsRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteRequestsRequestFilters) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteRequestsRequestFilters) SetHostnames(v []*string) *GetOriginFastRouteRequestsRequestFilters {
  s.Hostnames = v
  return s
}

type GetOriginFastRouteRequestsResponseHeader struct {
}

func (s GetOriginFastRouteRequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteRequestsResponseHeader) GoString() string {
  return s.String()
}

type GetOriginFastRouteRequestsResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetOriginFastRouteRequestsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginFastRouteRequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteRequestsResponse) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteRequestsResponse) SetDataName(v string) *GetOriginFastRouteRequestsResponse {
  s.DataName = &v
  return s
}

func (s *GetOriginFastRouteRequestsResponse) SetDataUnit(v string) *GetOriginFastRouteRequestsResponse {
  s.DataUnit = &v
  return s
}

func (s *GetOriginFastRouteRequestsResponse) SetDataSeries(v []*GetOriginFastRouteRequestsResponseDataSeries) *GetOriginFastRouteRequestsResponse {
  s.DataSeries = v
  return s
}

type GetOriginFastRouteRequestsResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetOriginFastRouteRequestsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteRequestsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteRequestsResponseDataSeries) SetTimestamp(v string) *GetOriginFastRouteRequestsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetOriginFastRouteRequestsResponseDataSeries) SetData(v int) *GetOriginFastRouteRequestsResponseDataSeries {
  s.Data = &v
  return s
}




type GetASummaryOfTrafficVolumePaths struct {
}

func (s GetASummaryOfTrafficVolumePaths) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumePaths) GoString() string {
  return s.String()
}

type GetASummaryOfTrafficVolumeParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z ", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time span. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-10-30T00:00:00Z. Due to latency associated with processing new traffic data, enddate should be no later than five minutes before the current time. This ensures you get the most accurate results.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。由于数据处理存在延迟，所指定的结束时间必须至少比当前时间早5分钟，否则返回的数据可能不准确。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: http,https 
  // Limit the results to the specified scheme.", "zh_CN": "取值范围: http,https 
  // 指定协议进行查询。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetASummaryOfTrafficVolumeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumeParameters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficVolumeParameters) SetStartdate(v string) *GetASummaryOfTrafficVolumeParameters {
  s.Startdate = &v
  return s
}

func (s *GetASummaryOfTrafficVolumeParameters) SetEnddate(v string) *GetASummaryOfTrafficVolumeParameters {
  s.Enddate = &v
  return s
}

func (s *GetASummaryOfTrafficVolumeParameters) SetScheme(v string) *GetASummaryOfTrafficVolumeParameters {
  s.Scheme = &v
  return s
}

type GetASummaryOfTrafficVolumeRequestHeader struct {
}

func (s GetASummaryOfTrafficVolumeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumeRequestHeader) GoString() string {
  return s.String()
}

type GetASummaryOfTrafficVolumeRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetASummaryOfTrafficVolumeRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
  // {"en" : "Range: <= 2 items 
  // You can group results using a combination of up to two of the following: 'hostnames', 'serverGroups',  'customerIds', 'propertyIds', and 'propertyHostnames'.", "zh_CN": "取值范围: <= 2 条目 
  // 指定分组依据对数据进行分组汇总。支持按'hostnames'，'serverGroups'单独进行分组汇总，也支持同时指定这2个参数进行分组汇总。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetASummaryOfTrafficVolumeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumeRequest) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficVolumeRequest) SetFilters(v *GetASummaryOfTrafficVolumeRequestFilters) *GetASummaryOfTrafficVolumeRequest {
  s.Filters = v
  return s
}

func (s *GetASummaryOfTrafficVolumeRequest) SetGroupBy(v []*string) *GetASummaryOfTrafficVolumeRequest {
  s.GroupBy = v
  return s
}

type GetASummaryOfTrafficVolumeRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Indicates one or more server groups.", "zh_CN": "指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
  // {"en" : "IDs of properties for which to return data. If unspecified, data for all properties will be returned.", "zh_CN": "指定加速项目ID进行查询。"}
  PropertyIds []*string `json:"propertyIds,omitempty" xml:"propertyIds,omitempty" type:"Repeated"`
  // {"en" : "A list of property hostnames for which to return data. These correspond to values in the hostnames field of a property configuration. These can differ from the value sent in an HTTP request. For example, if you create a property to accelerate '*.test.com', a client could request content and pass a Host header of 'abc.test.com'. In this case, '*.test.com' is a property hostname you can filter by using this field while 'abc.test.com' is a hostname you can filter by using the hostnames field. By default, data for all property hostnames is returned.", "zh_CN": "指定加速域名进行查询。上述hostnames参数将匹配用户请求时访问的hostnames进行查询，而此处的propertyHostnames将匹配加速项目中定义的hostnames进行查询。当加速项目中定义的域名包含泛域名，可使用propertyHostnames参数进行查询。"}
  PropertyHostnames []*string `json:"propertyHostnames,omitempty" xml:"propertyHostnames,omitempty" type:"Repeated"`
}

func (s GetASummaryOfTrafficVolumeRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumeRequestFilters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficVolumeRequestFilters) SetHostnames(v []*string) *GetASummaryOfTrafficVolumeRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetASummaryOfTrafficVolumeRequestFilters) SetServerGroups(v []*string) *GetASummaryOfTrafficVolumeRequestFilters {
  s.ServerGroups = v
  return s
}

func (s *GetASummaryOfTrafficVolumeRequestFilters) SetPropertyIds(v []*string) *GetASummaryOfTrafficVolumeRequestFilters {
  s.PropertyIds = v
  return s
}

func (s *GetASummaryOfTrafficVolumeRequestFilters) SetPropertyHostnames(v []*string) *GetASummaryOfTrafficVolumeRequestFilters {
  s.PropertyHostnames = v
  return s
}

type GetASummaryOfTrafficVolumeResponseHeader struct {
}

func (s GetASummaryOfTrafficVolumeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumeResponseHeader) GoString() string {
  return s.String()
}

type GetASummaryOfTrafficVolumeResponse struct {
  // {"en" : "This object contains fields describing the data returned in the groups object.", "zh_CN": "此对象包含的字段是对groups对象中返回数据的描述。"}
  MetaData *GetASummaryOfTrafficVolumeResponseMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" require:"true" type:"Struct"`
  // {"en" : "This object contains the breakdown of traffic by group. ", "zh_CN": "每个分组及其流量值。"}
  Groups []*GetASummaryOfTrafficVolumeResponseGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetASummaryOfTrafficVolumeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumeResponse) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficVolumeResponse) SetMetaData(v *GetASummaryOfTrafficVolumeResponseMetaData) *GetASummaryOfTrafficVolumeResponse {
  s.MetaData = v
  return s
}

func (s *GetASummaryOfTrafficVolumeResponse) SetGroups(v []*GetASummaryOfTrafficVolumeResponseGroups) *GetASummaryOfTrafficVolumeResponse {
  s.Groups = v
  return s
}

type GetASummaryOfTrafficVolumeResponseMetaData struct {
  // {"en" : "RFC 3339 date indicating the beginning of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的起始时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en" : "RFC 3339 date indicating the end of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en" : "The response can contain up to 10000 groups. If there are more groups, isComplete will be false.", "zh_CN": "该接口最多返回10000个分组的数据。如果实际分组数量大于10000，则isComplete值为false。"}
  IsComplete *bool `json:"isComplete,omitempty" xml:"isComplete,omitempty"`
  // {"en" : "Indicates the type of data returned. 
  // <table><tr><th>Data Name</th><th>Description</th></tr><tr><td>edge response</td><td>edge traffic to visitors</td></tr><tr><td>edge request</td><td>traffic from visitors to edge servers</td></tr><tr><td>intermediate response</td><td>traffic from intermediate CDN Pro servers to other CDN Pro servers</td></tr><tr><td>intermediate request</td><td>traffic of requests to intermediate CDN Pro servers</td></tr><tr><td>origin response</td><td>traffic from origin servers</td></tr><tr><td>origin request</td><td>traffic to origin servers</td></tr><tr><td>fast route response</td><td>traffic from your origin accelerated through the HDT product</td></tr><tr><td>fast route request</td><td>traffic to your origin accelerated through the HDT product</td></tr></table>
  // The order of the entries in dataNames array corresponds to the order of values returned in the data data array in the response.", "zh_CN": "返回的数据类型。<table><tr><th>数据名称</th><th>描述</th></tr><tr><td>edge response</td><td>边缘服务器的响应流量</td></tr><tr><td>edge request</td><td>从访客到边缘服务器的请求流量</td></tr><tr><td>intermediate response</td><td>CDN Pro中间层服务器的响应流量</td></tr><tr><td>intermediate request</td><td>发往CDN Pro中间层服务器的请求流量</td></tr><tr><td>origin response</td><td>源站的响应流量</td></tr><tr><td>origin request</td><td>回源的请求流量</td></tr><tr><td>fast route response</td><td>快速回源产生的响应流量</td></tr><tr><td>fast route request</td><td>快速回源产生的请求流量</td></tr></table>
  // dataNames数组中条目的顺序与groups[].data数组返回值的顺序一一对应。"}
  DataNames []*string `json:"dataNames,omitempty" xml:"dataNames,omitempty" type:"Repeated"`
  // {"en" : "Indicates the unit of measurement of the returned values.", "zh_CN": "指示返回值的计量单位。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
}

func (s GetASummaryOfTrafficVolumeResponseMetaData) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumeResponseMetaData) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficVolumeResponseMetaData) SetStartTime(v string) *GetASummaryOfTrafficVolumeResponseMetaData {
  s.StartTime = &v
  return s
}

func (s *GetASummaryOfTrafficVolumeResponseMetaData) SetEndTime(v string) *GetASummaryOfTrafficVolumeResponseMetaData {
  s.EndTime = &v
  return s
}

func (s *GetASummaryOfTrafficVolumeResponseMetaData) SetIsComplete(v bool) *GetASummaryOfTrafficVolumeResponseMetaData {
  s.IsComplete = &v
  return s
}

func (s *GetASummaryOfTrafficVolumeResponseMetaData) SetDataNames(v []*string) *GetASummaryOfTrafficVolumeResponseMetaData {
  s.DataNames = v
  return s
}

func (s *GetASummaryOfTrafficVolumeResponseMetaData) SetDataUnit(v string) *GetASummaryOfTrafficVolumeResponseMetaData {
  s.DataUnit = &v
  return s
}

type GetASummaryOfTrafficVolumeResponseGroups struct     {
  // {"en" : "Name of the group. '__all__' is a special group encompassing all groups.
  // ", "zh_CN": "分组名称。'__all__' 是一个特殊分组，表示总流量。"}
  Group *string `json:"group,omitempty" xml:"group,omitempty"`
  // {"en" : "Data values. The units of measurement are determined by the dataUnit field.", "zh_CN": "每个分组的流量值。注意：每个分组的值，是将每个分组各自查询得到的数据按保留小数点后5位、四舍五入处理后得到的结果。'__all__' 组的值也是按这种方式得到的结果，而不是将其它分组的值直接累加得出和。如果发现'__all__' 组的值不完全等于其他组的值累加的和，属于预期内的结果。此外，当分组条件包含serverGroups时，极个别情况下，可能会出现'__all__' 组的值明显大于其它组累加的和。这是由于存在未知原因导致某些流量无法映射到serverGroup所致。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetASummaryOfTrafficVolumeResponseGroups) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfTrafficVolumeResponseGroups) GoString() string {
  return s.String()
}

func (s *GetASummaryOfTrafficVolumeResponseGroups) SetGroup(v string) *GetASummaryOfTrafficVolumeResponseGroups {
  s.Group = &v
  return s
}

func (s *GetASummaryOfTrafficVolumeResponseGroups) SetData(v []*string) *GetASummaryOfTrafficVolumeResponseGroups {
  s.Data = v
  return s
}




type GetTheEdgeUploadTrafficVolumePaths struct {
}

func (s GetTheEdgeUploadTrafficVolumePaths) String() string {
  return tea.Prettify(s)
}

func (s GetTheEdgeUploadTrafficVolumePaths) GoString() string {
  return s.String()
}

type GetTheEdgeUploadTrafficVolumeParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes,hourly,daily,monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8  means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Range: [ 0 .. 5 ] characters 
  // Enum: http,https,all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetTheEdgeUploadTrafficVolumeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTheEdgeUploadTrafficVolumeParameters) GoString() string {
  return s.String()
}

func (s *GetTheEdgeUploadTrafficVolumeParameters) SetStartdate(v string) *GetTheEdgeUploadTrafficVolumeParameters {
  s.Startdate = &v
  return s
}

func (s *GetTheEdgeUploadTrafficVolumeParameters) SetEnddate(v string) *GetTheEdgeUploadTrafficVolumeParameters {
  s.Enddate = &v
  return s
}

func (s *GetTheEdgeUploadTrafficVolumeParameters) SetType(v string) *GetTheEdgeUploadTrafficVolumeParameters {
  s.Type = &v
  return s
}

func (s *GetTheEdgeUploadTrafficVolumeParameters) SetScheme(v string) *GetTheEdgeUploadTrafficVolumeParameters {
  s.Scheme = &v
  return s
}

type GetTheEdgeUploadTrafficVolumeRequestHeader struct {
}

func (s GetTheEdgeUploadTrafficVolumeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTheEdgeUploadTrafficVolumeRequestHeader) GoString() string {
  return s.String()
}

type GetTheEdgeUploadTrafficVolumeRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetTheEdgeUploadTrafficVolumeRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetTheEdgeUploadTrafficVolumeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTheEdgeUploadTrafficVolumeRequest) GoString() string {
  return s.String()
}

func (s *GetTheEdgeUploadTrafficVolumeRequest) SetFilters(v *GetTheEdgeUploadTrafficVolumeRequestFilters) *GetTheEdgeUploadTrafficVolumeRequest {
  s.Filters = v
  return s
}

type GetTheEdgeUploadTrafficVolumeRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "items Enum: standard,premium,deluxe,ultra,nearChina,ChinaStandard,ChinaPremium 
  // Indicates one or more server groups.", "zh_CN": "取值范围: standard,premium,deluxe,ultra,nearChina,ChinaStandard,ChinaPremium 
  // 指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetTheEdgeUploadTrafficVolumeRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetTheEdgeUploadTrafficVolumeRequestFilters) GoString() string {
  return s.String()
}

func (s *GetTheEdgeUploadTrafficVolumeRequestFilters) SetHostnames(v []*string) *GetTheEdgeUploadTrafficVolumeRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetTheEdgeUploadTrafficVolumeRequestFilters) SetServerGroups(v []*string) *GetTheEdgeUploadTrafficVolumeRequestFilters {
  s.ServerGroups = v
  return s
}

type GetTheEdgeUploadTrafficVolumeResponseHeader struct {
}

func (s GetTheEdgeUploadTrafficVolumeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTheEdgeUploadTrafficVolumeResponseHeader) GoString() string {
  return s.String()
}

type GetTheEdgeUploadTrafficVolumeResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetTheEdgeUploadTrafficVolumeResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTheEdgeUploadTrafficVolumeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTheEdgeUploadTrafficVolumeResponse) GoString() string {
  return s.String()
}

func (s *GetTheEdgeUploadTrafficVolumeResponse) SetDataName(v string) *GetTheEdgeUploadTrafficVolumeResponse {
  s.DataName = &v
  return s
}

func (s *GetTheEdgeUploadTrafficVolumeResponse) SetDataUnit(v string) *GetTheEdgeUploadTrafficVolumeResponse {
  s.DataUnit = &v
  return s
}

func (s *GetTheEdgeUploadTrafficVolumeResponse) SetDataSeries(v []*GetTheEdgeUploadTrafficVolumeResponseDataSeries) *GetTheEdgeUploadTrafficVolumeResponse {
  s.DataSeries = v
  return s
}

type GetTheEdgeUploadTrafficVolumeResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetTheEdgeUploadTrafficVolumeResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTheEdgeUploadTrafficVolumeResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetTheEdgeUploadTrafficVolumeResponseDataSeries) SetTimestamp(v string) *GetTheEdgeUploadTrafficVolumeResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTheEdgeUploadTrafficVolumeResponseDataSeries) SetData(v int) *GetTheEdgeUploadTrafficVolumeResponseDataSeries {
  s.Data = &v
  return s
}




type GetASummaryOfStatusCodesReturnedByEdgeServersPaths struct {
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersPaths) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersPaths) GoString() string {
  return s.String()
}

type GetASummaryOfStatusCodesReturnedByEdgeServersParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2021-09-05T00:00:00Z ", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2021-09-05T10:00:00Z Due to latency associated with new traffic data, enddate should be no later than five minutes before the current time. This ensures you get the most accurate results.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。由于数据处理存在延迟，所指定的结束时间必须至少比当前时间早5分钟，否则返回的数据可能不准确。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: http,https,all 
  // Default: all 
  // Limit the results to the specified scheme. By default, data from HTTPS and HTTP requests is returned.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersParameters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersParameters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersParameters) SetStartdate(v string) *GetASummaryOfStatusCodesReturnedByEdgeServersParameters {
  s.Startdate = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersParameters) SetEnddate(v string) *GetASummaryOfStatusCodesReturnedByEdgeServersParameters {
  s.Enddate = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersParameters) SetScheme(v string) *GetASummaryOfStatusCodesReturnedByEdgeServersParameters {
  s.Scheme = &v
  return s
}

type GetASummaryOfStatusCodesReturnedByEdgeServersRequestHeader struct {
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersRequestHeader) GoString() string {
  return s.String()
}

type GetASummaryOfStatusCodesReturnedByEdgeServersRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
  // {"en" : "Range: <= 2 items 
  // You can group results using any combination of up to two of 'hostnames', 'serverGroups', and 'customerIds'.", "zh_CN": "取值范围: <= 2 条目 
  // 指定分组依据对数据进行分组汇总。支持按'hostnames'，'serverGroups'单独进行分组汇总，也支持同时指定这2个参数进行分组汇总。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersRequest) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersRequest) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersRequest) SetFilters(v *GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters) *GetASummaryOfStatusCodesReturnedByEdgeServersRequest {
  s.Filters = v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersRequest) SetGroupBy(v []*string) *GetASummaryOfStatusCodesReturnedByEdgeServersRequest {
  s.GroupBy = v
  return s
}

type GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Indicates one or more server groups.", "zh_CN": "指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters) SetHostnames(v []*string) *GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters) SetServerGroups(v []*string) *GetASummaryOfStatusCodesReturnedByEdgeServersRequestFilters {
  s.ServerGroups = v
  return s
}

type GetASummaryOfStatusCodesReturnedByEdgeServersResponseHeader struct {
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersResponseHeader) GoString() string {
  return s.String()
}

type GetASummaryOfStatusCodesReturnedByEdgeServersResponse struct {
  // {"en" : "This object contains fields describing the data returned in the groups object.", "zh_CN": "此对象包含的字段是对groups对象中返回数据的描述。"}
  MetaData *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" require:"true" type:"Struct"`
  // {"en" : "This object contains the breakdown of requests by group.", "zh_CN": "每个分组及其状态码。"}
  Groups []*GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersResponse) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersResponse) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponse) SetMetaData(v *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData) *GetASummaryOfStatusCodesReturnedByEdgeServersResponse {
  s.MetaData = v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponse) SetGroups(v []*GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups) *GetASummaryOfStatusCodesReturnedByEdgeServersResponse {
  s.Groups = v
  return s
}

type GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData struct {
  // {"en" : "RFC 3339 date indicating the beginning of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的起始时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en" : "RFC 3339 date indicating the end of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en" : "The response can contain up to 10000 groups. If there are more groups, isComplete will be false.", "zh_CN": "返回最多10000个分组的数据。如果实际组数超过10000，则isComplete将为false。"}
  IsComplete *bool `json:"isComplete,omitempty" xml:"isComplete,omitempty" require:"true"`
  // {"en" : "A list of HTTP status codes represented as strings. The order of the entries in the dataNames array corresponds to the order of values returned in the data array in the response.", "zh_CN": "HTTP状态码列表。dataNames数组中条目的顺序与groups[].data中返回值的顺序一一对应。"}
  DataNames []*string `json:"dataNames,omitempty" xml:"dataNames,omitempty" type:"Repeated"`
  // {"en" : "Indicates the unit of measurement of the returned values.", "zh_CN": "表示返回值的计量单位。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData) SetStartTime(v string) *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData {
  s.StartTime = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData) SetEndTime(v string) *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData {
  s.EndTime = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData) SetIsComplete(v bool) *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData {
  s.IsComplete = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData) SetDataNames(v []*string) *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData {
  s.DataNames = v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData) SetDataUnit(v string) *GetASummaryOfStatusCodesReturnedByEdgeServersResponseMetaData {
  s.DataUnit = &v
  return s
}

type GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups struct     {
  // {"en" : "Name of a group.  '__all__' is a special group encompassing all groups.", "zh_CN": "分组名称。'__all__' 是一个特殊分组，表示每个状态码的总数。"}
  Group *string `json:"group,omitempty" xml:"group,omitempty"`
  // {"en" : "Data values. The units of measurement are determined by the dataUnit field.", "zh_CN": "每个分组的状态码数量。注意：当分组条件包含serverGroups时，极个别情况下，可能会出现'__all__' 组的值明显大于其它组累加的和。这是由于存在未知原因导致某些请求无法映射到serverGroup所致。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups) SetGroup(v string) *GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups {
  s.Group = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups) SetData(v []*string) *GetASummaryOfStatusCodesReturnedByEdgeServersResponseGroups {
  s.Data = v
  return s
}




type GetOriginFastRouteTrafficVolumePaths struct {
}

func (s GetOriginFastRouteTrafficVolumePaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteTrafficVolumePaths) GoString() string {
  return s.String()
}

type GetOriginFastRouteTrafficVolumeParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8 means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes, hourly, daily, monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "[ 0 .. 5 ] 字符 
  // 取值范围: http, https, all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetOriginFastRouteTrafficVolumeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteTrafficVolumeParameters) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteTrafficVolumeParameters) SetStartdate(v string) *GetOriginFastRouteTrafficVolumeParameters {
  s.Startdate = &v
  return s
}

func (s *GetOriginFastRouteTrafficVolumeParameters) SetEnddate(v string) *GetOriginFastRouteTrafficVolumeParameters {
  s.Enddate = &v
  return s
}

func (s *GetOriginFastRouteTrafficVolumeParameters) SetType(v string) *GetOriginFastRouteTrafficVolumeParameters {
  s.Type = &v
  return s
}

func (s *GetOriginFastRouteTrafficVolumeParameters) SetScheme(v string) *GetOriginFastRouteTrafficVolumeParameters {
  s.Scheme = &v
  return s
}

type GetOriginFastRouteTrafficVolumeRequestHeader struct {
}

func (s GetOriginFastRouteTrafficVolumeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteTrafficVolumeRequestHeader) GoString() string {
  return s.String()
}

type GetOriginFastRouteTrafficVolumeRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetOriginFastRouteTrafficVolumeRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetOriginFastRouteTrafficVolumeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteTrafficVolumeRequest) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteTrafficVolumeRequest) SetFilters(v *GetOriginFastRouteTrafficVolumeRequestFilters) *GetOriginFastRouteTrafficVolumeRequest {
  s.Filters = v
  return s
}

type GetOriginFastRouteTrafficVolumeRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Enum: standard premium deluxe ultra 
  // Indicates one or more server groups.", "zh_CN": "取值范围: standard, premium, deluxe, ultra 
  // 指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetOriginFastRouteTrafficVolumeRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteTrafficVolumeRequestFilters) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteTrafficVolumeRequestFilters) SetHostnames(v []*string) *GetOriginFastRouteTrafficVolumeRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetOriginFastRouteTrafficVolumeRequestFilters) SetServerGroups(v []*string) *GetOriginFastRouteTrafficVolumeRequestFilters {
  s.ServerGroups = v
  return s
}

type GetOriginFastRouteTrafficVolumeResponseHeader struct {
}

func (s GetOriginFastRouteTrafficVolumeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteTrafficVolumeResponseHeader) GoString() string {
  return s.String()
}

type GetOriginFastRouteTrafficVolumeResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Default: megabytes 
  // Unit of measurement.", "zh_CN": "默认值: megabytes 
  // 计量单位。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points for the time period include the upload and download traffic.", "zh_CN": "指定时间段的流量，包括上行和下行流量。"}
  DataSeries []*GetOriginFastRouteTrafficVolumeResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginFastRouteTrafficVolumeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteTrafficVolumeResponse) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteTrafficVolumeResponse) SetDataName(v string) *GetOriginFastRouteTrafficVolumeResponse {
  s.DataName = &v
  return s
}

func (s *GetOriginFastRouteTrafficVolumeResponse) SetDataUnit(v string) *GetOriginFastRouteTrafficVolumeResponse {
  s.DataUnit = &v
  return s
}

func (s *GetOriginFastRouteTrafficVolumeResponse) SetDataSeries(v []*GetOriginFastRouteTrafficVolumeResponseDataSeries) *GetOriginFastRouteTrafficVolumeResponse {
  s.DataSeries = v
  return s
}

type GetOriginFastRouteTrafficVolumeResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of  a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "RFC 3339格式的日期，表示每个时间段的开始时间，始终采用<b>UTC</b>时间。例如：'timestamp':'2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "Default: 0 >= 0 
  // Indicates the amount of traffic in megabytes uploaded to the origin using the fast route feature.", "zh_CN": "默认值: 0 取值范围 >= 0 
  // 表示通过快速回源传输的上行流量（以MB为单位）。"}
  DataUp *int `json:"dataUp,omitempty" xml:"dataUp,omitempty"`
  // {"en" : "Default: 0 >= 0 
  // Indicates the amount of traffic in megabytes downloaded from the origin using the fast route feature.", "zh_CN": "默认值: 0 取值范围 >= 0 
  // 表示通过快速回源传输的下行流量（以MB为单位）。"}
  DataDown *int `json:"dataDown,omitempty" xml:"dataDown,omitempty"`
}

func (s GetOriginFastRouteTrafficVolumeResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetOriginFastRouteTrafficVolumeResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetOriginFastRouteTrafficVolumeResponseDataSeries) SetTimestamp(v string) *GetOriginFastRouteTrafficVolumeResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetOriginFastRouteTrafficVolumeResponseDataSeries) SetDataUp(v int) *GetOriginFastRouteTrafficVolumeResponseDataSeries {
  s.DataUp = &v
  return s
}

func (s *GetOriginFastRouteTrafficVolumeResponseDataSeries) SetDataDown(v int) *GetOriginFastRouteTrafficVolumeResponseDataSeries {
  s.DataDown = &v
  return s
}




type GetEdgeRequestsPaths struct {
}

func (s GetEdgeRequestsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeRequestsPaths) GoString() string {
  return s.String()
}

type GetEdgeRequestsParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter endtdate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes hourly daily monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8 means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes, hourly, daily, monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "[ 0 .. 5 ] characters 
  // Enum: http https all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "[ 0 .. 5 ] 字符 
  // 取值范围: http, https, all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetEdgeRequestsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeRequestsParameters) GoString() string {
  return s.String()
}

func (s *GetEdgeRequestsParameters) SetStartdate(v string) *GetEdgeRequestsParameters {
  s.Startdate = &v
  return s
}

func (s *GetEdgeRequestsParameters) SetEnddate(v string) *GetEdgeRequestsParameters {
  s.Enddate = &v
  return s
}

func (s *GetEdgeRequestsParameters) SetType(v string) *GetEdgeRequestsParameters {
  s.Type = &v
  return s
}

func (s *GetEdgeRequestsParameters) SetScheme(v string) *GetEdgeRequestsParameters {
  s.Scheme = &v
  return s
}

type GetEdgeRequestsRequestHeader struct {
}

func (s GetEdgeRequestsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeRequestsRequestHeader) GoString() string {
  return s.String()
}

type GetEdgeRequestsRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetEdgeRequestsRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetEdgeRequestsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeRequestsRequest) GoString() string {
  return s.String()
}

func (s *GetEdgeRequestsRequest) SetFilters(v *GetEdgeRequestsRequestFilters) *GetEdgeRequestsRequest {
  s.Filters = v
  return s
}

type GetEdgeRequestsRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "items Enum: standard premium deluxe ultra nearChina ChinaStandard ChinaPremium 
  // Indicates one or more server groups.", "zh_CN": "取值范围: standard, premium, deluxe, ultra, nearChina, ChinaStandard, ChinaPremium 
  // 指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetEdgeRequestsRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeRequestsRequestFilters) GoString() string {
  return s.String()
}

func (s *GetEdgeRequestsRequestFilters) SetHostnames(v []*string) *GetEdgeRequestsRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetEdgeRequestsRequestFilters) SetServerGroups(v []*string) *GetEdgeRequestsRequestFilters {
  s.ServerGroups = v
  return s
}

type GetEdgeRequestsResponseHeader struct {
}

func (s GetEdgeRequestsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeRequestsResponseHeader) GoString() string {
  return s.String()
}

type GetEdgeRequestsResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetEdgeRequestsResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetEdgeRequestsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeRequestsResponse) GoString() string {
  return s.String()
}

func (s *GetEdgeRequestsResponse) SetDataName(v string) *GetEdgeRequestsResponse {
  s.DataName = &v
  return s
}

func (s *GetEdgeRequestsResponse) SetDataUnit(v string) *GetEdgeRequestsResponse {
  s.DataUnit = &v
  return s
}

func (s *GetEdgeRequestsResponse) SetDataSeries(v []*GetEdgeRequestsResponseDataSeries) *GetEdgeRequestsResponse {
  s.DataSeries = v
  return s
}

type GetEdgeRequestsResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetEdgeRequestsResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeRequestsResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetEdgeRequestsResponseDataSeries) SetTimestamp(v string) *GetEdgeRequestsResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetEdgeRequestsResponseDataSeries) SetData(v int) *GetEdgeRequestsResponseDataSeries {
  s.Data = &v
  return s
}




type GetASummaryOfStatusCodesReturnedByOriginServersPaths struct {
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersPaths) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersPaths) GoString() string {
  return s.String()
}

type GetASummaryOfStatusCodesReturnedByOriginServersParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2021-09-05T00:00:00Z ", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2021-09-05T10:00:00Z Due to latency associated with new traffic data, enddate should be no later than five minutes before the current time. This ensures you get the most accurate results.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。由于数据处理存在延迟，所指定的结束时间必须至少比当前时间早5分钟，否则返回的数据可能不准确。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: http,https,all 
  // Default: all 
  // Limit the results to the specified scheme. By default, data from HTTPS and HTTP requests is returned.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersParameters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersParameters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersParameters) SetStartdate(v string) *GetASummaryOfStatusCodesReturnedByOriginServersParameters {
  s.Startdate = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersParameters) SetEnddate(v string) *GetASummaryOfStatusCodesReturnedByOriginServersParameters {
  s.Enddate = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersParameters) SetScheme(v string) *GetASummaryOfStatusCodesReturnedByOriginServersParameters {
  s.Scheme = &v
  return s
}

type GetASummaryOfStatusCodesReturnedByOriginServersRequestHeader struct {
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersRequestHeader) GoString() string {
  return s.String()
}

type GetASummaryOfStatusCodesReturnedByOriginServersRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
  // {"en" : "Range: <= 2 items 
  // You can group results using a combination of up to two of the following: 'hostnames', 'serverGroups', and 'customerIds'.", "zh_CN": "取值范围: <= 2 条目 
  // 指定分组依据对数据进行分组汇总。支持按'hostnames'，'serverGroups'单独进行分组汇总，也支持同时指定这2个参数进行分组汇总。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersRequest) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersRequest) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersRequest) SetFilters(v *GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters) *GetASummaryOfStatusCodesReturnedByOriginServersRequest {
  s.Filters = v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersRequest) SetGroupBy(v []*string) *GetASummaryOfStatusCodesReturnedByOriginServersRequest {
  s.GroupBy = v
  return s
}

type GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Indicates one or more server groups.", "zh_CN": "指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters) SetHostnames(v []*string) *GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters) SetServerGroups(v []*string) *GetASummaryOfStatusCodesReturnedByOriginServersRequestFilters {
  s.ServerGroups = v
  return s
}

type GetASummaryOfStatusCodesReturnedByOriginServersResponseHeader struct {
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersResponseHeader) GoString() string {
  return s.String()
}

type GetASummaryOfStatusCodesReturnedByOriginServersResponse struct {
  // {"en" : "This object contains fields describing the data returned in the groups object.", "zh_CN": "此对象包含的字段是对groups对象中返回数据的描述。"}
  MetaData *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" require:"true" type:"Struct"`
  // {"en" : "This object contains the breakdown of requests by group.", "zh_CN": "分组汇总数据。"}
  Groups []*GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersResponse) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersResponse) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponse) SetMetaData(v *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData) *GetASummaryOfStatusCodesReturnedByOriginServersResponse {
  s.MetaData = v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponse) SetGroups(v []*GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups) *GetASummaryOfStatusCodesReturnedByOriginServersResponse {
  s.Groups = v
  return s
}

type GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData struct {
  // {"en" : "RFC 3339 date indicating the beginning of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的起始时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
  // {"en" : "RFC 3339 date indicating the end of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // {"en" : "The response can contain up to 10000 groups. If there are more groups, isComplete will be false.", "zh_CN": "返回最多10000个分组的数据。如果实际组数超过10000，则isComplete将为false。"}
  IsComplete *bool `json:"isComplete,omitempty" xml:"isComplete,omitempty" require:"true"`
  // {"en" : "A list of HTTP status codes represented as strings. The order of the entries in the dataNames array corresponds to the order of values returned in the data array in the response.", "zh_CN": "HTTP状态码列表。dataNames数组中条目的顺序与groups[].data中返回值的顺序一一对应。"}
  DataNames []*string `json:"dataNames,omitempty" xml:"dataNames,omitempty" type:"Repeated"`
  // {"en" : "Indicates the unit of measurement of the returned values.", "zh_CN": "表示返回值的计量单位。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData) SetStartTime(v string) *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData {
  s.StartTime = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData) SetEndTime(v string) *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData {
  s.EndTime = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData) SetIsComplete(v bool) *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData {
  s.IsComplete = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData) SetDataNames(v []*string) *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData {
  s.DataNames = v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData) SetDataUnit(v string) *GetASummaryOfStatusCodesReturnedByOriginServersResponseMetaData {
  s.DataUnit = &v
  return s
}

type GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups struct     {
  // {"en" : "Name of a group.  '__all__' is a special group encompassing all groups.", "zh_CN": "分组名称。'__all__' 是一个特殊分组，包含其它所有分组的数据。"}
  Group *string `json:"group,omitempty" xml:"group,omitempty"`
  // {"en" : "Data values. The units of measurement are determined by the dataUnit field.", "zh_CN": "状态码数量。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups) GoString() string {
  return s.String()
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups) SetGroup(v string) *GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups {
  s.Group = &v
  return s
}

func (s *GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups) SetData(v []*string) *GetASummaryOfStatusCodesReturnedByOriginServersResponseGroups {
  s.Data = v
  return s
}




type GetOriginTrafficVolumePaths struct {
}

func (s GetOriginTrafficVolumePaths) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficVolumePaths) GoString() string {
  return s.String()
}

type GetOriginTrafficVolumeParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes,hourly,daily,monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14. For example, type=daily+8 means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Range: [ 0 .. 5 ] characters 
  // Enum: http,https,all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetOriginTrafficVolumeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficVolumeParameters) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficVolumeParameters) SetStartdate(v string) *GetOriginTrafficVolumeParameters {
  s.Startdate = &v
  return s
}

func (s *GetOriginTrafficVolumeParameters) SetEnddate(v string) *GetOriginTrafficVolumeParameters {
  s.Enddate = &v
  return s
}

func (s *GetOriginTrafficVolumeParameters) SetType(v string) *GetOriginTrafficVolumeParameters {
  s.Type = &v
  return s
}

func (s *GetOriginTrafficVolumeParameters) SetScheme(v string) *GetOriginTrafficVolumeParameters {
  s.Scheme = &v
  return s
}

type GetOriginTrafficVolumeRequestHeader struct {
}

func (s GetOriginTrafficVolumeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficVolumeRequestHeader) GoString() string {
  return s.String()
}

type GetOriginTrafficVolumeRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetOriginTrafficVolumeRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetOriginTrafficVolumeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficVolumeRequest) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficVolumeRequest) SetFilters(v *GetOriginTrafficVolumeRequestFilters) *GetOriginTrafficVolumeRequest {
  s.Filters = v
  return s
}

type GetOriginTrafficVolumeRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "items Enum: standard,premium,deluxe,ultra,nearChina,ChinaStandard,ChinaPremium 
  // Indicates one or more server groups.", "zh_CN": "取值范围: standard,premium,deluxe,ultra,nearChina,ChinaStandard,ChinaPremium 
  // 指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetOriginTrafficVolumeRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficVolumeRequestFilters) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficVolumeRequestFilters) SetHostnames(v []*string) *GetOriginTrafficVolumeRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetOriginTrafficVolumeRequestFilters) SetServerGroups(v []*string) *GetOriginTrafficVolumeRequestFilters {
  s.ServerGroups = v
  return s
}

type GetOriginTrafficVolumeResponseHeader struct {
}

func (s GetOriginTrafficVolumeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficVolumeResponseHeader) GoString() string {
  return s.String()
}

type GetOriginTrafficVolumeResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetOriginTrafficVolumeResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetOriginTrafficVolumeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficVolumeResponse) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficVolumeResponse) SetDataName(v string) *GetOriginTrafficVolumeResponse {
  s.DataName = &v
  return s
}

func (s *GetOriginTrafficVolumeResponse) SetDataUnit(v string) *GetOriginTrafficVolumeResponse {
  s.DataUnit = &v
  return s
}

func (s *GetOriginTrafficVolumeResponse) SetDataSeries(v []*GetOriginTrafficVolumeResponseDataSeries) *GetOriginTrafficVolumeResponse {
  s.DataSeries = v
  return s
}

type GetOriginTrafficVolumeResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetOriginTrafficVolumeResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetOriginTrafficVolumeResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetOriginTrafficVolumeResponseDataSeries) SetTimestamp(v string) *GetOriginTrafficVolumeResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetOriginTrafficVolumeResponseDataSeries) SetData(v int) *GetOriginTrafficVolumeResponseDataSeries {
  s.Data = &v
  return s
}




type GetTheIntermediateTrafficVolumePaths struct {
}

func (s GetTheIntermediateTrafficVolumePaths) String() string {
  return tea.Prettify(s)
}

func (s GetTheIntermediateTrafficVolumePaths) GoString() string {
  return s.String()
}

type GetTheIntermediateTrafficVolumeParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes,hourly,daily,monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14.  For example, type=daily+8  means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes,hourly,daily,monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "Range: [ 0 .. 5 ] characters 
  // Enum: http,https,all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetTheIntermediateTrafficVolumeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetTheIntermediateTrafficVolumeParameters) GoString() string {
  return s.String()
}

func (s *GetTheIntermediateTrafficVolumeParameters) SetStartdate(v string) *GetTheIntermediateTrafficVolumeParameters {
  s.Startdate = &v
  return s
}

func (s *GetTheIntermediateTrafficVolumeParameters) SetEnddate(v string) *GetTheIntermediateTrafficVolumeParameters {
  s.Enddate = &v
  return s
}

func (s *GetTheIntermediateTrafficVolumeParameters) SetType(v string) *GetTheIntermediateTrafficVolumeParameters {
  s.Type = &v
  return s
}

func (s *GetTheIntermediateTrafficVolumeParameters) SetScheme(v string) *GetTheIntermediateTrafficVolumeParameters {
  s.Scheme = &v
  return s
}

type GetTheIntermediateTrafficVolumeRequestHeader struct {
}

func (s GetTheIntermediateTrafficVolumeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTheIntermediateTrafficVolumeRequestHeader) GoString() string {
  return s.String()
}

type GetTheIntermediateTrafficVolumeRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetTheIntermediateTrafficVolumeRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetTheIntermediateTrafficVolumeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetTheIntermediateTrafficVolumeRequest) GoString() string {
  return s.String()
}

func (s *GetTheIntermediateTrafficVolumeRequest) SetFilters(v *GetTheIntermediateTrafficVolumeRequestFilters) *GetTheIntermediateTrafficVolumeRequest {
  s.Filters = v
  return s
}

type GetTheIntermediateTrafficVolumeRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "items Enum: standard,premium,deluxe,ultra,nearChina,ChinaStandard,ChinaPremium 
  // Indicates one or more server groups.", "zh_CN": "取值范围: standard,premium,deluxe,ultra,nearChina,ChinaStandard,ChinaPremium 
  // 指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetTheIntermediateTrafficVolumeRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetTheIntermediateTrafficVolumeRequestFilters) GoString() string {
  return s.String()
}

func (s *GetTheIntermediateTrafficVolumeRequestFilters) SetHostnames(v []*string) *GetTheIntermediateTrafficVolumeRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetTheIntermediateTrafficVolumeRequestFilters) SetServerGroups(v []*string) *GetTheIntermediateTrafficVolumeRequestFilters {
  s.ServerGroups = v
  return s
}

type GetTheIntermediateTrafficVolumeResponseHeader struct {
}

func (s GetTheIntermediateTrafficVolumeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetTheIntermediateTrafficVolumeResponseHeader) GoString() string {
  return s.String()
}

type GetTheIntermediateTrafficVolumeResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetTheIntermediateTrafficVolumeResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetTheIntermediateTrafficVolumeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetTheIntermediateTrafficVolumeResponse) GoString() string {
  return s.String()
}

func (s *GetTheIntermediateTrafficVolumeResponse) SetDataName(v string) *GetTheIntermediateTrafficVolumeResponse {
  s.DataName = &v
  return s
}

func (s *GetTheIntermediateTrafficVolumeResponse) SetDataUnit(v string) *GetTheIntermediateTrafficVolumeResponse {
  s.DataUnit = &v
  return s
}

func (s *GetTheIntermediateTrafficVolumeResponse) SetDataSeries(v []*GetTheIntermediateTrafficVolumeResponseDataSeries) *GetTheIntermediateTrafficVolumeResponse {
  s.DataSeries = v
  return s
}

type GetTheIntermediateTrafficVolumeResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetTheIntermediateTrafficVolumeResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetTheIntermediateTrafficVolumeResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetTheIntermediateTrafficVolumeResponseDataSeries) SetTimestamp(v string) *GetTheIntermediateTrafficVolumeResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetTheIntermediateTrafficVolumeResponseDataSeries) SetData(v int) *GetTheIntermediateTrafficVolumeResponseDataSeries {
  s.Data = &v
  return s
}




type GetEdgeTrafficVolumePaths struct {
}

func (s GetEdgeTrafficVolumePaths) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeTrafficVolumePaths) GoString() string {
  return s.String()
}

type GetEdgeTrafficVolumeParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-10-30T00:00:00Z Your startdate may be rounded down to the nearest minute, hour, or day depending on the type parameter. For example, if you enter startdate=2019-09-05T03:14:01Z&type=hourly, the response includes data beginning 2019-09-05T03:00:00Z.", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。 根据type参数对应的粒度，您指定的开始时间可能被向前取整为最近的分钟、小时或天。例如，如果您指定 startdate=2019-09-05T03:14:01Z&type=hourly，则返回从2019-09-05T03:00:00Z开始的数据。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2019-11-14T00:00:00Z Your enddate may be rounded up to the nearest minute, hour, or day depending on the type parameter. For example, if you enter enddate=2019-09-05T03:14:01Z&type=hourly, the response includes data ending 2019-09-05T04:00:00Z.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。根据type参数对应的粒度，您指定的结束时间可能被向后取整为最近的分钟、小时或天。例如，如果您指定 enddate=2019-09-05T03:14:01Z&type=hourly，则返回截止到2019-09-05T04:00:00Z的数据。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: fiveminutes hourly daily monthly 
  // Indicates the granularity of returned data. By default, 00:00:00 in UTC is used as the beginning of a day. If you wish to use a different offset from UTC, you can append -12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, +1, +2, +3, +4, +5, +6, +7, +8, +9, +10, +11, +12, +13, +14. For example, type=daily+8 means the day as defined in UTC+8 time.", "zh_CN": "取值范围: fiveminutes, hourly, daily, monthly 
  // 指定返回数据的粒度，支持5分钟，小时，日，月粒度。默认情况下，我们以UTC 00:00:00作为一天的开始。如果您希望指定不同的时区，可以附加时区标识，即-12、-11、-10、-9、-8、-7、-6、-5、-4、-3、-2、-1、+1、+2、+3、+4、+5、+6、+7、+8、+9、+10、+11、+12。例如，type=daily+8表示使用UTC+8时区指定一天的开始时间。"}
  Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
  // {"en" : "[ 0 .. 5 ] characters 
  // Enum: http https all 
  // Default: all 
  // Choose whether to include HTTP and HTTPS traffic in the results. The default 'all' includes both types of traffic.", "zh_CN": "[ 0 .. 5 ] 字符 
  // 取值范围: http, https, all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetEdgeTrafficVolumeParameters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeTrafficVolumeParameters) GoString() string {
  return s.String()
}

func (s *GetEdgeTrafficVolumeParameters) SetStartdate(v string) *GetEdgeTrafficVolumeParameters {
  s.Startdate = &v
  return s
}

func (s *GetEdgeTrafficVolumeParameters) SetEnddate(v string) *GetEdgeTrafficVolumeParameters {
  s.Enddate = &v
  return s
}

func (s *GetEdgeTrafficVolumeParameters) SetType(v string) *GetEdgeTrafficVolumeParameters {
  s.Type = &v
  return s
}

func (s *GetEdgeTrafficVolumeParameters) SetScheme(v string) *GetEdgeTrafficVolumeParameters {
  s.Scheme = &v
  return s
}

type GetEdgeTrafficVolumeRequestHeader struct {
}

func (s GetEdgeTrafficVolumeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeTrafficVolumeRequestHeader) GoString() string {
  return s.String()
}

type GetEdgeTrafficVolumeRequest struct {
  // {"en" : "Specify conditions to filter report data.", "zh_CN": "指定查询条件过滤报表数据。"}
  Filters *GetEdgeTrafficVolumeRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
}

func (s GetEdgeTrafficVolumeRequest) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeTrafficVolumeRequest) GoString() string {
  return s.String()
}

func (s *GetEdgeTrafficVolumeRequest) SetFilters(v *GetEdgeTrafficVolumeRequestFilters) *GetEdgeTrafficVolumeRequest {
  s.Filters = v
  return s
}

type GetEdgeTrafficVolumeRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "items Enum: standard premium deluxe ultra nearChina ChinaStandard ChinaPremium 
  // Indicates one or more server groups.", "zh_CN": "取值范围: standard, premium, deluxe, ultra, nearChina, ChinaStandard, ChinaPremium 
  // 指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetEdgeTrafficVolumeRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeTrafficVolumeRequestFilters) GoString() string {
  return s.String()
}

func (s *GetEdgeTrafficVolumeRequestFilters) SetHostnames(v []*string) *GetEdgeTrafficVolumeRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetEdgeTrafficVolumeRequestFilters) SetServerGroups(v []*string) *GetEdgeTrafficVolumeRequestFilters {
  s.ServerGroups = v
  return s
}

type GetEdgeTrafficVolumeResponseHeader struct {
}

func (s GetEdgeTrafficVolumeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeTrafficVolumeResponseHeader) GoString() string {
  return s.String()
}

type GetEdgeTrafficVolumeResponse struct {
  // {"en" : "A brief description of the data.", "zh_CN": "对返回数据的简要描述。"}
  DataName *string `json:"dataName,omitempty" xml:"dataName,omitempty" require:"true"`
  // {"en" : "Unit of measurement. This will depend on the report API.", "zh_CN": "计量单位。不同报表类型计量单位不一样。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
  // {"en" : "The data points.", "zh_CN": "数据点。"}
  DataSeries []*GetEdgeTrafficVolumeResponseDataSeries `json:"dataSeries,omitempty" xml:"dataSeries,omitempty" require:"true" type:"Repeated"`
}

func (s GetEdgeTrafficVolumeResponse) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeTrafficVolumeResponse) GoString() string {
  return s.String()
}

func (s *GetEdgeTrafficVolumeResponse) SetDataName(v string) *GetEdgeTrafficVolumeResponse {
  s.DataName = &v
  return s
}

func (s *GetEdgeTrafficVolumeResponse) SetDataUnit(v string) *GetEdgeTrafficVolumeResponse {
  s.DataUnit = &v
  return s
}

func (s *GetEdgeTrafficVolumeResponse) SetDataSeries(v []*GetEdgeTrafficVolumeResponseDataSeries) *GetEdgeTrafficVolumeResponse {
  s.DataSeries = v
  return s
}

type GetEdgeTrafficVolumeResponseDataSeries struct     {
  // {"en" : "An RFC 3339 format date representing the beginning of a time interval. It is always in <b>UTC</b> time. For example:  'timestamp': '2019-10-29T01:00:00Z'", "zh_CN": "每个时间段的起始时间，以RFC 3339日期格式表示。始终采用UTC时区。例如：'timestamp': '2019-10-29T01:00:00Z'"}
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  // {"en" : "A value at that timestamp. Refer to the dataUnit field for the unit of measurement.", "zh_CN": "该时间段对应的值。计量单位，由dataUnit字段指定。"}
  Data *int `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetEdgeTrafficVolumeResponseDataSeries) String() string {
  return tea.Prettify(s)
}

func (s GetEdgeTrafficVolumeResponseDataSeries) GoString() string {
  return s.String()
}

func (s *GetEdgeTrafficVolumeResponseDataSeries) SetTimestamp(v string) *GetEdgeTrafficVolumeResponseDataSeries {
  s.Timestamp = &v
  return s
}

func (s *GetEdgeTrafficVolumeResponseDataSeries) SetData(v int) *GetEdgeTrafficVolumeResponseDataSeries {
  s.Data = &v
  return s
}




type GetASummaryOfCpuUsagePaths struct {
}

func (s GetASummaryOfCpuUsagePaths) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsagePaths) GoString() string {
  return s.String()
}

type GetASummaryOfCpuUsageParameters struct {
  // {"en" : "RFC 3339 date indicating the beginning of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: startdate=2019-11-14T00:00:00Z", "zh_CN": "查询范围的开始时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：startdate=2019-10-30T00:00:00Z。"}
  Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the time period. The time must be specified using the UTC timezone; it cannot be an offset. Example: enddate=2021-09-05T10:00:00Z Due to latency associated with new traffic data, enddate should be no later than five minutes before the current time. This ensures you get the most accurate results.", "zh_CN": "查询范围的结束时间，以RFC 3339日期格式表示。必须使用UTC时区指定时间。示例：enddate=2019-11-14T00:00:00Z。由于数据处理存在延迟，所指定的结束时间必须至少比当前时间早5分钟，否则返回的数据可能不准确。"}
  Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty" require:"true"`
  // {"en" : "Enum: all,http,https 
  // Default: all 
  // Limit the results to the specified scheme. By default, data from HTTPS and HTTP requests is returned.", "zh_CN": "取值范围: [ 0 .. 5 ] 字符 
  // 取值范围: http,https,all 
  // 默认值: all 
  // 指定查询HTTP与/或HTTPS协议的数据。默认查询全部2种协议的数据。"}
  Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
}

func (s GetASummaryOfCpuUsageParameters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsageParameters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfCpuUsageParameters) SetStartdate(v string) *GetASummaryOfCpuUsageParameters {
  s.Startdate = &v
  return s
}

func (s *GetASummaryOfCpuUsageParameters) SetEnddate(v string) *GetASummaryOfCpuUsageParameters {
  s.Enddate = &v
  return s
}

func (s *GetASummaryOfCpuUsageParameters) SetScheme(v string) *GetASummaryOfCpuUsageParameters {
  s.Scheme = &v
  return s
}

type GetASummaryOfCpuUsageRequestHeader struct {
}

func (s GetASummaryOfCpuUsageRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsageRequestHeader) GoString() string {
  return s.String()
}

type GetASummaryOfCpuUsageRequest struct {
  // {"en" : "", "zh_CN": ""}
  Filters *GetASummaryOfCpuUsageRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Struct"`
  // {"en" : "Range: <= 2 items 
  // You can group results using a combination of up to two of the following: 'hostnames', 'serverGroups'", "zh_CN": "取值范围: <= 2 条目 
  // 指定分组依据对数据进行分组汇总。支持按'hostnames'，'serverGroups'单独进行分组汇总，也支持同时指定这2个参数进行分组汇总。"}
  GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
}

func (s GetASummaryOfCpuUsageRequest) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsageRequest) GoString() string {
  return s.String()
}

func (s *GetASummaryOfCpuUsageRequest) SetFilters(v *GetASummaryOfCpuUsageRequestFilters) *GetASummaryOfCpuUsageRequest {
  s.Filters = v
  return s
}

func (s *GetASummaryOfCpuUsageRequest) SetGroupBy(v []*string) *GetASummaryOfCpuUsageRequest {
  s.GroupBy = v
  return s
}

type GetASummaryOfCpuUsageRequestFilters struct {
  // {"en" : "List of hostnames for which to return data. Wildcard hostnames such as *.domain.com are also permitted. If unspecified, data from all hostnames will be returned.", "zh_CN": "指定加速域名进行查询。可使用泛域名，如*.domain.com。如果未指定，将返回所有加速域名的数据。"}
  Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
  // {"en" : "Indicates one or more server groups.", "zh_CN": "指定serverGroups（节点组）进行查询。"}
  ServerGroups []*string `json:"serverGroups,omitempty" xml:"serverGroups,omitempty" type:"Repeated"`
}

func (s GetASummaryOfCpuUsageRequestFilters) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsageRequestFilters) GoString() string {
  return s.String()
}

func (s *GetASummaryOfCpuUsageRequestFilters) SetHostnames(v []*string) *GetASummaryOfCpuUsageRequestFilters {
  s.Hostnames = v
  return s
}

func (s *GetASummaryOfCpuUsageRequestFilters) SetServerGroups(v []*string) *GetASummaryOfCpuUsageRequestFilters {
  s.ServerGroups = v
  return s
}

type GetASummaryOfCpuUsageResponseHeader struct {
}

func (s GetASummaryOfCpuUsageResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsageResponseHeader) GoString() string {
  return s.String()
}

type GetASummaryOfCpuUsageResponse struct {
  // {"en" : "This object contains fields describing the data returned in the groups object.", "zh_CN": "此对象包含的字段是对groups对象中返回数据的描述。"}
  MetaData *GetASummaryOfCpuUsageResponseMetaData `json:"metaData,omitempty" xml:"metaData,omitempty" require:"true" type:"Struct"`
  // {"en" : "This object contains the breakdown of CPU usage by group. ", "zh_CN": "每个分组及其cpu时间。"}
  Groups []*GetASummaryOfCpuUsageResponseGroups `json:"groups,omitempty" xml:"groups,omitempty" require:"true" type:"Repeated"`
}

func (s GetASummaryOfCpuUsageResponse) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsageResponse) GoString() string {
  return s.String()
}

func (s *GetASummaryOfCpuUsageResponse) SetMetaData(v *GetASummaryOfCpuUsageResponseMetaData) *GetASummaryOfCpuUsageResponse {
  s.MetaData = v
  return s
}

func (s *GetASummaryOfCpuUsageResponse) SetGroups(v []*GetASummaryOfCpuUsageResponseGroups) *GetASummaryOfCpuUsageResponse {
  s.Groups = v
  return s
}

type GetASummaryOfCpuUsageResponseMetaData struct {
  // {"en" : "RFC 3339 date indicating the beginning of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的起始时间。"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en" : "RFC 3339 date indicating the end of the period.", "zh_CN": "RFC 3339格式的日期，表示查询的结束时间。"}
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty" require:"true"`
  // {"en" : "The response can contain up to 10000 groups. If there are more groups, isComplete will be false.", "zh_CN": "该接口最多返回10000个分组的数据。如果实际分组数量大于10000，则isComplete将为false。"}
  IsComplete *bool `json:"isComplete,omitempty" xml:"isComplete,omitempty" require:"true"`
  // {"en" : "Indicates the type of data returned. 
  // <table><tr><th>Data Name</th><th>Description</th></tr><tr><td>edge cpu time</td><td>CPU usage by edge servers</td></tr><tr><td>intermediate cpu time</td><td>CPU usage by intermediate CDN Pro servers</td></tr></table>
  // The order of the entries in dataNames array corresponds to the order of values returned in the data data array in the response.", "zh_CN": "表示返回的数据类型。
  // <table><tr><th>数据名称</th><th>描述</th></tr><tr><td>edge cpu time</td><td>边缘服务器的CPU使用时间</td></tr><tr><td>intermediate cpu time</td><td>中间层服务器的CPU使用时间</td></tr></table>
  // dataNames数组中条目的顺序与groups[].data中返回值的顺序一一对应。"}
  DataNames []*string `json:"dataNames,omitempty" xml:"dataNames,omitempty" require:"true" type:"Repeated"`
  // {"en" : "Indicates the unit of measurement of the returned values.", "zh_CN": "表示返回值的计量单位。"}
  DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty" require:"true"`
}

func (s GetASummaryOfCpuUsageResponseMetaData) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsageResponseMetaData) GoString() string {
  return s.String()
}

func (s *GetASummaryOfCpuUsageResponseMetaData) SetStartTime(v string) *GetASummaryOfCpuUsageResponseMetaData {
  s.StartTime = &v
  return s
}

func (s *GetASummaryOfCpuUsageResponseMetaData) SetEndTime(v string) *GetASummaryOfCpuUsageResponseMetaData {
  s.EndTime = &v
  return s
}

func (s *GetASummaryOfCpuUsageResponseMetaData) SetIsComplete(v bool) *GetASummaryOfCpuUsageResponseMetaData {
  s.IsComplete = &v
  return s
}

func (s *GetASummaryOfCpuUsageResponseMetaData) SetDataNames(v []*string) *GetASummaryOfCpuUsageResponseMetaData {
  s.DataNames = v
  return s
}

func (s *GetASummaryOfCpuUsageResponseMetaData) SetDataUnit(v string) *GetASummaryOfCpuUsageResponseMetaData {
  s.DataUnit = &v
  return s
}

type GetASummaryOfCpuUsageResponseGroups struct     {
  // {"en" : "Name of the group. '__all__' is a special group encompassing all groups.
  // ", "zh_CN": "分组名称。'__all__' 是一个特殊分组，表示总cpu时间。"}
  Group *string `json:"group,omitempty" xml:"group,omitempty"`
  // {"en" : "Data values. The units of measurement are determined by the dataUnit field.", "zh_CN": "每个分组的cpu时间。注意：每个分组的值，是将每个分组各自查询得到的数据按保留小数点后5位、四舍五入处理后得到的结果。'__all__' 组的值也是按这种方式得到的结果，而不是将其它分组的值直接累加得出和。如果发现'__all__' 组的值不完全等于其他组的值累加的和，属于预期内的结果。此外，当分组条件包含serverGroups时，极个别情况下，可能会出现'__all__' 组的值明显大于其它组累加的和。这是由于存在未知原因导致某些请求无法映射到serverGroup所致。"}
  Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetASummaryOfCpuUsageResponseGroups) String() string {
  return tea.Prettify(s)
}

func (s GetASummaryOfCpuUsageResponseGroups) GoString() string {
  return s.String()
}

func (s *GetASummaryOfCpuUsageResponseGroups) SetGroup(v string) *GetASummaryOfCpuUsageResponseGroups {
  s.Group = &v
  return s
}

func (s *GetASummaryOfCpuUsageResponseGroups) SetData(v []*string) *GetASummaryOfCpuUsageResponseGroups {
  s.Data = v
  return s
}




