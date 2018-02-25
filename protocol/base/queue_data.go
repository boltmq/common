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
package base

import (
	"fmt"
	"strings"
)

type QueueData struct {
	BrokerName     string `json:"brokerName"`
	ReadQueueNums  int    `json:"readQueueNums"`
	WriteQueueNums int    `json:"writeQueueNums"`
	Perm           int    `json:"perm"`
	TopicSynFlag   int    `json:"topicSynFlag"`
}

func NewQueueData(brokerName string, topicConfig *TopicConfig) *QueueData {
	queueData := &QueueData{
		BrokerName:     brokerName,
		WriteQueueNums: int(topicConfig.WriteQueueNums),
		ReadQueueNums:  int(topicConfig.ReadQueueNums),
		Perm:           topicConfig.Perm,
		TopicSynFlag:   topicConfig.TopicSysFlag,
	}
	return queueData
}

func (qd *QueueData) String() string {
	if qd == nil {
		return "<nil>"
	}

	return fmt.Sprintf("QueueData {brokerName=%s, readQueueNums=%d, writeQueueNums=%d, perm=%d, topicSynFlag=%d}",
		qd.BrokerName, qd.ReadQueueNums, qd.WriteQueueNums, qd.Perm, qd.TopicSynFlag)
}

func (qd *QueueData) Equals(v interface{}) bool {
	if v == nil {
		return false
	}

	qData1, ok := v.(QueueData)
	var qData2 *QueueData
	if !ok {
		qData2, ok = v.(*QueueData)
	}
	if !ok {
		return ok
	}
	if qData2 == nil {
		if !strings.EqualFold(qd.BrokerName, qData1.BrokerName) {
			return false
		}
		if qd.Perm != qData1.Perm {
			return false
		}
		if qd.WriteQueueNums != qData1.WriteQueueNums {
			return false
		}
		if qd.ReadQueueNums != qData1.ReadQueueNums {
			return false
		}
		if qd.TopicSynFlag != qData1.TopicSynFlag {
			return false
		}
	} else {
		if !strings.EqualFold(qd.BrokerName, qData2.BrokerName) {
			return false
		}
		if qd.Perm != qData2.Perm {
			return false
		}
		if qd.WriteQueueNums != qData2.WriteQueueNums {
			return false
		}
		if qd.ReadQueueNums != qData2.ReadQueueNums {
			return false
		}
		if qd.TopicSynFlag != qData2.TopicSynFlag {
			return false
		}
	}
	return true
}

type QueueDatas []*QueueData

func (qds QueueDatas) Less(i, j int) bool {
	iq := qds[i]
	jq := qds[j]

	if iq.BrokerName < jq.BrokerName {
		return true
	} else if iq.BrokerName > jq.BrokerName {
		return false
	}
	return false
}

func (qds QueueDatas) Swap(i, j int) {
	qds[i], qds[j] = qds[j], qds[i]
}

func (qds QueueDatas) Len() int {
	return len(qds)
}
