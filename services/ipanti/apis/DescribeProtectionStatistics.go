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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type DescribeProtectionStatisticsRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeProtectionStatisticsRequest(
    regionId string,
) *DescribeProtectionStatisticsRequest {

	return &DescribeProtectionStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/describeProtectionStatistics",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 */
func NewDescribeProtectionStatisticsRequestWithAllParams(
    regionId string,
) *DescribeProtectionStatisticsRequest {

    return &DescribeProtectionStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeProtectionStatistics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeProtectionStatisticsRequestWithoutParam() *DescribeProtectionStatisticsRequest {

    return &DescribeProtectionStatisticsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeProtectionStatistics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DescribeProtectionStatisticsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeProtectionStatisticsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeProtectionStatisticsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeProtectionStatisticsResult `json:"result"`
}

type DescribeProtectionStatisticsResult struct {
    Data ipanti.ProtectionStatistics `json:"data"`
}