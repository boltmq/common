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
package heartbeat

import (
	"github.com/pquerna/ffjson/ffjson"
)

// HeartbeatData 客户端与broker心跳结构体
// Author: yintongqiang
// Since:  2017/8/8
type HeartbeatData struct {
	ClientID      string         `json:"clientID"`
	ProducerDatas []ProducerData `json:"producerDataSet"`
	ConsumerDatas []ConsumerData `json:"consumerDataSet"`
}

// ProducerData 生产者心跳信息
// Author: yintongqiang
// Since:  2017/8/9
type ProducerData struct {
	GroupName string `json:"groupName"` // 生产者组名称(ProducerGroupName)
}

type HeartbeatDataPlus struct {
	ClientID        string             `json:"clientID"`
	ProducerDataSet []ProducerData     `json:"producerDataSet"`
	ConsumerDataSet []ConsumerDataPlus `json:"consumerDataSet"`
}

// NewHeartbeatData 初始化broker心跳结构体
// Author: rongzhihong
// Since:  2017/9/21
func NewHeartbeatData() *HeartbeatData {
	return &HeartbeatData{}
}

func (heartbeatData *HeartbeatData) Encode() []byte {
	bytes, _ := ffjson.Marshal(heartbeatData)
	return bytes
}

func (heartbeatData *HeartbeatData) Decode(byte []byte) *HeartbeatData {
	ffjson.Unmarshal(byte, heartbeatData)
	return heartbeatData
}

func (heartbeatDataPlus *HeartbeatDataPlus) Decode(byte []byte) *HeartbeatDataPlus {
	ffjson.Unmarshal(byte, heartbeatDataPlus)
	return heartbeatDataPlus
}
