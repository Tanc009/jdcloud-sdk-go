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


type CabinetOverview struct {

    /* 机柜总数目 (Optional) */
    Sum int `json:"sum"`

    /* 已开通机柜数目 (Optional) */
    Enabled int `json:"enabled"`

    /* 未开通机柜数目 (Optional) */
    Disabled int `json:"disabled"`

    /* 开通中机柜数目 (Optional) */
    Enabling int `json:"enabling"`

    /* 关闭中机柜数目 (Optional) */
    Disabling int `json:"disabling"`
}