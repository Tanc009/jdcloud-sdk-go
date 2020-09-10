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
    cloudsign "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudsign/models"
)

type DescribeTemplateListRequest struct {

    core.JDCloudRequest

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小, 默认为10, 取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 合同模板名称或者标题 (Optional) */
    TemplateNameOrTitle *string `json:"templateNameOrTitle"`

    /* 模板类型 pdf,word,pdf-auto(不传查所有类型) (Optional) */
    TemplateType *string `json:"templateType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTemplateListRequest(
) *DescribeTemplateListRequest {

	return &DescribeTemplateListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/template",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小, 默认为10, 取值范围[10, 100] (Optional)
 * param templateNameOrTitle: 合同模板名称或者标题 (Optional)
 * param templateType: 模板类型 pdf,word,pdf-auto(不传查所有类型) (Optional)
 */
func NewDescribeTemplateListRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    templateNameOrTitle *string,
    templateType *string,
) *DescribeTemplateListRequest {

    return &DescribeTemplateListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/template",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        TemplateNameOrTitle: templateNameOrTitle,
        TemplateType: templateType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTemplateListRequestWithoutParam() *DescribeTemplateListRequest {

    return &DescribeTemplateListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/template",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeTemplateListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小, 默认为10, 取值范围[10, 100](Optional) */
func (r *DescribeTemplateListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param templateNameOrTitle: 合同模板名称或者标题(Optional) */
func (r *DescribeTemplateListRequest) SetTemplateNameOrTitle(templateNameOrTitle string) {
    r.TemplateNameOrTitle = &templateNameOrTitle
}

/* param templateType: 模板类型 pdf,word,pdf-auto(不传查所有类型)(Optional) */
func (r *DescribeTemplateListRequest) SetTemplateType(templateType string) {
    r.TemplateType = &templateType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTemplateListRequest) GetRegionId() string {
    return ""
}

type DescribeTemplateListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTemplateListResult `json:"result"`
}

type DescribeTemplateListResult struct {
    TemplateList []cloudsign.TemplateInfo `json:"templateList"`
    TotalCount int `json:"totalCount"`
}