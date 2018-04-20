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

type ModifyInstanceIpBlackListRequest struct {

    JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* ip黑名单列表  */
    IpBlackList []string `json:"ipBlackList"`
}

/*
 * param regionId: Region ID 
 * param instanceId: 实例id 
 * param ipBlackList: ip黑名单列表 
 */
func NewModifyInstanceIpBlackListRequest(
    regionId string,
    instanceId string,
    ipBlackList []string,
) *ModifyInstanceIpBlackListRequest {

	return &ModifyInstanceIpBlackListRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/setIpBlackList",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        IpBlackList: ipBlackList,
	}
}

func (r *ModifyInstanceIpBlackListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *ModifyInstanceIpBlackListRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *ModifyInstanceIpBlackListRequest) SetIpBlackList(ipBlackList []string) {
    r.IpBlackList = ipBlackList
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceIpBlackListRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type ModifyInstanceIpBlackListResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result ModifyInstanceIpBlackListResult `json:"result"`
}

type ModifyInstanceIpBlackListResult struct {
}