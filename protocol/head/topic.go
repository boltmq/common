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

import (
	"github.com/boltmq/common/basis"
	"github.com/boltmq/common/protocol/base"
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

func NewCreateTopicRequestHeader(topicWithProjectGroup, defaultTopic string, topicConfig *base.TopicConfig) *CreateTopicRequestHeader {
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

// DeleteTopicRequestHeader 删除Topic
// Author gaoyanlei
// Since 2017/8/25
type DeleteTopicRequestHeader struct {
	Topic string `json:"topic"`
}

func (header *DeleteTopicRequestHeader) CheckFields() error {
	return nil
}

// GetAllTopicConfigResponseHeader 获得Topic配置的返回头
// Author rongzhihong
// Since 2017/9/19
type GetAllTopicConfigResponseHeader struct {
}

func (header *GetAllTopicConfigResponseHeader) CheckFields() error {
	return nil
}

// QueryTopicConsumeByWhoRequestHeader 查询Topic被哪些消费者消费的请求头
// Author rongzhihong
// Since 2017/9/19
type QueryTopicConsumeByWhoRequestHeader struct {
	Topic string
}

func (header *QueryTopicConsumeByWhoRequestHeader) CheckFields() error {
	return nil
}
