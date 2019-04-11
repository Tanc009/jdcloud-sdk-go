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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    rms "github.com/jdcloud-api/jdcloud-sdk-go/services/rms/apis"
    "encoding/json"
    "errors"
)

type RmsClient struct {
    core.JDCloudClient
}

func NewRmsClient(credential *core.Credential) *RmsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("rms.jdcloud-api.com")

    return &RmsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "rms",
            Revision:    "1.0.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *RmsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *RmsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 增加富媒体短信内容接口 */
func (c *RmsClient) AddTemplate(request *rms.AddTemplateRequest) (*rms.AddTemplateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rms.AddTemplateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 指定短信Id群发短信 */
func (c *RmsClient) SendBatchMsg(request *rms.SendBatchMsgRequest) (*rms.SendBatchMsgResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rms.SendBatchMsgResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询一个富媒体短信内容接口 */
func (c *RmsClient) QueryOneTemplate(request *rms.QueryOneTemplateRequest) (*rms.QueryOneTemplateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rms.QueryOneTemplateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询富媒体短信内容列表接口 */
func (c *RmsClient) QueryTemplateList(request *rms.QueryTemplateListRequest) (*rms.QueryTemplateListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rms.QueryTemplateListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取发送状态 */
func (c *RmsClient) QuerySendStatus(request *rms.QuerySendStatusRequest) (*rms.QuerySendStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &rms.QuerySendStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

