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

type SetAuthConfigRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 是否开启鉴权[on,off] (Optional) */
    EnableUrlAuth *string `json:"enableUrlAuth"`

    /* 鉴权key (Optional) */
    AuthKey *string `json:"authKey"`

    /* 鉴权时间戳过期时间，默认为0 (Optional) */
    Age *int `json:"age"`

    /* 鉴权参数加密算法，默认为md5且只支持md5 (Optional) */
    EncAlgorithm *string `json:"encAlgorithm"`

    /* 时间戳格式[hex,dec] (Optional) */
    TimeFormat *string `json:"timeFormat"`

    /* 加密算法版本[dash,dashv2,video],默认dashv2 (Optional) */
    UriType *string `json:"uriType"`

    /* 鉴权key生成顺序 (Optional) */
    Rule *string `json:"rule"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetAuthConfigRequest(
    domain string,
) *SetAuthConfigRequest {

	return &SetAuthConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/setAuthConfig",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param enableUrlAuth: 是否开启鉴权[on,off] (Optional)
 * param authKey: 鉴权key (Optional)
 * param age: 鉴权时间戳过期时间，默认为0 (Optional)
 * param encAlgorithm: 鉴权参数加密算法，默认为md5且只支持md5 (Optional)
 * param timeFormat: 时间戳格式[hex,dec] (Optional)
 * param uriType: 加密算法版本[dash,dashv2,video],默认dashv2 (Optional)
 * param rule: 鉴权key生成顺序 (Optional)
 */
func NewSetAuthConfigRequestWithAllParams(
    domain string,
    enableUrlAuth *string,
    authKey *string,
    age *int,
    encAlgorithm *string,
    timeFormat *string,
    uriType *string,
    rule *string,
) *SetAuthConfigRequest {

    return &SetAuthConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/setAuthConfig",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        EnableUrlAuth: enableUrlAuth,
        AuthKey: authKey,
        Age: age,
        EncAlgorithm: encAlgorithm,
        TimeFormat: timeFormat,
        UriType: uriType,
        Rule: rule,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetAuthConfigRequestWithoutParam() *SetAuthConfigRequest {

    return &SetAuthConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/setAuthConfig",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetAuthConfigRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param enableUrlAuth: 是否开启鉴权[on,off](Optional) */
func (r *SetAuthConfigRequest) SetEnableUrlAuth(enableUrlAuth string) {
    r.EnableUrlAuth = &enableUrlAuth
}

/* param authKey: 鉴权key(Optional) */
func (r *SetAuthConfigRequest) SetAuthKey(authKey string) {
    r.AuthKey = &authKey
}

/* param age: 鉴权时间戳过期时间，默认为0(Optional) */
func (r *SetAuthConfigRequest) SetAge(age int) {
    r.Age = &age
}

/* param encAlgorithm: 鉴权参数加密算法，默认为md5且只支持md5(Optional) */
func (r *SetAuthConfigRequest) SetEncAlgorithm(encAlgorithm string) {
    r.EncAlgorithm = &encAlgorithm
}

/* param timeFormat: 时间戳格式[hex,dec](Optional) */
func (r *SetAuthConfigRequest) SetTimeFormat(timeFormat string) {
    r.TimeFormat = &timeFormat
}

/* param uriType: 加密算法版本[dash,dashv2,video],默认dashv2(Optional) */
func (r *SetAuthConfigRequest) SetUriType(uriType string) {
    r.UriType = &uriType
}

/* param rule: 鉴权key生成顺序(Optional) */
func (r *SetAuthConfigRequest) SetRule(rule string) {
    r.Rule = &rule
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetAuthConfigRequest) GetRegionId() string {
    return ""
}

type SetAuthConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetAuthConfigResult `json:"result"`
}

type SetAuthConfigResult struct {
    TaskId string `json:"taskId"`
}