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


type AppQueryResultItem struct {

    /* 应用 (Optional) */
    ClientId string `json:"clientId"`

    /* 应用名 (Optional) */
    ClientName string `json:"clientName"`

    /* tokenEndpointAuthMethod (Optional) */
    TokenEndpointAuthMethod string `json:"tokenEndpointAuthMethod"`

    /* grantTypes (Optional) */
    GrantTypes string `json:"grantTypes"`

    /* responseTypes (Optional) */
    ResponseTypes string `json:"responseTypes"`

    /* redirectUris (Optional) */
    RedirectUris string `json:"redirectUris"`

    /* clientUri (Optional) */
    ClientUri string `json:"clientUri"`

    /* logoUri (Optional) */
    LogoUri string `json:"logoUri"`

    /* tosUri (Optional) */
    TosUri string `json:"tosUri"`

    /* policyUri (Optional) */
    PolicyUri string `json:"policyUri"`

    /* scope (Optional) */
    Scope string `json:"scope"`

    /* jwksUri (Optional) */
    JwksUri string `json:"jwksUri"`

    /* jwks (Optional) */
    Jwks string `json:"jwks"`

    /* contacts (Optional) */
    Contacts string `json:"contacts"`

    /* extension (Optional) */
    Extension string `json:"extension"`

    /* accessTokenValiditySeconds (Optional) */
    AccessTokenValiditySeconds int `json:"accessTokenValiditySeconds"`

    /* refreshTokenValiditySeconds (Optional) */
    RefreshTokenValiditySeconds int `json:"refreshTokenValiditySeconds"`

    /* multiTenant (Optional) */
    MultiTenant bool `json:"multiTenant"`

    /* secretUpdateTime (Optional) */
    SecretUpdateTime int64 `json:"secretUpdateTime"`

    /* updateTime (Optional) */
    UpdateTime int64 `json:"updateTime"`

    /* createTime (Optional) */
    CreateTime int64 `json:"createTime"`

    /* account (Optional) */
    Account string `json:"account"`

    /* userType (Optional) */
    UserType string `json:"userType"`

    /* state (Optional) */
    State string `json:"state"`
}