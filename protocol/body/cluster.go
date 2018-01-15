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

	"github.com/boltmq/common/protocol/base"
	set "github.com/deckarep/golang-set"
)

// ClusterInfo 协议中传输对象，内容为集群信息
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
type ClusterInfo struct {
	BrokerAddrTable  map[string]*base.BrokerData `json:"brokerAddrTable"`  // brokerName[BrokerData]
	ClusterAddrTable map[string]set.Set          `json:"clusterAddrTable"` // clusterName[set<brokerName>]
}

// ClusterPlusInfo 协议中传输对象，内容为集群信息
//
// 注意: set.Set类型在反序列化过程无法解析，因此额外设置ClusterPlusInfo类型来解析
//
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
type ClusterPlusInfo struct {
	BrokerAddrTable  map[string]*base.BrokerData `json:"brokerAddrTable"`  // brokerName[BrokerData]
	ClusterAddrTable map[string][]string         `json:"clusterAddrTable"` // clusterName[set<brokerName>]
}

// ClusterBrokerInfo cluster与broker包装器
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/15
type ClusterBrokerWapper struct {
	ClusterName string `json:"clusterName"`
	BrokerName  string `json:"brokerName"`
	BrokerAddr  string `json:"brokerAddr"`
	BrokerId    int    `json:"brokerId"`
}

// NewClusterBrokerWapper 初始化
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/15
func NewClusterBrokerWapper(clusterName, brokerName, brokerAddr string, brokerId int) *ClusterBrokerWapper {
	clusterBrokerWapper := &ClusterBrokerWapper{
		ClusterName: clusterName,
		BrokerName:  brokerName,
		BrokerAddr:  brokerAddr,
		BrokerId:    brokerId,
	}
	return clusterBrokerWapper
}

// String 格式化ClusterBrokerWapper数据
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/15
func (wapper *ClusterBrokerWapper) String() string {
	return fmt.Sprintf("ClusterBrokerWapper {clusterName=%s, brokerName=%s, brokerAddr=%s, brokerId=%d}",
		wapper.ClusterName, wapper.BrokerName, wapper.BrokerAddr, wapper.BrokerId)
}

// NewClusterInfo 初始化
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
func NewClusterInfo() *ClusterInfo {
	clusterInfo := &ClusterInfo{
		BrokerAddrTable:  make(map[string]*base.BrokerData),
		ClusterAddrTable: make(map[string]set.Set),
	}
	return clusterInfo
}

// NewClusterPlusInfo 初始化
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
func NewClusterPlusInfo() *ClusterPlusInfo {
	clusterPlusInfo := &ClusterPlusInfo{
		BrokerAddrTable:  make(map[string]*base.BrokerData),
		ClusterAddrTable: make(map[string][]string),
	}
	return clusterPlusInfo
}

// String 格式化
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/15
func (plus *ClusterPlusInfo) String() string {
	if plus == nil {
		return "<nil>"
	}

	return ""
}

// RetrieveAllClusterNames 处理所有brokerName名称
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
func (ci *ClusterInfo) RetrieveAllClusterNames() []string {
	if ci.ClusterAddrTable == nil || len(ci.ClusterAddrTable) == 0 {
		return []string{}
	}

	brokerNames := make([]string, 0, len(ci.ClusterAddrTable))
	for _, v := range ci.ClusterAddrTable {
		for value := range v.Iterator().C {
			if brokerName, ok := value.(string); ok {
				brokerNames = append(brokerNames, brokerName)
			}
		}
	}

	return brokerNames
}

// RetrieveAllAddrByCluster 处理所有brokerAddr地址
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
func (ci *ClusterInfo) RetrieveAllAddrByCluster(clusterName string) []string {
	if ci.ClusterAddrTable == nil || len(ci.ClusterAddrTable) == 0 {
		return []string{}
	}

	brokerAddrs := make([]string, 0)
	if v, ok := ci.ClusterAddrTable[clusterName]; ok {
		for itor := range v.Iterator().C {
			if brokerName, ok := itor.(string); ok {
				if brokerData, ok := ci.BrokerAddrTable[brokerName]; ok {
					if brokerData.BrokerAddrs != nil && len(brokerData.BrokerAddrs) > 0 {
						for _, val := range brokerData.BrokerAddrs {
							brokerAddrs = append(brokerAddrs, val)
						}
					}
				}
			}
		}
	}
	return brokerAddrs
}

