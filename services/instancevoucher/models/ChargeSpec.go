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


type ChargeSpec struct {

    /* 计费单位，取值为：month、year，默认为month  */
    ChargeUnit string `json:"chargeUnit"`

    /* 计费时长。当chargeUnit为month时取值为：1 ~ 9，当chargeUnit为year时取值为：1、2、3  */
    ChargeDuration int `json:"chargeDuration"`

    /* True=：OPEN——开通自动续费、False=CLOSE—— 不开通自动续费，默认为CLOSE (Optional) */
    AutoRenew *bool `json:"autoRenew"`
}