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
package basis

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync/atomic"
)

// 压缩
func Compress(src []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(src)
	w.Flush()
	return b.Bytes()
}

// 解压
func UnCompress(src []byte) []byte {
	r, _ := gzip.NewReader(bytes.NewBuffer(src))
	defer r.Close()
	data, _ := ioutil.ReadAll(r)
	return data
}

func CompareAndIncreaseOnly(target *int64, value int64) bool {
	if value > *target {
		updated := atomic.CompareAndSwapInt64(target, *target, value)
		if updated {
			return true
		}
	}
	return false
}
