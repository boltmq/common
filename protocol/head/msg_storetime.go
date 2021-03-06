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

// GetEarliestMsgStoretimeResponseHeader 获得早期消息存储时间的返回头
// Author rongzhihong
// Since 2017/9/19
type GetEarliestMsgStoretimeResponseHeader struct {
	Timestamp int64
}

func (header *GetEarliestMsgStoretimeResponseHeader) CheckFields() error {
	return nil
}

// GetEarliestMsgStoretimeRequestHeader 获得早期消息存储时间的请求头
// Author rongzhihong
// Since 2017/9/19
type GetEarliestMsgStoretimeRequestHeader struct {
	Topic   string
	QueueId int32
}

func (header *GetEarliestMsgStoretimeRequestHeader) CheckFields() error {
	return nil
}
