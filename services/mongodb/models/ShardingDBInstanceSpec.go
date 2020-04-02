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


type ShardingDBInstanceSpec struct {

    /* 实例名称，名称只支持中文、数字、大小写字母及英文下划线“_”及中划线“-”，2-32个字符。 (Optional) */
    InstanceName *string `json:"instanceName"`

    /* 数据库类型，默认为MongoDB。 (Optional) */
    Engine *string `json:"engine"`

    /* 数据库版本，支持：3.4, 3.6；默认为3.6。 (Optional) */
    EngineVersion *string `json:"engineVersion"`

    /* mongos节点规格代码，必填。  */
    MongosNodeType string `json:"mongosNodeType"`

    /* mongos节点数量，支持2-32个，必填。 (Optional) */
    MongosNodeNumber *int `json:"mongosNodeNumber"`

    /* configserve节点规格代码，默认为mongo.m1.small。 (Optional) */
    ConfigserverNodeType *string `json:"configserverNodeType"`

    /* shard节点规格代码，必填。  */
    ShardNodeType string `json:"shardNodeType"`

    /* shard节点存储空间，单位GB，取值10-1000,10的倍数，必填。  */
    ShardNodeStorageGB int `json:"shardNodeStorageGB"`

    /* mongos节点数量，支持2-32个，必填。  */
    ShardNodeNumber int `json:"shardNodeNumber"`

    /* 是否选择多可用区部署，默认为否。 (Optional) */
    MultiAZ *bool `json:"multiAZ"`

    /* 必填。单可用区部署，填写1个可用区；多可用区部署，依次填写每个mongos所在可用区，数量与mognos节点数量一致。  */
    MongosAzId []string `json:"mongosAzId"`

    /* 必填。单可用区部署，填写1个可用区；多可用区部署，需填写3个可用区，依次为副本集的primary、secondary、hidden所在的可用区，将应用到分片集群的shard节点和configserver节点的全部副本集。  */
    ShardAzId []string `json:"shardAzId"`

    /* VPCID  */
    VpcId string `json:"vpcId"`

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* 密码，必须包含且只支持字母及数字，不少于8字符不超过16字符。 (Optional) */
    Password *string `json:"password"`

    /* 基于一个实例的备份创建新实例，如填写则restoreTime也需要填写。 (Optional) */
    OriginDBInstanceId *string `json:"originDBInstanceId"`

    /* 用户指定备份保留周期内的任意时间点，如2011-06-11T16:00:00Z。 (Optional) */
    RestoreTime *string `json:"restoreTime"`

    /* 存储类型，支持：LOCAL_SSD -本地盘SSD、EBS_SSD -云盘；默认值为：LOCAL_SSD。 (Optional) */
    InstanceStorageType *string `json:"instanceStorageType"`

    /* 实例数据加密(存储类型为云硬盘才支持数据加密). false：不加密; true：加密. 缺省为false. (Optional) */
    StorageEncrypted *bool `json:"storageEncrypted"`

    /* 跨域服务ID，用于跨地域按时间点创建实例 (Optional) */
    ServiceId *string `json:"serviceId"`
}
