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

// GetMaxOffsetRequestHeader: 获取队列最大offset
// Author: yintongqiang
// Since:  2017/8/23
type GetMaxOffsetRequestHeader struct {
	Topic   string `json:"topic"`
	QueueId int    `json:"queueId"`
}

func NewGetMaxOffsetRequestHeader() *GetMaxOffsetRequestHeader {
	header := new(GetMaxOffsetRequestHeader)
	return header
}

func (header *GetMaxOffsetRequestHeader) CheckFields() error {
	return nil
}

// GetMaxOffsetResponseHeader: 最大偏移响应头
// Author: yintongqiang
// Since:  2017/8/23
type GetMaxOffsetResponseHeader struct {
	Offset int64 `json:"offset"`
}

func (header *GetMaxOffsetResponseHeader) CheckFields() error {
	return nil
}

// SearchOffsetRequestHeader 查询偏移量的请求头
// Author rongzhihong
// Since 2017/9/19
type SearchOffsetRequestHeader struct {
	Topic     string
	QueueId   int32
	Timestamp int64
}

func (header *SearchOffsetRequestHeader) CheckFields() error {
	return nil
}

// SearchOffsetResponseHeader 查询偏移量的返回头
// Author rongzhihong
// Since 2017/9/19
type SearchOffsetResponseHeader struct {
	Offset int64
}

func (header *SearchOffsetResponseHeader) CheckFields() error {
	return nil
}

// GetMinOffsetRequestHeader 获得最小偏移量的请求头
// Author rongzhihong
// Since 2017/9/19
type GetMinOffsetRequestHeader struct {
	Topic   string `json:"topic"`
	QueueId int32  `json:"queueId"`
}

func (header *GetMinOffsetRequestHeader) CheckFields() error {
	return nil
}

// GetMinOffsetResponseHeader 获得最小偏移量的返回头
// Author rongzhihong
// Since 2017/9/19
type GetMinOffsetResponseHeader struct {
	Offset int64 `json:"offset"`
}

func (header *GetMinOffsetResponseHeader) CheckFields() error {
	return nil
}

// ResetOffsetRequestHeader 重置偏移量的请求头
// Author rongzhihong
// Since 2017/9/18
type ResetOffsetRequestHeader struct {
	Topic     string
	Group     string
	Timestamp int64
	IsForce   bool
}

func (req *ResetOffsetRequestHeader) CheckFields() error {
	return nil
}

// QueryCorrectionOffsetRequestHeader 查找被修正 offset (转发组件）的请求头
// Author rongzhihong
// Since 2017/9/19
type QueryCorrectionOffsetRequestHeader struct {
	FilterGroups string `json:"filterGroups"`
	CompareGroup string `json:"compareGroup"`
	Topic        string `json:"topic"`
}

func (header *QueryCorrectionOffsetRequestHeader) CheckFields() error {
	return nil
}

// CloneGroupOffsetRequestHeader 克隆cloneGroupOffset的请求头
// Author rongzhihong
// Since 2017/9/19
type CloneGroupOffsetRequestHeader struct {
	SrcGroup  string `json:"srcGroup"`
	DestGroup string `json:"destGroup"`
	Topic     string `json:"topic"`
	Offline   bool   `json:"offline"`
}

func (header *CloneGroupOffsetRequestHeader) CheckFields() error {
	return nil
}

// NewCloneGroupOffsetRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/1
func NewCloneGroupOffsetRequestHeader(srcGroup, destGroup, topic string, isOffline bool) *CloneGroupOffsetRequestHeader {
	cloneGroupOffsetRequestHeader := &CloneGroupOffsetRequestHeader{
		SrcGroup:  srcGroup,
		DestGroup: destGroup,
		Topic:     topic,
		Offline:   isOffline,
	}
	return cloneGroupOffsetRequestHeader
}

// QueryConsumerOffsetRequestHeader: 查询消费offset
// Author: yintongqiang
// Since:  2017/8/24
type QueryConsumerOffsetRequestHeader struct {
	ConsumerGroup string `json:"consumerGroup"`
	Topic         string `json:"topic"`
	QueueId       int32  `json:"queueId"`
}

func (header *QueryConsumerOffsetRequestHeader) CheckFields() error {
	return nil
}

// QueryConsumerOffsetResponseHeader: 查询offset响应头
// Author: yintongqiang
// Since:  2017/8/24
type QueryConsumerOffsetResponseHeader struct {
	Offset int64 `json:"offset"`
}

func (header *QueryConsumerOffsetResponseHeader) CheckFields() error {
	return nil
}

// UpdateConsumerOffsetRequestHeader: 更新消费offset的请求头
// Author: yintongqiang
// Since:  2017/8/11
type UpdateConsumerOffsetRequestHeader struct {
	ConsumerGroup string `json:"consumerGroup"`
	Topic         string `json:"topic"`
	QueueId       int    `json:"queueId"`
	CommitOffset  int64  `json:"commitOffset"`
}

func (header *UpdateConsumerOffsetRequestHeader) CheckFields() error {
	return nil
}
