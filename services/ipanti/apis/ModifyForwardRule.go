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

type ModifyForwardRuleRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* 转发规则id  */
    ForwardRuleId string `json:"forwardRuleId"`

    /* 非网站类规则参数  */
    ForwardRuleSpec *ipanti.ForwardRuleSpec `json:"forwardRuleSpec"`
}

/*
 * param regionId: Region ID 
 * param instanceId: 实例id 
 * param forwardRuleId: 转发规则id 
 * param forwardRuleSpec: 非网站类规则参数 
 */
func NewModifyForwardRuleRequest(
    regionId string,
    instanceId string,
    forwardRuleId string,
    forwardRuleSpec *ipanti.ForwardRuleSpec,
) *ModifyForwardRuleRequest {

	return &ModifyForwardRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/forwardRules/{forwardRuleId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ForwardRuleId: forwardRuleId,
        ForwardRuleSpec: forwardRuleSpec,
	}
}

func (r *ModifyForwardRuleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ModifyForwardRuleRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *ModifyForwardRuleRequest) SetForwardRuleId(forwardRuleId string) {
    r.ForwardRuleId = forwardRuleId
}

func (r *ModifyForwardRuleRequest) SetForwardRuleSpec(forwardRuleSpec *ipanti.ForwardRuleSpec) {
    r.ForwardRuleSpec = forwardRuleSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyForwardRuleRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyForwardRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyForwardRuleResult `json:"result"`
}

type ModifyForwardRuleResult struct {
}