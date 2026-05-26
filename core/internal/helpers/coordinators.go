// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

// Package helpers - Common utilities.
// The helpers subsystem provides common utilities that can be used by all subsystems. This includes utilities for
// coordinators to start and stop modules, as well as Kafka and Zookeeper client implementations. There are also a
// number of mocks that are provided for testing purposes only, and should not be used in normal code.
package helpers

import (
	"regexp"
	"time"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"

	"github.com/linkedin/Burrow/core/protocol"
)

// StartCoordinatorModules is a helper func for coordinators to start a list of modules. Given a map of protocol.Module,
// it calls the Start func on each one. If any module returns an error, it immediately stops and returns that error
func StartCoordinatorModules(modules map[string]protocol.Module) error {
	_ = "STUB: not implemented"
	// Start all the modules, returning an error if any fail to start
	return nil
}

// StopCoordinatorModules is a helper func for coordinators to stop a list of modules. Given a map of protocol.Module,
// it calls the Stop func on each one. Any errors that are returned are ignored.
func StopCoordinatorModules(modules map[string]protocol.Module) {
	_ = "STUB: not implemented"
	// Stop all the modules passed in
	return
}

// MockModule is a mock of protocol.Module that also satisfies the various subsystem Module variants, and is used in
// tests. It should never be used in the normal code.
type MockModule struct {
	mock.Mock
}

// Configure mocks the protocol.Module Configure func
func (m *MockModule) Configure(name, configRoot string) { _ = "STUB: not implemented"; return }

// Start mocks the protocol.Module Start func
func (m *MockModule) Start() error { _ = "STUB: not implemented"; return nil }

// Stop mocks the protocol.Module Stop func
func (m *MockModule) Stop() error { _ = "STUB: not implemented"; return nil }

// GetName mocks the notifier.Module GetName func
func (m *MockModule) GetName() string { _ = "STUB: not implemented"; return "" }

// GetGroupAllowlist mocks the notifier.Module GetGroupAllowlist func
func (m *MockModule) GetGroupAllowlist() *regexp.Regexp { _ = "STUB: not implemented"; return nil }

// GetGroupDenylist mocks the notifier.Module GetGroupDenylist func
func (m *MockModule) GetGroupDenylist() *regexp.Regexp { _ = "STUB: not implemented"; return nil }

// GetLogger mocks the notifier.Module GetLogger func
func (m *MockModule) GetLogger() *zap.Logger { _ = "STUB: not implemented"; return nil }

// AcceptConsumerGroup mocks the notifier.Module AcceptConsumerGroup func
func (m *MockModule) AcceptConsumerGroup(status *protocol.ConsumerGroupStatus) bool {
	_ = "STUB: not implemented"
	return false
}

// Notify mocks the notifier.Module Notify func
func (m *MockModule) Notify(status *protocol.ConsumerGroupStatus, eventID string, startTime time.Time, stateGood bool) {
	_ = "STUB: not implemented"
	return
}
