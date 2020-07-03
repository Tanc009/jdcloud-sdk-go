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
    iotcore "github.com/jdcloud-api/jdcloud-sdk-go/services/iotcore/models"
)

type DescribeThingTypeListRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    InstanceId string `json:"instanceId"`

    /* 设备型号标识 (Optional) */
    DeviceMetaId *string `json:"deviceMetaId"`

    /* 设备型号名称 (Optional) */
    DeviceMetaName *string `json:"deviceMetaName"`

    /* 节点类型 (Optional) */
    NodeType *int `json:"nodeType"`

    /* 页码 (Optional) */
    PageNo *int `json:"pageNo"`

    /* 每页显示条数 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeThingTypeListRequest(
    regionId string,
    instanceId string,
) *DescribeThingTypeListRequest {

	return &DescribeThingTypeListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/coreinstances/{instanceId}/thingType:list",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例Id (Required)
 * param deviceMetaId: 设备型号标识 (Optional)
 * param deviceMetaName: 设备型号名称 (Optional)
 * param nodeType: 节点类型 (Optional)
 * param pageNo: 页码 (Optional)
 * param pageSize: 每页显示条数 (Optional)
 */
func NewDescribeThingTypeListRequestWithAllParams(
    regionId string,
    instanceId string,
    deviceMetaId *string,
    deviceMetaName *string,
    nodeType *int,
    pageNo *int,
    pageSize *int,
) *DescribeThingTypeListRequest {

    return &DescribeThingTypeListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/thingType:list",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DeviceMetaId: deviceMetaId,
        DeviceMetaName: deviceMetaName,
        NodeType: nodeType,
        PageNo: pageNo,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeThingTypeListRequestWithoutParam() *DescribeThingTypeListRequest {

    return &DescribeThingTypeListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/coreinstances/{instanceId}/thingType:list",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *DescribeThingTypeListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例Id(Required) */
func (r *DescribeThingTypeListRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param deviceMetaId: 设备型号标识(Optional) */
func (r *DescribeThingTypeListRequest) SetDeviceMetaId(deviceMetaId string) {
    r.DeviceMetaId = &deviceMetaId
}

/* param deviceMetaName: 设备型号名称(Optional) */
func (r *DescribeThingTypeListRequest) SetDeviceMetaName(deviceMetaName string) {
    r.DeviceMetaName = &deviceMetaName
}

/* param nodeType: 节点类型(Optional) */
func (r *DescribeThingTypeListRequest) SetNodeType(nodeType int) {
    r.NodeType = &nodeType
}

/* param pageNo: 页码(Optional) */
func (r *DescribeThingTypeListRequest) SetPageNo(pageNo int) {
    r.PageNo = &pageNo
}

/* param pageSize: 每页显示条数(Optional) */
func (r *DescribeThingTypeListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeThingTypeListRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeThingTypeListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeThingTypeListResult `json:"result"`
}

type DescribeThingTypeListResult struct {
    PageSize int `json:"pageSize"`
    CurrentPage int `json:"currentPage"`
    TotalCount int `json:"totalCount"`
    List []iotcore.ThingTypeInfoVO `json:"list"`
}