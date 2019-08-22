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
    elite "github.com/jdcloud-api/jdcloud-sdk-go/services/elite/apis"
    "encoding/json"
    "errors"
)

type EliteClient struct {
    core.JDCloudClient
}

func NewEliteClient(credential *core.Credential) *EliteClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("elite.jdcloud-api.com")

    return &EliteClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "elite",
            Revision:    "1.0.6",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *EliteClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *EliteClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询价格 */
func (c *EliteClient) JdxQueryPrice(request *elite.JdxQueryPriceRequest) (*elite.JdxQueryPriceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.JdxQueryPriceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 分页查询交付单信息 */
func (c *EliteClient) ListSaleService(request *elite.ListSaleServiceRequest) (*elite.ListSaleServiceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.ListSaleServiceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询交付信息接口 */
func (c *EliteClient) JdxQueryDeliveryInfo(request *elite.JdxQueryDeliveryInfoRequest) (*elite.JdxQueryDeliveryInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.JdxQueryDeliveryInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取云存服务信息 */
func (c *EliteClient) GetStoreService(request *elite.GetStoreServiceRequest) (*elite.GetStoreServiceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.GetStoreServiceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 上报订单 */
func (c *EliteClient) JdxReportOrder(request *elite.JdxReportOrderRequest) (*elite.JdxReportOrderResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.JdxReportOrderResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 下单接口 */
func (c *EliteClient) JdxCreateOrder(request *elite.JdxCreateOrderRequest) (*elite.JdxCreateOrderResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.JdxCreateOrderResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据交付单号查询交付单信息 */
func (c *EliteClient) GetSaleServiceByDeliverNumber(request *elite.GetSaleServiceByDeliverNumberRequest) (*elite.GetSaleServiceByDeliverNumberResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.GetSaleServiceByDeliverNumberResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 输出商品接口 */
func (c *EliteClient) JdxQueryProduct(request *elite.JdxQueryProductRequest) (*elite.JdxQueryProductResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.JdxQueryProductResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 确认交付 */
func (c *EliteClient) ConfirmSaleServiceDelivery(request *elite.ConfirmSaleServiceDeliveryRequest) (*elite.ConfirmSaleServiceDeliveryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &elite.ConfirmSaleServiceDeliveryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

