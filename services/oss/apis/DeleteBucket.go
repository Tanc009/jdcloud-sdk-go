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

type DeleteBucketRequest struct {

    JDCloudRequest

    /* Region ID，例如：cn-north-1  */
    RegionId string `json:"regionId"`

    /* bucket名字，例如：test-bucket  */
    Bucketname string `json:"bucketname"`
}

/*
 * param regionId: Region ID，例如：cn-north-1 
 * param bucketname: bucket名字，例如：test-bucket 
 */
func NewDeleteBucketRequest(
    regionId string,
    bucketname string,
) *DeleteBucketRequest {

	return &DeleteBucketRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/buckets/{bucketname}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Bucketname: bucketname,
	}
}

func (r *DeleteBucketRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DeleteBucketRequest) SetBucketname(bucketname string) {
    r.Bucketname = bucketname
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteBucketRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DeleteBucketResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result DeleteBucketResult `json:"result"`
}

type DeleteBucketResult struct {
}