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
package subscription

import (
	"sync"

	"github.com/boltmq/common/basis"
	concurrent "github.com/fanliao/go-concurrentMap"
)

type SubscriptionGroupTable struct {
	Tables      map[string]*SubscriptionGroupConfig `json:"tables"` // key:group
	DataVersion basis.DataVersion                   `json:"dataVersion"`
	lock        sync.RWMutex                        `json:"-"`
}

func NewSubscriptionGroupTable() *SubscriptionGroupTable {
	table := &SubscriptionGroupTable{
		Tables:      make(map[string]*SubscriptionGroupConfig, 1024),
		DataVersion: *basis.NewDataVersion(),
	}
	return table
}

func (table *SubscriptionGroupTable) Size() int {
	table.lock.RLock()
	defer table.lock.RUnlock()

	return len(table.Tables)
}

func (table *SubscriptionGroupTable) Put(k string, v *SubscriptionGroupConfig) *SubscriptionGroupConfig {
	table.lock.Lock()
	defer table.lock.Unlock()
	old := table.Tables[k]
	table.Tables[k] = v
	return old
}

func (table *SubscriptionGroupTable) Get(k string) *SubscriptionGroupConfig {
	table.lock.RLock()
	defer table.lock.RUnlock()

	v, ok := table.Tables[k]
	if !ok {
		return nil
	}

	return v
}

func (table *SubscriptionGroupTable) Remove(k string) *SubscriptionGroupConfig {
	table.lock.Lock()
	defer table.lock.Unlock()

	v, ok := table.Tables[k]
	if !ok {
		return nil
	}

	delete(table.Tables, k)
	return v
}

func (table *SubscriptionGroupTable) Foreach(fn func(k string, v *SubscriptionGroupConfig)) {
	table.lock.RLock()
	defer table.lock.RUnlock()

	for k, v := range table.Tables {
		fn(k, v)
	}
}

// Clear 清空
// Author rongzhihong
// Since 2017/9/18
func (table *SubscriptionGroupTable) Clear() {
	table.lock.Lock()
	defer table.lock.Unlock()

	table.Tables = make(map[string]*SubscriptionGroupConfig, 1024)
}

// syncTopicConfig 同步Topic配置文件
// Author rongzhihong
// Since 2017/9/18
func (table *SubscriptionGroupTable) PutAll(offsetMap *concurrent.ConcurrentMap) {
	table.lock.Lock()
	defer table.lock.Unlock()

	if offsetMap == nil {
		return
	}

	for iter := offsetMap.Iterator(); iter.HasNext(); {
		key, value, _ := iter.Next()
		if groupName, ok := key.(string); ok && groupName != "" {
			if subscriptionGroupConfig, ok := value.(*SubscriptionGroupConfig); ok {
				table.Tables[groupName] = subscriptionGroupConfig
			}
		}
	}
}

// ClearAndPutAll 清空map,再PutAll
// Author rongzhihong
// Since 2017/9/18
func (table *SubscriptionGroupTable) ClearAndPutAll(offsetMap *concurrent.ConcurrentMap) {
	table.lock.Lock()
	defer table.lock.Unlock()

	table.Tables = make(map[string]*SubscriptionGroupConfig, 1024)
	if offsetMap == nil {
		return
	}

	for iter := offsetMap.Iterator(); iter.HasNext(); {
		key, value, _ := iter.Next()
		if groupName, ok := key.(string); ok && groupName != "" {
			if subscriptionGroupConfig, ok := value.(*SubscriptionGroupConfig); ok {
				table.Tables[groupName] = subscriptionGroupConfig
			}
		}
	}
}
