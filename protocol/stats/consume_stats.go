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
package stats

import (
	"github.com/boltmq/common/message"
)

// OffsetWrapper 消费者统计
// Author rongzhihong
// Since 2017/9/19
type ConsumeStats struct {
	ConsumeTps  float64                                  `json:"consumeTps"`
	OffsetTable map[*message.MessageQueue]*OffsetWrapper `json:"offsetTable"`
}

// NewConsumeStats 初始化
// Author rongzhihong
// Since 2017/9/19
func NewConsumeStats() *ConsumeStats {
	consumeStats := new(ConsumeStats)
	consumeStats.OffsetTable = make(map[*message.MessageQueue]*OffsetWrapper)
	return consumeStats
}

// ComputeTotalDiff 偏移量差值
// Author rongzhihong
// Since 2017/9/19
func (consumeStats *ConsumeStats) ComputeTotalDiff() int64 {
	diffTotal := int64(0)
	if consumeStats == nil || consumeStats.OffsetTable == nil {
		return diffTotal
	}
	for _, wrapper := range consumeStats.OffsetTable {
		diff := wrapper.BrokerOffset - wrapper.ConsumerOffset
		diffTotal += diff
	}
	return diffTotal
}

// OffsetWrapper 消费者统计
// Author rongzhihong
// Since 2017/9/19
type ConsumeStatsPlus struct {
	ConsumeTps  float64                   `json:"consumeTps"`
	OffsetTable map[string]*OffsetWrapper `json:"offsetTable"` // key: Topic@BrokerName@QueueId
}

// NewConsumeStats 初始化
// Author rongzhihong
// Since 2017/9/19
func NewConsumeStatsPlus() *ConsumeStatsPlus {
	consumeStats := new(ConsumeStatsPlus)
	consumeStats.OffsetTable = make(map[string]*OffsetWrapper)
	return consumeStats
}
