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


type VmInfo struct {

    /* 资源ID，如果为空，则执行创建操作，否则执行修改操作 (Optional) */
    Id string `json:"id"`

    /* 可用区,根据各云平台规范填写 (Optional) */
    Region string `json:"region"`

    /* 云主机所属的可用区 (Optional) */
    Az string `json:"az"`

    /* 云主机名称 (Optional) */
    Name string `json:"name"`

    /* 云主机 (Optional) */
    HostName string `json:"hostName"`

    /*  (Optional) */
    ImageType ImageType `json:"imageType"`

    /*  (Optional) */
    InstanceType InstanceType `json:"instanceType"`

    /* 云主机描述 (Optional) */
    Description string `json:"description"`

    /* 子网ID (Optional) */
    SubnetId string `json:"subnetId"`

    /*  (Optional) */
    Tags []Tag `json:"tags"`

    /* 所属云提供商ID (Optional) */
    CloudID string `json:"cloudID"`

    /* 密钥对名称,jd当前只支持传入一个 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 主网卡主IP绑定弹性IP的地址 (Optional) */
    ElasticIpAddress string `json:"elasticIpAddress"`

    /* 私有ip地址 (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* 云主机状态 (Optional) */
    Status string `json:"status"`

    /* 创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 镜像ID (Optional) */
    ImageId string `json:"imageId"`

    /* 安全组ID (Optional) */
    SecurityGroupIds []string `json:"securityGroupIds"`
}
