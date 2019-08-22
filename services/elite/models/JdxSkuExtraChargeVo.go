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


type JdxSkuExtraChargeVo struct {

    /* 额外计费项名称 (Optional) */
    ExtraChargeName string `json:"extraChargeName"`

    /* 额外计费项单位 (Optional) */
    ExtraChargeUnit string `json:"extraChargeUnit"`

    /* 售价 (Optional) */
    SellingPrice int `json:"sellingPrice"`

    /* 1、范围 2、枚举 (Optional) */
    NumType int `json:"numType"`

    /* 1,100逗号分隔,numType=1表示可购买数量的范围,numType=2表示只支持购买特定数量 (Optional) */
    Num string `json:"num"`
}
