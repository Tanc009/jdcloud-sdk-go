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


type ConstraintsConf struct {

    /* 磁盘类型，zbs表示SSD云硬盘， local_ssd表示本地SSD盘 (Optional) */
    StorageType string `json:"storageType"`

    /* 是否在售 (Optional) */
    OnSale bool `json:"onSale"`

    /* 规格代码，规格代码详情参见：https://docs.jdcloud.com/cn/jcs-for-elasticsearch/specifications (Optional) */
    ClassCode []string `json:"classCode"`

    /* 节点数最小值 (Optional) */
    MinCount int `json:"minCount"`

    /* 节点数最大值 (Optional) */
    MaxCount int `json:"maxCount"`

    /* 节点数默认值 (Optional) */
    DefaultCount int `json:"defaultCount"`

    /* 节点数步长 (Optional) */
    StepCount int `json:"stepCount"`

    /* 是否包含存储 (Optional) */
    StorageScale bool `json:"storageScale"`

    /* 存储最小值 (Optional) */
    MinStorageGB int `json:"minStorageGB"`

    /* 存储最大值 (Optional) */
    MaxStorageGB int `json:"maxStorageGB"`

    /* 存储默认值 (Optional) */
    DefaultStorageGB int `json:"defaultStorageGB"`

    /* 存储步长 (Optional) */
    StepStorageGB int `json:"stepStorageGB"`
}
