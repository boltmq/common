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
package protocol

import (
	"sync"
	"testing"
)

func TestInrcOpaque(t *testing.T) {
	var (
		i int32
	)

	resetOpaque()

	i = inrcOpaque()
	if i != 1 {
		t.Errorf("Test failed, %d incorrect, except[%d]", i, 1)
		return
	}
	i = inrcOpaque()
	if i != 2 {
		t.Errorf("Test failed, %d incorrect, except[%d]", i, 2)
		return
	}
	i = inrcOpaque()
	if i != 3 {
		t.Errorf("Test failed, %d incorrect, except[%d]", i, 3)
		return
	}
}

func TestThreadInrcOpaque(t *testing.T) {
	var (
		wg sync.WaitGroup
		o  int32
		c  = 3
	)

	resetOpaque()

	wg.Add(c)
	fn := func() {
		o = inrcOpaque()
		wg.Done()
	}

	for i := 0; i < c; i++ {
		go fn()
	}

	wg.Wait()
	if o != 3 {
		t.Errorf("Test failed, %d incorrect, except[%d]", o, 3)
		return
	}
}
