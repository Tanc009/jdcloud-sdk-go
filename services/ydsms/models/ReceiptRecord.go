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


type ReceiptRecord struct {

    /* 短信回执量 (Optional) */
    ReceiptCount int64 `json:"receiptCount"`

    /* 短信未回执量 (Optional) */
    UnReceiptCount int64 `json:"unReceiptCount"`

    /* 短信回执率 (Optional) */
    ReceiptRate string `json:"receiptRate"`

    /* 短信回执成功量 (Optional) */
    ReceiptSuccessCount int64 `json:"receiptSuccessCount"`

    /* 短信回执失败量 (Optional) */
    ReceiptFailCount int64 `json:"receiptFailCount"`

    /* 短信回执成功率率 (Optional) */
    ReceiptSuccessRate string `json:"receiptSuccessRate"`

    /* 聚合开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 聚合截止时间 (Optional) */
    EndTime string `json:"endTime"`
}
