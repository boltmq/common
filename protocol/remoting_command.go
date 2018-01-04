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
package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	"github.com/go-errors/errors"
	"github.com/pquerna/ffjson/ffjson"
)

// RemotingCommand remoting command
// Author: luoji, <gunsluo@gmail.com>
// Since: 2017-08-22
type RemotingCommand struct {
	Code         int32               `json:"code"`
	Language     string              `json:"language"`
	Version      int32               `json:"version"`
	Opaque       int32               `json:"opaque"`
	Flag         int32               `json:"flag"`
	Remark       string              `json:"remark"`
	ExtFields    map[string]string   `json:"extFields"` // 请求拓展字段
	CustomHeader CommandCustomHeader `json:"-"`         // 修改字段类型,"CustomHeader"字段不序列化 2017/8/24 Modify by luoji, <gunsluo@gmail.com>
	Body         []byte              `json:"-"`         // body字段不会被Encode()并进行网络传输
}

// CreateResponseCommand 只有通信层内部会调用，业务不会调用
func CreateDefaultResponseCommand(customHeader ...CommandCustomHeader) *RemotingCommand {
	cmd := CreateResponseCommand(SYSTEM_ERROR, "not set any response code")
	// 设置头信息
	if customHeader != nil && len(customHeader) > 0 {
		cmd.CustomHeader = customHeader[0]
	}
	return cmd
}

// CreateResponseCommand
func CreateResponseCommand(code int32, remark string) *RemotingCommand {
	cmd := &RemotingCommand{
		Code:      code,
		Remark:    remark,
		ExtFields: make(map[string]string),
		Language:  GOLANG.String(),
	}
	// 设置为响应报文
	cmd.MarkResponseType()
	// 设置版本信息
	cmd.setCMDVersion()

	return cmd
}

// CreateRequestCommand 创建客户端请求信息 2017/8/16 Add by yintongqiang
func CreateRequestCommand(code int32, customHeader ...CommandCustomHeader) *RemotingCommand {
	cmd := &RemotingCommand{
		Code:      code,
		ExtFields: make(map[string]string),
		Language:  GOLANG.String(),
	}
	if customHeader != nil && len(customHeader) > 0 {
		cmd.CustomHeader = customHeader[0]
	}
	cmd.Opaque = inrcOpaque() // 标识自增，请求唯一标识
	cmd.setCMDVersion()       // 设置版本信息
	return cmd
}

// Author: luoji, <gunsluo@gmail.com>
// Since: 2017-08-22
func (rc *RemotingCommand) setCMDVersion() {
	if configVersion >= 0 {
		rc.Version = configVersion
		return
	}

	version := os.Getenv(RemotingVersionKey)
	if version == "" {
		return
	}

	v, e := strconv.Atoi(version)
	if e == nil {
		rc.Version = int32(v)
	}
}

// IsOnewayRPC is oneway rpc, return bool
func (rc *RemotingCommand) IsOnewayRPC() bool {
	var bits int32
	bits = 1 << rpcOneway
	return (rc.Flag & bits) == bits
}

// MarkResponseType mark response type
func (rc *RemotingCommand) MarkResponseType() {
	var bits int32
	bits = 1 << rpcType
	rc.Flag |= bits
}

// MarkOnewayRPC mark oneway type
func (rc *RemotingCommand) MarkOnewayRPC() {
	var bits int32
	bits = 1 << rpcOneway
	rc.Flag |= bits
}

// IsResponseType is response type, return bool
func (rc *RemotingCommand) IsResponseType() bool {
	var bits int32
	bits = 1 << rpcType
	return (rc.Flag & bits) == bits
}

// EncodeHeader 编码头部
func (rc *RemotingCommand) EncodeHeader() []byte {
	var (
		length       int32 = 4
		headerLength int32
	)
	// 构建头部报文
	headerData := rc.buildHeader()
	headerLength = int32(len(headerData))
	length += headerLength

	if rc.Body != nil {
		length += int32(len(rc.Body))
	}

	buf := bytes.NewBuffer([]byte{})
	// 写入报文长度
	binary.Write(buf, binary.BigEndian, length)
	// 写入报文头部长度
	binary.Write(buf, binary.BigEndian, headerLength)
	// 写入报文头部
	buf.Write(headerData)

	return buf.Bytes()
}

