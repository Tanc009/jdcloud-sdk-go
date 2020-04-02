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


type InstanceSpec struct {

    /* 私有网络vpcId  */
    VpcId string `json:"vpcId"`

    /* 子网subnetId  */
    SubnetId string `json:"subnetId"`

    /* ipVersion，空(代表v4)或者v4&v6 (Optional) */
    IpVersion *string `json:"ipVersion"`

    /* kafka版本，当前支持1.0.2  */
    InstanceVersion string `json:"instanceVersion"`

    /* kafka集群名称，不可为空，只支持大小写字母、数字、英文下划线或者中划线，以字母开头且不能超过32位  */
    InstanceName string `json:"instanceName"`

    /* azId  */
    AzId []string `json:"azId"`

    /* 集群规格配置  */
    InstanceClassSpec []InstanceClassSpec `json:"instanceClassSpec"`

    /* 扩展配置 (Optional) */
    Extension *ReqExtension `json:"extension"`
}
