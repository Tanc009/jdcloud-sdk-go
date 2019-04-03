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


type AggReceiptList struct {

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 应用Id (Optional) */
    AppId string `json:"appId"`

    /* 发送量 (Optional) */
    SendCount int `json:"sendCount"`

    /* 回执量 (Optional) */
    ReceiptCount int `json:"receiptCount"`

    /* 未回执量 (Optional) */
    NoReceiptCount int `json:"noReceiptCount"`

    /* 回执成功量 (Optional) */
    ReceiptSuccessCount int `json:"receiptSuccessCount"`

    /* 回执失败量 (Optional) */
    ReceiptFailureCount int `json:"receiptFailureCount"`

    /* 回执率 (Optional) */
    ReceiptRate string `json:"receiptRate"`

    /* 回执成功率 (Optional) */
    ReceiptSuccessRate string `json:"receiptSuccessRate"`

    /* 统计日期 (Optional) */
    AggTime string `json:"aggTime"`
}