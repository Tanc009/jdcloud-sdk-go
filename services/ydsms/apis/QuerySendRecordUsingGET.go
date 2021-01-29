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
    ydsms "github.com/jdcloud-api/jdcloud-sdk-go/services/ydsms/models"
)

type QuerySendRecordUsingGETRequest struct {

    core.JDCloudRequest

    /* 应用id  */
    AppId string `json:"appId"`

    /* 发送状态，0 全部状态 1 发送成功 2 发送失败 3 已发送未回执， 默认为0 (Optional) */
    SendStatus *int `json:"sendStatus"`

    /* 手机号码 (Optional) */
    SendNumber *string `json:"sendNumber"`

    /* 模板id (Optional) */
    TemplateId *string `json:"templateId"`

    /* 签名id (Optional) */
    Sign_id *string `json:"sign_id"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"  */
    EndTime string `json:"endTime"`

    /* 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"  */
    StartTime string `json:"startTime"`
}

/*
 * param appId: 应用id (Required)
 * param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 * param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQuerySendRecordUsingGETRequest(
    appId string,
    endTime string,
    startTime string,
) *QuerySendRecordUsingGETRequest {

	return &QuerySendRecordUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsApps/{appId}:querySendRecord",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        EndTime: endTime,
        StartTime: startTime,
	}
}

/*
 * param appId: 应用id (Required)
 * param sendStatus: 发送状态，0 全部状态 1 发送成功 2 发送失败 3 已发送未回执， 默认为0 (Optional)
 * param sendNumber: 手机号码 (Optional)
 * param templateId: 模板id (Optional)
 * param sign_id: 签名id (Optional)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 分页大小 (Optional)
 * param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 * param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 */
func NewQuerySendRecordUsingGETRequestWithAllParams(
    appId string,
    sendStatus *int,
    sendNumber *string,
    templateId *string,
    sign_id *string,
    pageNumber *int,
    pageSize *int,
    endTime string,
    startTime string,
) *QuerySendRecordUsingGETRequest {

    return &QuerySendRecordUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsApps/{appId}:querySendRecord",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        SendStatus: sendStatus,
        SendNumber: sendNumber,
        TemplateId: templateId,
        Sign_id: sign_id,
        PageNumber: pageNumber,
        PageSize: pageSize,
        EndTime: endTime,
        StartTime: startTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQuerySendRecordUsingGETRequestWithoutParam() *QuerySendRecordUsingGETRequest {

    return &QuerySendRecordUsingGETRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsApps/{appId}:querySendRecord",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用id(Required) */
func (r *QuerySendRecordUsingGETRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param sendStatus: 发送状态，0 全部状态 1 发送成功 2 发送失败 3 已发送未回执， 默认为0(Optional) */
func (r *QuerySendRecordUsingGETRequest) SetSendStatus(sendStatus int) {
    r.SendStatus = &sendStatus
}

/* param sendNumber: 手机号码(Optional) */
func (r *QuerySendRecordUsingGETRequest) SetSendNumber(sendNumber string) {
    r.SendNumber = &sendNumber
}

/* param templateId: 模板id(Optional) */
func (r *QuerySendRecordUsingGETRequest) SetTemplateId(templateId string) {
    r.TemplateId = &templateId
}

/* param sign_id: 签名id(Optional) */
func (r *QuerySendRecordUsingGETRequest) SetSign_id(sign_id string) {
    r.Sign_id = &sign_id
}

/* param pageNumber: 页码(Optional) */
func (r *QuerySendRecordUsingGETRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小(Optional) */
func (r *QuerySendRecordUsingGETRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"(Required) */
func (r *QuerySendRecordUsingGETRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"(Required) */
func (r *QuerySendRecordUsingGETRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QuerySendRecordUsingGETRequest) GetRegionId() string {
    return ""
}

type QuerySendRecordUsingGETResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QuerySendRecordUsingGETResult `json:"result"`
}

type QuerySendRecordUsingGETResult struct {
    SendRecords []ydsms.SendRecord `json:"sendRecords"`
    TotalCount int64 `json:"totalCount"`
}