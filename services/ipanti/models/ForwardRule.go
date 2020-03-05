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


type ForwardRule struct {

    /* 规则 Id (Optional) */
    Id string `json:"id"`

    /* 实例 Id (Optional) */
    InstanceId string `json:"instanceId"`

    /* TCP 或 UDP (Optional) */
    Protocol string `json:"protocol"`

    /* 规则的 CNAME (Optional) */
    Cname string `json:"cname"`

    /* 回源类型: ip 或者 domain (Optional) */
    OriginType string `json:"originType"`

    /* 高防 IP (Optional) */
    ServiceIp string `json:"serviceIp"`

    /* 端口号 (Optional) */
    Port int `json:"port"`

    /* 转发规则. <br>- wrr: 带权重的轮询<br>- rr:  不带权重的轮询<br>- sh:  源地址hash (Optional) */
    Algorithm string `json:"algorithm"`

    /*  (Optional) */
    OriginAddr []OriginAddrItem `json:"originAddr"`

    /* 备用的回源地址列表 (Optional) */
    OnlineAddr []string `json:"onlineAddr"`

    /* 回源域名 (Optional) */
    OriginDomain string `json:"originDomain"`

    /* 回源端口号 (Optional) */
    OriginPort int `json:"originPort"`

    /* 0: 防御状态<br>1: 回源状态 (Optional) */
    Status int `json:"status"`
}
