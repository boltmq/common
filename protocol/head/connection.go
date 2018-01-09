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

// GetConsumerConnectionsRequestHeader 获得Toipc统计信息的请求头
// Author rongzhihong
// Since 2017/9/19
type GetConsumerConnectionsRequestHeader struct {
	ConsumerGroup string `json:"consumerGroup"`
}

func (header *GetConsumerConnectionsRequestHeader) CheckFields() error {
	return nil
}

// NewGetConsumerConnectionsRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/1
func NewGetConsumerConnectionsRequestHeader(consumerGroup string) *GetConsumerConnectionsRequestHeader {
	ccrh := &GetConsumerConnectionsRequestHeader{
		ConsumerGroup: consumerGroup,
	}
	return ccrh
}

// GetProducerConnectionsRequestHeader 获得生产者连接信息请求头
// Author rongzhihong
// Since 2017/9/19
type GetProducerConnectionsRequestHeader struct {
	ProducerGroup string `json:"producerGroup"`
}

func (header *GetProducerConnectionsRequestHeader) CheckFields() error {
	return nil
}

// NewGetProducerConnectionsRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/6
func NewGetProducerConnectionsRequestHeader(producerGroup string) *GetProducerConnectionsRequestHeader {
	pcrh := &GetProducerConnectionsRequestHeader{
		ProducerGroup: producerGroup,
	}
	return pcrh
}
