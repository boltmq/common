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
	"fmt"
	"strings"

	"github.com/boltmq/common/message"
)

// TopicStatsTable Topic统计信息
// Author rongzhihong
// Since 2017/9/19
type TopicStatsTable struct {
	OffsetTable map[*message.MessageQueue]*TopicOffset `json:"offsetTable"`
}

// NewTopicStatsTable 初始化Topic统计信息
// Author rongzhihong
// Since 2017/9/19
func NewTopicStatsTable() *TopicStatsTable {
	topicStatsTable := new(TopicStatsTable)
	topicStatsTable.OffsetTable = make(map[*message.MessageQueue]*TopicOffset)
	return topicStatsTable
}

// TopicStatsTablePlus 因key为struct，Encode报错，修改结构
// Author rongzhihong
// Since 2017/9/19
type TopicStatsTablePlus struct {
	OffsetTable map[string]*TopicOffset `json:"offsetTable"` // key Topic@BrokerName@QueueId
}

// NewTopicStatsTablePlus 初始化Topic统计信息
// Author rongzhihong
// Since 2017/9/19
func NewTopicStatsTablePlus() *TopicStatsTablePlus {
	topicStatsTablePlus := new(TopicStatsTablePlus)
	topicStatsTablePlus.OffsetTable = make(map[string]*TopicOffset)
	return topicStatsTablePlus
}

// String 格式化TopicStatsTablePlus内容
// Author: tianyuliang
// Since: 2017/11/9
func (plus *TopicStatsTablePlus) String() string {
	if plus == nil || plus.OffsetTable == nil {
		return "<nil>"
	}

	var datas []string
	for mqKey, topicOffet := range plus.OffsetTable {
		data := fmt.Sprintf("[messageQueue=%s, topicOffet=%s]", mqKey, topicOffet)
		datas = append(datas, data)
	}
	return fmt.Sprintf("TopicStatsTablePlus {%s} ", strings.Join(datas, ", "))
}
