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


type InstanceClass struct {

    /* 实例规格代码,参见实例规格代码表 (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* cpu (Optional) */
    Cpu int `json:"cpu"`

    /* 内存 (Optional) */
    MemoryMB int `json:"memoryMB"`

    /* 磁盘 (Optional) */
    DiskGB int `json:"diskGB"`

    /* 最大链接数 (Optional) */
    MaxConnetction int `json:"maxConnetction"`

    /* 带宽 (Optional) */
    BandwidthMbps int `json:"bandwidthMbps"`
}
