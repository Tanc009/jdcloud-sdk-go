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
    . "github.com/jdcloud-api/jdcloud-sdk-go/services/mps/apis"
    "encoding/json"
    "errors"
)

type MpsClient struct {
    JDCloudClient
}

func NewMpsClient(credential *Credential) *MpsClient {
    if credential == nil {
        return nil
    }

    config := NewConfig()
    config.SetEndpoint("mps.jdcloud-api.com")

    return &MpsClient{
        JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "mps",
            Revision:    "0.3.2",
            Logger:      NewDefaultLogger(LOG_INFO),
        }}
}

func (c *MpsClient) SetConfig(config *Config) {
    c.Config = *config
}

func (c *MpsClient) SetLogger(logger Logger) {
    c.Logger = logger
}

/* 设置媒体处理通知, 在设置Notification时会对endpoint进行校验，设置时会对endpoint发一条SubscriptionConfirmation(x-jdcloud-message-type头)的通知，要求把Message内容进行base64编码返回给系统(body)进行校验 */
func (c *MpsClient) SetNotification(request *SetNotificationRequest) (*SetNotificationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &SetNotificationResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 获取截图通知 */
func (c *MpsClient) GetNotification(request *GetNotificationRequest) (*GetNotificationResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &GetNotificationResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 创建截图任务 */
func (c *MpsClient) CreateThumbnailTask(request *CreateThumbnailTaskRequest) (*CreateThumbnailTaskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &CreateThumbnailTaskResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 获取截图任务 */
func (c *MpsClient) GetThumbnailTask(request *GetThumbnailTaskRequest) (*GetThumbnailTaskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &GetThumbnailTaskResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

/* 查询截图任务 */
func (c *MpsClient) ListThumbnailTask(request *ListThumbnailTaskRequest) (*ListThumbnailTaskResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ListThumbnailTaskResponse{}
    err = json.Unmarshal(resp, jdResp)
    return jdResp, err
}

