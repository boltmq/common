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

// MessageSysFlag: 消息flag
// Author: yintongqiang
// Since:  2017/8/10
const (
	// SysFlag
	CompressedFlag = 0x1 << 0
	MultiTagsFlag  = 0x1 << 1

	// 7 6 5 4 3 2 1 0<br>
	// SysFlag 事务相关，从左属，2与3
	TransactionNotType      = 0x0 << 2
	TransactionPreparedType = 0x1 << 2
	TransactionCommitType   = 0x2 << 2
	TransactionRollbackType = 0x3 << 2
)

func GetTransactionValue(flag int) int {
	return flag & TransactionRollbackType
}

func ResetTransactionValue(flag int, ty int) int {
	return (flag & (0xFFFFFFFF ^ TransactionRollbackType)) | ty
}

func ClearCompressedFlag(flag int) int {
	return flag & (0xFFFFFFFF ^ CompressedFlag)
}
