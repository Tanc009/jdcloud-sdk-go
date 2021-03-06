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


type CommonTrafficData struct {

    /*  (Optional) */
    Avgbandwidth float64 `json:"avgbandwidth"`

    /*  (Optional) */
    BandUnit string `json:"bandUnit"`

    /*  (Optional) */
    Flow float64 `json:"flow"`

    /*  (Optional) */
    FlowUnit string `json:"flowUnit"`

    /*  (Optional) */
    Oriflow float64 `json:"oriflow"`

    /*  (Optional) */
    OriflowUnit string `json:"oriflowUnit"`

    /*  (Optional) */
    Pv float64 `json:"pv"`

    /*  (Optional) */
    PvUnit string `json:"pvUnit"`

    /*  (Optional) */
    Oripv float64 `json:"oripv"`

    /*  (Optional) */
    OripvUnit string `json:"oripvUnit"`

    /*  (Optional) */
    TopTimeStamp int64 `json:"topTimeStamp"`

    /*  (Optional) */
    OriFlowPercent string `json:"oriFlowPercent"`

    /*  (Optional) */
    OriPvPercent string `json:"oriPvPercent"`
}
