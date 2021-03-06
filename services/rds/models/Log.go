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


type Log struct {

    /* 日志文件id (Optional) */
    Id string `json:"id"`

    /* 日志文件名称 (Optional) */
    Name string `json:"name"`

    /* 日志文件大小，单位Byte (Optional) */
    SizeByte int `json:"sizeByte"`

    /* 日志最后更改时间 (Optional) */
    LastModified string `json:"lastModified"`

    /* 公网下载链接 (Optional) */
    PublicURL string `json:"publicURL"`

    /* 内网下载链接 (Optional) */
    InternalURL string `json:"internalURL"`
}
