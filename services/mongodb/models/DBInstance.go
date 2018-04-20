// Copyright 2018-2025 JDCLOUD.COM
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

import . "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type DBInstance struct {

    /* 实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 数据库类型 (Optional) */
    Engine string `json:"engine"`

    /* 数据库版本 (Optional) */
    EngineVersion string `json:"engineVersion"`

    /* 实例规格代码 (Optional) */
    InstanceClass string `json:"instanceClass"`

    /* 存储空间 (Optional) */
    InstanceStorageGB int `json:"instanceStorageGB"`

    /* CPU核数 (Optional) */
    InstanceCPU int `json:"instanceCPU"`

    /* 内存，单位GB (Optional) */
    InstanceMemoryGB int `json:"instanceMemoryGB"`

    /* 可取区ID，依次为主、从、隐藏节点所在可用区 (Optional) */
    AzId []string `json:"azId"`

    /* VPCID (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网ID (Optional) */
    SubnetId string `json:"subnetId"`

    /* 副本集名称 (Optional) */
    ReplicaSetName string `json:"replicaSetName"`

    /* 域名 (Optional) */
    InstanceDomain string `json:"instanceDomain"`

    /* 默认库名 (Optional) */
    DBName string `json:"dBName"`

    /* 默认用户名 (Optional) */
    AccountName string `json:"accountName"`

    /* 应用访问端口 (Optional) */
    InstancePort string `json:"instancePort"`

    /* 实例状态.RUNNING：运行, ERROR：错误 ,BUILDING：创建中, DELETING：删除中, RESTORING：恢复中, RESIZING：变配中 (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 自动备份保留时间 (Optional) */
    BackupRetentionPeriod int `json:"backupRetentionPeriod"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 自动备份时间，如：00:00-02:00，表示0点到2点进行数据库自动备份 (Optional) */
    PreferredBackupWindow string `json:"preferredBackupWindow"`

    /* 系统维护时间，如：00:00-02:00，表示0点到2点进行系统维护 (Optional) */
    PreferredmaintenanceWindow string `json:"preferredmaintenanceWindow"`

    /* 计费信息 (Optional) */
    Charge Charge `json:"charge"`
}
