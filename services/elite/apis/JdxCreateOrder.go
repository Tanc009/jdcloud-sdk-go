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
    elite "github.com/jdcloud-api/jdcloud-sdk-go/services/elite/models"
)

type JdxCreateOrderRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 下单信息  */
    CreateOrderInfo *elite.CreateOrderInfo `json:"createOrderInfo"`
}

/*
 * param regionId: 地域ID (Required)
 * param createOrderInfo: 下单信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewJdxCreateOrderRequest(
    regionId string,
    createOrderInfo *elite.CreateOrderInfo,
) *JdxCreateOrderRequest {

	return &JdxCreateOrderRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/jdxCreateOrder",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CreateOrderInfo: createOrderInfo,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param createOrderInfo: 下单信息 (Required)
 */
func NewJdxCreateOrderRequestWithAllParams(
    regionId string,
    createOrderInfo *elite.CreateOrderInfo,
) *JdxCreateOrderRequest {

    return &JdxCreateOrderRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jdxCreateOrder",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CreateOrderInfo: createOrderInfo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewJdxCreateOrderRequestWithoutParam() *JdxCreateOrderRequest {

    return &JdxCreateOrderRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/jdxCreateOrder",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *JdxCreateOrderRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param createOrderInfo: 下单信息(Required) */
func (r *JdxCreateOrderRequest) SetCreateOrderInfo(createOrderInfo *elite.CreateOrderInfo) {
    r.CreateOrderInfo = createOrderInfo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r JdxCreateOrderRequest) GetRegionId() string {
    return r.RegionId
}

type JdxCreateOrderResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result JdxCreateOrderResult `json:"result"`
}

type JdxCreateOrderResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
    Data elite.CreateOrderResultVo `json:"data"`
}