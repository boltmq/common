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
package body

import (
	set "github.com/deckarep/golang-set"
)

// UnlockBatchRequest 解锁队列响应头
// Author rongzhihong
// Since 2017/9/19
type UnlockBatchRequest struct {
	ConsumerGroup string  `json:"consumerGroup"`
	ClientId      string  `json:"clientId"`
	MQSet         set.Set `json:"mqSet"`
}

func NewUnlockBatchRequest() *UnlockBatchRequest {
	body := new(UnlockBatchRequest)
	body.MQSet = set.NewSet()
	return body
}

// LockBatchRequest 锁队列请求头
// Author rongzhihong
// Since 2017/9/19
type LockBatchRequest struct {
	ConsumerGroup string  `json:"consumerGroup"`
	ClientId      string  `json:"clientId"`
	MQSet         set.Set `json:"mq_set"`
}

func NewLockBatchRequest() *LockBatchRequest {
	body := new(LockBatchRequest)
	body.MQSet = set.NewSet()
	return body
}

// LockBatchResponse 锁队列响应头
// Author rongzhihong
// Since 2017/9/19
type LockBatchResponse struct {
	LockOKMQSet set.Set `json:"lockOKMQSet"`
}

func NewLockBatchResponse() *LockBatchResponse {
	body := new(LockBatchResponse)
	body.LockOKMQSet = set.NewSet()
	return body
}
