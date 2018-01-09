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

// GetConsumerConnectionListRequestHeader 获得Toipc统计信息的请求头
// Author rongzhihong
// Since 2017/9/19
type GetConsumerConnectionListRequestHeader struct {
	ConsumerGroup string `json:"consumerGroup"`
}

func (header *GetConsumerConnectionListRequestHeader) CheckFields() error {
	return nil
}

// NewGetConsumerConnectionListRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/1
func NewGetConsumerConnectionListRequestHeader(consumerGroup string) *GetConsumerConnectionListRequestHeader {
	consumerConnectionListRequestHeader := &GetConsumerConnectionListRequestHeader{
		ConsumerGroup: consumerGroup,
	}
	return consumerConnectionListRequestHeader
}
