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
    rms "github.com/jdcloud-api/jdcloud-sdk-go/services/rms/models"
)

type QueryOneTemplateRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* templateId参数 (Optional) */
    TemplateId *string `json:"templateId"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryOneTemplateRequest(
    regionId string,
) *QueryOneTemplateRequest {

	return &QueryOneTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/queryOneTemplate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param templateId: templateId参数 (Optional)
 */
func NewQueryOneTemplateRequestWithAllParams(
    regionId string,
    templateId *string,
) *QueryOneTemplateRequest {

    return &QueryOneTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/queryOneTemplate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TemplateId: templateId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryOneTemplateRequestWithoutParam() *QueryOneTemplateRequest {

    return &QueryOneTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/queryOneTemplate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *QueryOneTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param templateId: templateId参数(Optional) */
func (r *QueryOneTemplateRequest) SetTemplateId(templateId string) {
    r.TemplateId = &templateId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryOneTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type QueryOneTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryOneTemplateResult `json:"result"`
}

type QueryOneTemplateResult struct {
    Data []rms.RespQueryOneTemplateResourceList `json:"data"`
    Description string `json:"description"`
    MediaLength string `json:"mediaLength"`
    Type string `json:"type"`
    Title string `json:"title"`
    IsTuiding string `json:"isTuiding"`
    Message string `json:"message"`
    Status string `json:"status"`
}