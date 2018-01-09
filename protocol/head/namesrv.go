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
	"fmt"
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

// RegisterBrokerRequestHeader 注册Broker-请求头
// Author gaoyanlei
// Since 2017/8/22
type RegisterBrokerRequestHeader struct {
	BrokerName   string // broker名称
	BrokerAddr   string // broker地址(ip:port)
	ClusterName  string // 集群名字
	HaServerAddr string // ha地址
	BrokerId     int64  // brokerId
}

func (header *RegisterBrokerRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.BrokerName) == "" {
		return fmt.Errorf("RegisterBrokerRequestHeader.BrokerName is empty")
	}
	if strings.TrimSpace(header.BrokerAddr) == "" {
		return fmt.Errorf("RegisterBrokerRequestHeader.BrokerAddr is empty")
	}
	if !verify.CheckIpAndPort(header.BrokerAddr) {
		return fmt.Errorf("RegisterBrokerRequestHeader.BrokerAddr[%s] is invalid.", header.BrokerAddr)
	}
	if strings.TrimSpace(header.ClusterName) == "" {
		return fmt.Errorf("RegisterBrokerRequestHeader.ClusterName is empty")
	}
	if strings.TrimSpace(header.HaServerAddr) == "" {
		return fmt.Errorf("RegisterBrokerRequestHeader.HaServerAddr is empty")
	}
	if header.BrokerId < 0 {
		return fmt.Errorf("RegisterBrokerRequestHeader.BrokerId[%d] is invalid", header.BrokerId)
	}
	return nil
}

func NewRegisterBrokerRequestHeader(clusterName, brokerAddr, brokerName, haServerAddr string, brokerId int64) *RegisterBrokerRequestHeader {
	registerBrokerRequestHeader := &RegisterBrokerRequestHeader{
		BrokerName:   brokerName,
		BrokerAddr:   brokerAddr,
		ClusterName:  clusterName,
		HaServerAddr: haServerAddr,
		BrokerId:     brokerId,
	}

	return registerBrokerRequestHeader
}

// RegisterBrokerResponseHeader 注册Broker-响应头
// Author gaoyanlei
// Since 2017/8/22
type RegisterBrokerResponseHeader struct {
	HaServerAddr string // broker备节点地址
	MasterAddr   string // broker主节点地址
}

func (header *RegisterBrokerResponseHeader) CheckFields() error {
	return nil
}

func NewRegisterBrokerResponseHeader(haServerAddr, masterAddr string) *RegisterBrokerResponseHeader {
	registerBrokerResponseHeader := &RegisterBrokerResponseHeader{
		HaServerAddr: haServerAddr,
		MasterAddr:   masterAddr,
	}
	return registerBrokerResponseHeader
}
