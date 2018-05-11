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

type ModifyInstanceAttributeRequest struct {

    JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 名称；名称和描述必传其中一个；不为空且只允许中文、数字、大小写字母、英文下划线“_”及中划线“-”，长度不超过32字符 (Optional) */
    Name *string `json:"name"`

    /* 描述；名称和描述必传其中一个；长度不超过256个字符 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID 
 * param instanceId: Instance ID 
 * param name: 名称；名称和描述必传其中一个；不为空且只允许中文、数字、大小写字母、英文下划线“_”及中划线“-”，长度不超过32字符 (Optional)
 * param description: 描述；名称和描述必传其中一个；长度不超过256个字符 (Optional)
 */
func NewModifyInstanceAttributeRequest(
    regionId string,
    instanceId string,
) *ModifyInstanceAttributeRequest {

	return &ModifyInstanceAttributeRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceAttribute",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

func (r *ModifyInstanceAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ModifyInstanceAttributeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *ModifyInstanceAttributeRequest) SetName(name string) {
    r.Name = &name
}

func (r *ModifyInstanceAttributeRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceAttributeRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type ModifyInstanceAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result ModifyInstanceAttributeResult `json:"result"`
}

type ModifyInstanceAttributeResult struct {
}