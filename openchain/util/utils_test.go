/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package util

import (
	"bytes"
	"testing"
	"time"
)

func TestComputeCryptoHash(t *testing.T) {
	if bytes.Compare(ComputeCryptoHash([]byte("foobar")), ComputeCryptoHash([]byte("foobar"))) != 0 {
		t.Fatalf("Expected hashes to match, but they did not match")
	}
	if bytes.Compare(ComputeCryptoHash([]byte("foobar1")), ComputeCryptoHash([]byte("foobar2"))) == 0 {
		t.Fatalf("Expected hashes to be different, but they match")
	}
}

func TestUUIDGeneration(t *testing.T) {
	uuid, err := GenerateUUID()
	if err != nil {
		t.Fatalf("Error generating UUID. Error: %s", err)
	}
	if len(uuid) != 36 {
		t.Fatalf("UUID length is not correct. Expected = 36, Got = %d", len(uuid))
	}
	uuid2, err2 := GenerateUUID()
	if err2 != nil {
		t.Fatalf("Error generating UUID. Error: %s", err)
	}
	if uuid == uuid2 {
		t.Fatalf("Two UUIDs are equal. This should never occur")
	}
}

func TestTimestamp(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("timestamp now: %v", CreateUtcTimestamp())
		time.Sleep(200 * time.Millisecond)
	}
}
