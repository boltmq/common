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
package heartbeat

// ConsumeType 消费类型枚举
// Author: yintongqiang
// Since:  2017/8/8
type ConsumeType int

const (
	CONSUME_ACTIVELY  ConsumeType = iota // 主动方式消费(常见于pull消费)
	CONSUME_PASSIVELY                    // 被动方式消费(常见于push消费)
)

// String 转化为对应的字符串
// Author: yintongqiang
// Since:  2017/8/8
func (ct ConsumeType) String() string {
	switch ct {
	case CONSUME_ACTIVELY:
		return "CONSUME_ACTIVELY"
	case CONSUME_PASSIVELY:
		return "CONSUME_PASSIVELY"
	default:
		return "Unknown"
	}
}

// MessageModel 消息类型枚举
// Author: yintongqiang
// Since:  2017/8/8
type MessageModel int

const (
	BROADCASTING MessageModel = iota // 广播消费模式
	CLUSTERING                       // 集群消费模式
)

// 消费类型枚举
// Author: yintongqiang
// Since:  2017/8/8
func (m MessageModel) String() string {
	switch m {
	case BROADCASTING:
		return "BROADCASTING"
	case CLUSTERING:
		return "CLUSTERING"
	default:
		return "Unknown"
	}
}

// ConsumeFromWhere 从哪里开始消费
// Author: yintongqiang
// Since:  2017/8/8
type ConsumeFromWhere int

const (
	// 一个新的订阅组第一次启动从队列的最后位置开始消费<br>
	// 后续再启动接着上次消费的进度开始消费
	CONSUME_FROM_LAST_OFFSET ConsumeFromWhere = iota

	// 一个新的订阅组第一次启动从队列的最前位置开始消费<br>
	// 后续再启动接着上次消费的进度开始消费
	CONSUME_FROM_FIRST_OFFSET

	// 一个新的订阅组第一次启动从指定时间点开始消费<br>
	// 后续再启动接着上次消费的进度开始消费<br>
	// 时间点设置参见DefaultMQPushConsumer.consumeTimestamp参数
	CONSUME_FROM_TIMESTAMP
)

// String 从哪里开始消费，转化为相应的字符串
// Author: yintongqiang
// Since:  2017/8/8
func (cfw ConsumeFromWhere) String() string {
	switch cfw {
	case CONSUME_FROM_LAST_OFFSET:
		return "CONSUME_FROM_LAST_OFFSET"
	case CONSUME_FROM_FIRST_OFFSET:
		return "CONSUME_FROM_FIRST_OFFSET"
	case CONSUME_FROM_TIMESTAMP:
		return "CONSUME_FROM_TIMESTAMP"
	default:
		return "Unknown"
	}
}
