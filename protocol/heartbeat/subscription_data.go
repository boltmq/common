// Copyright 2017 yintongqiang

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package heartbeat

import (
	"fmt"
	"strings"

	set "github.com/deckarep/golang-set"
)

// SubscriptionData: 订阅信息结构体
// Author: yintongqiang
// Since:  2017/8/9
type SubscriptionData struct {
	SUB_ALL         string  `json:"-"`
	ClassFilterMode bool    `json:"classFilterMode"`
	Topic           string  `json:"topic"`
	SubString       string  `json:"subString"`
	TagsSet         set.Set `json:"tagsSet"`
	CodeSet         set.Set `json:"codeSet"`
	SubVersion      int     `json:"subVersion"`
}

type SubscriptionDataPlus struct {
	SUB_ALL         string   `json:"-"`
	ClassFilterMode bool     `json:"classFilterMode"`
	Topic           string   `json:"topic"`
	SubString       string   `json:"subString"`
	TagsSet         []string `json:"tagsSet"`
	CodeSet         []int32  `json:"codeSet"`
	SubVersion      int      `json:"subVersion"`
}

// String 格式化订阅信息结构体的内容
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/6
func (sd *SubscriptionData) String() string {
	if sd == nil {
		return "SubscriptionData is nil"
	}
	format := "SubscriptionData {topic=%s, subString=%s, tagsSet=%s, codeSet=%s, subVersion=%d, classFilterMode=%t}"
	return fmt.Sprintf(format, sd.Topic, sd.SubString, sd.TagsSet.String(), sd.CodeSet.String(), sd.SubVersion, sd.ClassFilterMode)
}

// String 格式化订阅信息结构体的内容
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/6
func (sdp *SubscriptionDataPlus) String() string {
	if sdp == nil {
		return "SubscriptionDataPlus is nil"
	}

	tags := strings.Join(sdp.TagsSet, ",")
	format := "SubscriptionDataPlus {Topic=%s, SubString=%s, TagsSet=[%s], CodeSet=[%v], SubVersion=%d, ClassFilterMode=%t}"
	return fmt.Sprintf(format, sdp.Topic, sdp.SubString, tags, sdp.CodeSet, sdp.SubVersion, sdp.ClassFilterMode)
}
