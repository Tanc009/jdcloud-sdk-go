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


type GetEsLogReq struct {

    /* 实例id，代表要查询的WAF实例，为空时表示当前用户下的所有实例 (Optional) */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名，为空时表示当前实例下的所有域名 (Optional) */
    Domain string `json:"domain"`

    /* 来源ip，检索字段 (Optional) */
    RemoteAddr string `json:"remoteAddr"`

    /* 动作，检索字段 (Optional) */
    Action string `json:"action"`

    /* 状态码，检索字段 (Optional) */
    Status string `json:"status"`

    /*  (Optional) */
    LogType []string `json:"logType"`

    /* 日志Id，检索字段 (Optional) */
    LogId string `json:"logId"`

    /* 开始时间戳，单位毫秒，时间间隔要求大于5分钟，小于30天。  */
    Start int `json:"start"`

    /* 结束时间戳，单位毫秒，时间间隔要求大于5分钟，小于30天。  */
    End int `json:"end"`

    /* 每页显示的个数，默认是10。 (Optional) */
    PageSize int `json:"pageSize"`

    /* 页数，默认是1。 (Optional) */
    PageIndex int `json:"pageIndex"`
}
