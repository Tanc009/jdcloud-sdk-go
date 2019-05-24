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

type DeleteLiveSnapshotsRequest struct {

    core.JDCloudRequest

    /* 需要删除的截图ID，多个时以逗号（,）分隔
  */
    ImgIds string `json:"imgIds"`
}

/*
 * param imgIds: 需要删除的截图ID，多个时以逗号（,）分隔
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveSnapshotsRequest(
    imgIds string,
) *DeleteLiveSnapshotsRequest {

	return &DeleteLiveSnapshotsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/snapshots",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        ImgIds: imgIds,
	}
}

/*
 * param imgIds: 需要删除的截图ID，多个时以逗号（,）分隔
 (Required)
 */
func NewDeleteLiveSnapshotsRequestWithAllParams(
    imgIds string,
) *DeleteLiveSnapshotsRequest {

    return &DeleteLiveSnapshotsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshots",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        ImgIds: imgIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLiveSnapshotsRequestWithoutParam() *DeleteLiveSnapshotsRequest {

    return &DeleteLiveSnapshotsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshots",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param imgIds: 需要删除的截图ID，多个时以逗号（,）分隔
(Required) */
func (r *DeleteLiveSnapshotsRequest) SetImgIds(imgIds string) {
    r.ImgIds = imgIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveSnapshotsRequest) GetRegionId() string {
    return ""
}

type DeleteLiveSnapshotsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveSnapshotsResult `json:"result"`
}

type DeleteLiveSnapshotsResult struct {
}