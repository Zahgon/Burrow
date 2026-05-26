// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

// Package zookeeper - Common Zookeeper subsystem.
// The zookeeper subsystem provides a Zookeeper client that is common across all of Burrow, and can be used by other
// subsystems to store metadata or coordinate operations between multiple Burrow instances. It is used primarily to
// assure that only one Burrow instance is sending notifications at any time.
package zookeeper

import (
	"sync"
	"time"

	"github.com/linkedin/go-zk"
	"go.uber.org/zap"

	"github.com/linkedin/Burrow/core/protocol"
)

// Coordinator (zookeeper) manages a single Zookeeper connection for other coordinators and modules to make use of in
// order to store metadata for Burrow itself. This is not required to connect to the same Zookeeper ensemble as any
// specific Kafka cluster. The ZookeeperClient is stored in the application context, as well as the root path that
// any modules should create their metadata underneath.
//
// The coordinator monitors the connection state transitions and signals when the session is expired, and then when it
// reconnects. Code that must be aware of session expirations, such as code that makes use of watches, should have a
// structure as in the example.
type Coordinator struct {
	App *protocol.ApplicationContext
	Log *zap.Logger

	servers     []string
	connectFunc func([]string, time.Duration, *zap.Logger) (protocol.ZookeeperClient, <-chan zk.Event, error)
	running     sync.WaitGroup
}

// Configure validates that the configuration has a list of servers provided for the Zookeeper ensemble, of the form
// host:port. It also checks the provided root path, using a default of "/burrow" if none has been provided.
func (zc *Coordinator) Configure() { _ = "STUB: not implemented"; return }

// if zookeeper.tls has been set, use the TLS connect function otherwise use default connect

// Set and check configs

// Start creates the connection to the Zookeeper ensemble, and assures that the root path exists. Once that is done,
// it sets the ZookeeperConnected flag in the application context to true, and creates the ZookeeperExpired condition
// flag. It then starts a main loop to watch for connection state changes.
func (zc *Coordinator) Start() error { _ = "STUB: not implemented"; return nil }

// This ZK client will be shared by other parts of Burrow for things like locks
// NOTE - samuel/go-zookeeper does not support chroot, so we pass along the configured root path in config

// Assure that our root path exists

// Stop closes the connection to the Zookeeper ensemble and waits for the connection state monitor to exit (which it
// will because the event channel will be closed).
func (zc *Coordinator) Stop() error { _ = "STUB: not implemented"; return nil }

// This will close the event channel, closing the mainLoop

func (zc *Coordinator) createRecursive(path string) error { _ = "STUB: not implemented"; return nil }

// If the rootpath exists, skip the Create process to avoid "zk: not authenticated" error

// Ignore when the node exists already

func (zc *Coordinator) mainLoop(eventChan <-chan zk.Event) { _ = "STUB: not implemented"; return }
