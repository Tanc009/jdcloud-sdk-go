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


type SetRiskPolicyReq struct {

    /* 规则id,新增时传0 (Optional) */
    Id *int `json:"id"`

    /* WAF实例id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* 规则名称  */
    Name string `json:"name"`

    /* eventCode  */
    EventCode string `json:"eventCode"`

    /* desc  */
    Desc string `json:"desc"`

    /* 0-使用中 1-禁用  */
    Disable int `json:"disable"`

    /* 策略编排逻辑, 格式：1&2&(3\|4\|5)  */
    Logic string `json:"logic"`

    /* 策略规则  */
    Rules *RiskPolicyRuleCfg `json:"rules"`

    /* 动作 支持 verify@captcha / verify@jscookie / forbidden / notice / redirect  */
    Action string `json:"action"`

    /* 跳转地址，Action为redirect时必须为当前实例下的域名的url，forbidden时为自定义页面名称 (Optional) */
    Redirection *string `json:"redirection"`
}
