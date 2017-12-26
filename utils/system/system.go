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
package system

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Home() string {
	if runtime.GOOS == "windows" {
		return WindowsHome()
	}

	return UnixHome()
}

func WindowsHome() string {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path

	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}

	return home
}

func UnixHome() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return ""
	}

	result := strings.TrimSpace(stdout.String())
	return result
}
