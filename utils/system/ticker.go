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

import "time"

type Ticker struct {
	tm    *time.Timer
	d     time.Duration
	delay time.Duration
	fn    func()
	isRun bool
	wait  bool
	over  chan interface{}
}

func NewTicker(wait bool, delay, period time.Duration, fn func()) *Ticker {
	return &Ticker{
		d:     period,
		delay: delay,
		fn:    fn,
		wait:  wait,
	}
}

func (t *Ticker) Start() {
	if t.isRun {
		return
	}

	go t.start()
}

func (t *Ticker) start() {
	if t.wait {
		t.over = make(chan interface{})
	}

	t.tm = time.NewTimer(t.d)
	t.isRun = true

	if t.delay >= 0 {
		time.Sleep(t.delay)
		t.fn()
	}

	for {
		select {
		case <-t.tm.C:
			t.fn()
			if !t.isRun {
				if t.over != nil {
					t.tm.Stop()
					close(t.over)
				}
				return
			}

			t.tm.Reset(t.d)
		}
	}
}

func (t *Ticker) flush() {
	if t.isRun {
		t.tm.Reset(t.d)
	}
}

// Stop stop ticker
func (t *Ticker) Stop() bool {
	if t.isRun == false {
		return true
	}

	// 等待正在执行的任务完成
	t.isRun = false
	if t.over != nil {
		<-t.over
	}
	return true
}
