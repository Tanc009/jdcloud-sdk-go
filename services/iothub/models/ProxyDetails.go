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


type ProxyDetails struct {

    /* IoT Hub Proxy实例编号 (Optional) */
    ProxyId string `json:"proxyId"`

    /* Proxy对应的用户Pin (Optional) */
    UserPin string `json:"userPin"`

    /* IoT Hub Proxy实例创建时间 (Optional) */
    CreateTime int `json:"createTime"`

    /* IoT Hub Proxy所在区域编号 (Optional) */
    RegionId string `json:"regionId"`

    /* IoT Hub Proxy所在区域名称 (Optional) */
    RegionName string `json:"regionName"`

    /* IoT Hub Proxy所在可用区编号 (Optional) */
    AzId string `json:"azId"`

    /* IoT Hub Proxy所在可用区名称 (Optional) */
    AzName string `json:"azName"`

    /* IoT Hub Proxy所在VPC编号 (Optional) */
    VpcId string `json:"vpcId"`

    /* IoT Hub Proxy所在VPC名称 (Optional) */
    VpcName string `json:"vpcName"`

    /* IoT Hub Proxy所在subnet编号 (Optional) */
    SubnetId string `json:"subnetId"`

    /* IoT Hub Proxy所在subnet名称 (Optional) */
    SubnetName string `json:"subnetName"`

    /* 内部创建JCQ对应的accessKey (Optional) */
    JcqAccessKey string `json:"jcqAccessKey"`

    /* 内部创建JCQ对应的Secret Access Key (Optional) */
    JcqSecretAccessKey string `json:"jcqSecretAccessKey"`

    /* 内部创建JCQ对应的ConsumerGroup (Optional) */
    JcqConsumerGroupId string `json:"jcqConsumerGroupId"`

    /* 内部创建JCQ对应的endpoint (Optional) */
    JcqEndpoint string `json:"jcqEndpoint"`

    /* 当前Proxy中负责处理的规则总数 (Optional) */
    TotalRuleNums string `json:"totalRuleNums"`
}
