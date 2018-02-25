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
	"strings"
	"sync"

	"github.com/boltmq/common/basis"
)

type BrokerData struct {
	BrokerName  string         `json:"brokerName"`
	BrokerAddrs map[int]string `json:"brokerAddrs"`
	lock        sync.RWMutex   `json:"-"`
}

func NewBrokerData(brokerName string) *BrokerData {
	brokerData := &BrokerData{
		BrokerName:  brokerName,
		BrokerAddrs: make(map[int]string),
	}
	return brokerData
}

func (bd *BrokerData) SelectBrokerAddr() string {
	value := bd.BrokerAddrs[basis.MASTER_ID]
	if strings.EqualFold(value, "") {
		for _, value := range bd.BrokerAddrs {
			return value
		}
	}
	return value
}

func (bd *BrokerData) CloneBrokerData() *BrokerData {
	brokerDataClone := &BrokerData{}
	brokerDataClone.BrokerName = bd.BrokerName
	brokerDataClone.BrokerAddrs = make(map[int]string, 0)
	for k, v := range bd.BrokerAddrs {
		brokerDataClone.BrokerAddrs[k] = v
	}

	return brokerDataClone
}

func (bd *BrokerData) String() string {
	if bd == nil {
		return "<nil>"
	}

	brokerAddrs := make([]string, 0, len(bd.BrokerAddrs))
	if bd.BrokerAddrs != nil && len(bd.BrokerAddrs) > 0 {
		for k, v := range bd.BrokerAddrs {
			brokerAddr := fmt.Sprintf("{brokerId=%d, brokerAddr=%s}", k, v)
			brokerAddrs = append(brokerAddrs, brokerAddr)
		}
	}

	return fmt.Sprintf("BrokerData {brokerName:%s, brokerAddrs:%s}",
		bd.BrokerName, strings.Join(brokerAddrs, ","))
}

func (brokerData *BrokerData) Equals(v interface{}) bool {
	if v == nil {
		return false
	}
	bData1, ok := v.(BrokerData)
	var bData2 *BrokerData
	if !ok {
		bData2, ok = v.(*BrokerData)
	}
	if !ok {
		return ok
	}
	if bData2 == nil {
		if !strings.EqualFold(brokerData.BrokerName, bData1.BrokerName) {
			return false
		}
		if len(brokerData.BrokerAddrs) != len(bData1.BrokerAddrs) {
			return false
		}
		var flag bool = true
		for brokerId, brokderAddr := range brokerData.BrokerAddrs {
			if !strings.EqualFold(brokderAddr, bData1.BrokerAddrs[brokerId]) {
				flag = false
				break
			}
		}
		return flag
	} else {
		if !strings.EqualFold(brokerData.BrokerName, bData2.BrokerName) {
			return false
		}
		if len(brokerData.BrokerAddrs) != len(bData2.BrokerAddrs) {
			return false
		}
		var flag bool = true
		for brokerId, brokderAddr := range brokerData.BrokerAddrs {
			if !strings.EqualFold(brokderAddr, bData2.BrokerAddrs[brokerId]) {
				flag = false
				break
			}
		}
		return flag
	}
}

type BrokerDatas []*BrokerData

func (bds BrokerDatas) Less(i, j int) bool {
	iq := bds[i]
	jq := bds[j]

	if iq.BrokerName < jq.BrokerName {
		return true
	} else if iq.BrokerName > jq.BrokerName {
		return false
	}
	return false
}

func (bds BrokerDatas) Swap(i, j int) {
	bds[i], bds[j] = bds[j], bds[i]
}

func (bds BrokerDatas) Len() int {
	return len(bds)
}
