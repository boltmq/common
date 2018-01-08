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
