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

	"github.com/boltmq/boltmq/net/core"
	"github.com/boltmq/common/basis"
	"github.com/boltmq/common/utils/system"
)

// BrokerLiveInfo 活动broker存储结构
type BrokerLiveInfo struct {
	LastUpdateTimestamp int64
	DataVersion         *basis.DataVersion
	Ctx                 core.Context
	HaServerAddr        string
}

// NewBrokerLiveInfo 初始化BrokerLiveInfo
func NewBrokerLiveInfo(dataVersion *basis.DataVersion, haServerAddr string, ctx core.Context) *BrokerLiveInfo {
	brokerLiveInfo := BrokerLiveInfo{
		LastUpdateTimestamp: system.CurrentTimeMillis(),
		DataVersion:         dataVersion,
		HaServerAddr:        haServerAddr,
		Ctx:                 ctx,
	}
	return &brokerLiveInfo
}

// String 打印
func (bli *BrokerLiveInfo) String() string {
	return fmt.Sprintf("BrokerLiveInfo {lastUpdateTimestamp=%d, %s, %s, haServerAddr=%s}",
		bli.LastUpdateTimestamp, bli.DataVersion, bli.Ctx, bli.HaServerAddr)
}
