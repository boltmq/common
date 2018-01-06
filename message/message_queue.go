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
package message

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/boltmq/common/utils/codec"
)

type MessageQueue struct {
	Topic      string `json:"topic"`
	BrokerName string `json:"brokerName"`
	QueueId    int    `json:"queueId"`
}

func NewMessageQueue(topic, brokerName string, queueId int) *MessageQueue {
	messageQueue := &MessageQueue{
		Topic:      topic,
		BrokerName: brokerName,
		QueueId:    queueId,
	}
	return messageQueue
}

func (mq *MessageQueue) HashBytes() []byte {
	return []byte(strconv.Itoa(mq.hashCode()))
}

func (mq *MessageQueue) Equals(v interface{}) bool {
	if v == nil {
		return false
	}
	mq1, ok := v.(MessageQueue)
	var mq2 *MessageQueue
	if !ok {
		mq2, ok = v.(*MessageQueue)
	}
	if mq2 == nil {
		return ok && (strings.EqualFold(mq.BrokerName, mq1.BrokerName) && strings.EqualFold(mq.Topic, mq1.Topic) && mq.QueueId == mq1.QueueId)
	} else {
		return ok && (strings.EqualFold(mq.BrokerName, mq2.BrokerName) && strings.EqualFold(mq.Topic, mq2.Topic) && mq.QueueId == mq2.QueueId)
	}
}

func (mq *MessageQueue) hashCode() int {
	var prime int = 31
	var result int = 1
	if strings.EqualFold(mq.BrokerName, "") {
		result = prime*result + 0
	} else {
		result = prime*result + int(codec.HashCode(mq.BrokerName))
	}
	result = prime*result + mq.QueueId
	if strings.EqualFold(mq.Topic, "") {
		result = prime*result + 0
	} else {
		result = prime*result + int(codec.HashCode(mq.Topic))
	}
	return result
}

func (mq *MessageQueue) String() string {
	format := "MessageQueue {topic=%s, brokerName=%s, queueId=%d}"
	return fmt.Sprintf(format, mq.Topic, mq.BrokerName, mq.QueueId)
}

func (mq *MessageQueue) Key() string {
	return fmt.Sprintf("%s@%s@%d", mq.Topic, mq.BrokerName, mq.QueueId)
}

func (mq *MessageQueue) Equal(other *MessageQueue) bool {
	return strings.EqualFold(other.Topic, mq.Topic) && strings.EqualFold(other.BrokerName, mq.BrokerName) && other.QueueId == mq.QueueId
}

type MessageQueues []*MessageQueue

func (mqs MessageQueues) Less(i, j int) bool {
	imq := mqs[i]
	jmq := mqs[j]

	if imq.Topic < jmq.Topic {
		return true
	} else if imq.Topic > jmq.Topic {
		return false
	}

	if imq.BrokerName < jmq.BrokerName {
		return true
	} else if imq.BrokerName > jmq.BrokerName {
		return false
	}

	if imq.QueueId < jmq.QueueId {
		return true
	} else {
		return false
	}
}

func (mqs MessageQueues) Swap(i, j int) {
	mqs[i], mqs[j] = mqs[j], mqs[i]
}

func (mqs MessageQueues) Len() int {
	return len(mqs)
}
