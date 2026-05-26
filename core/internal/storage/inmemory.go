// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package storage

import (
	"container/ring"
	"regexp"
	"sync"

	"go.uber.org/zap"

	"github.com/linkedin/Burrow/core/protocol"
)

// InMemoryStorage is a storage module that maintains the entire data set in memory in a series of maps. It has a
// configurable number of worker goroutines to service requests, and for requests that are group-specific, the group
// and cluster name are used to hash the request to a consistent worker. This assures that requests for a group are
// processed in order.
type InMemoryStorage struct {
	// App is a pointer to the application context. This stores the channel to the storage subsystem
	App *protocol.ApplicationContext

	// Log is a logger that has been configured for this module to use. Normally, this means it has been set up with
	// fields that are appropriate to identify this coordinator
	Log *zap.Logger

	name        string
	intervals   int
	numWorkers  int
	expireGroup int64
	minDistance int64
	queueDepth  int

	requestChannel chan *protocol.StorageRequest
	workersRunning sync.WaitGroup
	mainRunning    sync.WaitGroup
	offsets        map[string]clusterOffsets
	groupAllowlist *regexp.Regexp
	groupDenylist  *regexp.Regexp
	workers        []chan *protocol.StorageRequest
}

type brokerOffset struct {
	Offset    int64
	Timestamp int64
}

type consumerPartition struct {
	offsets  *ring.Ring
	owner    string
	clientID string
}

type consumerGroup struct {
	// This lock is held when using the individual group, either for read or write
	lock       *sync.RWMutex
	topics     map[string][]*consumerPartition
	lastCommit int64
}

type clusterOffsets struct {
	broker   map[string][]*ring.Ring
	consumer map[string]*consumerGroup

	// This lock is used when modifying broker topics or offsets
	brokerLock *sync.RWMutex

	// This lock is used when modifying the overall consumer list
	// It does not need to be held for modifying an individual group
	consumerLock *sync.RWMutex
}

// Represents the destination of adding an offset into
// the consumer offsets ring buffer.
// `insertDest`: the destination for an insert
// `extendDest`: the next slot in the ring, which will be overridden
//
// If only `extendDest` is set, we're appending to the ring.
// If only `insertDest` is set, we're replacing an existing slot.
// If both are set, we're inserting before some existing items,
// i.e. shifting all items past `insertDest` forward by one,
type offsetRingDestination struct {
	insertDest *ring.Ring
	extendDest *ring.Ring
}

func (destination *offsetRingDestination) destinationSlot() *ring.Ring {
	_ = "STUB: not implemented"
	return nil
}

func (destination *offsetRingDestination) isAppend() bool { _ = "STUB: not implemented"; return false }

func (destination *offsetRingDestination) isShift() bool { _ = "STUB: not implemented"; return false }

// Configure validates the configuration for the module, creates a channel to receive requests on, and sets up the
// storage map. If no expiration time for groups is set, a default value of 7 days is used. If no interval count is
// set, a default of 10 intervals is used. If no worker count is set, a default of 20 workers is used.
func (module *InMemoryStorage) Configure(name, configRoot string) {
	_ = "STUB: not implemented"
	return
}

// Set defaults for configs if needed

// Check for disallowed config values

// GetCommunicationChannel returns the RequestChannel that has been setup for this module.
func (module *InMemoryStorage) GetCommunicationChannel() chan *protocol.StorageRequest {
	_ = "STUB: not implemented"
	return nil
}

// Start sets up the rest of the storage map for each configured cluster. It then starts the configured number of
// worker routines to handle requests. Finally, it starts a main loop which will receive requests and hash them to the
// correct worker.
func (module *InMemoryStorage) Start() error { _ = "STUB: not implemented"; return nil }

// Start the appropriate number of workers, with a channel for each

// Stop closes the incoming request channel, which will close the main loop. It then closes each of the worker
// channels, to close the workers, and waits for all goroutines to exit before returning.
func (module *InMemoryStorage) Stop() error { _ = "STUB: not implemented"; return nil }

func (module *InMemoryStorage) requestWorker(workerNum int, requestChannel chan *protocol.StorageRequest) {
	_ = "STUB: not implemented"
	return
}

// Using a map for the request types avoids a bit of complexity below

func (module *InMemoryStorage) mainLoop() { _ = "STUB: not implemented"; return }

// Send to any worker

// Hash to a consistent worker

