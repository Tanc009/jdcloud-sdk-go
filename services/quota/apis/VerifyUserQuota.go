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
    quota "github.com/jdcloud-api/jdcloud-sdk-go/services/quota/models"
)

type VerifyUserQuotaRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 站点类型 (Optional) */
    SiteType *int `json:"siteType"`

    /* 业务线 (Optional) */
    AppCode *string `json:"appCode"`

    /* 资源产品线 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 地域 (Optional) */
    Region *string `json:"region"`

    /* 父层资源id（针对有两层结构的服务） (Optional) */
    ParentResourceId *string `json:"parentResourceId"`

    /* 业务唯一键 (Optional) */
    Uid *string `json:"uid"`

    /* 配额数量 (Optional) */
    QuotaAmount *int `json:"quotaAmount"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewVerifyUserQuotaRequest(
    regionId string,
) *VerifyUserQuotaRequest {

	return &VerifyUserQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/userQuota:verify",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Optional)
 * param siteType: 站点类型 (Optional)
 * param appCode: 业务线 (Optional)
 * param serviceCode: 资源产品线 (Optional)
 * param region: 地域 (Optional)
 * param parentResourceId: 父层资源id（针对有两层结构的服务） (Optional)
 * param uid: 业务唯一键 (Optional)
 * param quotaAmount: 配额数量 (Optional)
 */
func NewVerifyUserQuotaRequestWithAllParams(
    regionId string,
    pin *string,
    siteType *int,
    appCode *string,
    serviceCode *string,
    region *string,
    parentResourceId *string,
    uid *string,
    quotaAmount *int,
) *VerifyUserQuotaRequest {

    return &VerifyUserQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userQuota:verify",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        Pin: pin,
        SiteType: siteType,
        AppCode: appCode,
        ServiceCode: serviceCode,
        Region: region,
        ParentResourceId: parentResourceId,
        Uid: uid,
        QuotaAmount: quotaAmount,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewVerifyUserQuotaRequestWithoutParam() *VerifyUserQuotaRequest {

    return &VerifyUserQuotaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/userQuota:verify",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *VerifyUserQuotaRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Optional) */
func (r *VerifyUserQuotaRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param siteType: 站点类型(Optional) */
func (r *VerifyUserQuotaRequest) SetSiteType(siteType int) {
    r.SiteType = &siteType
}

/* param appCode: 业务线(Optional) */
func (r *VerifyUserQuotaRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: 资源产品线(Optional) */
func (r *VerifyUserQuotaRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param region: 地域(Optional) */
func (r *VerifyUserQuotaRequest) SetRegion(region string) {
    r.Region = &region
}

/* param parentResourceId: 父层资源id（针对有两层结构的服务）(Optional) */
func (r *VerifyUserQuotaRequest) SetParentResourceId(parentResourceId string) {
    r.ParentResourceId = &parentResourceId
}

/* param uid: 业务唯一键(Optional) */
func (r *VerifyUserQuotaRequest) SetUid(uid string) {
    r.Uid = &uid
}

/* param quotaAmount: 配额数量(Optional) */
func (r *VerifyUserQuotaRequest) SetQuotaAmount(quotaAmount int) {
    r.QuotaAmount = &quotaAmount
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r VerifyUserQuotaRequest) GetRegionId() string {
    return r.RegionId
}

type VerifyUserQuotaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result VerifyUserQuotaResult `json:"result"`
}

type VerifyUserQuotaResult struct {
    Success bool `json:"success"`
    Code int `json:"code"`
    Message string `json:"message"`
    Data quota.VerifyQuotaResVo `json:"data"`
    RequestId string `json:"requestId"`
}