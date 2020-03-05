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


type AccountInfo struct {

    /* 组织名称  */
    OrgName string `json:"orgName"`

    /* 身份证号码 (Optional) */
    IdCard *string `json:"idCard"`

    /* 银行卡号  */
    BankCardNum string `json:"bankCardNum"`

    /* 银行名称  */
    BankName string `json:"bankName"`

    /* 支行名称  */
    BranchBankName string `json:"branchBankName"`

    /* 银行代码 (Optional) */
    BankCode *string `json:"bankCode"`

    /* 城市代码 (Optional) */
    CityCode *string `json:"cityCode"`

    /* 省份代码 (Optional) */
    ProvinceCode *string `json:"provinceCode"`
}
