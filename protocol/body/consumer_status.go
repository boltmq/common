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

import (
	"github.com/boltmq/common/message"
)

// GetConsumerStatus  获得消费状态的body
// Author rongzhihong
// Since 2017/9/19
type GetConsumerStatus struct {
	MessageQueueTable map[*message.MessageQueue]int64            `json:"messageQueueTable"`
	ConsumerTable     map[string]map[*message.MessageQueue]int64 `json:"consumerTable"`
}

func NewGetConsumerStatus() *GetConsumerStatus {
	body := new(GetConsumerStatus)
	body.MessageQueueTable = make(map[*message.MessageQueue]int64)
	body.ConsumerTable = make(map[string]map[*message.MessageQueue]int64)
	return body
}
