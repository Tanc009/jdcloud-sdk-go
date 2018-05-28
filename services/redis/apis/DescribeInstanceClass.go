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
    redis "github.com/jdcloud-api/jdcloud-sdk-go/services/redis/models"
)

type DescribeInstanceClassRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: Region ID 
 */
func NewDescribeInstanceClassRequest(
    regionId string,
) *DescribeInstanceClassRequest {

	return &DescribeInstanceClassRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceClass",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *DescribeInstanceClassRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceClassRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceClassResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceClassResult `json:"result"`
}

type DescribeInstanceClassResult struct {
    InstanceClasses []redis.InstanceClass `json:"instanceClasses"`
    TotalCount int `json:"totalCount"`
}