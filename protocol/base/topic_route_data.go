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

	"github.com/pquerna/ffjson/ffjson"
)

type TopicRouteData struct {
	OrderTopicConf    string              `json:"orderTopicConf"`
	QueueDatas        []*QueueData        `json:"queueDatas"`
	BrokerDatas       []*BrokerData       `json:"brokerDatas"`
	FilterServerTable map[string][]string `json:"filterServerTable"`
}

func NewTopicRouteData() *TopicRouteData {
	topicRouteData := &TopicRouteData{
		QueueDatas:        make([]*QueueData, 0),
		BrokerDatas:       make([]*BrokerData, 0),
		FilterServerTable: make(map[string][]string, 0),
	}
	return topicRouteData
}

func (trd *TopicRouteData) Decode(data []byte) error {
	return ffjson.Unmarshal(data, trd)
}

func (trd *TopicRouteData) Encode() ([]byte, error) {
	return ffjson.Marshal(trd)
}

func (trd *TopicRouteData) CloneTopicRouteData() *TopicRouteData {
	queueDatas := []*QueueData{}
	brokerDatas := []*BrokerData{}
	for _, queueData := range trd.QueueDatas {
		queueDatas = append(queueDatas, &QueueData{BrokerName: queueData.BrokerName, Perm: queueData.Perm,
			WriteQueueNums: queueData.WriteQueueNums, ReadQueueNums: queueData.ReadQueueNums, TopicSynFlag: queueData.TopicSynFlag})
	}
	for _, brokerData := range trd.BrokerDatas {
		brokerDatas = append(brokerDatas, &BrokerData{BrokerName: brokerData.BrokerName, BrokerAddrs: brokerData.BrokerAddrs})
	}
	return &TopicRouteData{
		OrderTopicConf: trd.OrderTopicConf,
		QueueDatas:     queueDatas,
		BrokerDatas:    brokerDatas,
	}
}

func (routeData *TopicRouteData) Equals(v interface{}) bool {
	if v == nil {
		return false
	}
	rData1, ok := v.(TopicRouteData)
	var rData2 *TopicRouteData
	if !ok {
		rData2, ok = v.(*TopicRouteData)
	}
	if !ok {
		return ok
	}
	if rData2 == nil {
		if !strings.EqualFold(routeData.OrderTopicConf, rData1.OrderTopicConf) {
			return false
		}
		if len(routeData.BrokerDatas) != len(rData1.BrokerDatas) {
			return false
		}
		if len(routeData.QueueDatas) != len(rData1.QueueDatas) {
			return false
		}
		var flagB bool = true
		for i := 0; i < len(routeData.BrokerDatas); i++ {
			if !routeData.BrokerDatas[i].Equals(rData1.BrokerDatas[i]) {
				flagB = false
				break
			}
		}
		if !flagB {
			return flagB
		}
		var flagQ bool = true
		for i := 0; i < len(routeData.QueueDatas); i++ {
			if !routeData.QueueDatas[i].Equals(rData1.QueueDatas[i]) {
				flagQ = false
				break
			}
		}
		if !flagQ {
			return flagQ
		}

	} else {
		if !strings.EqualFold(routeData.OrderTopicConf, rData2.OrderTopicConf) {
			return false
		}
		if len(routeData.BrokerDatas) != len(rData2.BrokerDatas) {
			return false
		}
		if len(routeData.QueueDatas) != len(rData2.QueueDatas) {
			return false
		}
		var flagB bool = true
		for i := 0; i < len(routeData.BrokerDatas); i++ {
			if !routeData.BrokerDatas[i].Equals(rData2.BrokerDatas[i]) {
				flagB = false
				break
			}
		}
		if !flagB {
			return flagB
		}
		var flagQ bool = true
		for i := 0; i < len(routeData.QueueDatas); i++ {
			if !routeData.QueueDatas[i].Equals(rData2.QueueDatas[i]) {
				flagQ = false
				break
			}
		}
		if !flagQ {
			return flagQ
		}
	}
	return true
}

func (trd *TopicRouteData) String() string {
	if trd == nil {
		return ""
	}

	data1 := ""
	queueDatas := make([]string, 0, len(trd.QueueDatas))
	if trd.QueueDatas != nil && len(trd.QueueDatas) > 0 {
		for _, v := range trd.QueueDatas {
			queueData := v.String()
			queueDatas = append(queueDatas, queueData)
		}
		data1 = strings.Join(queueDatas, ",")
	}

	data2 := ""
	brokerDatas := make([]string, 0, len(trd.BrokerDatas))
	if trd.BrokerDatas != nil && len(trd.BrokerDatas) > 0 {
		for _, v := range trd.BrokerDatas {
			brokerData := v.String()
			brokerDatas = append(brokerDatas, brokerData)
		}
		data2 = strings.Join(brokerDatas, ",")
	}

	vals := make([]string, 0, len(trd.FilterServerTable))
	if trd.FilterServerTable != nil && len(trd.FilterServerTable) > 0 {
		for brokerAddr, filterServer := range trd.FilterServerTable {
			filterServerList := ""
			if filterServer != nil && len(filterServer) > 0 {
				filterServerList = strings.Join(filterServer, ",")
			}
			val := fmt.Sprintf("brokerAddr=%s, filterServer=[%s]", brokerAddr, filterServerList)
			vals = append(vals, val)
		}
	}
	filterServerTable := strings.Join(vals, ",")

	format := "TopicRouteData {orderTopicConf=%s, %s, %s, %s}"
	info := fmt.Sprintf(format, trd.OrderTopicConf, data1, data2, filterServerTable)
	return info
}
