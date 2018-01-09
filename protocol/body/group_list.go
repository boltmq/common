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

import (
	"fmt"

	"github.com/boltmq/common/protocol"
	set "github.com/deckarep/golang-set"
)

// GroupList 分组集合
// Author rongzhihong
// Since 2017/9/19
type GroupList struct {
	GroupList set.Set `json:"groupList"`
	protocol.RemotingSerializable
}

// NewGroupList 初始化
// Author: tianyuliang
// Since: 2017/11/1
func NewGroupList() *GroupList {
	groupList := new(GroupList)
	groupList.GroupList = set.NewSet()
	return groupList
}

// String 打印结构信息
// Author: tianyuliang
// Since: 2017/11/1
func (gl *GroupList) String() string {
	if gl == nil || gl.GroupList == nil {
		return "<nil>"
	}

	var values []string
	for itor := range gl.GroupList.Iterator().C {
		values = append(values, itor.(string))
	}
	return fmt.Sprintf("GroupList %s.", values)
}
