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
package header

// ConsumerSendMsgBackRequestHeader: 消费消息头
// Author: yintongqiang
// Since:  2017/8/17
type ConsumerSendMsgBackRequestHeader struct {
	Offset      int64  `json:"offset"`
	Group       string `json:"group"`
	DelayLevel  int32  `json:"delayLevel"`
	OriginMsgId string `json:"originMsgId"`
	OriginTopic string `json:"originTopic"`
	UnitMode    bool   `json:"unitMode"`
}

func (header *ConsumerSendMsgBackRequestHeader) CheckFields() error {
	return nil
}

// 初始化 ConsumerSendMsgBackRequestHeader
// Author gaoyanlei
// Since 2017/8/17
func NewConsumerSendMsgBackRequestHeader() *ConsumerSendMsgBackRequestHeader {
	return &ConsumerSendMsgBackRequestHeader{
		UnitMode: false,
	}
}
