// Copyright 2017 luoji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package body

// BrokerStatsItem Broker统计最小数据单元
// Author rongzhihong
// Since 2017/9/19
type BrokerStatsItem struct {
	Sum   int64   `json:"sum"`
	Tps   float64 `json:"tps"`
	Avgpt float64 `json:"avgpt"`
}

// BrokerStatsData Broker统计数据
// Author rongzhihong
// Since 2017/9/19
type BrokerStatsData struct {
	StatsMinute *BrokerStatsItem `json:"statsMinute"`
	StatsHour   *BrokerStatsItem `json:"statsHour"`
	StatsDay    *BrokerStatsItem `json:"statsDay"`
}

// BrokerStatsData Broker统计数据
// Author rongzhihong
// Since 2017/9/19
func NewBrokerStatsData() *BrokerStatsData {
	statsData := new(BrokerStatsData)
	statsData.StatsMinute = new(BrokerStatsItem)
	statsData.StatsHour = new(BrokerStatsItem)
	statsData.StatsDay = new(BrokerStatsItem)
	return statsData
}
