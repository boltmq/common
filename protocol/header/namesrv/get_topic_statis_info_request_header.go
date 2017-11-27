package namesrv

// GetTopicStatisInfoRequestHeader 获取topic统计信息头
// Author: jerrylou, <gunsluo@gmail.com>
// Since: 2017-08-25
type GetTopicStatisInfoRequestHeader struct {
	Topic string
}

func (header *GetTopicStatisInfoRequestHeader) CheckFields() error {
	return nil
}
