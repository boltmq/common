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

import (
	"github.com/boltmq/common/basis"
	"github.com/boltmq/common/protocol"
)

// CreateTopicRequestHeader: 创建topic头信息
// Author: yintongqiang
// Since:  2017/8/17
type CreateTopicRequestHeader struct {
	Topic           string // 真正的topic名称是位于topicConfig.Topic字段
	DefaultTopic    string // 表示创建topic的key值
	ReadQueueNums   int32
	WriteQueueNums  int32
	Perm            int
	TopicFilterType basis.TopicFilterType
	TopicSysFlag    int
	Order           bool
}

func (header *CreateTopicRequestHeader) CheckFields() error {
	return nil
}

func NewCreateTopicRequestHeader(topicWithProjectGroup, defaultTopic string, topicConfig *protocol.TopicConfig) *CreateTopicRequestHeader {
	createTopicRequestHeader := &CreateTopicRequestHeader{
		Topic:           topicWithProjectGroup,
		DefaultTopic:    defaultTopic,
		ReadQueueNums:   topicConfig.ReadQueueNums,
		WriteQueueNums:  topicConfig.WriteQueueNums,
		TopicFilterType: topicConfig.TpFilterType,
		TopicSysFlag:    topicConfig.TopicSysFlag,
		Order:           topicConfig.Order,
		Perm:            topicConfig.Perm,
	}
	return createTopicRequestHeader
}
