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

// LockBatchRequestBody 锁队列响应头
// Author rongzhihong
// Since 2017/9/19
type LockBatchResponseBody struct {
	LockOKMQSet set.Set `json:"lockOKMQSet"`
}

func NewLockBatchResponseBody() *LockBatchResponseBody {
	body := new(LockBatchResponseBody)
	body.LockOKMQSet = set.NewSet()
	return body
}
