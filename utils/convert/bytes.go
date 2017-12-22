// Copyright 2017 tantexian

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
	"bytes"
	"encoding/binary"
)

// Int32ToBytes 整形转换成字节
// Author: tantexian, <tantexian@qq.com>
// Since: 2017/8/7
func Int32ToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

// Int64ToBytes
func Int64ToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

// Int16ToBytes
func Int16ToBytes(n int16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

// Int8ToBytes
func Int8ToBytes(n int8) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

// Int32ToBytes 字节转换成整形
// Author: tantexian, <tantexian@qq.com>
// Since: 2017/8/7
func BytesToInt32(b []byte) int32 {
	var tmp int32

	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

// BytesToInt64
func BytesToInt64(b []byte) int64 {
	var tmp int64

	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

// BytesToInt16
func BytesToInt16(b []byte) int16 {
	var tmp int16

	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}

// BytesToInt8
func BytesToInt8(b []byte) int8 {
	var tmp int8

	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return tmp
}
