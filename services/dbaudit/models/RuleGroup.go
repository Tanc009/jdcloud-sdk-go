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


type RuleGroup struct {

    /* 规则组ID(新建规则组时不需要传递此值) (Optional) */
    RuleGroupId *string `json:"ruleGroupId"`

    /* 规则组名称 (Optional) */
    RuleGroupName *string `json:"ruleGroupName"`

    /* 数据库用户是否区分大小写 (Optional) */
    DbUserCase *bool `json:"dbUserCase"`

    /* 操作系统用户是否区分大小写 (Optional) */
    OsUserCase *bool `json:"osUserCase"`

    /* 客户端程序是否区分大小写 (Optional) */
    ProgramCase *bool `json:"programCase"`

    /* 规则组是否启用 (Optional) */
    Enabled *bool `json:"enabled"`

    /* 标识从哪个规则组复制而来 (Optional) */
    CopyFromId *int `json:"copyFromId"`
}
