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


type RemoteQuota struct {

    /* 产品线 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 资源名称 (Optional) */
    Name string `json:"name"`

    /* 配额总量 (Optional) */
    TotalInventory string `json:"totalInventory"`

    /* 已用配额 (Optional) */
    UsedAmount string `json:"usedAmount"`

    /* 父资源id (Optional) */
    ParentResourceId string `json:"parentResourceId"`

    /* 可用配额 (Optional) */
    AvailableAmount string `json:"availableAmount"`
}
