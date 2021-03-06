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


type SendRecord struct {

    /* 手机号码 (Optional) */
    SendNumber string `json:"sendNumber"`

    /* 短信内容 (Optional) */
    SmsContent string `json:"smsContent"`

    /* 短信字数 (Optional) */
    ContentLength int `json:"contentLength"`

    /* 折成条数 (Optional) */
    ChargeCount int `json:"chargeCount"`

    /* 短信类型 短信类型，1 通道短信 2 官方短信 (Optional) */
    PackageType int `json:"packageType"`

    /* 发送时间 (Optional) */
    SendTime string `json:"sendTime"`

    /* 发送状态 (Optional) */
    SendStatus string `json:"sendStatus"`
}
