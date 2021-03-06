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

package models


type Ip struct {

    /* 机房英文标识 (Optional) */
    Idc string `json:"idc"`

    /* 机房名称 (Optional) */
    IdcName string `json:"idcName"`

    /* 公网IP实例ID (Optional) */
    IpId string `json:"ipId"`

    /* IP地址段 (Optional) */
    CidrAddr string `json:"cidrAddr"`

    /* IP数量 (Optional) */
    IpQuantity string `json:"ipQuantity"`

    /* IP类型 IPV4/IPV6 (Optional) */
    IpType string `json:"ipType"`

    /* 网络位地址 (Optional) */
    NetworkAddr string `json:"networkAddr"`

    /* 网关地址 (Optional) */
    GatewayAddr string `json:"gatewayAddr"`

    /* 广播地址 (Optional) */
    BroadcastAddr string `json:"broadcastAddr"`

    /* 线路类型 dynamicBGP:动态BGP thirdLineBGP:三线BGP telecom:电信单线 unicom:联通单线 mobile:移动单线 (Optional) */
    LineType string `json:"lineType"`

    /* 状态 normal:正常 abnormal:异常 (Optional) */
    Status string `json:"status"`
}
