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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type RestoreDatabaseFromBackupRequest struct {

    core.JDCloudRequest

    /* 区域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 库名称  */
    DbName string `json:"dbName"`

    /* 备份ID  */
    BackupId string `json:""`

    /* 指定该备份中用于恢复数据库的文件名称  */
    BackupFileName string `json:""`
}

/*
 * param regionId: 区域代码 
 * param instanceId: 实例ID 
 * param dbName: 库名称 
 * param backupId: 备份ID 
 * param backupFileName: 指定该备份中用于恢复数据库的文件名称 
 */
func NewRestoreDatabaseFromBackupRequest(
    regionId string,
    instanceId string,
    dbName string,
    backupId string,
    backupFileName string,
) *RestoreDatabaseFromBackupRequest {

	return &RestoreDatabaseFromBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/databases/{dbName}:restoreDatabaseFromBackup",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DbName: dbName,
        BackupId: backupId,
        BackupFileName: backupFileName,
	}
}

func (r *RestoreDatabaseFromBackupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *RestoreDatabaseFromBackupRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *RestoreDatabaseFromBackupRequest) SetDbName(dbName string) {
    r.DbName = dbName
}

func (r *RestoreDatabaseFromBackupRequest) SetBackupId(backupId string) {
    r.BackupId = backupId
}

func (r *RestoreDatabaseFromBackupRequest) SetBackupFileName(backupFileName string) {
    r.BackupFileName = backupFileName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RestoreDatabaseFromBackupRequest) GetRegionId() string {
    return r.RegionId
}

type RestoreDatabaseFromBackupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RestoreDatabaseFromBackupResult `json:"result"`
}

type RestoreDatabaseFromBackupResult struct {
}