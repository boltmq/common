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
package protocol

import (
	"fmt"
	"strings"
	"sync"

	"github.com/boltmq/common/constant"
)

const (
	defaultReadQueueNums  = 16
	defaultWriteQueueNums = 16
	separator             = " "
	perm                  = constant.PERM_READ | constant.PERM_WRITE
	topicFilterType       = TopicFilterType(SINGLE_TAG)
	topicSysFlag          = 0
)

type TopicConfig struct {
	SEPARATOR      string
	TopicName      string          `json:"topicName"`
	ReadQueueNums  int32           `json:"readQueueNums"`
	WriteQueueNums int32           `json:"writeQueueNums"`
	Perm           int             `json:"perm"`
	TpFilterType   TopicFilterType `json:"topicFilterType"`
	TopicSysFlag   int             `json:"topicSysFlag"`
	Order          bool            `json:"order"`
}

func NewTopicConfig(topicName string) *TopicConfig {
	topicConfig := &TopicConfig{
		TopicName:      topicName,
		WriteQueueNums: defaultWriteQueueNums,
		ReadQueueNums:  defaultReadQueueNums,
		SEPARATOR:      separator,
		Perm:           perm,
		TpFilterType:   topicFilterType,
		TopicSysFlag:   topicSysFlag,
	}
	return topicConfig
}

func NewDefaultTopicConfig(topicName string, readQueueNums, writeQueueNums int32, perm int, filterType TopicFilterType) *TopicConfig {
	topicConfig := &TopicConfig{
		TopicName:      topicName,
		WriteQueueNums: writeQueueNums,
		ReadQueueNums:  readQueueNums,
		SEPARATOR:      separator,
		Perm:           perm,
		TpFilterType:   filterType,
		TopicSysFlag:   topicSysFlag,
	}
	return topicConfig
}

func NewCustomTopicConfig(topicName string, readQueueNums, writeQueueNums int32, topicSysFlag int, filterType ...TopicFilterType) *TopicConfig {
	topicConfig := &TopicConfig{
		TopicName:      topicName,
		WriteQueueNums: writeQueueNums,
		ReadQueueNums:  readQueueNums,
		SEPARATOR:      separator,
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

type TopicConfigTable struct {
	TopicConfigs map[string]*TopicConfig `json:"topicConfigs"`
	lock         sync.RWMutex            `json:"-"`
}

func NewTopicConfigTable() *TopicConfigTable {
	topicConfigTable := &TopicConfigTable{
		TopicConfigs: make(map[string]*TopicConfig),
	}
	return topicConfigTable
}

func (table *TopicConfigTable) Size() int {
	table.lock.RLock()
	defer table.lock.RUnlock()

	return len(table.TopicConfigs)
}

func (table *TopicConfigTable) Keys() []string {
	table.lock.RLock()
	defer table.lock.RUnlock()
	if table == nil || table.TopicConfigs == nil || len(table.TopicConfigs) == 0 {
		return []string{}
	}

	topics := make([]string, 0, len(table.TopicConfigs))
	for topic, _ := range table.TopicConfigs {
		topics = append(topics, topic)
	}
	return topics
}

func (table *TopicConfigTable) Put(k string, v *TopicConfig) *TopicConfig {
	table.lock.Lock()
	defer table.lock.Unlock()
	oldV, _ := table.TopicConfigs[k]
	table.TopicConfigs[k] = v
	return oldV
}

func (table *TopicConfigTable) Get(k string) *TopicConfig {
	table.lock.RLock()
	defer table.lock.RUnlock()

	v, ok := table.TopicConfigs[k]
	if !ok {
		return nil
	}

	return v
}

func (table *TopicConfigTable) Remove(k string) *TopicConfig {
	table.lock.Lock()
	defer table.lock.Unlock()

	v, ok := table.TopicConfigs[k]
	if !ok {
		return nil
	}

	delete(table.TopicConfigs, k)
	return v
}

func (table *TopicConfigTable) Foreach(fn func(k string, v *TopicConfig)) {
	table.lock.RLock()
	defer table.lock.RUnlock()

	for k, v := range table.TopicConfigs {
		fn(k, v)
	}
}

func (table *TopicConfigTable) ForeachUpdate(fn func(k string, v *TopicConfig)) {
	table.lock.Lock()
	defer table.lock.Unlock()

	for k, v := range table.TopicConfigs {
		fn(k, v)
	}
}

// Clear 清空map
// author rongzhihong
// since 2017/9/18
func (table *TopicConfigTable) Clear() {
	table.lock.Lock()
	defer table.lock.Unlock()
	table.TopicConfigs = make(map[string]*TopicConfig)
}

// PutAll put all
// author rongzhihong
// since 2017/9/18
func (table *TopicConfigTable) PutAll(topicConfigTable map[string]*TopicConfig) {
	table.lock.Lock()
	defer table.lock.Unlock()

	for k, v := range topicConfigTable {
		table.TopicConfigs[k] = v
	}
}

// String 打印TopicConfigTable结构体的数据
// Author: tianyuliang
// Since: 2017/10/3
func (table *TopicConfigTable) String() string {
	if table == nil || table.TopicConfigs == nil || table.Size() == 0 {
		return ""
	}

	table.lock.RLock()
	defer table.lock.RUnlock()

	infos := make([]string, 0, table.Size())
	for topic, topicConfig := range table.TopicConfigs {
		info := fmt.Sprintf("[topic=%s, %s]", topic, topicConfig)
		infos = append(infos, info)
	}

	return fmt.Sprintf("TopicConfigTable [%s]", strings.Join(infos, ","))
}

// ClearAndPutAll 清空map后，再putAll
// author rongzhihong
// since 2017/9/18
func (table *TopicConfigTable) ClearAndPutAll(topicConfigTable map[string]*TopicConfig) {
	table.lock.Lock()
	defer table.lock.Unlock()
	table.TopicConfigs = make(map[string]*TopicConfig)

	for k, v := range topicConfigTable {
		table.TopicConfigs[k] = v
	}
}

// TopicConfigSerializeWrapper topic
// Author gaoyanlei
// Since 2017/8/11
type TopicConfigSerializeWrapper struct {
	TpConfigTable *TopicConfigTable `json:"topicConfigTable"`
	DataVersion   *DataVersion      `json:"dataVersion"`
	RemotingSerializable
}

// NewTopicConfigSerializeWrapper 格式化
// Author: tianyuliang
// Since: 2017/10/21
func NewTopicConfigSerializeWrapper(dataVersion ...*DataVersion) *TopicConfigSerializeWrapper {
	topicConfigSerializeWrapper := &TopicConfigSerializeWrapper{
		TpConfigTable: NewTopicConfigTable(),
	}

	topicConfigSerializeWrapper.DataVersion = NewDataVersion()
	if dataVersion != nil && len(dataVersion) > 0 {
		topicConfigSerializeWrapper.DataVersion = dataVersion[0]
	}
	return topicConfigSerializeWrapper
}

// String 格式化
// Author: tianyuliang
// Since: 2017/10/21
func (wrap *TopicConfigSerializeWrapper) String() string {
	return fmt.Sprintf("TopicConfigSerializeWrapper [%s, %s]", wrap.DataVersion, wrap.TpConfigTable)
}
