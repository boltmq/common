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

import (
	"strings"

	"github.com/go-errors/errors"
)

type GetTopicsByClusterRequestHeader struct {
	Cluster string `json:"cluster"` // 集群名称
	Extra   bool   `json:"extra"`   // 是否额外查询topic、cluster对应关系
}

func (header *GetTopicsByClusterRequestHeader) CheckFields() error {
	return nil
}

// DeleteTopicInNamesrvRequestHeader 删除Topic-请求头
// Author: tianyuliang
// Since: 2017/9/4
type DeleteTopicInNamesrvRequestHeader struct {
	Topic string
}

func (header *DeleteTopicInNamesrvRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.Topic) == "" {
		return errors.Errorf("topic is empty")
	}
	return nil
}

// WipeWritePermOfBrokerResponseHeader 优雅地向Broker写数据-响应头
// Author: tianyuliang
// Since: 2017/9/4
type WipeWritePermOfBrokerResponseHeader struct {
	WipeTopicCount int
}

func (header *WipeWritePermOfBrokerResponseHeader) CheckFields() error {
	return nil
}

// WipeWritePermOfBrokerRequestHeader 优雅地向Broker写数据-请求头
// Author: tianyuliang
// Since: 2017/9/4
type WipeWritePermOfBrokerRequestHeader struct {
	BrokerName string // broker名称
}

func (header *WipeWritePermOfBrokerRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.BrokerName) == "" {
		return errors.Errorf("broker name is empty")
	}
	return nil
}

// GetRouteInfoRequestHeader: 获取topic路由信息头
// Author: yintongqiang
// Since:  2017/8/16
type GetRouteInfoRequestHeader struct {
	Topic string
}

func (header *GetRouteInfoRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.Topic) == "" {
		return errors.Errorf("topic is empty")
	}
	return nil
}
