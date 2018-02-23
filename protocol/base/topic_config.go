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

	"github.com/boltmq/common/basis"
	"github.com/boltmq/common/constant"
)

const (
	defaultReadQueueNums  = 16
	defaultWriteQueueNums = 16
	separator             = " "
	perm                  = constant.PERM_READ | constant.PERM_WRITE
	topicFilterType       = basis.TopicFilterType(basis.SINGLE_TAG)
	topicSysFlag          = 0
)

type TopicConfig struct {
	Separator      string
	TopicName      string                `json:"topicName"`
	ReadQueueNums  int32                 `json:"readQueueNums"`
	WriteQueueNums int32                 `json:"writeQueueNums"`
	Perm           int                   `json:"perm"`
	TpFilterType   basis.TopicFilterType `json:"topicFilterType"`
	TopicSysFlag   int                   `json:"topicSysFlag"`
	Order          bool                  `json:"order"`
}

func NewTopicConfig(topicName string) *TopicConfig {
	topicConfig := &TopicConfig{
		TopicName:      topicName,
		WriteQueueNums: defaultWriteQueueNums,
		ReadQueueNums:  defaultReadQueueNums,
		Separator:      separator,
		Perm:           perm,
		TpFilterType:   topicFilterType,
		TopicSysFlag:   topicSysFlag,
	}
	return topicConfig
}

func NewDefaultTopicConfig(topicName string, readQueueNums, writeQueueNums int32, perm int, filterType basis.TopicFilterType) *TopicConfig {
	topicConfig := &TopicConfig{
		TopicName:      topicName,
		WriteQueueNums: writeQueueNums,
		ReadQueueNums:  readQueueNums,
		Separator:      separator,
		Perm:           perm,
		TpFilterType:   filterType,
		TopicSysFlag:   topicSysFlag,
	}
	return topicConfig
}

func NewCustomTopicConfig(topicName string, readQueueNums, writeQueueNums int32, topicSysFlag int, filterType ...basis.TopicFilterType) *TopicConfig {
	topicConfig := &TopicConfig{
		TopicName:      topicName,
		WriteQueueNums: writeQueueNums,
		ReadQueueNums:  readQueueNums,
		Separator:      separator,
		Perm:           perm,
		TopicSysFlag:   topicSysFlag,
	}
	if filterType != nil && len(filterType) > 0 {
		topicConfig.TpFilterType = filterType[0]
	} else {
		topicConfig.TpFilterType = topicFilterType
	}

	return topicConfig
}

func (tc *TopicConfig) String() string {
	if tc == nil {
		return "<nil>"
	}

	return fmt.Sprintf("TopicConfig [topicName=%s, readQueueNums=%d, writeQueueNums=%d, perm=%s, topicFilterType=%d, topicSysFlag=%d, order=%t]",
		tc.TopicName, tc.ReadQueueNums, tc.WriteQueueNums, tc.PermString(), int(tc.TpFilterType), tc.TopicSysFlag, tc.Order)
}

func (tc *TopicConfig) PermString() string {
	if tc == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%d(%s)", tc.Perm, constant.Perm2String(tc.Perm))
}