func (rc *RemotingCommand) buildHeader() []byte {
	// 处理custom header
	rc.makeCustomHeaderToNet()

	// json 编码
	buf, err := ffjson.Marshal(rc)
	if err != nil {
		return nil
	}
	return buf
}

func (rc *RemotingCommand) makeCustomHeaderToNet() {
	if rc.CustomHeader == nil {
		return
	}

	extFields := encodeCommandCustomHeader(rc.CustomHeader)
	for k, v := range extFields {
		rc.ExtFields[k] = v
	}
}

// Type return remoting command type
func (rc *RemotingCommand) Type() RemotingCommandType {
	if rc.IsResponseType() {
		return RESPONSE_COMMAND
	}

	return REQUEST_COMMAND
}

func (rc *RemotingCommand) DecodeCommandCustomHeader(commandCustomHeader CommandCustomHeader) error {
	if commandCustomHeader == nil {
		return nil
	}

	if rc.ExtFields == nil {
		return nil
	}

	err := decodeCommandCustomHeader(rc.ExtFields, commandCustomHeader)
	if err != nil {
		return err
	}

	return commandCustomHeader.CheckFields()
}

// DecodeRemotingCommand 解析返回RemotingCommand
func DecodeRemotingCommand(buf *bytes.Buffer) (*RemotingCommand, error) {
	var (
		length       int32
		headerLength int32
		bodyLength   int32
		body         []byte
	)

	// step 1 读取报文长度
	if buf.Len() < 4 {
		return nil, errors.Errorf("DecodeRemotingCommand: buffer-length[%d] incorrect，minimal is 4", buf.Len())
	}

	err := binary.Read(buf, binary.BigEndian, &length)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	// step 2 读取报文头长度
	if buf.Len() < 4 {
		return nil, errors.Errorf("DecodeRemotingCommand: buffer-length[%d] < header-length[%d](length is 4).", buf.Len(), 4)
	}

	err = binary.Read(buf, binary.BigEndian, &headerLength)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	// step 3 读取报文头数据
	if buf.Len() == 0 || buf.Len() < int(headerLength) {
		return nil, errors.Errorf("DecodeRemotingCommand: buffer-length[%d] < attribute header-data[%d] from buffer.", buf.Len(), headerLength)
	}

	header := make([]byte, headerLength)
	_, err = buf.Read(header)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	// step 4 读取报文Body
	bodyLength = length - 4 - headerLength
	if buf.Len() < int(bodyLength) {
		return nil, errors.Errorf("DecodeRemotingCommand: buffer-length[%d] < attribute body[%d] from buffer.", buf.Len(), bodyLength)
	}

	if bodyLength > 0 {
		body = make([]byte, bodyLength)
		_, err = buf.Read(body)
		if err != nil {
			return nil, errors.Wrap(err, 0)
		}
	}

	// 解码
	return decodeRemotingCommand(header, body)
}

// decodeRemotingCommand 解析header
func decodeRemotingCommand(header, body []byte) (*RemotingCommand, error) {
	cmd := &RemotingCommand{}
	cmd.ExtFields = make(map[string]string)
	err := ffjson.Unmarshal(header, cmd)
	if err != nil {
		return nil, err
	}
	cmd.Body = body
	return cmd, nil
}

// Bytes 实现Serirable接口
func (rc *RemotingCommand) Bytes() []byte {
	var (
		header []byte
		packet []byte
	)

	// 头部进行编码
	header = rc.EncodeHeader()
	if rc.Body != nil && len(rc.Body) > 0 {
		packet = append(header, rc.Body...)
	} else {
		packet = header
	}

	return packet
}

// String
func (rc *RemotingCommand) String() string {
	if rc == nil {
		return "<nil>"
	}

	extFields := "{}"
	if bf, err := ffjson.Marshal(rc.ExtFields); err == nil {
		extFields = string(bf)
	}

	return fmt.Sprintf("[code=%d(%s), language=%s, version=%d, opaque=%d, flag(B)=%b, extFields=%s, body=%v, remark=%s]",
		rc.Code, ParseRequestCode(rc.Code), rc.Language, rc.Version, rc.Opaque, rc.Flag, extFields, rc.Body, rc.Remark)
}
