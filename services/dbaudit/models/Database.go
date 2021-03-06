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


type Database struct {

    /* 数据库ID (Optional) */
    DbId string `json:"dbId"`

    /* 数据库名称,长度限制32字节,允许英文字母,数字,下划线,中划线和中文 (Optional) */
    DbName string `json:"dbName"`

    /* 数据库的描述,长度限制128字节 (Optional) */
    DbDesc string `json:"dbDesc"`

    /* 数据库端口 (Optional) */
    DbPort int `json:"dbPort"`

    /* 数据库版本 (Optional) */
    DbVersion string `json:"dbVersion"`

    /* 数据库类型: 
0->Oracle
1->SQLServer
2->DB2
3->MySQL
4->Cache
5->Sybase
6->DM7
7->Informix
9->人大金仓
10->Teradata
11->Postgresql
12->Gbase
16->Hive
17->MongoDB
 (Optional) */
    DbType int `json:"dbType"`

    /* 审计端口 (Optional) */
    AgentPort int `json:"agentPort"`

    /* 数据库地址，如 ip:port 或 域名:port (Optional) */
    DbAddr string `json:"dbAddr"`

    /* 数据库启用状态, 0 停用 1 运行中 2 创建中 3 创建失败 (Optional) */
    State int `json:"state"`
}
