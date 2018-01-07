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
package filter

import (
	"errors"
	"strings"

	"github.com/boltmq/common/protocol/heartbeat"
	"github.com/boltmq/common/utils/codec"
	set "github.com/deckarep/golang-set"
)

// FilterAPI: filter api
// Author: yintongqiang
// Since:  2017/8/11
type FilterAPI struct {
}

func BuildSubscriptionData(consumerGroup string, topic string, subString string) (*heartbeat.SubscriptionData, error) {
	subscriptionData := &heartbeat.SubscriptionData{Topic: topic, SubString: subString, TagsSet: set.NewSet(), CodeSet: set.NewSet()}
	if strings.EqualFold(subString, "") || strings.EqualFold(subString, "*") {
		subscriptionData.SubString = "*"
	} else {
		tags := strings.Split(subString, "||")
		for _, tag := range tags {
			trimTag := strings.TrimSpace(tag)
			if !strings.EqualFold(trimTag, "") {
				subscriptionData.TagsSet.Add(trimTag)
				subscriptionData.CodeSet.Add(codec.HashCode(trimTag))
			} else {
				return subscriptionData, errors.New("subString split error")
			}
		}
	}
	return subscriptionData, nil
}

func BuildSubscriptionData4Ponit(consumerGroup string, topic string, subString string) (subscription *heartbeat.SubscriptionData, err error) {
	subscriptionData := &heartbeat.SubscriptionData{Topic: topic, SubString: subString, TagsSet: set.NewSet(), CodeSet: set.NewSet()}
	if strings.EqualFold(subString, "") || strings.EqualFold(subString, "*") {
		subscriptionData.SubString = "*"
	} else {
		tags := strings.Split(subString, "||")
		for _, tag := range tags {
			trimTag := strings.TrimSpace(tag)
			if !strings.EqualFold(trimTag, "") {
				subscriptionData.TagsSet.Add(trimTag)
				subscriptionData.CodeSet.Add(codec.HashCode(trimTag))
			} else {
				return subscriptionData, errors.New("subString split error")
			}
		}
	}
	return subscriptionData, nil
}
