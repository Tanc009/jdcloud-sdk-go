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


type SystemLogHistorySpec struct {

    /*   */
    AppName string `json:"appName"`

    /* 精确匹配，true 或者 false (Optional) */
    ExactMatch bool `json:"exactMatch"`

    /*   */
    Instance string `json:"instance"`

    /* 查询关键字 (Optional) */
    Keyword string `json:"keyword"`

    /*   */
    LogType string `json:"logType"`

    /* 排序，取值"ASC"或"DESC"，默认"ASC" (Optional) */
    Order string `json:"order"`

    /* 页数，从1开始 (Optional) */
    Page int64 `json:"page"`

    /* 每页日志条数 (Optional) */
    Size int64 `json:"size"`

    /*   */
    Time TimestampRange `json:"time"`
}
