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

type DeleteLiveRecordingsRequest struct {

    core.JDCloudRequest

    /* 需要删除的录制文件在oss的url
  */
    FileUrl string `json:"fileUrl"`

    /* 是否深度删除所有的ts文件，仅对.m3u8录制文件生效。默认: true
 (Optional) */
    Completely *bool `json:"completely"`
}

/*
 * param fileUrl: 需要删除的录制文件在oss的url
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteLiveRecordingsRequest(
    fileUrl string,
) *DeleteLiveRecordingsRequest {

	return &DeleteLiveRecordingsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/recordings:delete",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        FileUrl: fileUrl,
	}
}

/*
 * param fileUrl: 需要删除的录制文件在oss的url
 (Required)
 * param completely: 是否深度删除所有的ts文件，仅对.m3u8录制文件生效。默认: true
 (Optional)
 */
func NewDeleteLiveRecordingsRequestWithAllParams(
    fileUrl string,
    completely *bool,
) *DeleteLiveRecordingsRequest {

    return &DeleteLiveRecordingsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordings:delete",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        FileUrl: fileUrl,
        Completely: completely,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteLiveRecordingsRequestWithoutParam() *DeleteLiveRecordingsRequest {

    return &DeleteLiveRecordingsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordings:delete",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param fileUrl: 需要删除的录制文件在oss的url
(Required) */
func (r *DeleteLiveRecordingsRequest) SetFileUrl(fileUrl string) {
    r.FileUrl = fileUrl
}

/* param completely: 是否深度删除所有的ts文件，仅对.m3u8录制文件生效。默认: true
(Optional) */
func (r *DeleteLiveRecordingsRequest) SetCompletely(completely bool) {
    r.Completely = &completely
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteLiveRecordingsRequest) GetRegionId() string {
    return ""
}

type DeleteLiveRecordingsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteLiveRecordingsResult `json:"result"`
}

type DeleteLiveRecordingsResult struct {
}