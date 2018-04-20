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

type CreateAccountRequest struct {

    JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 用户名  */
    AccountName string `json:"accountName"`

    /* 用户密码  */
    AccountPassword string `json:"accountPassword"`
}

/*
 * param regionId: 地域代码 
 * param instanceId: 实例ID 
 * param accountName: 用户名 
 * param accountPassword: 用户密码 
 */
func NewCreateAccountRequest(
    regionId string,
    instanceId string,
    accountName string,
    accountPassword string,
) *CreateAccountRequest {

	return &CreateAccountRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/accounts",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        AccountName: accountName,
        AccountPassword: accountPassword,
	}
}

func (r *CreateAccountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *CreateAccountRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *CreateAccountRequest) SetAccountName(accountName string) {
    r.AccountName = accountName
}

func (r *CreateAccountRequest) SetAccountPassword(accountPassword string) {
    r.AccountPassword = accountPassword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAccountRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type CreateAccountResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result CreateAccountResult `json:"result"`
}

type CreateAccountResult struct {
}