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

import . "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/models"

type InstanceDiskAttachment struct {

    /* 磁盘分类，取值范围{local, cloud} (Optional) */
    DiskCategory string `json:"diskCategory"`

    /* 自动删除，删除主机时自动删除此磁盘，默认为True (Optional) */
    AutoDelete bool `json:"autoDelete"`

    /* 本地磁盘 (Optional) */
    LocalDisk LocalDisk `json:"localDisk"`

    /* 云硬盘 (Optional) */
    CloudDisk Disk `json:"cloudDisk"`

    /* 数据盘逻辑挂载点vdb,vdc,vdd,vde,vdf,vdg,vdh (Optional) */
    DeviceName string `json:"deviceName"`
}
