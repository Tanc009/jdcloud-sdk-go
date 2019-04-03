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
    sms "github.com/jdcloud-api/jdcloud-sdk-go/services/sms/models"
)

type PullMtMsgRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 拉取发送短信请求参数  */
    PullMtMsgSpec *sms.PullMtMsgSpec `json:"pullMtMsgSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param pullMtMsgSpec: 拉取发送短信请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPullMtMsgRequest(
    regionId string,
    pullMtMsgSpec *sms.PullMtMsgSpec,
) *PullMtMsgRequest {

	return &PullMtMsgRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pullMtMsg",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PullMtMsgSpec: pullMtMsgSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pullMtMsgSpec: 拉取发送短信请求参数 (Required)
 */
func NewPullMtMsgRequestWithAllParams(
    regionId string,
    pullMtMsgSpec *sms.PullMtMsgSpec,
) *PullMtMsgRequest {

    return &PullMtMsgRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pullMtMsg",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PullMtMsgSpec: pullMtMsgSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPullMtMsgRequestWithoutParam() *PullMtMsgRequest {

    return &PullMtMsgRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pullMtMsg",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *PullMtMsgRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pullMtMsgSpec: 拉取发送短信请求参数(Required) */
func (r *PullMtMsgRequest) SetPullMtMsgSpec(pullMtMsgSpec *sms.PullMtMsgSpec) {
    r.PullMtMsgSpec = pullMtMsgSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PullMtMsgRequest) GetRegionId() string {
    return r.RegionId
}

type PullMtMsgResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PullMtMsgResult `json:"result"`
}

type PullMtMsgResult struct {
    Data []sms.PullMtMsg `json:"data"`
}