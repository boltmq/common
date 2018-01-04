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
package namesrv

import "github.com/boltmq/common/protocol"

// RegisterBrokerBody 注册Broker-请求/响应体
// Author gaoyanlei
// Since 2017/8/22
type RegisterBrokerBody struct {
	TpConfigSerializeWrapper *protocol.TopicConfigSerializeWrapper `json:"topicConfigSerializeWrapper"`
	FilterServerList         []string                              `json:"filterServerList"`
	protocol.RemotingSerializable
}

func NewRegisterBrokerBody(topicConfigWrapper *protocol.TopicConfigSerializeWrapper, filterServerList []string) *RegisterBrokerBody {
	body := new(RegisterBrokerBody)
	body.FilterServerList = filterServerList
	body.TpConfigSerializeWrapper = topicConfigWrapper
	return body
}
