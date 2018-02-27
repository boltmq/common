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
package basis

const (
	MASTER_ID                   = 0
	SELF_TEST_TOPIC             = "SELF_TEST_TOPIC"
	DEFAULT_TOPIC               = "MY_DEFAULT_TOPIC"
	BENCHMARK_TOPIC             = "BenchmarkTest"
	OFFSET_MOVED_EVENT          = "OFFSET_MOVED_EVENT"
	DEFAULT_PRODUCER_GROUP      = "DEFAULT_PRODUCER"
	DEFAULT_CONSUMER_GROUP      = "DEFAULT_CONSUMER"
	TOOLS_CONSUMER_GROUP        = "TOOLS_CONSUMER"
	FILTERSRV_CONSUMER_GROUP    = "FILTERSRV_CONSUMER"
	CLIENT_INNER_PRODUCER_GROUP = "CLIENT_INNER_PRODUCER"
	SELF_TEST_CONSUMER_GROUP    = "SELF_TEST_C_GROUP"
	RETRY_GROUP_TOPIC_PREFIX    = "%RETRY%"
	DLQ_GROUP_TOPIC_PREFIX      = "%DLQ%"
)
