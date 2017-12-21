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
	"testing"
)

func TestRemotingCommandEncodeHeader(t *testing.T) {
	remotingCommand := CreateRequestCommand(2, nil)
	buffer := remotingCommand.EncodeHeader()

	buf := bytes.NewBuffer(buffer)
	nRemotingCommand, err := DecodeRemotingCommand(buf)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if nRemotingCommand == nil {
		t.Errorf("Test failed: Decode return nil", err)
	}
}

type testCustomHeader struct {
	Topic   string
	QueueId int
}

func (header *testCustomHeader) CheckFields() error {
	return nil
}

func TestCustomHeaderRemotingCommandEncodeHeader(t *testing.T) {

	tCustomHeader := &testCustomHeader{
		Topic:   "testTopic",
		QueueId: 20,
	}

	remotingCommand := CreateRequestCommand(2, tCustomHeader)
	buffer := remotingCommand.EncodeHeader()

	buf := bytes.NewBuffer(buffer)
	nRemotingCommand, err := DecodeRemotingCommand(buf)
	if err != nil {
		t.Errorf("Test failed: %s", err)
	}

	if nRemotingCommand == nil {
		t.Errorf("Test failed: Decode return nil", err)
	}

	if nRemotingCommand.ExtFields["Topic"] != "testTopic" {
		t.Errorf("Test failed: ExtFields value invalid")
	}

	if nRemotingCommand.ExtFields["QueueId"] != "20" {
		t.Errorf("Test failed: ExtFields value invalid")
	}
}
