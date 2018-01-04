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
package constant

// PERM 读写权限
// Author gaoyanlei
// Since 2017/8/9
const (
	PERM_PRIORITY = 0x1 << 3
	PERM_READ     = 0x1 << 2
	PERM_WRITE    = 0x1 << 1
	PERM_INHERIT  = 0x1 << 0
)

// IsReadable 是否可读
func IsReadable(perm int) bool {
	return (perm & PERM_READ) == PERM_READ
}

// IsWriteable 是否可写
func IsWriteable(perm int) bool {
	return (perm & PERM_WRITE) == PERM_WRITE
}

// IsInherited 是否可继承
func IsInherited(perm int) bool {
	return (perm & PERM_INHERIT) == PERM_INHERIT
}

// Perm2String 转化为友好的展示效果：可读、可写、可继承
func Perm2String(perm int) string {
	str := "---"
	if IsReadable(perm) {
		str = "R" + str[1:]
	}

	if IsWriteable(perm) {
		str = str[:1] + "W" + str[2:]
	}

	if IsInherited(perm) {
		str = str[:2] + "X"
	}
	return str
}
