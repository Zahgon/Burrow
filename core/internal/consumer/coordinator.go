// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

// Package consumer - Kafka consumer subsystem.
// The consumer subsystem is responsible for getting consumer offset information and sending that information to the
// storage subsystem. This consumer information could be stored in a variety of places, and each module supports a
// different type of repository.
//
// # Modules
//
// Currently, the following modules are provided:
//
// * kafka - Consume a Kafka cluster's __consumer_offsets topic to get consumer information (new consumer)
//
// * kafka_zk - Parse the /consumers tree of a Kafka cluster's metadata to get consumer information (old consumer)
package consumer

import (
	"go.uber.org/zap"

	"github.com/linkedin/Burrow/core/protocol"
)

// The consumer module is responsible for fetching information about consumer group status from some external system
// and forwarding it to the storage module. Each consumer module is associated with a single cluster.

// Coordinator manages all consumer modules, making sure they are configured, started, and stopped at the appropriate
// time.
type Coordinator struct {
	// App is a pointer to the application context. This stores the channel to the storage subsystem
	App *protocol.ApplicationContext

	// Log is a logger that has been configured for this module to use. Normally, this means it has been set up with
	// fields that are appropriate to identify this coordinator
	Log *zap.Logger

	modules map[string]protocol.Module
}

// getModuleForClass returns the correct module based on the passed className. As part of the Configure steps, if there
// is any error, it will panic with an appropriate message describing the problem.
func getModuleForClass(app *protocol.ApplicationContext, moduleName, className string) protocol.Module {
	_ = "STUB: not implemented"
	return *new(protocol.Module)
}

// Configure is called to create each of the configured consumer modules and call their Configure funcs to validate
// their individual configurations and set them up. If there are any problems, it is expected that these funcs will
// panic with a descriptive error message, as configuration failures are not recoverable errors.
func (cc *Coordinator) Configure() { _ = "STUB: not implemented"; return }

// Create all configured cluster modules, add to list of clusters

// Start calls each of the configured consumer modules' underlying Start funcs. As the coordinator itself has no ongoing
// work to do, it does not start any other goroutines. If any module Start returns an error, this func stops immediately
// and returns that error to the caller. No further modules will be loaded after that.
func (cc *Coordinator) Start() error { _ = "STUB: not implemented"; return nil }

// Start Consumer modules

// All consumers started, Burrow is ready to serve requests
// set the readiness probe

// Stop calls each of the configured consumer modules' underlying Stop funcs. It is expected that the module Stop will
// not return until the module has been completely stopped. While an error can be returned, this func always returns no
// error, as a failure during stopping is not a critical failure
func (cc *Coordinator) Stop() error { _ = "STUB: not implemented"; return nil }

// The individual consumer modules can choose whether or not to implement a wait in the Stop routine
