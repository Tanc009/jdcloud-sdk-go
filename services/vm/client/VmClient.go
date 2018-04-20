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

package client

import (
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    . "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
    "encoding/json"
    "errors"
)

type VmClient struct {
    JDCloudClient
}

func NewVmClient(credential *Credential) *VmClient {
    if credential == nil {
        return nil
    }

    config := NewConfig()
    config.SetEndpoint("vm.jdcloud-api.com")

    return &VmClient{
        JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "vm",
            Revision:    "0.6.0",
            Logger:      NewDefaultLogger(LOG_INFO),
        }}
}

func (c *VmClient) SetConfig(config *Config) {
    c.Config = *config
}

func (c *VmClient) SetLogger(logger Logger) {
    c.Logger = logger
}

/* 查询（虚机、镜像、密钥、模板）配额 */
func (c *VmClient) DescribeQuotas(request *DescribeQuotasRequest) (*DescribeQuotasResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeQuotasResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 云主机挂载硬盘，主机和云盘没有未完成的任务时才可挂载，一个主机上最多可挂载4块数据盘 */
func (c *VmClient) AttachDisk(request *AttachDiskRequest) (*AttachDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &AttachDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* "虚机创建私有镜像"
"虚机状态必须为stopped"
"如果虚机上有挂载数据盘，默认会将数据盘创建快照，生成打包镜像"
"主机没有未完成的任务才可制作镜像"
 */
func (c *VmClient) CreateImage(request *CreateImageRequest) (*CreateImageResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &CreateImageResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询云主机列表 */
func (c *VmClient) DescribeInstances(request *DescribeInstancesRequest) (*DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 创建一台或多台指定配置的实例 */
func (c *VmClient) CreateInstances(request *CreateInstancesRequest) (*CreateInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &CreateInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询镜像信息 */
func (c *VmClient) DescribeImage(request *DescribeImageRequest) (*DescribeImageResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeImageResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询镜像共享帐户列表，不能操作非私有镜像 */
func (c *VmClient) DescribeImageMembers(request *DescribeImageMembersRequest) (*DescribeImageMembersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeImageMembersResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 修改主机密码，主机没有未完成的任务时才可操作 */
func (c *VmClient) ModifyInstancePassword(request *ModifyInstancePasswordRequest) (*ModifyInstancePasswordResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ModifyInstancePasswordResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 取消共享镜像，不能操作非私有镜像 */
func (c *VmClient) UnShareImage(request *UnShareImageRequest) (*UnShareImageResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &UnShareImageResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 修改主机信息 */
func (c *VmClient) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (*ModifyInstanceAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ModifyInstanceAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 重启单个实例，只能重启running状态的实例，主机没有未完成的任务才可重启 */
func (c *VmClient) RebootInstance(request *RebootInstanceRequest) (*RebootInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &RebootInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* "共享镜像，最多可共享给20个帐户"
"打包镜像暂不支持共享"
"不能操作非私有镜像"
"不能共享给自己"
 */
func (c *VmClient) ShareImage(request *ShareImageRequest) (*ShareImageResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ShareImageResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询镜像资源信息列表 */
func (c *VmClient) DescribeImages(request *DescribeImagesRequest) (*DescribeImagesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeImagesResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 云主机解绑公网IP 解绑的是主网卡、主内网IP对应的弹性IP */
func (c *VmClient) DisassociateElasticIp(request *DisassociateElasticIpRequest) (*DisassociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DisassociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 停止单个实例，只能停止running状态的实例，主机没有未完成的任务才可停止 */
func (c *VmClient) StopInstance(request *StopInstanceRequest) (*StopInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &StopInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询镜像限制 */
func (c *VmClient) DescribeImageConstraints(request *DescribeImageConstraintsRequest) (*DescribeImageConstraintsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeImageConstraintsResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询主机vnc */
func (c *VmClient) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (*DescribeInstanceVncUrlResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeInstanceVncUrlResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 启动单个实例，只能启动stopped状态的实例，主机没有未完成的任务才可启动 */
func (c *VmClient) StartInstance(request *StartInstanceRequest) (*StartInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &StartInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 删除私有镜像 */
func (c *VmClient) DeleteImage(request *DeleteImageRequest) (*DeleteImageResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DeleteImageResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 云主机绑定公网IP 绑定的是主网卡、主内网IP对应的弹性IP */
func (c *VmClient) AssociateElasticIp(request *AssociateElasticIpRequest) (*AssociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &AssociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* "删除单个实例"
"主机状态必须为停止状态、同时主机没有未完成的任务才可删除"
"包年包月未到期的主机不能删除"
"如果主机中挂载了数据盘，并且设置了AutoDelete属性为True，那么数据盘会随主机一起删除"
 */
func (c *VmClient) DeleteInstance(request *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DeleteInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询云主机详情 */
func (c *VmClient) DescribeInstance(request *DescribeInstanceRequest) (*DescribeInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 云主机缷载硬盘，主机和云盘没有未完成的任务时才可缷载 */
func (c *VmClient) DetachDisk(request *DetachDiskRequest) (*DetachDiskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DetachDiskResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询实例类型资源信息列表 */
func (c *VmClient) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (*DescribeInstanceTypesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &DescribeInstanceTypesResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

