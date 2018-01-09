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
