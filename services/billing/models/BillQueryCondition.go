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


type BillQueryCondition struct {

    /* 查询类别   1：资源账单   2：消费记录  */
    QueryType int `json:"queryType"`

    /* 用户pin  */
    Pin string `json:"pin"`

    /* appCode  */
    AppCode string `json:"appCode"`

    /* serviceCode  */
    ServiceCode string `json:"serviceCode"`

    /* 计费类型  */
    BillingType int `json:"billingType"`

    /* 支付类型  */
    PayType int `json:"payType"`

    /* 支付状态  */
    PayState int `json:"payState"`

    /* 1按账期、2按消费时间  */
    TimeType int `json:"timeType"`

    /* 开始时间  */
    StartTime string `json:"startTime"`

    /* 结束时间  */
    EndTime string `json:"endTime"`

    /* 是否忽略0元账单  */
    IgnoreZero int `json:"ignoreZero"`

    /* 资源id  */
    ResourceId string `json:"resourceId"`

    /* 站点  */
    Site int `json:"site"`

    /* 角色  */
    Role int `json:"role"`

    /* 区域  */
    Region string `json:"region"`
}
