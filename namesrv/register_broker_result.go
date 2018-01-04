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
	"fmt"
	"strings"

	"github.com/boltmq/common/protocol"
)

type RegisterBrokerResult struct {
	HaServerAddr string
	MasterAddr   string
	KvTable      *protocol.KVTable
}

// NewRegisterBrokerResult 初始化RegisterBrokerResult
// Author: tianyuliang
// Since: 2017/9/22
func NewRegisterBrokerResult(haServerAddr, masterAddr string) *RegisterBrokerResult {
	result := &RegisterBrokerResult{
		HaServerAddr: haServerAddr,
		MasterAddr:   masterAddr,
		KvTable:      protocol.NewKVTable(),
	}
	return result
}

func (result *RegisterBrokerResult) String() string {
	if result == nil {
		return "<nil>"
	}

	datas := make([]string, 0, len(result.KvTable.Table))
	if result.KvTable != nil && result.KvTable.Table != nil {
		for key, value := range result.KvTable.Table {
			kv := fmt.Sprintf("[key=%s, value=%s]", key, value)
			datas = append(datas, kv)
		}
	}

	return fmt.Sprintf("registerBrokerResult {haServerAddr=%s, masterAddr=%s, kvTable=%s}",
		result.HaServerAddr, result.MasterAddr, strings.Join(datas, ","))
}