func (module *InMemoryStorage) addBrokerOffset(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Ignore offsets for clusters that we don't know about - should never happen anyways

// The partition count has increased. Append enough extra partitions, with offset rings, to our slice

// Advance to the next ring entry (this means the pointer is always at the most recent entry, rather than the
// oldest entry)

func (module *InMemoryStorage) getBrokerOffset(clusterMap *clusterOffsets, topic string, partition int32, requestLogger *zap.Logger) (int64, int32) {
	_ = "STUB: not implemented"
	return 0, 0
}

// We don't know about this topic from the brokers yet - skip consumer offsets for now

// This should never happen, but if it does, log an warning with the offset information for review

// We know about the topic, but partitions have been expanded and we haven't seen that from the broker yet

// We know about the topic and partition, but we haven't actually gotten the broker offset yet

func (module *InMemoryStorage) getConsumerPartition(consumerMap *consumerGroup, topic string, partition, partitionCount int32, requestLogger *zap.Logger) *consumerPartition {
	_ = "STUB: not implemented"
	// Get or create the topic for the consumer
	return nil
}

// Get the partition specified

// The partition count must have increased. Append enough extra partitions to our slice

// Get or create the offsets ring for this partition

func (module *InMemoryStorage) acceptConsumerGroup(group string) bool {
	_ = "STUB: not implemented"
	return false
}

func (module *InMemoryStorage) addConsumerOffset(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Ignore offsets for clusters that we don't know about - should never happen anyways

// Get the broker offset for this partition, as well as the partition count

// If the returned partitionCount is zero, there was an error that was already logged. Just stop processing

// Make the consumer group if it does not yet exist

// For the rest of this, we need the write lock for the consumer group

// Get the offset ring for this partition - it always points to the earliest offset (or where to insert a new value)

// Calculate the lag against the brokerOffset

// Little bit of a hack - because we only get broker offsets periodically, it's possible the consumer offset could be ahead of where we think the broker
// is. If the broker appears behind, just treat it as zero lag.

// Given a consumer offset ring and a storage request, find the destination
// based on the order of commits.
func findConsumerOffsetDestination(offsetRing *ring.Ring, request *protocol.StorageRequest, requestLogger *zap.Logger) *offsetRingDestination {
	_ = "STUB: not implemented"
	return nil
}

// nonempty ring, prevSlot is the latest offset

// full ring, offsetRing is the oldest

// the request is not the newest, so we need to insert it somewhere

// Reached a blank slot or the oldest slot in the ring, just replace it

// Lookback one previous commit

// Insert here

// else Request is older than prevSlot; keep searching

// If we haven't returned by now we are just appending

// If the offset commit is faster than we are allowing (less than the min-distance config), replace the previous
// commit with this one. This lets us store the new offset commit without dropping an old one
func (module *InMemoryStorage) mergeFrequentCommitIntoPrevious(destination *offsetRingDestination, request *protocol.StorageRequest, requestLogger *zap.Logger) *offsetRingDestination {
	_ = "STUB: not implemented"
	return nil
}

// We also set the timestamp for the request to the previous timestamp. The reason for this is that if we
// update the timestamp to the new timestamp, we may never create a new offset in the ring (consider the
// case where someone is committing with a frequency lower than min-distance)

func (module *InMemoryStorage) storeConsumerOffset(consumerPartition *consumerPartition, destination *offsetRingDestination, request *protocol.StorageRequest, partitionLag *protocol.Lag) {
	_ = "STUB: not implemented"
	return
}

// Shift each item past destination.insertDest forward (so we end up occupying extendDest)

// Write into the destination

// We've extended the ring by either appending or shifting, update the ring pointer

func (module *InMemoryStorage) addConsumerOwner(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Ignore offsets for clusters that we don't know about - should never happen anyways

// Make the consumer group if it does not yet exist

// Get the partition count for this partition (we don't need the actual broker offset)

// If the returned partitionCount is zero, there was an error that was already logged. Just stop processing

// For the rest of this, we need the write lock for the consumer group

// Get the consumer partition state for this partition - we don't need it, but it will properly create the topic and partitions for us

// Write the owner for the given topic/partition

func (module *InMemoryStorage) clearConsumerOwners(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Ignore metadata for clusters that we don't know about - should never happen anyways

// Make the consumer group if it does not yet exist

// Consumer group doesn't exist, so we can't clear owners for it

// For the rest of this, we need the write lock for the consumer group

func (module *InMemoryStorage) deleteTopic(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Work backwards - remove the topic from consumer groups first

// No need to check for existence

// Now remove the topic from the broker list

func (module *InMemoryStorage) deleteGroup(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// The consumer group consumes other topics, thus we need to keep its metrics

// only a specific topic was deleted and the consumer group still exists, thus we delete only a subset of the metrics

func (module *InMemoryStorage) fetchClusterList(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

func (module *InMemoryStorage) fetchTopicList(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

func (module *InMemoryStorage) fetchConsumerList(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

func (module *InMemoryStorage) fetchTopic(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

func getConsumerTopicList(consumerMap *consumerGroup) protocol.ConsumerTopics {
	_ = "STUB: not implemented"
	return *new(protocol.ConsumerTopics)
}

// Make a copy so that we can release the lock and be safe

func (module *InMemoryStorage) fetchConsumer(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Lazily purge consumers that haven't committed in longer than the defined interval. Return as a 404

// Swap for a write lock

// Calculate the current lag for each now. We do this separate from getting the consumer info so we can avoid
// locking both the consumers and the brokers at the same time

// The topic may have just been deleted, so we'll skip this part and just return the consumer data we have

// Build the slice of broker offsets to return

// nolint:scopelint

// Little bit of a hack - because we only get broker offsets periodically, it's possible the consumer offset could be ahead of where we think the broker
// is. In this case, just mark it as zero lag.

func (module *InMemoryStorage) fetchConsumersForTopicList(request *protocol.StorageRequest, requestLogger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}
