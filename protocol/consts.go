package protocol

const (
	RemotingVersionKey = "boltmq.remoting.version"
	rpcType            = 0
	rpcOneway          = 1
)

var (
	// 当前MQ配置版本号
	configVersion int32 = 1
)
