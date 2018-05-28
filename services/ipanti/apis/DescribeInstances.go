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

type DescribeInstancesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 实例名称，可模糊匹配 (Optional) */
    Name *string `json:"name"`
}

/*
 * param regionId: Region ID 
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 * param name: 实例名称，可模糊匹配 (Optional)
 */
func NewDescribeInstancesRequest(
    regionId string,
) *DescribeInstancesRequest {

	return &DescribeInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *DescribeInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeInstancesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

func (r *DescribeInstancesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

func (r *DescribeInstancesRequest) SetName(name string) {
    r.Name = &name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstancesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstancesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstancesResult `json:"result"`
}

type DescribeInstancesResult struct {
    Instances []ipanti.Instance `json:"instances"`
    TotalCount int `json:"totalCount"`
}