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


type WarningCfg struct {

    /* 规则id (Optional) */
    Id int `json:"id"`

    /* 用户名 (Optional) */
    UserPin string `json:"userPin"`

    /* 告警类型 (Optional) */
    WarnType string `json:"warnType"`

    /* 是否生效，0-不生效，1-生效 (Optional) */
    Enable int `json:"enable"`

    /* 检测周期 (Optional) */
    DetectSpan int `json:"detectSpan"`

    /* 告警阈值 (Optional) */
    DetectThreshold int `json:"detectThreshold"`

    /* 告警方式 (Optional) */
    ContactWays []string `json:"contactWays"`

    /* 告警通知人 (Optional) */
    ContactorPersons []ContactPerson `json:"contactorPersons"`

    /* 告警通知群组 (Optional) */
    ContactorGroups []ContactGroup `json:"contactorGroups"`
}
