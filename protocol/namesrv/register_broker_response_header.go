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

// RegisterBrokerRequestHeader 注册Broker-响应头
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
