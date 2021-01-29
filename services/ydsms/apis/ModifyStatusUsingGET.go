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
)

type ModifyStatusUsingGETRequest struct {

    core.JDCloudRequest

    /* 应用id  */
    AppId string `json:"appId"`

    /* 应用状态,0 停用 1 启用 (Optional) */
    Status *int `json:"status"`
}

/*
 * param appId: 应用id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyStatusUsingGETRequest(
    appId string,
) *ModifyStatusUsingGETRequest {

	return &ModifyStatusUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsApps/{appId}:modifyStatus",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
	}
}

/*
 * param appId: 应用id (Required)
 * param status: 应用状态,0 停用 1 启用 (Optional)
 */
func NewModifyStatusUsingGETRequestWithAllParams(
    appId string,
    status *int,
) *ModifyStatusUsingGETRequest {

    return &ModifyStatusUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsApps/{appId}:modifyStatus",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyStatusUsingGETRequestWithoutParam() *ModifyStatusUsingGETRequest {

    return &ModifyStatusUsingGETRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsApps/{appId}:modifyStatus",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用id(Required) */
func (r *ModifyStatusUsingGETRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param status: 应用状态,0 停用 1 启用(Optional) */
func (r *ModifyStatusUsingGETRequest) SetStatus(status int) {
    r.Status = &status
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyStatusUsingGETRequest) GetRegionId() string {
    return ""
}

type ModifyStatusUsingGETResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyStatusUsingGETResult `json:"result"`
}

type ModifyStatusUsingGETResult struct {
    AppId string `json:"appId"`
}