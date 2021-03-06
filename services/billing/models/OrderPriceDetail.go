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


type OrderPriceDetail struct {

    /* 折扣前总价 (Optional) */
    Price int `json:"price"`

    /* 四位小数价格 (Optional) */
    PriceScale4 int `json:"priceScale4"`

    /* 折扣金额 (Optional) */
    Discount int `json:"discount"`

    /* 折扣后订单金额 (Optional) */
    DiscountedPrice int `json:"discountedPrice"`

    /* 订单原价 包年时 一年原价为12个月价格，totalPrice为10个月价格 (Optional) */
    OriginalPrice int `json:"originalPrice"`

    /* 资源id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 业务线 (Optional) */
    AppCode string `json:"appCode"`

    /* 产品线 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 站点  0:主站  其他:专有云 (Optional) */
    Site int `json:"site"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 计费类型1:按配置2:按用量3:包年包月 (Optional) */
    BillingType int `json:"billingType"`

    /* 时长 (Optional) */
    TimeSpan int `json:"timeSpan"`

    /* 时长类型 1:小时2:天3:月4:年 (Optional) */
    TimeUnit int `json:"timeUnit"`

    /* 网络类型 0:non1:非BGP2:BGP (Optional) */
    NetworkOperator int `json:"networkOperator"`

    /* 配置信息 (Optional) */
    Formula []Formula `json:"formula"`

    /* FavorableInfo转成json后的字符串 (Optional) */
    FavorableInfo string `json:"favorableInfo"`

    /* 价格快照 (Optional) */
    PriceSnapShot string `json:"priceSnapShot"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 自然单列表 (Optional) */
    TaskId string `json:"taskId"`

    /* 开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 变配明细（1-升配补差价，2-降配延时，3-临时升配） (Optional) */
    ProcessType int `json:"processType"`

    /* 交易单模块sourceId (Optional) */
    SourceId string `json:"sourceId"`
}
