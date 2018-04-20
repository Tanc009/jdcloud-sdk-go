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

type ModifyInstanceNameRequest struct {

    JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* 新的实例名称  */
    Name string `json:"name"`
}

/*
 * param regionId: Region ID 
 * param instanceId: 实例id 
 * param name: 新的实例名称 
 */
func NewModifyInstanceNameRequest(
    regionId string,
    instanceId string,
    name string,
) *ModifyInstanceNameRequest {

	return &ModifyInstanceNameRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/rename",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Name: name,
	}
}

func (r *ModifyInstanceNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ModifyInstanceNameRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *ModifyInstanceNameRequest) SetName(name string) {
    r.Name = name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceNameRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type ModifyInstanceNameResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result ModifyInstanceNameResult `json:"result"`
}

type ModifyInstanceNameResult struct {
}