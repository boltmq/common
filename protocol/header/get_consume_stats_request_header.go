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

// GetConsumeStatsRequestHeader 获得消费者统计信息的请求头
// Author rongzhihong
// Since 2017/9/19
type GetConsumeStatsRequestHeader struct {
	Topic         string `json:"topic"`
	ConsumerGroup string `json:"consumerGroup"`
}

func (header *GetConsumeStatsRequestHeader) CheckFields() error {
	return nil
}

// NewGetConsumeStatsRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/1
func NewGetConsumeStatsRequestHeader(consumerGroup, topic string) *GetConsumeStatsRequestHeader {
	consumeStatsRequestHeader := &GetConsumeStatsRequestHeader{
		Topic:         topic,
		ConsumerGroup: consumerGroup,
	}
	return consumeStatsRequestHeader
}
