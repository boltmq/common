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
	"github.com/boltmq/common/message"
	"github.com/boltmq/common/utils/system"
)

// QueueTimeSpan 查询时间宽度
// Author rongzhihong
// Since 2017/9/19
type QueueTimeSpan struct {
	MessageQueue     *message.MessageQueue `json:"messageQueue"`
	MinTimeStamp     int64                 `json:"minTimeStamp"`
	MaxTimeStamp     int64                 `json:"maxTimeStamp"`
	ConsumeTimeStamp int64                 `json:"consumeTimeStamp"`
}

// GetMinTimeStampStr 起始时间
// Author rongzhihong
// Since 2017/9/19
func (timespan *QueueTimeSpan) GetMinTimeStampStr() string {
	return system.MilliTime2String(timespan.MinTimeStamp)
}

// GetMaxTimeStampStr 终止时间
// Author rongzhihong
// Since 2017/9/19
func (timespan *QueueTimeSpan) GetMaxTimeStampStr() string {
	return system.MilliTime2String(timespan.MaxTimeStamp)
}

// GetConsumeTimeStampStr 消费时间
// Author rongzhihong
// Since 2017/9/19
func (timespan *QueueTimeSpan) GetConsumeTimeStampStr() string {
	return system.MilliTime2String(timespan.ConsumeTimeStamp)
}
