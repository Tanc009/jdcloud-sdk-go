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

type DeleteAlarmsRequest struct {

    JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 待删除的规则id，用"|"间隔  */
    Ids string `json:"ids"`
}

/*
 * param regionId: 地域 Id 
 * param ids: 待删除的规则id，用"|"间隔 
 */
func NewDeleteAlarmsRequest(
    regionId string,
    ids string,
) *DeleteAlarmsRequest {

	return &DeleteAlarmsRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/alarms",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Ids: ids,
	}
}

func (r *DeleteAlarmsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DeleteAlarmsRequest) SetIds(ids string) {
    r.Ids = ids
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteAlarmsRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DeleteAlarmsResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result DeleteAlarmsResult `json:"result"`
}

type DeleteAlarmsResult struct {
}