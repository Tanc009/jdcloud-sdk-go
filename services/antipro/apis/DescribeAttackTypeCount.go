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
    antipro "github.com/jdcloud-api/jdcloud-sdk-go/services/antipro/models"
)

type DescribeAttackTypeCountRequest struct {

    core.JDCloudRequest

    /* 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ  */
    StartTime string `json:"startTime"`

    /* 结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ  */
    EndTime string `json:"endTime"`

    /* 防护包实例 Id (Optional) */
    InstanceId *string `json:"instanceId"`

    /* DDoS 防护包已防护的公网 IP
- 使用 <a href="http://docs.jdcloud.com/anti-ddos-protection-package/api/describeprotectediplist">describeProtectedIpList</a> 接口查询 DDoS 防护包已防护的公网 IP
 (Optional) */
    Ip []string `json:"ip"`
}

/*
 * param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAttackTypeCountRequest(
    startTime string,
    endTime string,
) *DescribeAttackTypeCountRequest {

	return &DescribeAttackTypeCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeAttackTypeCount",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param instanceId: 防护包实例 Id (Optional)
 * param ip: DDoS 防护包已防护的公网 IP
- 使用 <a href="http://docs.jdcloud.com/anti-ddos-protection-package/api/describeprotectediplist">describeProtectedIpList</a> 接口查询 DDoS 防护包已防护的公网 IP
 (Optional)
 */
func NewDescribeAttackTypeCountRequestWithAllParams(
    startTime string,
    endTime string,
    instanceId *string,
    ip []string,
) *DescribeAttackTypeCountRequest {

    return &DescribeAttackTypeCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeAttackTypeCount",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        InstanceId: instanceId,
        Ip: ip,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAttackTypeCountRequestWithoutParam() *DescribeAttackTypeCountRequest {

    return &DescribeAttackTypeCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeAttackTypeCount",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 开始时间, 只能查询最近 90 天以内的数据, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeAttackTypeCountRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeAttackTypeCountRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param instanceId: 防护包实例 Id(Optional) */
func (r *DescribeAttackTypeCountRequest) SetInstanceId(instanceId string) {
    r.InstanceId = &instanceId
}

/* param ip: DDoS 防护包已防护的公网 IP
- 使用 <a href="http://docs.jdcloud.com/anti-ddos-protection-package/api/describeprotectediplist">describeProtectedIpList</a> 接口查询 DDoS 防护包已防护的公网 IP
(Optional) */
func (r *DescribeAttackTypeCountRequest) SetIp(ip []string) {
    r.Ip = ip
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAttackTypeCountRequest) GetRegionId() string {
    return ""
}

type DescribeAttackTypeCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAttackTypeCountResult `json:"result"`
}

type DescribeAttackTypeCountResult struct {
    DataList []antipro.AttackTypeCount `json:"dataList"`
}