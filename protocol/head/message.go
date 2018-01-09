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

// PullMessageRequestHeader: 拉取消息请求头信息
// Author: yintongqiang
// Since:  2017/8/14
type PullMessageRequestHeader struct {
	ConsumerGroup        string `json:"consumerGroup"`
	Topic                string `json:"topic"`
	QueueId              int32  `json:"queueId"`
	QueueOffset          int64  `json:"queueOffset"`
	MaxMsgNums           int    `json:"maxMsgNums"`
	SysFlag              int    `json:"sysFlag"`
	CommitOffset         int64  `json:"commitOffset"`
	SuspendTimeoutMillis int    `json:"suspendTimeoutMillis"`
	Subscription         string `json:"subscription"`
	SubVersion           int    `json:"subVersion"`
}

func (header *PullMessageRequestHeader) CheckFields() error {
	return nil
}

// PullMessageResponseHeader: 拉取消息响应头
// Author: yintongqiang
// Since:  2017/8/16
type PullMessageResponseHeader struct {
	SuggestWhichBrokerId int64 `json:"suggestWhichBrokerId"`
	NextBeginOffset      int64 `json:"nextBeginOffset"`
	MinOffset            int64 `json:"minOffset"`
	MaxOffset            int64 `json:"maxOffset"`
}

func (header *PullMessageResponseHeader) CheckFields() error {
	return nil
}

// QueryMessageResponseHeader 查询消息请求头
// Author rongzhihong
// Since 2017/9/18
type QueryMessageRequestHeader struct {
	Topic          string `json:"topic"`
	Key            string `json:"key"`
	MaxNum         int32  `json:"maxNum"`
	BeginTimestamp int64  `json:"beginTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp"`
}

func (query QueryMessageRequestHeader) CheckFields() error {
	return nil
}

// QueryMessageResponseHeader 查询消息返回头
// Author rongzhihong
// Since 2017/9/18
type QueryMessageResponseHeader struct {
	IndexLastUpdateTimestamp int64 `json:"indexLastUpdateTimestamp"`
	IndexLastUpdatePhyoffset int64 `json:"indexLastUpdatePhyoffset"`
}

func (query QueryMessageResponseHeader) CheckFields() error {
	return nil
}

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

// SendMessageResponseHeader: 发送消息响应头
// Author: yintongqiang
// Since:  2017/8/16
type SendMessageResponseHeader struct {
	MsgId         string `json:"msgId"`
	QueueId       int32  `json:"queueId"`
	QueueOffset   int64  `json:"queueOffset"`
	TransactionId string `json:"transactionId"`
}

func (header *SendMessageResponseHeader) CheckFields() error {
	return nil
}

// ViewMessageRequestHeader 根据MsgId查询消息的请求头
// Author rongzhihong
// Since 2017/9/18
type ViewMessageRequestHeader struct {
	Offset uint64 `json:"offset"`
}

func (header *ViewMessageRequestHeader) CheckFields() error {
	return nil
}

// ConsumerSendMsgBackRequestHeader: 消费消息头
// Author: yintongqiang
// Since:  2017/8/17
type ConsumerSendMsgBackRequestHeader struct {
	Offset      int64  `json:"offset"`
	Group       string `json:"group"`
	DelayLevel  int32  `json:"delayLevel"`
	OriginMsgId string `json:"originMsgId"`
	OriginTopic string `json:"originTopic"`
	UnitMode    bool   `json:"unitMode"`
}

func (header *ConsumerSendMsgBackRequestHeader) CheckFields() error {
	return nil
}

// 初始化 ConsumerSendMsgBackRequestHeader
// Author gaoyanlei
// Since 2017/8/17
func NewConsumerSendMsgBackRequestHeader() *ConsumerSendMsgBackRequestHeader {
	return &ConsumerSendMsgBackRequestHeader{
		UnitMode: false,
	}
}
