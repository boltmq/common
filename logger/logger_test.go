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
package logger

import (
	"testing"
)

func TestConsoleLog(t *testing.T) {
	Trace("this is a test")
	Debug("this is a test")
	Info("this is a test")
	Warn("this is a test")
	Error("this is a test")
	Fatal("this is a test")

	Tracef("this is a %s", "test")
	Debugf("this is a %s", "test")
	Infof("this is a %s", "test")
	Warnf("this is a %s", "test")
	Errorf("this is a %s", "test")
	Fatalf("this is a %s", "test")
}

func TestFileLog(t *testing.T) {
	ConfigAsFile("seelog.xml")

	Trace("this is a test")
	Debug("this is a test")
	Info("this is a test")
	Warn("this is a test")
	Error("this is a test")
	Fatal("this is a test")

	Tracef("this is a %s", "test")
	Debugf("this is a %s", "test")
	Infof("this is a %s", "test")
	Warnf("this is a %s", "test")
	Errorf("this is a %s", "test")
	Fatalf("this is a %s", "test")
}

var frontLogConfig = `
<seelog>
    <outputs>
        <console/>
    </outputs>
</seelog>
`

func TestBytesLog(t *testing.T) {
	ConfigAsBytes([]byte(frontLogConfig))

	Trace("this is a test")
	Debug("this is a test")
	Info("this is a test")
	Warn("this is a test")
	Error("this is a test")
	Fatal("this is a test")

	Tracef("this is a %s", "test")
	Debugf("this is a %s", "test")
	Infof("this is a %s", "test")
	Warnf("this is a %s", "test")
	Errorf("this is a %s", "test")
	Fatalf("this is a %s", "test")
}
