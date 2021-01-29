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

type ModifyInstanceCustomPageDefaultRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 实例 ID  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 实例 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceCustomPageDefaultRequest(
    regionId string,
    instanceId string,
) *ModifyInstanceCustomPageDefaultRequest {

	return &ModifyInstanceCustomPageDefaultRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceCustomPageDefault",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 实例 ID (Required)
 */
func NewModifyInstanceCustomPageDefaultRequestWithAllParams(
    regionId string,
    instanceId string,
) *ModifyInstanceCustomPageDefaultRequest {

    return &ModifyInstanceCustomPageDefaultRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceCustomPageDefault",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceCustomPageDefaultRequestWithoutParam() *ModifyInstanceCustomPageDefaultRequest {

    return &ModifyInstanceCustomPageDefaultRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceCustomPageDefault",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *ModifyInstanceCustomPageDefaultRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例 ID(Required) */
func (r *ModifyInstanceCustomPageDefaultRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceCustomPageDefaultRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceCustomPageDefaultResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceCustomPageDefaultResult `json:"result"`
}

type ModifyInstanceCustomPageDefaultResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}