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


type TaskInfo struct {

    /* 报表任务ID (Optional) */
    TaskId string `json:"taskId"`

    /* 报表任务名称，长度限制32字节，允许英文字母,数字,下划线,中划线和中文 (Optional) */
    TaskName string `json:"taskName"`

    /* 报表任务描述，长度限制128字节 (Optional) */
    TaskDesc string `json:"taskDesc"`

    /* 报表任务状态(0 停止 1 运行中 2 一次性任务) (Optional) */
    TaskState int `json:"taskState"`

    /* 数据库审计实例ID (Optional) */
    InsId string `json:"insId"`

    /* 审计数据库ID(默认为空,代表全部数据据库) (Optional) */
    DbId string `json:"dbId"`

    /* 0,1,2,3,4,5,6,7,8(0:立即实行,1-7为每周特定日期执行,8为每天执行) (Optional) */
    ExecDate int `json:"execDate"`

    /* 报表创建时间 (Optional) */
    CreateTime string `json:"createTime"`
}
