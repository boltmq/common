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
package system

import (
	"sync"
	"time"
)

// WaitNotify 用来做线程之间异步通知
// Author zhoufei
// Since 2017/10/23
type WaitNotify struct {
	waitingThreadTable map[int64]*Notify
	hasNotified        bool
	notify             *Notify
	mutex              *sync.Mutex
}

func NewWaitNotify() *WaitNotify {
	return &WaitNotify{
		waitingThreadTable: make(map[int64]*Notify),
		hasNotified:        false,
		notify:             CreateNotify(),
		mutex:              new(sync.Mutex),
	}
}

func (wno *WaitNotify) Wakeup(interval int64) {
	wno.mutex.Lock()
	defer wno.mutex.Unlock()

	if !wno.hasNotified {
		wno.hasNotified = true
	}
}

func (wno *WaitNotify) WaitForRunning(interval int64) {
	wno.mutex.Lock()
	defer wno.mutex.Unlock()

	if wno.hasNotified {
		wno.hasNotified = false
		return
	}

	wno.notify.WaitTimeout(time.Duration(interval) * time.Millisecond)
	wno.hasNotified = false
}
