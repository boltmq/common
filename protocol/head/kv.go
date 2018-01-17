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
package head

import (
	"strings"

	"github.com/go-errors/errors"
)

// DeleteKVConfigRequestHeader 删除KV配置项-请求头
// Author: tianyuliang
// Since: 2017/9/4
type DeleteKVConfigRequestHeader struct {
	Namespace string
	Key       string
}

func (header *DeleteKVConfigRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.Key) == "" {
		return errors.Errorf("key is empty")
	}
	if strings.TrimSpace(header.Namespace) == "" {
		return errors.Errorf("namespace is empty")
	}
	return nil
}

type GetKVListByNamespaceRequestHeader struct {
	Namespace string
}

func (header *GetKVListByNamespaceRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.Namespace) == "" {
		return errors.Errorf("namespace is empty")
	}
	return nil
}

// PutKVConfigRequestHeader 向Namesrv追加KV配置-请求头
// Author: tianyuliang
// Since: 2017/9/4
type PutKVConfigRequestHeader struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func (header *PutKVConfigRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.Namespace) == "" {
		return errors.Errorf("namespace is empty")
	}
	if strings.TrimSpace(header.Key) == "" {
		return errors.Errorf("key is empty")
	}
	if strings.TrimSpace(header.Value) == "" {
		return errors.Errorf("value is empty")
	}

	return nil
}

// GetKVConfigRequestHeader: 创建头请求信息
// Author: yintongqiang
// Since:  2017/8/23
type GetKVConfigRequestHeader struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
}

func (header *GetKVConfigRequestHeader) CheckFields() error {
	if strings.TrimSpace(header.Key) == "" {
		return errors.Errorf("key is empty")
	}
	if strings.TrimSpace(header.Namespace) == "" {
		return errors.Errorf("namespace is empty")
	}
	return nil
}

// GetKVConfigResponseHeader: 响应头
// Author: yintongqiang
// Since:  2017/8/23
type GetKVConfigResponseHeader struct {
	Value string `json:"value"`
}

func (header *GetKVConfigResponseHeader) CheckFields() error {
	return nil
}
