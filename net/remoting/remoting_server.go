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
package remoting

import (
	"github.com/boltmq/common/net/core"
	"github.com/boltmq/common/protocol"
)

// RemotingServer remoting server define
type RemotingServer interface {
	InvokeSync(ctx core.Context, request *protocol.RemotingCommand, timeoutMillis int64) (*protocol.RemotingCommand, error)
	InvokeAsync(ctx core.Context, request *protocol.RemotingCommand, timeoutMillis int64, invokeCallback InvokeCallback) error
	InvokeOneway(ctx core.Context, request *protocol.RemotingCommand, timeoutMillis int64) error
	RegisterProcessor(requestCode int32, processor RequestProcessor)
	SetDefaultProcessor(processor RequestProcessor)
	RegisterRPCHook(rpcHook RPCHook)
	SetContextEventListener(contextEventListener ContextEventListener)
	ListenPort() int
	Start()
	Shutdown()
}
