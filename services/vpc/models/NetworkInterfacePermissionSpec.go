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


type NetworkInterfacePermissionSpec struct {

    /* 弹性网卡ID  */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 授信用户,需要存在于京东云许可的服务账号名单中 (Optional) */
    TrustUser string `json:"trustUser"`

    /* 授权策略, 授权后，该弹性网卡可以关联的服务端账号的资源类型，取值范围，instance-attach：可以关联服务端账号的实例资源，eip-associate：可以关联服务端账号的弹性公网IP资源 (Optional) */
    Policy []string `json:"policy"`
}
