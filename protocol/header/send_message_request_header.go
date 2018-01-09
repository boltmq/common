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
package header

// SendMessageRequestHeader: 发送消息请求头信息
// Author: yintongqiang
// Since:  2017/8/10
type SendMessageRequestHeader struct {
	ProducerGroup         string `json:"producerGroup"`
	Topic                 string `json:"topic"`
	DefaultTopic          string `json:"defaultTopic"`
	DefaultTopicQueueNums int32  `json:"defaultTopicQueueNums"`
	QueueId               int32  `json:"queueId"`
	SysFlag               int32  `json:"sysFlag"`
	BornTimestamp         int64  `json:"bornTimestamp"`
	Flag                  int32  `json:"flag"`
	Properties            string `json:"properties"`
	ReconsumeTimes        int32  `json:"reconsumeTimes"`
	UnitMode              bool   `json:"unitMode"`
}

func (header *SendMessageRequestHeader) CheckFields() error {
	return nil
}

// SendMessageRequestHeaderV2: 为减少网络传输数量准备
// Author: yintongqiang
// Since:  2017/8/10
type SendMessageRequestHeaderV2 struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D int32  `json:"d"`
	E int32  `json:"e"`
	F int32  `json:"f"`
	G int64  `json:"g"`
	H int32  `json:"h"`
	I string `json:"i"`
	J int32  `json:"j"`
	K bool   `json:"k"`
}

func (header *SendMessageRequestHeaderV2) CheckFields() error {
	return nil
}

func CreateSendMessageRequestHeaderV2(v1 *SendMessageRequestHeader) *SendMessageRequestHeaderV2 {
	v2 := &SendMessageRequestHeaderV2{}
	v2.A = v1.ProducerGroup
	v2.B = v1.Topic
	v2.C = v1.DefaultTopic
	v2.D = v1.DefaultTopicQueueNums
	v2.E = v1.QueueId
	v2.F = v1.SysFlag
	v2.G = v1.BornTimestamp
	v2.H = v1.Flag
	v2.I = v1.Properties
	v2.J = v1.ReconsumeTimes
	v2.K = v1.UnitMode
	return v2
}

// CreateSendMessageRequestHeaderV1 v2转v1
// Author gaoyanlei
// Since 2017/8/15
func CreateSendMessageRequestHeaderV1(v2 *SendMessageRequestHeaderV2) *SendMessageRequestHeader {
	v1 := new(SendMessageRequestHeader)
	v1.ProducerGroup = v2.A
	v1.Topic = v2.B
	v1.DefaultTopic = v2.C
	v1.DefaultTopicQueueNums = v2.D
	v1.QueueId = v2.E
	v1.SysFlag = v2.F
	v1.BornTimestamp = v2.G
	v1.Flag = v2.H
	v1.Properties = v2.I
	v1.ReconsumeTimes = v2.J
	v1.UnitMode = v2.K
	return v1
}
