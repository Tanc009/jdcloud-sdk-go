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
    live "github.com/jdcloud-api/jdcloud-sdk-go/services/live/models"
)

type AddLiveRecordTaskRequest struct {

    core.JDCloudRequest

    /* 推流域名  */
    PublishDomain string `json:"publishDomain"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 流名称  */
    StreamName string `json:"streamName"`

    /* 录制时间集合
- 最大支持10段,多段合并成一个文件
- 多段时间跨度最小不能小于10s
- 多段时间跨度最大不能超过8小时
  */
    RecordTimes []live.RecordTime `json:"recordTimes"`

    /* 存储桶  */
    SaveBucket string `json:"saveBucket"`

    /* 存储地址  */
    SaveEndpoint string `json:"saveEndpoint"`

    /* 录制文件类型:
- 取值: ts,flv,mp4 (多种类型之间用;隔开)
- 不区分大小写
  */
    RecordFileType string `json:"recordFileType"`

    /* 录制文件存储Object:
- 默认: /record/{Date}/{ServerId}/{AppName}/{StreamName}/{StartTime}_{EndTime}
 (Optional) */
    SaveObject *string `json:"saveObject"`

    /* 打点录制任务外部id(可传入您的id,在回调时会在该字段返回)
 (Optional) */
    TaskExternalId *string `json:"taskExternalId"`
}

/*
 * param publishDomain: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param streamName: 流名称 (Required)
 * param recordTimes: 录制时间集合
- 最大支持10段,多段合并成一个文件
- 多段时间跨度最小不能小于10s
- 多段时间跨度最大不能超过8小时
 (Required)
 * param saveBucket: 存储桶 (Required)
 * param saveEndpoint: 存储地址 (Required)
 * param recordFileType: 录制文件类型:
- 取值: ts,flv,mp4 (多种类型之间用;隔开)
- 不区分大小写
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddLiveRecordTaskRequest(
    publishDomain string,
    appName string,
    streamName string,
    recordTimes []live.RecordTime,
    saveBucket string,
    saveEndpoint string,
    recordFileType string,
) *AddLiveRecordTaskRequest {

	return &AddLiveRecordTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/records/{publishDomain}/appNames/{appName}/streamNames/{streamName}/task",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        RecordTimes: recordTimes,
        SaveBucket: saveBucket,
        SaveEndpoint: saveEndpoint,
        RecordFileType: recordFileType,
	}
}

/*
 * param publishDomain: 推流域名 (Required)
 * param appName: 应用名称 (Required)
 * param streamName: 流名称 (Required)
 * param recordTimes: 录制时间集合
- 最大支持10段,多段合并成一个文件
- 多段时间跨度最小不能小于10s
- 多段时间跨度最大不能超过8小时
 (Required)
 * param saveBucket: 存储桶 (Required)
 * param saveEndpoint: 存储地址 (Required)
 * param recordFileType: 录制文件类型:
- 取值: ts,flv,mp4 (多种类型之间用;隔开)
- 不区分大小写
 (Required)
 * param saveObject: 录制文件存储Object:
- 默认: /record/{Date}/{ServerId}/{AppName}/{StreamName}/{StartTime}_{EndTime}
 (Optional)
 * param taskExternalId: 打点录制任务外部id(可传入您的id,在回调时会在该字段返回)
 (Optional)
 */
func NewAddLiveRecordTaskRequestWithAllParams(
    publishDomain string,
    appName string,
    streamName string,
    recordTimes []live.RecordTime,
    saveBucket string,
    saveEndpoint string,
    recordFileType string,
    saveObject *string,
    taskExternalId *string,
) *AddLiveRecordTaskRequest {

    return &AddLiveRecordTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/records/{publishDomain}/appNames/{appName}/streamNames/{streamName}/task",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        RecordTimes: recordTimes,
        SaveBucket: saveBucket,
        SaveEndpoint: saveEndpoint,
        RecordFileType: recordFileType,
        SaveObject: saveObject,
        TaskExternalId: taskExternalId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddLiveRecordTaskRequestWithoutParam() *AddLiveRecordTaskRequest {

    return &AddLiveRecordTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/records/{publishDomain}/appNames/{appName}/streamNames/{streamName}/task",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 推流域名(Required) */
func (r *AddLiveRecordTaskRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param appName: 应用名称(Required) */
func (r *AddLiveRecordTaskRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param streamName: 流名称(Required) */
func (r *AddLiveRecordTaskRequest) SetStreamName(streamName string) {
    r.StreamName = streamName
}

/* param recordTimes: 录制时间集合
- 最大支持10段,多段合并成一个文件
- 多段时间跨度最小不能小于10s
- 多段时间跨度最大不能超过8小时
(Required) */
func (r *AddLiveRecordTaskRequest) SetRecordTimes(recordTimes []live.RecordTime) {
    r.RecordTimes = recordTimes
}

/* param saveBucket: 存储桶(Required) */
func (r *AddLiveRecordTaskRequest) SetSaveBucket(saveBucket string) {
    r.SaveBucket = saveBucket
}

/* param saveEndpoint: 存储地址(Required) */
func (r *AddLiveRecordTaskRequest) SetSaveEndpoint(saveEndpoint string) {
    r.SaveEndpoint = saveEndpoint
}

/* param recordFileType: 录制文件类型:
- 取值: ts,flv,mp4 (多种类型之间用;隔开)
- 不区分大小写
(Required) */
func (r *AddLiveRecordTaskRequest) SetRecordFileType(recordFileType string) {
    r.RecordFileType = recordFileType
}

/* param saveObject: 录制文件存储Object:
- 默认: /record/{Date}/{ServerId}/{AppName}/{StreamName}/{StartTime}_{EndTime}
(Optional) */
func (r *AddLiveRecordTaskRequest) SetSaveObject(saveObject string) {
    r.SaveObject = &saveObject
}

/* param taskExternalId: 打点录制任务外部id(可传入您的id,在回调时会在该字段返回)
(Optional) */
func (r *AddLiveRecordTaskRequest) SetTaskExternalId(taskExternalId string) {
    r.TaskExternalId = &taskExternalId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddLiveRecordTaskRequest) GetRegionId() string {
    return ""
}

type AddLiveRecordTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddLiveRecordTaskResult `json:"result"`
}

type AddLiveRecordTaskResult struct {
}