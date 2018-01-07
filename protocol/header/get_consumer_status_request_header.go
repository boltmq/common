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

// GetConsumerStatusRequestHeader 获得消费者状态的请求头
// Author rongzhihong
// Since 2017/9/19
type GetConsumerStatusRequestHeader struct {
	Topic      string `json:"topic"`
	Group      string `json:"group"`
	ClientAddr string `json:"clientAddr"`
}

func (header *GetConsumerStatusRequestHeader) CheckFields() error {
	return nil
}

// NewGetConsumerStatusRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/6
func NewGetConsumerStatusRequestHeader(topic, group, clientAddr string) *GetConsumerStatusRequestHeader {
	consumerStatusRequestHeader := &GetConsumerStatusRequestHeader{
		Topic:      topic,
		Group:      group,
		ClientAddr: clientAddr,
	}
	return consumerStatusRequestHeader
}
