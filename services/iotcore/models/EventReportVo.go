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


type EventReportVo struct {

    /* 设备标识id (Optional) */
    DeviceId string `json:"deviceId"`

    /* 事件名称 (Optional) */
    DisplayName string `json:"displayName"`

    /* 事件key (Optional) */
    Key string `json:"key"`

    /* 消息id (Optional) */
    MessageId string `json:"messageId"`

    /* 事件上报内容 json格式 (Optional) */
    Parameters string `json:"parameters"`

    /* 事件上报时间 (Optional) */
    Timestamp int64 `json:"timestamp"`
}
