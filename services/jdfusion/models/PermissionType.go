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


type PermissionType struct {

    /* 云注册信息ID (Optional) */
    CloudID string `json:"cloudID"`

    /* IP协议 (Optional) */
    IpProtocol string `json:"ipProtocol"`

    /* 端口范围 (Optional) */
    PortRange string `json:"portRange"`

    /* 描述信息 (Optional) */
    Description string `json:"description"`

    /* 源IP地址段，用于入方向授权 (Optional) */
    SourceCidrIp string `json:"sourceCidrIp"`

    /* 源安全组，用于入方向授权 (Optional) */
    SourceGroupId string `json:"sourceGroupId"`

    /* 源安全组所属阿里云账户Id (Optional) */
    SourceGroupOwnerAccount string `json:"sourceGroupOwnerAccount"`

    /* 目标IP地址段，用于出方向授权 (Optional) */
    DestCidrIp string `json:"destCidrIp"`

    /* 目标安全组，用于出方向授权 (Optional) */
    DestGroupId string `json:"destGroupId"`

    /* 目标安全组所属阿里云账户Id (Optional) */
    DestGroupOwnerAccount string `json:"destGroupOwnerAccount"`

    /* 授权策略 (Optional) */
    Policy string `json:"policy"`

    /* 网络类型 (Optional) */
    NicType string `json:"nicType"`

    /* 规则优先级 (Optional) */
    Priority string `json:"priority"`

    /* 授权方向 (Optional) */
    Direction string `json:"direction"`
}
