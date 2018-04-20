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
)

type EnableInstanceUrlWhiteListRequest struct {

    JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: Region ID 
 * param instanceId: 实例id 
 */
func NewEnableInstanceUrlWhiteListRequest(
    regionId string,
    instanceId string,
) *EnableInstanceUrlWhiteListRequest {

	return &EnableInstanceUrlWhiteListRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/enableUrlWhiteList",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

func (r *EnableInstanceUrlWhiteListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *EnableInstanceUrlWhiteListRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableInstanceUrlWhiteListRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type EnableInstanceUrlWhiteListResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result EnableInstanceUrlWhiteListResult `json:"result"`
}

type EnableInstanceUrlWhiteListResult struct {
}