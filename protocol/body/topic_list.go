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
	"github.com/boltmq/common/protocol/base"
	set "github.com/deckarep/golang-set"
)

// TopicList topic列表
// Author: tianyuliang
// Since: 2017/9/16
type TopicList struct {
	TopicList  set.Set `json:"topicList"`  // topic列表
	BrokerAddr string  `json:"brokerAddr"` // broker地址
}

// NewTopicList 初始化
// Author: tianyuliang
// Since: 2017/9/16
func NewTopicList() *TopicList {
	topicList := &TopicList{
		TopicList: set.NewSet(),
	}
	return topicList
}

// TopicPlusList 拓展Topic列表
// Author: tianyuliang
// Since: 2017/9/16
type TopicPlusList struct {
	TopicList        []string                     `json:"topicList"`        // topic列表
	BrokerAddr       string                       `json:"brokerAddr"`       // broker地址
	TopicQueueTable  map[string][]*base.QueueData `json:"topicQueueTable"`  // 额外增加字段 topic<*route.QueueData>
	ClusterAddrTable map[string][]string          `json:"clusterAddrTable"` // clusterName[set<brokerName>]
}

// NewTopicPlusList 初始化
// Author: tianyuliang
// Since: 2017/9/16
func NewTopicPlusList() *TopicPlusList {
	topicPlusList := &TopicPlusList{
		TopicList:        make([]string, 0),
		TopicQueueTable:  make(map[string][]*base.QueueData),
		ClusterAddrTable: make(map[string][]string),
	}
	return topicPlusList
}

type TopicUpdateConfigWapper struct {
	TopicName      string `json:"topicName"`
	ClusterName    string `json:"clusterName"`
	Order          bool   `json:"order"`
	WriteQueueNums int    `json:"writeQueueNums"`
	ReadQueueNums  int    `json:"readQueueNums"`
	BrokerAddr     string `json:"brokerAddr"`
	BrokerId       int    `json:"brokerId"`
	BrokerName     string `json:"brokerName"`
	Unit           bool   `json:"unit"`
	Perm           int    `json:"perm"`
	TopicSynFlag   int    `json:"topicSynFlag"`
}

type TopicBrokerClusterWapper struct {
	ClusterName          string                   `json:"clusterName"`
	TopicName            string                   `json:"topic"`
	TpUpdateConfigWapper *TopicUpdateConfigWapper `json:"topicConfig"`
}

func NewTopicBrokerClusterWapper(clusterName, topicName string, queueData *base.QueueData) *TopicBrokerClusterWapper {
	topicBrokerClusterWapper := &TopicBrokerClusterWapper{
		ClusterName:          clusterName,
		TopicName:            topicName,
		TpUpdateConfigWapper: NewTopicUpdateConfigWapper(clusterName, topicName, queueData),
	}
	return topicBrokerClusterWapper
}

func NewTopicUpdateConfigWapper(clusterName, topicName string, queueData *base.QueueData) *TopicUpdateConfigWapper {
	topicUpdateConfig := &TopicUpdateConfigWapper{
		ClusterName:    clusterName,
		TopicName:      topicName,
		Order:          false,
		Unit:           false,
		WriteQueueNums: queueData.WriteQueueNums,
		ReadQueueNums:  queueData.ReadQueueNums,
		BrokerName:     queueData.BrokerName,
		Perm:           queueData.Perm,
		TopicSynFlag:   queueData.TopicSynFlag,
	}
	return topicUpdateConfig
}
