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
package verify

import (
	"net"
	"strconv"
)

// CheckIpAndPort 校验ip:port是否有效
func CheckIpAndPort(addr string) bool {
	if addr == "" {
		return false
	}

	host, sport, err := net.SplitHostPort(addr)
	if err != nil {
		return false
	}

	if ip := net.ParseIP(host); ip == nil {
		return false
	}

	port, err := strconv.Atoi(sport)
	if err != nil {
		return false
	}

	if port < 1 || port > 65535 {
		return false
	}

	return true
}
