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


type FieldSpec struct {

    /* 字段名称  */
    FieldName string `json:"fieldName"`

    /* 加密字段  */
    EncryptField string `json:"encryptField"`

    /* 索引字段  */
    IndexField string `json:"indexField"`

    /* 是否保留明文字段,true:保留明文字段  */
    KeepPlainText bool `json:"keepPlainText"`
}
