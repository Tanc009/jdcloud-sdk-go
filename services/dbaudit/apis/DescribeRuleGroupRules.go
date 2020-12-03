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
    dbaudit "github.com/jdcloud-api/jdcloud-sdk-go/services/dbaudit/models"
)

type DescribeRuleGroupRulesRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 审计实例ID  */
    InsId string `json:"insId"`

    /* 规则组ID  */
    RuleGroupId string `json:"ruleGroupId"`

    /* 数据库ID (Optional) */
    DbId *string `json:"dbId"`
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param ruleGroupId: 规则组ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRuleGroupRulesRequest(
    regionId string,
    insId string,
    ruleGroupId string,
) *DescribeRuleGroupRulesRequest {

	return &DescribeRuleGroupRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InsId: insId,
        RuleGroupId: ruleGroupId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param insId: 审计实例ID (Required)
 * param ruleGroupId: 规则组ID (Required)
 * param dbId: 数据库ID (Optional)
 */
func NewDescribeRuleGroupRulesRequestWithAllParams(
    regionId string,
    insId string,
    ruleGroupId string,
    dbId *string,
) *DescribeRuleGroupRulesRequest {

    return &DescribeRuleGroupRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InsId: insId,
        RuleGroupId: ruleGroupId,
        DbId: dbId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRuleGroupRulesRequestWithoutParam() *DescribeRuleGroupRulesRequest {

    return &DescribeRuleGroupRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{insId}/ruleGroups/{ruleGroupId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeRuleGroupRulesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param insId: 审计实例ID(Required) */
func (r *DescribeRuleGroupRulesRequest) SetInsId(insId string) {
    r.InsId = insId
}

/* param ruleGroupId: 规则组ID(Required) */
func (r *DescribeRuleGroupRulesRequest) SetRuleGroupId(ruleGroupId string) {
    r.RuleGroupId = ruleGroupId
}

/* param dbId: 数据库ID(Optional) */
func (r *DescribeRuleGroupRulesRequest) SetDbId(dbId string) {
    r.DbId = &dbId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRuleGroupRulesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeRuleGroupRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRuleGroupRulesResult `json:"result"`
}

type DescribeRuleGroupRulesResult struct {
    TotalCount int `json:"totalCount"`
    List []dbaudit.Rule `json:"list"`
}