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
package body

import (
	"fmt"

	"github.com/boltmq/common/protocol/heartbeat"
	set "github.com/deckarep/golang-set"
	concurrent "github.com/fanliao/go-concurrentMap"
)

// Connection 连接信息
// Author rongzhihong
// Since 2017/9/19
type Connection struct {
	ClientId   string `json:"clientId"`   // 客户端实例
	ClientAddr string `json:"clientAddr"` // 客户端地址
	Language   string `json:"language"`   // 开发语言
	Version    int32  `json:"version"`    // mq发行版本号
}

// NewConnection 初始化Connection
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/16
func NewConnection(clientId, clientAddr, language string, version int32) *Connection {
	conn := &Connection{
		ClientId:   clientId,
		ClientAddr: clientAddr,
		Language:   language,
		Version:    version,
	}
	return conn
}

// String 格式化Connection结构体的数据
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/16
func (conn *Connection) String() string {
	if conn == nil {
		return "<nil>"
	}

	return fmt.Sprintf("Connection {clientAddr=%s, clientId=%s, language=%s, version=%d}",
		conn.ClientId, conn.ClientAddr, conn.Language, conn.Version)
}

// ConsumerConnection 消费者连接信息
// Author rongzhihong
// Since 2017/9/19
type ConsumerConnection struct {
	ConnectionSet     set.Set                    `json:"connectionSet"`     // type: Connection
	SubscriptionTable *concurrent.ConcurrentMap  `json:"subscriptionTable"` // topic<*SubscriptionDataPlus>
	ConsumeType       heartbeat.ConsumeType      `json:"consumeType"`
	MessageModel      heartbeat.MessageModel     `json:"messageModel"`
	ConsumeFromWhere  heartbeat.ConsumeFromWhere `json:"consumeFromWhere"`
}

// ConsumerConnectionPlus 消费者连接信息(处理set集合无法反序列化问题)
// Author rongzhihong
// Since 2017/9/19
type ConsumerConnectionPlus struct {
	ConnectionSet     []*Connection                              `json:"connectionSet"`     // type: Connection
	SubscriptionTable map[string]*heartbeat.SubscriptionDataPlus `json:"subscriptionTable"` // topic<*SubscriptionDataPlus>
	ConsumeType       heartbeat.ConsumeType                      `json:"consumeType"`
	MessageModel      heartbeat.MessageModel                     `json:"messageModel"`
	ConsumeFromWhere  heartbeat.ConsumeFromWhere                 `json:"consumeFromWhere"`
}

// NewConsumerConnection 初始化
// Author rongzhihong
// Since 2017/9/19
func NewConsumerConnectionPlus() *ConsumerConnectionPlus {
	connect := new(ConsumerConnectionPlus)
	connect.ConnectionSet = make([]*Connection, 0)
	connect.SubscriptionTable = make(map[string]*heartbeat.SubscriptionDataPlus)
	return connect
}

// ToConsumerConnection 转化为ConsumerConnection
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/13
func (plus *ConsumerConnectionPlus) ToConsumerConnection() *ConsumerConnection {
	consumerConnection := &ConsumerConnection{
		SubscriptionTable: concurrent.NewConcurrentMap(),
		ConsumeType:       plus.ConsumeType,
		MessageModel:      plus.MessageModel,
		ConsumeFromWhere:  plus.ConsumeFromWhere,
		ConnectionSet:     set.NewSet(),
	}

	if plus.SubscriptionTable != nil {
		for k, v := range plus.SubscriptionTable {
			consumerConnection.SubscriptionTable.Put(k, v)
		}
	}

	if plus.ConnectionSet != nil && len(plus.ConnectionSet) > 0 {
		for _, connect := range plus.ConnectionSet {
			consumerConnection.ConnectionSet.Add(connect)
		}
	}
	return consumerConnection
}

// NewConsumerConnection 初始化
// Author rongzhihong
// Since 2017/9/19
func NewConsumerConnection() *ConsumerConnection {
	connect := new(ConsumerConnection)
	connect.ConnectionSet = set.NewSet()
	connect.SubscriptionTable = concurrent.NewConcurrentMap()
	return connect
}

// ComputeMinVersion 计算最小版本号
// Author rongzhihong
// Since 2017/9/19
func (consumerConn *ConsumerConnection) ComputeMinVersion() int32 {
	minVersion := int32(2147483647)
	iterator := consumerConn.ConnectionSet.Iterator()
	for item := range iterator.C {
		if c, ok := item.(*Connection); ok {
			if c.Version < minVersion {
				minVersion = c.Version
			}
		}
	}
	return minVersion
}
