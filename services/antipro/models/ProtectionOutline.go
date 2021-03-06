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


type ProtectionOutline struct {

    /* 已防护天数 (Optional) */
    ProtectedDay int64 `json:"protectedDay"`

    /* 已防护 IP 数量 (Optional) */
    ProtectedIpCount int64 `json:"protectedIpCount"`

    /* 7 日攻击次数 (Optional) */
    WeekAttackCount int64 `json:"weekAttackCount"`

    /* 7 日攻击流量峰值 (Optional) */
    WeekAttackPeak float64 `json:"weekAttackPeak"`

    /* 7 日攻击流量单位 (Optional) */
    WeekAttackUnit string `json:"weekAttackUnit"`

    /* 30 日攻击次数 (Optional) */
    MonthAttackCount int64 `json:"monthAttackCount"`

    /* 30 日攻击流量峰值 (Optional) */
    MonthAttackPeak float64 `json:"monthAttackPeak"`

    /* 30 日攻击流量单位 (Optional) */
    MonthAttackUnit string `json:"monthAttackUnit"`

    /* 当前攻击数量 (Optional) */
    CurrentAttackCount int64 `json:"currentAttackCount"`
}
