package admin

import "fmt"

type TopicOffset struct {
	MinOffset           int64 `json:"minOffset"`
	MaxOffset           int64 `json:"maxOffset"`
	LastUpdateTimestamp int64 `json:"lastUpdateTimestamp"`
}

func NewTopicOffset() *TopicOffset {
	topicOffset := new(TopicOffset)
	return topicOffset
}

func (to *TopicOffset) ToString() string {
	if to == nil {
		return "<nil>"
	}

	return fmt.Sprintf("{minOffset=%d, maxOffset=%d, lastUpdateTimestamp=%d}",
		to.MinOffset, to.MaxOffset, to.LastUpdateTimestamp)
}
