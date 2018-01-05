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
	"strings"

	"github.com/boltmq/common/utils/verify"
	"github.com/juju/errors"
)

// UnRegisterBrokerRequestHeader 注销broker-请求头信息
// Author gaoyanlei
// Since 2017/8/22
type UnRegisterBrokerRequestHeader struct {
	BrokerName  string // broker名字
	BrokerAddr  string // broker地址
	ClusterName string // 集群名字
	BrokerId    int    // brokerId
}

func NewUnRegisterBrokerRequestHeader(brokerName, brokerAddr, clusterName string, brokerId int) *UnRegisterBrokerRequestHeader {
	unRegisterBrokerRequestHeader := &UnRegisterBrokerRequestHeader{
		BrokerName:  brokerName,
		BrokerAddr:  brokerAddr,
		ClusterName: clusterName,
		BrokerId:    brokerId,
	}
	return unRegisterBrokerRequestHeader
}

func (header *UnRegisterBrokerRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.BrokerName) == "" {
		return errors.Errorf("UnRegisterBrokerRequestHeader.BrokerName is empty")
	}
	if strings.TrimSpace(header.BrokerAddr) == "" {
		return errors.Errorf("UnRegisterBrokerRequestHeader.BrokerAddr is empty")
	}
	if !verify.CheckIpAndPort(header.BrokerAddr) {
		return errors.Errorf("UnRegisterBrokerRequestHeader.BrokerAddr[%s] is invalid.", header.BrokerAddr)
	}
	if strings.TrimSpace(header.ClusterName) == "" {
		return errors.Errorf("UnRegisterBrokerRequestHeader.ClusterName is empty")
	}
	if header.BrokerId < 0 {
		return errors.Errorf("UnRegisterBrokerRequestHeader.BrokerId[%d] is invalid", header.BrokerId)
	}
	return nil
}