// RetrieveAllAddrByCluster 处理所有brokerAddr地址
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
func (cpi *ClusterPlusInfo) RetrieveAllAddrByCluster(clusterName string) ([]string, []*ClusterBrokerWapper) {
	clusterBrokerWappers := make([]*ClusterBrokerWapper, 0)
	brokerAddrs := make([]string, 0)
	if cpi.ClusterAddrTable == nil || len(cpi.ClusterAddrTable) == 0 {
		return brokerAddrs, clusterBrokerWappers
	}

	if brokerNames, ok := cpi.ClusterAddrTable[clusterName]; ok && brokerNames != nil {
		for _, brokerName := range brokerNames {
			brokerData, ok := cpi.BrokerAddrTable[brokerName]
			if ok && brokerData.BrokerAddrs != nil && len(brokerData.BrokerAddrs) > 0 {
				for brokerId, brokerAddr := range brokerData.BrokerAddrs {
					brokerAddrs = append(brokerAddrs, brokerAddr)

					wapper := NewClusterBrokerWapper(clusterName, brokerName, brokerAddr, brokerId)
					clusterBrokerWappers = append(clusterBrokerWappers, wapper)
				}
			}
		}
	}
	return brokerAddrs, clusterBrokerWappers
}

// RetrieveAllAddrByCluster 处理所有brokerAddr地址
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
func (cpi *ClusterPlusInfo) ResolveClusterBrokerWapper() ([]string, []*ClusterBrokerWapper) {
	clusterBrokerWappers := make([]*ClusterBrokerWapper, 0)
	brokerAddrs := make([]string, 0)
	if cpi.ClusterAddrTable == nil || len(cpi.ClusterAddrTable) == 0 {
		return brokerAddrs, clusterBrokerWappers
	}
	for clusterName, _ := range cpi.ClusterAddrTable {
		if brokerNames, ok := cpi.ClusterAddrTable[clusterName]; ok && brokerNames != nil {
			for _, brokerName := range brokerNames {
				brokerData, ok := cpi.BrokerAddrTable[brokerName]
				if ok && brokerData.BrokerAddrs != nil && len(brokerData.BrokerAddrs) > 0 {
					for brokerId, brokerAddr := range brokerData.BrokerAddrs {
						brokerAddrs = append(brokerAddrs, brokerAddr)

						wapper := NewClusterBrokerWapper(clusterName, brokerName, brokerAddr, brokerId)
						clusterBrokerWappers = append(clusterBrokerWappers, wapper)
					}
				}
			}
		}
	}
	return brokerAddrs, clusterBrokerWappers
}

// ClusterInfo 转化为 ClusterInfo 类型
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/11/8
func (plus *ClusterPlusInfo) ClusterInfo() *ClusterInfo {
	if plus == nil {
		return nil
	}
	clusterInfo := &ClusterInfo{}
	clusterInfo.ClusterAddrTable = make(map[string]set.Set)

	clusterInfo.BrokerAddrTable = plus.BrokerAddrTable
	if plus.ClusterAddrTable != nil {
		for clusterName, brokerNames := range plus.ClusterAddrTable {
			if brokerNames != nil && len(brokerNames) > 0 {
				brokerNameSet := set.NewSet()
				for _, brokerName := range brokerNames {
					brokerNameSet.Add(brokerName)
				}
				clusterInfo.ClusterAddrTable[clusterName] = brokerNameSet
			}
		}
	}
	return clusterInfo
}

// RetrieveAllClusterNames 处理所有brokerName名称
// Author: tianyuliang, <tianyuliang@gome.com.cn>
// Since: 2017/9/4
func (cpi *ClusterPlusInfo) RetrieveAllClusterNames() []string {
	if cpi.ClusterAddrTable == nil || len(cpi.ClusterAddrTable) == 0 {
		return make([]string, 0)
	}

	result := make([]string, 0, len(cpi.ClusterAddrTable))
	for _, brokerNames := range cpi.ClusterAddrTable {
		if brokerNames != nil {
			for _, brokerName := range brokerNames {
				result = append(result, brokerName)
			}
		}
	}

	return result
}
