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
	"fmt"
	"sync/atomic"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

type DataVersion struct {
	Timestamp int64 `json:"timestamp"`
	Counter   int64 `json:"counter"`
}

func NewDataVersion(timestamp ...int64) *DataVersion {
	dataVersion := new(DataVersion)
	dataVersion.Timestamp = time.Now().UnixNano()
	if timestamp != nil && len(timestamp) > 0 {
		dataVersion.Timestamp = timestamp[0]
	}
	dataVersion.Counter = atomic.AddInt64(&dataVersion.Counter, 0)
	return dataVersion
}

func (dv *DataVersion) AssignNewOne(dataVersion DataVersion) {
	dv.Timestamp = dataVersion.Timestamp
	dv.Counter = dataVersion.Counter
}

func (dv *DataVersion) Equals(dataVersion *DataVersion) bool {
	if dv == nil && dataVersion == nil {
		return true
	}
	if (dv == nil && dataVersion != nil) || (dv != nil && dataVersion == nil) {
		return false
	}

	return dv.Timestamp == dataVersion.Timestamp && dv.Counter == dataVersion.Counter
}

func (dv *DataVersion) NextVersion() {
	dv.Timestamp = time.Now().UnixNano()
	dv.Counter = atomic.AddInt64(&dv.Counter, 1)
}

func (dv *DataVersion) String() string {
	return fmt.Sprintf("dataVersion [timestamp=%d, counter=%d]", dv.Timestamp, dv.Counter)
}

func (dv *DataVersion) Json() string {
	buf, err := ffjson.Marshal(dv)
	if err != nil {
		return err.Error()
	}

	return string(buf)
}
