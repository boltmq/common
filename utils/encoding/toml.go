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
package encoding

import (
	"github.com/BurntSushi/toml"
)

// DecodeToml is Decode toml file, except it will automatically read the
// contents of the file at `path` and decode it for you.
func DecodeToml(path string, v interface{}) error {
	if _, err := toml.DecodeFile(path, v); err != nil {
		return err
	}

	return nil
}
