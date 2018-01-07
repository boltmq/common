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
package sysflag

// PullSysFlag: 拉取消费flag
// Author: yintongqiang
// Since:  2017/8/14
const (
	FLAG_COMMIT_OFFSET = 0x1 << 0
	FLAG_SUSPEND       = 0x1 << 1
	FLAG_SUBSCRIPTION  = 0x1 << 2
	FLAG_CLASS_FILTER  = 0x1 << 3
)

func BuildSysFlag(commitOffset bool, suspend bool, subscription bool, classFilter bool) int {
	flag := 0
	if commitOffset {
		flag |= FLAG_COMMIT_OFFSET
	}

	if suspend {
		flag |= FLAG_SUSPEND
	}

	if subscription {
		flag |= FLAG_SUBSCRIPTION
	}

	if classFilter {
		flag |= FLAG_CLASS_FILTER
	}

	return flag
}
func ClearCommitOffsetFlag(sysFlag int) int {
	return sysFlag & (0xFFFFFFFF ^ FLAG_COMMIT_OFFSET)
}

func HasCommitOffsetFlag(sysFlag int) bool {
	return (sysFlag & FLAG_COMMIT_OFFSET) == FLAG_COMMIT_OFFSET
}

func HasSuspendFlag(sysFlag int) bool {
	return (sysFlag & FLAG_SUSPEND) == FLAG_SUSPEND
}

func HasSubscriptionFlag(sysFlag int) bool {
	return (sysFlag & FLAG_SUBSCRIPTION) == FLAG_SUBSCRIPTION
}

func HasClassFilterFlag(sysFlag int) bool {
	return (sysFlag & FLAG_CLASS_FILTER) == FLAG_CLASS_FILTER
}
