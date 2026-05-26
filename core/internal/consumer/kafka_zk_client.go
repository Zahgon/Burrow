// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package consumer

import (
	"regexp"
	"sync"
	"time"

	"github.com/linkedin/go-zk"
	"go.uber.org/zap"

	"github.com/linkedin/Burrow/core/protocol"
)

type topicList struct {
	topics map[string]*partitionCount
	lock   *sync.RWMutex
}
type partitionCount struct {
	count int32
	lock  *sync.Mutex
}

// KafkaZkClient is a consumer module which connects to the Zookeeper ensemble where an Apache Kafka cluster maintains
// metadata, and reads consumer group information from the /consumers tree (older ZK-based consumers). It uses watches
// to monitor every group and offset, and the information is forwarded to the storage subsystem for use in evaluations.
type KafkaZkClient struct {
	// App is a pointer to the application context. This stores the channel to the storage subsystem
	App *protocol.ApplicationContext

	// Log is a logger that has been configured for this module to use. Normally, this means it has been set up with
	// fields that are appropriate to identify this coordinator
	Log *zap.Logger

	name             string
	cluster          string
	servers          []string
	zookeeperTimeout int
	zookeeperPath    string

	zk             protocol.ZookeeperClient
	areWatchesSet  bool
	running        *sync.WaitGroup
	groupLock      *sync.RWMutex
	groupList      map[string]*topicList
	groupAllowlist *regexp.Regexp
	groupDenylist  *regexp.Regexp
	connectFunc    func([]string, time.Duration, *zap.Logger) (protocol.ZookeeperClient, <-chan zk.Event, error)
}

// Configure validates the configuration for the consumer. At minimum, there must be a cluster name to which these
// consumers belong, as well as a list of servers provided for the Zookeeper ensemble, of the form host:port. If not
// explicitly configured, it is assumed that the Kafka cluster metadata is present in the ensemble root path. If the
// cluster name is unknown, or if the server list is missing or invalid, this func will panic.
func (module *KafkaZkClient) Configure(name, configRoot string) { _ = "STUB: not implemented"; return }

// Set defaults for configs if needed, and get them

// Check for disallowed config values

// Start connects to the Zookeeper ensemble configured. Any error connecting to the cluster is returned to the caller.
// Once the client is set up, the consumer group list is enumerated and watches are set up for each group, topic,
// partition, and offset. A goroutine is also started to monitor the Zookeeper connection state, and reset the watches
// in the case the the session expires.
func (module *KafkaZkClient) Start() error { _ = "STUB: not implemented"; return nil }

// Set up all groups initially (we can't count on catching the first CONNECTED event

// Start up a func to watch for connection state changes and reset all the watches when needed

// Stop closes the Zookeeper client.
func (module *KafkaZkClient) Stop() error { _ = "STUB: not implemented"; return nil }

// Closing the ZK client will invalidate all the watches, which will close all the running goroutines

func (module *KafkaZkClient) connectionStateWatcher(eventChan <-chan zk.Event) {
	_ = "STUB: not implemented"
	return
}

func (module *KafkaZkClient) acceptConsumerGroup(group string) bool {
	_ = "STUB: not implemented"
	return false
}

// This is a simple goroutine that will wait for an event on a watch channnel and then exit. It's here so that when
// we set a watch that we don't care about (from an ExistsW on a node that already exists), we can drain it properly.
func drainEventChannel(eventChan <-chan zk.Event) { _ = "STUB: not implemented"; return }

func (module *KafkaZkClient) waitForNodeToExist(zkPath string, logger *zap.Logger) bool {
	_ = "STUB: not implemented"
	return false
}

// This is a real error (since NoNode will not return an error)

// The node already exists, just drain the data watch that got created whenever it fires

// Wait for the node to exist

// Watch is gone, so we're gone too

func (module *KafkaZkClient) watchGroupList(eventChan <-chan zk.Event) {
	_ = "STUB: not implemented"
	return
}

// We're done here

func (module *KafkaZkClient) resetGroupListWatchAndAdd(resetOnly bool) {
	_ = "STUB: not implemented"
	return
}

// Get the current group list and reset our watch

// Can't read the consumers path. Bail for now

// Check for any new groups and create the watches for them

func (module *KafkaZkClient) watchTopicList(group string, eventChan <-chan zk.Event) {
	_ = "STUB: not implemented"
	return
}

// We're done here

func (module *KafkaZkClient) resetTopicListWatchAndAdd(group string, resetOnly bool) {
	_ = "STUB: not implemented"
	return
}

// Wait for the offsets znode for this group to exist. We need to do this because the previous child watch
// fires on /consumers/(group) existing, but here we try to read /consumers/(group)/offsets (which might not exist
// yet)

// There was an error checking node existence, so we can't continue

// Get the current group topic list and reset our watch

// Check for any new topics and create the watches for them

func (module *KafkaZkClient) watchPartitionList(group, topic string, eventChan <-chan zk.Event) {
	_ = "STUB: not implemented"
	return
}

// We're done here

func (module *KafkaZkClient) resetPartitionListWatchAndAdd(group, topic string, resetOnly bool) {
	_ = "STUB: not implemented"
	return
}

// Get the current topic partition list and reset our watch

// Can't read the partition list path. Bail for now

// Check for any new partitions and create the watches for them

func (module *KafkaZkClient) watchOffset(group, topic string, partition int32, eventChan <-chan zk.Event) {
	_ = "STUB: not implemented"
	return
}

// We're done here

func (module *KafkaZkClient) resetOffsetWatchAndSend(group, topic string, partition int32, resetOnly bool) {
	_ = "STUB: not implemented"
	return
}

// Get the current offset and reset our watch

// Get the current owner of the partition
// nolint:dogsled

// Can't read the partition offset path. Bail for now

// Badly formatted offset

// Send the offset to the storage module

// Send the owner of the current partition to the storage module
