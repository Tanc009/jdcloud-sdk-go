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


type AudioDetailList struct {

    /* 总时长,单位为分钟 (Optional) */
    Audio_time []int `json:"audio_time"`

    /* 正常时长,单位为分钟 (Optional) */
    Audio_pass_time []int `json:"audio_pass_time"`

    /* 疑似时长，单位为分钟 (Optional) */
    Audio_review_time []int `json:"audio_review_time"`

    /* 违规时长，单位为分钟 (Optional) */
    Audio_block_time []int `json:"audio_block_time"`
}
