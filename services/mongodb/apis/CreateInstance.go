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
    mongodb "github.com/jdcloud-api/jdcloud-sdk-go/services/mongodb/models"
    charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"
)

type CreateInstanceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例规格  */
    InstanceSpec *mongodb.DBInstanceSpec `json:"instanceSpec"`

    /* 付费方式 (Optional) */
    ChargeSpec *charge.ChargeSpec `json:"chargeSpec"`
}

/*
 * param regionId: Region ID 
 * param instanceSpec: 实例规格 
 * param chargeSpec: 付费方式 (Optional)
 */
func NewCreateInstanceRequest(
    regionId string,
    instanceSpec *mongodb.DBInstanceSpec,
) *CreateInstanceRequest {

	return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceSpec: instanceSpec,
	}
}

func (r *CreateInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *CreateInstanceRequest) SetInstanceSpec(instanceSpec *mongodb.DBInstanceSpec) {
    r.InstanceSpec = instanceSpec
}

func (r *CreateInstanceRequest) SetChargeSpec(chargeSpec *charge.ChargeSpec) {
    r.ChargeSpec = chargeSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateInstanceResult `json:"result"`
}

type CreateInstanceResult struct {
    InstanceId string `json:"instanceId"`
    OrderId string `json:"orderId"`
}