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

type DeleteForwardRuleRequest struct {

    core.JDCloudRequest

    /* 区域 ID, 高防不区分区域, 传 cn-north-1 即可  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId string `json:"instanceId"`

    /* 转发规则 Id  */
    ForwardRuleId string `json:"forwardRuleId"`
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleId: 转发规则 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteForwardRuleRequest(
    regionId string,
    instanceId string,
    forwardRuleId string,
) *DeleteForwardRuleRequest {

	return &DeleteForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
	}
}

/*
 * param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可 (Required)
 * param instanceId: 高防实例 Id (Required)
 * param forwardRuleId: 转发规则 Id (Required)
 */
func NewDeleteForwardRuleRequestWithAllParams(
    regionId string,
    instanceId string,
    forwardRuleId string,
) *DeleteForwardRuleRequest {

    return &DeleteForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteForwardRuleRequestWithoutParam() *DeleteForwardRuleRequest {

    return &DeleteForwardRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 ID, 高防不区分区域, 传 cn-north-1 即可(Required) */
func (r *DeleteForwardRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DeleteForwardRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param forwardRuleId: 转发规则 Id(Required) */
func (r *DeleteForwardRuleRequest) SetForwardRuleId(forwardRuleId string) {
    r.ForwardRuleId = forwardRuleId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteForwardRuleRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteForwardRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteForwardRuleResult `json:"result"`
}

type DeleteForwardRuleResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}