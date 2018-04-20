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

package models


type VpcPeering struct {

    /* VpcPeering的Id (Optional) */
    VpcPeeringId string `json:"vpcPeeringId"`

    /* VpcPeering名称，同账号下不允许重名，取值范围：1-32个中文、英文大小写的字母、数字和下划线分隔符 (Optional) */
    VpcPeeringName string `json:"vpcPeeringName"`

    /* 状态，取值为Connected，Disconnected，Initiated (Optional) */
    VpcPeeringState string `json:"vpcPeeringState"`

    /* VpcPeering 描述，可为空值，取值范围：0-256个中文、英文大小写的字母、数字和下划线分隔符 (Optional) */
    Description string `json:"description"`

    /* 发起VpcPeering的Vpc信息 (Optional) */
    VpcInfo VpcPeeringVpcInfo `json:"vpcInfo"`

    /* 对端的Vpc信息 (Optional) */
    RemoteVpcInfo VpcPeeringVpcInfo `json:"remoteVpcInfo"`

    /* VpcPeering创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`
}
