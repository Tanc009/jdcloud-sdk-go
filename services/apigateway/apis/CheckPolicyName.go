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

type CheckPolicyNameRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    PolicyName string `json:"policyName"`
}

/*
 * param regionId: 地域ID (Required)
 * param policyName:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckPolicyNameRequest(
    regionId string,
    policyName string,
) *CheckPolicyNameRequest {

	return &CheckPolicyNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rateLimitPolicies:checkPolicyNameExist",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PolicyName: policyName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param policyName:  (Required)
 */
func NewCheckPolicyNameRequestWithAllParams(
    regionId string,
    policyName string,
) *CheckPolicyNameRequest {

    return &CheckPolicyNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rateLimitPolicies:checkPolicyNameExist",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PolicyName: policyName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckPolicyNameRequestWithoutParam() *CheckPolicyNameRequest {

    return &CheckPolicyNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rateLimitPolicies:checkPolicyNameExist",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CheckPolicyNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param policyName: (Required) */
func (r *CheckPolicyNameRequest) SetPolicyName(policyName string) {
    r.PolicyName = policyName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckPolicyNameRequest) GetRegionId() string {
    return r.RegionId
}

type CheckPolicyNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckPolicyNameResult `json:"result"`
}

type CheckPolicyNameResult struct {
    PolicyId string `json:"policyId"`
}