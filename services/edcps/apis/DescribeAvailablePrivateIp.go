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
)

type DescribeAvailablePrivateIpRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 分布式云物理服务器ID  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param instanceId: 分布式云物理服务器ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAvailablePrivateIpRequest(
    regionId string,
    instanceId string,
) *DescribeAvailablePrivateIpRequest {

	return &DescribeAvailablePrivateIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/availablePrivateIps",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param instanceId: 分布式云物理服务器ID (Required)
 */
func NewDescribeAvailablePrivateIpRequestWithAllParams(
    regionId string,
    instanceId string,
) *DescribeAvailablePrivateIpRequest {

    return &DescribeAvailablePrivateIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/availablePrivateIps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAvailablePrivateIpRequestWithoutParam() *DescribeAvailablePrivateIpRequest {

    return &DescribeAvailablePrivateIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/availablePrivateIps",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *DescribeAvailablePrivateIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 分布式云物理服务器ID(Required) */
func (r *DescribeAvailablePrivateIpRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAvailablePrivateIpRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAvailablePrivateIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAvailablePrivateIpResult `json:"result"`
}

type DescribeAvailablePrivateIpResult struct {
}