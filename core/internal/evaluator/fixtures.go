// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package evaluator

import (
	"github.com/linkedin/Burrow/core/internal/storage"
)

// StorageAndEvaluatorCoordinatorsWithOffsets sets up a Coordinator with a single caching module defined. In order to do
// this, it also calls the storage subsystem fixture to get a configured storage.Coordinator with offsets for a test
// cluster and group. This func should never be called in normal code. It is only provided to facilitate testing by
// other subsystems.
func StorageAndEvaluatorCoordinatorsWithOffsets() (*Coordinator, *storage.Coordinator) {
	_ = "STUB: not implemented"
	return nil, nil
}
