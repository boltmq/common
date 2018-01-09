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
package head

// GetTopicStatsInfoRequestHeader 获得Topic统计信息的请求头
// Author rongzhihong
// Since 2017/9/19
type GetTopicStatsInfoRequestHeader struct {
	Topic string `json:"topic"`
}

func (header *GetTopicStatsInfoRequestHeader) CheckFields() error {
	return nil
}

// NewGetTopicStatsInfoRequestHeader 初始化
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/6
func NewGetTopicStatsInfoRequestHeader(topic string) *GetTopicStatsInfoRequestHeader {
	header := &GetTopicStatsInfoRequestHeader{
		Topic: topic,
	}
	return header
}
