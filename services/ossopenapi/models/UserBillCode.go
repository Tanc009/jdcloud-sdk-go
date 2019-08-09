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


type UserBillCode struct {

    /* 用户的AppCode (Optional) */
    AppCode string `json:"appCode"`

    /* 用户的ServiceCode (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 用户的StorageResourceType (Optional) */
    StorageResourceType string `json:"storageResourceType"`

    /* 用户的ProcessResourceType (Optional) */
    ProcessResourceType string `json:"processResourceType"`
}