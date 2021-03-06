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


type DataResult struct {

    /* 文本类型为检测内容，图片类型为图片短链 (Optional) */
    Content string `json:"content"`

    /* taskId (Optional) */
    TaskId string `json:"taskId"`

    /* dataId (Optional) */
    DataId string `json:"dataId"`

    /* 送审时间 2019-12-18 16:02:19（北京时间UTC+8） (Optional) */
    Time string `json:"time"`

    /* 日志落盘时间 2019-12-18 16:02:19（北京时间UTC+8） (Optional) */
    LogTime string `json:"logTime"`

    /* 识别结果 (Optional) */
    Result string `json:"result"`

    /* 场景_结果 格式。 (Optional) */
    Details []string `json:"details"`

    /* 响应结果 (Optional) */
    Response string `json:"response"`

    /* 返回时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 状态码 (Optional) */
    Code string `json:"code"`

    /* 图片/音频/视频的url (Optional) */
    Url string `json:"url"`

    /* 视频截帧 (Optional) */
    Frame FrameCfg `json:"frame"`

    /* 人工审核结果，空表示没有审核 (Optional) */
    FbSuggestion string `json:"fbSuggestion"`
}
