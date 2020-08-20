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


type RsConfig struct {

    /* 回源地址  */
    RsAddr []string `json:"rsAddr"`

    /* http回源端口 (Optional) */
    HttpRsPort []string `json:"httpRsPort"`

    /* https回源端口 (Optional) */
    HttpsRsPort []string `json:"httpsRsPort"`

    /* 回源地址类型，0为ip，1为域名 (Optional) */
    RsType *int `json:"rsType"`
}
