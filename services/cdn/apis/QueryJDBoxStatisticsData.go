// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryJDBoxStatisticsDataRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,时间戳 (Optional) */
    StartTime *int64 `json:"startTime"`

    /* 查询截止时间,时间戳 (Optional) */
    EndTime *int64 `json:"endTime"`

    /* 查询的字段，取值范围(avgbandwidth,pv,flow)。多个用逗号分隔。默认为空，表示查询带宽流量 (Optional) */
    Fields *string `json:"fields"`

    /*  (Optional) */
    Area *string `json:"area"`

    /*  (Optional) */
    Isp *string `json:"isp"`

    /* 查询周期，当前取值范围：“oneMin,fiveMin”，分别表示1min，5min。默认为空，表示fiveMin (Optional) */
    Period *string `json:"period"`

    /* 业务类型 (Optional) */
    Category *string `json:"category"`

    /* 设备id (Optional) */
    MacAddr *string `json:"macAddr"`

    /* 插件pin (Optional) */
    PluginPin *string `json:"pluginPin"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryJDBoxStatisticsDataRequest(
) *QueryJDBoxStatisticsDataRequest {

	return &QueryJDBoxStatisticsDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/jdBoxStatistics",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,时间戳 (Optional)
 * param endTime: 查询截止时间,时间戳 (Optional)
 * param fields: 查询的字段，取值范围(avgbandwidth,pv,flow)。多个用逗号分隔。默认为空，表示查询带宽流量 (Optional)
 * param area:  (Optional)
 * param isp:  (Optional)
 * param period: 查询周期，当前取值范围：“oneMin,fiveMin”，分别表示1min，5min。默认为空，表示fiveMin (Optional)
 * param category: 业务类型 (Optional)
 * param macAddr: 设备id (Optional)
 * param pluginPin: 插件pin (Optional)
 */
func NewQueryJDBoxStatisticsDataRequestWithAllParams(
    startTime *int64,
    endTime *int64,
    fields *string,
    area *string,
    isp *string,
    period *string,
    category *string,
    macAddr *string,
    pluginPin *string,
) *QueryJDBoxStatisticsDataRequest {

    return &QueryJDBoxStatisticsDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/jdBoxStatistics",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Fields: fields,
        Area: area,
        Isp: isp,
        Period: period,
        Category: category,
        MacAddr: macAddr,
        PluginPin: pluginPin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryJDBoxStatisticsDataRequestWithoutParam() *QueryJDBoxStatisticsDataRequest {

    return &QueryJDBoxStatisticsDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/jdBoxStatistics",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,时间戳(Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetStartTime(startTime int64) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,时间戳(Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetEndTime(endTime int64) {
    r.EndTime = &endTime
}

/* param fields: 查询的字段，取值范围(avgbandwidth,pv,flow)。多个用逗号分隔。默认为空，表示查询带宽流量(Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetFields(fields string) {
    r.Fields = &fields
}

/* param area: (Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: (Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param period: 查询周期，当前取值范围：“oneMin,fiveMin”，分别表示1min，5min。默认为空，表示fiveMin(Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetPeriod(period string) {
    r.Period = &period
}

/* param category: 业务类型(Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetCategory(category string) {
    r.Category = &category
}

/* param macAddr: 设备id(Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetMacAddr(macAddr string) {
    r.MacAddr = &macAddr
}

/* param pluginPin: 插件pin(Optional) */
func (r *QueryJDBoxStatisticsDataRequest) SetPluginPin(pluginPin string) {
    r.PluginPin = &pluginPin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryJDBoxStatisticsDataRequest) GetRegionId() string {
    return ""
}

type QueryJDBoxStatisticsDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryJDBoxStatisticsDataResult `json:"result"`
}

type QueryJDBoxStatisticsDataResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Statistics []cdn.StatisticsDataItem `json:"statistics"`
}