package admin

import (
	"github.com/boltmq/common/message"
	"github.com/boltmq/common/protocol"
)

// OffsetWrapper 消费者统计
// Author rongzhihong
// Since 2017/9/19
type ConsumeStats struct {
	ConsumeTps  float64                                  `json:"consumeTps"`
	OffsetTable map[*message.MessageQueue]*OffsetWrapper `json:"offsetTable"`
	protocol.RemotingSerializable
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
	protocol.RemotingSerializable
}

// NewConsumeStats 初始化
// Author rongzhihong
// Since 2017/9/19
func NewConsumeStatsPlus() *ConsumeStatsPlus {
	consumeStats := new(ConsumeStatsPlus)
	consumeStats.OffsetTable = make(map[string]*OffsetWrapper)
	return consumeStats
}
