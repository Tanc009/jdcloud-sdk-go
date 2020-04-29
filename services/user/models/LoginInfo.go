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


type LoginInfo struct {

    /* 身份类型： 1-主账号；2-子账号；3-角色；0-无(用户未登陆) (Optional) */
    LoginType int `json:"loginType"`

    /* 当前登录用户（type=1 或2时，此字段生效） (Optional) */
    Pin string `json:"pin"`

    /* 当前登录用户的主账号（type=2时，此字段生效） (Optional) */
    AdminPin string `json:"adminPin"`

    /* 角色身份信息(type=3时，此字段生效) (Optional) */
    CredentialInfo CredentialInfo `json:"credentialInfo"`

    /* 产品线跳转的url（type=0时，此字段生效） (Optional) */
    LoginUrl string `json:"loginUrl"`

    /* 主账号登录名（type=1，此字段不为空） (Optional) */
    LoginName string `json:"loginName"`
}