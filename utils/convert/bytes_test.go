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
package convert

import (
	"testing"
)

func TestInt32ToBytes(t *testing.T) {
	var (
		v      int32 = 10
		expect       = []byte{0, 0, 0, 10}
	)

	bytes := Int32ToBytes(v)

	for i, b := range bytes {
		if b != expect[i] {
			t.Errorf("Int32ToBytes(%d)=%v not equal expect%v.", v, bytes, expect)
			return
		}
	}
}

func TestInt64ToBytes(t *testing.T) {
	var (
		v      int64 = 10
		expect       = []byte{0, 0, 0, 0, 0, 0, 0, 10}
	)

	bytes := Int64ToBytes(v)

	for i, b := range bytes {
		if b != expect[i] {
			t.Errorf("Int64ToBytes(%d)=%v not equal expect%v.", v, bytes, expect)
			return
		}
	}
}

func TestInt8ToBytes(t *testing.T) {
	var (
		v      int8 = 10
		expect      = []byte{10}
	)

	bytes := Int8ToBytes(v)

	for i, b := range bytes {
		if b != expect[i] {
			t.Errorf("Int8ToBytes(%d)=%v not equal expect%v.", v, bytes, expect)
			return
		}
	}
}

func TestInt16ToBytes(t *testing.T) {
	var (
		v      int16 = 10
		expect       = []byte{0, 10}
	)

	bytes := Int16ToBytes(v)

	for i, b := range bytes {
		if b != expect[i] {
			t.Errorf("Int16ToBytes(%d)=%v not equal expect%v.", v, bytes, expect)
			return
		}
	}
}

func TestBytesToInt32(t *testing.T) {
	var (
		bytes        = []byte{0, 0, 0, 10}
		expect int32 = 10
	)

	v := BytesToInt32(bytes)

	if v != expect {
		t.Errorf("BytesToInt32(%v)=%d not equal expect[%d].", bytes, v, expect)
	}
}

func TestBytesToInt64(t *testing.T) {
	var (
		bytes        = []byte{0, 0, 0, 0, 0, 0, 0, 10}
		expect int64 = 10
	)

	v := BytesToInt64(bytes)

	if v != expect {
		t.Errorf("BytesToInt64(%v)=%d not equal expect[%d].", bytes, v, expect)
	}
}

func TestBytesToInt16(t *testing.T) {
	var (
		bytes        = []byte{0, 10}
		expect int16 = 10
	)

	v := BytesToInt16(bytes)

	if v != expect {
		t.Errorf("BytesToInt16(%v)=%d not equal expect[%d].", bytes, v, expect)
	}
}

func TestBytesToInt8(t *testing.T) {
	var (
		bytes       = []byte{10}
		expect int8 = 10
	)

	v := BytesToInt8(bytes)

	if v != expect {
		t.Errorf("BytesToInt8(%v)=%d not equal expect[%d].", bytes, v, expect)
	}
}
