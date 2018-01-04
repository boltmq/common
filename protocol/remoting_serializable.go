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

import "github.com/pquerna/ffjson/ffjson"

// RemotingSerializable 序列化
// Author gaoyanlei
// Since 2017/8/22
type RemotingSerializable struct {
}

//// Encode 只会序列化“RemotingSerializable”结构体的字段 ??
//func (serial *RemotingSerializable) Encode() []byte {
//	if value, err := ffjson.Marshal(serial); err == nil {
//		return value
//	}
//	return nil
//}
//
//// Decode 反序列化为serial对应的结构体
//func (serial *RemotingSerializable) Decode(data []byte) error {
//	return ffjson.Unmarshal(data, serial)
//}

// CustomEncode 默认序列化参数v对应的结构体
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/15
func (serial *RemotingSerializable) CustomEncode(v interface{}) []byte {
	if value, err := ffjson.Marshal(v); err == nil {
		return value
	}
	return nil
}

// CustomDecode 反序列化为传入参数v结构体
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/15
func (serial *RemotingSerializable) CustomDecode(data []byte, v interface{}) error {
	return ffjson.Unmarshal(data, v)
}
