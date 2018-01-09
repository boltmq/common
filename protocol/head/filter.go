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

// RegisterFilterServerRequestHeader 注册过滤器的请求头
// Author rongzhihong
// Since 2017/9/19
type RegisterFilterServerRequestHeader struct {
	FilterServerAddr string `json:"filterServerAddr"`
}

func (header *RegisterFilterServerRequestHeader) CheckFields() error {
	return nil
}

// RegisterFilterServerResponseHeader 注册过滤器的返回头
// Author rongzhihong
// Since 2017/9/19
type RegisterFilterServerResponseHeader struct {
	BrokerName string `json:"brokerName"`
	BrokerId   int64  `json:"brokerId"`
}

func (header *RegisterFilterServerResponseHeader) CheckFields() error {
	return nil
}
