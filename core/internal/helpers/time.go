// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package helpers

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// Ticker is a generic interface for a channel that delivers `ticks' of a clock at intervals.
type Ticker interface {
	// Start sending ticks over the channel
	Start()

	// Stop sending ticks over the channel
	Stop()

	// Return the channel that ticks will be sent over
	GetChannel() <-chan time.Time
}

// PausableTicker is an implementation of Ticker which can be stopped and restarted without changing the underlying
// channel. This is useful for cases where you may need to stop performing actions for a while (such as sending
// notifications), but you do not want to tear down everything.
type PausableTicker struct {
	channel     chan time.Time
	duration    time.Duration
	ticker      *time.Ticker
	quitChannel chan struct{}
}

// NewPausableTicker returns a Ticker that has not yet been started, but the channel is ready to use. This ticker can be
// started and stopped multiple times without needing to swap the ticker channel
func NewPausableTicker(d time.Duration) Ticker { _ = "STUB: not implemented"; return *new(Ticker) }

// Start begins sending ticks over the channel at the interval that has already been configured. If the ticker is
// already sending ticks, this func has no effect.
func (ticker *PausableTicker) Start() { _ = "STUB: not implemented"; return }

// Don't restart a ticker that's already running

// Channel to be able to close the goroutine

// Start the ticker

// This goroutine will forward the ticker ticks to our exposed channel

// Stop stops ticks from being sent over the channel. If the ticker is not currently sending ticks, this func has no
// effect
func (ticker *PausableTicker) Stop() { _ = "STUB: not implemented"; return }

// Don't stop an already stopped ticker

// Stop the underlying ticker

// Tell our goroutine to quit

// GetChannel returns the channel over which ticks will be sent. This channel can be used over multiple Start/Stop
// cycles, and will not be closed.
func (ticker *PausableTicker) GetChannel() <-chan time.Time { _ = "STUB: not implemented"; return nil }

// MockTicker is a mock Ticker interface that can be used for testing. It should not be used in normal code.
type MockTicker struct {
	mock.Mock
}

// Start mocks Ticker.Start
func (m *MockTicker) Start() {
	_ = "STUB: not implemented"

	// Stop mocks Ticker.Stop
	return
}

func (m *MockTicker) Stop() {
	_ = "STUB: not implemented"

	// GetChannel mocks Ticker.GetChannel
	return
}

func (m *MockTicker) GetChannel() <-chan time.Time { _ = "STUB: not implemented"; return nil }
