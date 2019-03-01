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


type DeploymentResourcesInfo struct {

    /*  (Optional) */
    Vms *interface{} `json:"vms"`

    /*  (Optional) */
    Eips *interface{} `json:"eips"`

    /*  (Optional) */
    Subnets *interface{} `json:"subnets"`

    /*  (Optional) */
    NetworkInterfaces *interface{} `json:"networkInterfaces"`

    /*  (Optional) */
    Slbs *interface{} `json:"slbs"`

    /*  (Optional) */
    SecurityGroups *interface{} `json:"securityGroups"`

    /*  (Optional) */
    Keypairs *interface{} `json:"keypairs"`

    /*  (Optional) */
    Disks *interface{} `json:"disks"`

    /*  (Optional) */
    Vpcs *interface{} `json:"vpcs"`

    /*  (Optional) */
    VserverGroups *interface{} `json:"vserverGroups"`

    /*  (Optional) */
    HttpListeners *interface{} `json:"httpListeners"`

    /*  (Optional) */
    DiskAttachment *interface{} `json:"diskAttachment"`

    /*  (Optional) */
    NetInterfaceAttachment *interface{} `json:"netInterfaceAttachment"`

    /*  (Optional) */
    EipAssociate *interface{} `json:"eipAssociate"`

    /*  (Optional) */
    Variables *interface{} `json:"variables"`
}