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
package body

// QueryCorrectionOffset 查找被修正 offset (转发组件）的返回内容
// Author rongzhihong
// Since 2017/9/19
type QueryCorrectionOffset struct {
	CorrectionOffsets map[int]int64 `json:"correctionOffsets"`
}

func NewQueryCorrectionOffset() *QueryCorrectionOffset {
	body := new(QueryCorrectionOffset)
	body.CorrectionOffsets = make(map[int]int64)
	return body
}
