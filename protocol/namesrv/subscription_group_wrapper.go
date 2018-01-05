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
package namesrv

import (
	"github.com/boltmq/common/basis"
	"github.com/boltmq/common/protocol"
	concurrent "github.com/fanliao/go-concurrentMap"
)

// SubscriptionGroupWrapper 订阅组配置，序列化包装
// Author gaoyanlei
// Since 2017/8/22
type SubscriptionGroupWrapper struct {
	SubscriptionGroupTable *concurrent.ConcurrentMap `json:"subscriptionGroupTable"`
	DataVersion            basis.DataVersion         `json:"dataVersion"`
	protocol.RemotingSerializable
}

func NewSubscriptionGroupWrapper() *SubscriptionGroupWrapper {
	wrapper := new(SubscriptionGroupWrapper)
	wrapper.SubscriptionGroupTable = concurrent.NewConcurrentMap()
	wrapper.DataVersion = *basis.NewDataVersion()
	return wrapper
}
