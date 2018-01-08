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
package header

// PullMessageResponseHeader: 拉取消息响应头
// Author: yintongqiang
// Since:  2017/8/16
type PullMessageResponseHeader struct {
	SuggestWhichBrokerId int64 `json:"suggestWhichBrokerId"`
	NextBeginOffset      int64 `json:"nextBeginOffset"`
	MinOffset            int64 `json:"minOffset"`
	MaxOffset            int64 `json:"maxOffset"`
}

func (header *PullMessageResponseHeader) CheckFields() error {
	return nil
}
