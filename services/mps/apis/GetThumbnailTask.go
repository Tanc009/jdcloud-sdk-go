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
    mps "github.com/jdcloud-api/jdcloud-sdk-go/services/mps/models"
)

type GetThumbnailTaskRequest struct {

    core.JDCloudRequest

    /* region id  */
    RegionId string `json:"regionId"`

    /* task id  */
    TaskId string `json:"taskId"`
}

/*
 * param regionId: region id 
 * param taskId: task id 
 */
func NewGetThumbnailTaskRequest(
    regionId string,
    taskId string,
) *GetThumbnailTaskRequest {

	return &GetThumbnailTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/thumbnail/{taskId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TaskId: taskId,
	}
}

func (r *GetThumbnailTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *GetThumbnailTaskRequest) SetTaskId(taskId string) {
    r.TaskId = taskId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetThumbnailTaskRequest) GetRegionId() string {
    return r.RegionId
}

type GetThumbnailTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetThumbnailTaskResult `json:"result"`
}

type GetThumbnailTaskResult struct {
    TaskID string `json:"taskID"`
    Status string `json:"status"`
    ErrorCode int `json:"errorCode"`
    CreatedTime string `json:"createdTime"`
    LastUpdatedTime string `json:"lastUpdatedTime"`
    Source mps.ThumbnailTaskSource `json:"source"`
    Target mps.ThumbnailTaskTarget `json:"target"`
    Rule mps.ThumbnailTaskRule `json:"rule"`
}