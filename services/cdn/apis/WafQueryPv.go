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

type WafQueryPvRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewWafQueryPvRequest(
) *WafQueryPvRequest {

	return &WafQueryPvRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/wafPvStatistic",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 */
func NewWafQueryPvRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
) *WafQueryPvRequest {

    return &WafQueryPvRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/wafPvStatistic",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewWafQueryPvRequestWithoutParam() *WafQueryPvRequest {

    return &WafQueryPvRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/wafPvStatistic",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *WafQueryPvRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *WafQueryPvRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *WafQueryPvRequest) SetDomain(domain string) {
    r.Domain = &domain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r WafQueryPvRequest) GetRegionId() string {
    return ""
}

type WafQueryPvResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result WafQueryPvResult `json:"result"`
}

type WafQueryPvResult struct {
    Pvs []cdn.PvItem `json:"pvs"`
    PeakAttackPv int `json:"peakAttackPv"`
    PeakTotalPv int `json:"peakTotalPv"`
}