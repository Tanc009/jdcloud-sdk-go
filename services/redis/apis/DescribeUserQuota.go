// Copyright 2018-2025 JDCLOUD.COM
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
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeUserQuotaRequest struct {

    JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前缓存Redis有华北、华南、华东区域，对应Region ID为cn-north-1、cn-south-1、cn-east-2 
 */
func NewDescribeUserQuotaRequest(
    regionId string,
) *DescribeUserQuotaRequest {

	return &DescribeUserQuotaRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/quota",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

func (r *DescribeUserQuotaRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUserQuotaRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeUserQuotaResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result DescribeUserQuotaResult `json:"result"`
}

type DescribeUserQuotaResult struct {
    Quota common.Quota `json:"quota"`
}