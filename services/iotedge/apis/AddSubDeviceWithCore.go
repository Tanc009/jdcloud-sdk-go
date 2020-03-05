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
    iotedge "github.com/jdcloud-api/jdcloud-sdk-go/services/iotedge/models"
)

type AddSubDeviceWithCoreRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* IoTCore实例编号  */
    InstanceId string `json:"instanceId"`

    /* Edge名称  */
    EdgeName string `json:"edgeName"`

    /* Edge的ProductKey  */
    ProductKey string `json:"productKey"`

    /* 待添加的子设备列表 (Optional) */
    Devices []iotedge.AddDevices `json:"devices"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoTCore实例编号 (Required)
 * param edgeName: Edge名称 (Required)
 * param productKey: Edge的ProductKey (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddSubDeviceWithCoreRequest(
    regionId string,
    instanceId string,
    edgeName string,
    productKey string,
) *AddSubDeviceWithCoreRequest {

	return &AddSubDeviceWithCoreRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/edges/{edgeName}:addSubDevice",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        EdgeName: edgeName,
        ProductKey: productKey,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoTCore实例编号 (Required)
 * param edgeName: Edge名称 (Required)
 * param productKey: Edge的ProductKey (Required)
 * param devices: 待添加的子设备列表 (Optional)
 */
func NewAddSubDeviceWithCoreRequestWithAllParams(
    regionId string,
    instanceId string,
    edgeName string,
    productKey string,
    devices []iotedge.AddDevices,
) *AddSubDeviceWithCoreRequest {

    return &AddSubDeviceWithCoreRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/edges/{edgeName}:addSubDevice",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        EdgeName: edgeName,
        ProductKey: productKey,
        Devices: devices,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddSubDeviceWithCoreRequestWithoutParam() *AddSubDeviceWithCoreRequest {

    return &AddSubDeviceWithCoreRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/products/{productKey}/edges/{edgeName}:addSubDevice",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *AddSubDeviceWithCoreRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: IoTCore实例编号(Required) */
func (r *AddSubDeviceWithCoreRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param edgeName: Edge名称(Required) */
func (r *AddSubDeviceWithCoreRequest) SetEdgeName(edgeName string) {
    r.EdgeName = edgeName
}

/* param productKey: Edge的ProductKey(Required) */
func (r *AddSubDeviceWithCoreRequest) SetProductKey(productKey string) {
    r.ProductKey = productKey
}

/* param devices: 待添加的子设备列表(Optional) */
func (r *AddSubDeviceWithCoreRequest) SetDevices(devices []iotedge.AddDevices) {
    r.Devices = devices
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddSubDeviceWithCoreRequest) GetRegionId() string {
    return r.RegionId
}

type AddSubDeviceWithCoreResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddSubDeviceWithCoreResult `json:"result"`
}

type AddSubDeviceWithCoreResult struct {
}