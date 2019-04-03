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


type DicDetailQuery struct {

    /* ID (Optional) */
    Id int `json:"id"`

    /* 字典类型 (Optional) */
    CodeType string `json:"codeType"`

    /* 字典编码 (Optional) */
    Code string `json:"code"`

    /* 字典编码名称 (Optional) */
    Name string `json:"name"`

    /* 字典编码值 (Optional) */
    Value string `json:"value"`

    /* null (Optional) */
    UseFlag bool `json:"useFlag"`

    /* 系统类型 (Optional) */
    SystemType string `json:"systemType"`

    /* 引用值 (Optional) */
    RefValue string `json:"refValue"`

    /* 顺序 (Optional) */
    Seq int `json:"seq"`

    /* 字典说明 (Optional) */
    Remark string `json:"remark"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 创建人 (Optional) */
    CreateUser string `json:"createUser"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 修改人 (Optional) */
    UpdateUser string `json:"updateUser"`

    /* 是否删除0未删除,1已删除 (Optional) */
    Yn int `json:"yn"`

    /* 当前页序号 (Optional) */
    PageIndex int `json:"pageIndex"`

    /* 每页结果数量 (Optional) */
    PageSize int `json:"pageSize"`

    /*  (Optional) */
    Offset int `json:"offset"`
}