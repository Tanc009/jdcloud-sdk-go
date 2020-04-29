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


type NodeConf struct {

    /* 磁盘类型，支持两种：zbs（SSD云硬盘）和local_ssd（本地SSD盘），默认为zbs (Optional) */
    StorageType []string `json:"storageType"`

    /* 限制条件 (Optional) */
    Constraints ConstraintsConf `json:"constraints"`
}