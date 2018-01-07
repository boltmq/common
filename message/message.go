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
package message

import (
	"strconv"
)

// Message: 消息结构体
// Author: yintongqiang
// Since:  2017/8/9
type Message struct {
	Topic      string            // 消息主题
	Flag       int32             // 消息标志，系统不做干预，完全由应用决定如何使用
	Properties map[string]string // 消息属性，都是系统属性，禁止应用设置
	Body       []byte            // 消息体
}

func NewMessage(topic string, tags string, body []byte) *Message {
	properties := make(map[string]string)
	properties[PROPERTY_TAGS] = tags
	return &Message{Topic: topic, Properties: properties, Body: body}
}

func (msg *Message) ClearProperty(name string) {
	delete(msg.Properties, name)
}

func (msg *Message) PutProperty(name string, value string) {
	if msg.Properties == nil {
		msg.Properties = make(map[string]string)
	}

	msg.Properties[name] = value
}

func (msg *Message) GetProperty(name string) string {
	if msg.Properties == nil {
		msg.Properties = make(map[string]string)
	}
	return msg.Properties[name]
}

func (msg *Message) SetTags(tags string) {
	msg.PutProperty(PROPERTY_TAGS, tags)
}

func (msg *Message) GetTags() string {
	return msg.GetProperty(PROPERTY_TAGS)
}

func (msg *Message) GetOriginMessageID() string {
	return msg.GetProperty(PROPERTY_ORIGIN_MESSAGE_ID)
}

func (msg *Message) SetWaitStoreMsgOK(waitStoreMsgOK bool) {
	msg.PutProperty(PROPERTY_WAIT_STORE_MSG_OK, strconv.FormatBool(waitStoreMsgOK))
}

func (msg *Message) SetDelayTimeLevel(level int) {
	msg.PutProperty(PROPERTY_DELAY_TIME_LEVEL, strconv.Itoa(level))
}

func (msg *Message) GetKeys() string {
	return msg.GetProperty(PROPERTY_KEYS)
}

func (msg *Message) SetKeys(keys string) {
	msg.PutProperty(PROPERTY_KEYS, keys)
}
