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


type DeviceOnlineInfo struct {

    /* 设备ID (Optional) */
    Device_id string `json:"device_id"`

    /* 在线时长 (Optional) */
    Online_time int `json:"online_time"`

    /* 上线次数 (Optional) */
    Count_time int `json:"count_time"`

    /* 连接码 (Optional) */
    Identifier string `json:"identifier"`

    /* 设备状态 (Optional) */
    Status int `json:"status"`

    /* 采集器类型 (Optional) */
    CollDeviceType string `json:"collDeviceType"`
}
