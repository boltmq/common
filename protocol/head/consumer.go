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

import "fmt"

// GetConsumerStatusRequestHeader 获得消费者状态的请求头
// Author rongzhihong
// Since 2017/9/19
type GetConsumerStatusRequestHeader struct {
	Topic      string `json:"topic"`
	Group      string `json:"group"`
	ClientAddr string `json:"clientAddr"`
}

func (header *GetConsumerStatusRequestHeader) CheckFields() error {
	return nil
}

// NewGetConsumerStatusRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/6
func NewGetConsumerStatusRequestHeader(topic, group, clientAddr string) *GetConsumerStatusRequestHeader {
	consumerStatusRequestHeader := &GetConsumerStatusRequestHeader{
		Topic:      topic,
		Group:      group,
		ClientAddr: clientAddr,
	}
	return consumerStatusRequestHeader
}

// QueryConsumerTimeSpanRequestHeader 根据 topic 和 group 获取消息的时间跨度的请求头
// Author rongzhihong
// Since 2017/9/19
type QueryConsumerTimeSpanRequestHeader struct {
	Topic string `json:"topic"`
	Group string `json:"group"`
}

func (header *QueryConsumerTimeSpanRequestHeader) CheckFields() error {
	return nil
}

// GetConsumerRunningInfoRequestHeader 获取Consumer内存数据结构的请求头
// Author rongzhihong
// Since 2017/9/19
type GetConsumerRunningInfoRequestHeader struct {
	ConsumerGroup string `json:"consumerGroup"`
	ClientId      string `json:"clientId"`
	JstackEnable  bool   `json:"jstackEnable"`
}

func (header *GetConsumerRunningInfoRequestHeader) CheckFields() error {
	return nil
}

// NewGetConsumerRunningInfoRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/6
func NewGetConsumerRunningInfoRequestHeader(consumerGroup, clientId string, jstackEnable bool) *GetConsumerRunningInfoRequestHeader {
	consumerRunningInfoRequestHeader := &GetConsumerRunningInfoRequestHeader{
		ConsumerGroup: consumerGroup,
		ClientId:      clientId,
		JstackEnable:  jstackEnable,
	}
	return consumerRunningInfoRequestHeader
}

// NotifyConsumerIdsChangedRequestHeader 通知请求头部
// Author: rongzhihong
// Since:  2017/9/11
type NotifyConsumerIdsChangedRequestHeader struct {
	ConsumerGroup string
}

// CheckFields
// Author: rongzhihong
// Since:  2017/9/11
func (notify *NotifyConsumerIdsChangedRequestHeader) CheckFields() error {
	return nil
}

// GetConsumersByGroupRequestHeader: 获取消费列表
// Author: yintongqiang
// Since:  2017/8/11
type GetConsumersByGroupRequestHeader struct {
	ConsumerGroup string
}

func (header *GetConsumersByGroupRequestHeader) CheckFields() error {
	return nil
}

type ConsumeMessageDirectlyResult struct {
	Order          bool
	AutoCommit     bool
	ConsumeResult  CMResult
	Remark         string
	SpentTimeMills int64
}

func (r *ConsumeMessageDirectlyResult) String() string {
	if r == nil {
		return "<nil>"
	}

	return fmt.Sprintf("{Order=%t, AutoCommit=%t, ConsumeResult=%s, SpentTimeMills=%d, Remark=%s}",
		r.Order, r.AutoCommit, r.ConsumeResult, r.SpentTimeMills, r.Remark)
}

type CMResult int

const (
	CR_SUCCESS CMResult = iota
	CR_LATER
	CR_ROLLBACK
	CR_COMMIT
	CR_THROW_EXCEPTION
	CR_RETURN_NULL
)

func (cr CMResult) String() string {
	switch cr {
	case CR_SUCCESS:
		return "CR_SUCCESS"
	case CR_LATER:
		return "CR_LATER"
	case CR_ROLLBACK:
		return "CR_ROLLBACK"
	case CR_COMMIT:
		return "CR_COMMIT"
	case CR_THROW_EXCEPTION:
		return "CR_THROW_EXCEPTION"
	case CR_RETURN_NULL:
		return "CR_RETURN_NULL"
	default:
		return "UnKnown"
	}
}
