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


type ClusterListNode struct {

    /* 集群ID (Optional) */
    Id string `json:"id"`

    /* 集群名称 (Optional) */
    Name string `json:"name"`

    /* 集群所属地域 (Optional) */
    DataCenter string `json:"dataCenter"`

    /* 集群ID (Optional) */
    RecordId int `json:"recordId"`

    /* 监控ID (Optional) */
    MonitorResourceId string `json:"monitorResourceId"`

    /* 集群状态 (Optional) */
    Status string `json:"status"`

    /* 错误信息 (Optional) */
    ErrorMessage string `json:"errorMessage"`

    /* 集群创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 集群收费类型 (Optional) */
    PayType string `json:"payType"`

    /* 集群运行时间 (Optional) */
    Duration string `json:"duration"`

    /* 公网Ip (Optional) */
    OuterIp string `json:"outerIp"`
}
