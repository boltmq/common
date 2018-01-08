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
package basis

import (
	"github.com/boltmq/common/sysflag"
	"github.com/boltmq/common/utils/codec"
)

// TopicFilterType Topic过滤方式，默认为单Tag过滤
// Author gaoyanlei
// Since 2017/8/9
type TopicFilterType int

const (
	SINGLE_TAG TopicFilterType = iota // 每个消息只能有一个Tag
	MULTI_TAG                         // (1)每个消息可以有多个Tag（暂时不支持，后续视情况支持)(2)为什么暂时不支持?(3)此功能可能会对用户造成困扰，且方案并不完美，所以暂不支持
)

// String
func (tft TopicFilterType) String() string {
	switch tft {
	case SINGLE_TAG:
		return "SINGLE_TAG"
	case MULTI_TAG:
		return "MULTI_TAG"
	default:
		return ""
	}
}

func ParseTopicFilterType(sysFlag int32) TopicFilterType {
	if (sysFlag & sysflag.MultiTagsFlag) == sysflag.MultiTagsFlag {
		return MULTI_TAG
	}

	return SINGLE_TAG
}

func TagsString2tagsCode(filterType TopicFilterType, tags string) int64 {
	if tags == "" || len(tags) == 0 {
		return 0
	}

	return codec.HashCode(tags)
}
