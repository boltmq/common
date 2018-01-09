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

const (
	TIMEFORMAT = "2006-01-02 15:04:05"
)

// CurrentTimeMillis 当前时间毫秒数
// Author rongzhihong
// Since 2017/9/5
func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

// ComputNextMorningTimeMillis 下一整点天（时、分、秒、毫秒置为0）
// Author rongzhihong
// Since 2017/9/5
func ComputNextMorningTimeMillis() int64 {
	currentTime := time.Now()
	nextMorning := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 0, 0, 0, 0, currentTime.Location())
	nextMorningTimeMillis := nextMorning.UnixNano() / 1000000
	return nextMorningTimeMillis
}

// ComputNextMinutesTimeMillis 下一整点分钟（秒、毫秒置为0）
// Author rongzhihong
// Since 2017/9/5
func ComputNextMinutesTimeMillis() int64 {
	currentTime := time.Now()
	nextMorning := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute()+1, 0, 0, currentTime.Location())

	nextMorningTimeMillis := nextMorning.UnixNano() / 1000000
	return nextMorningTimeMillis
}

// ComputNextHourTimeMillis 下一整点小时（分、秒、毫秒置为0）
// Author rongzhihong
// Since 2017/9/5
func ComputNextHourTimeMillis() int64 {
	currentTime := time.Now()
	nextMorning := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour()+1, 0, 0, 0, currentTime.Location())

	nextMorningTimeMillis := nextMorning.UnixNano() / int64(time.Millisecond)
	return nextMorningTimeMillis
}

// MillsTime2String 将毫秒时间转为字符时间
// Author: rongzhihong
// Since: 2017/9/19
func MilliTime2String(millisecond int64) string {
	secondTime := millisecond / 1000
	return time.Unix(secondTime, 0).Format(TIMEFORMAT)
}
