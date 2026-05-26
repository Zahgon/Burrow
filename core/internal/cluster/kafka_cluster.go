// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package cluster

import (
	"sync"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"

	"github.com/linkedin/Burrow/core/internal/helpers"
	"github.com/linkedin/Burrow/core/protocol"
)

// KafkaCluster is a cluster module which connects to a single Apache Kafka cluster and manages the broker topic and
// partition information. It periodically updates a list of all topics and partitions, and also fetches the broker
// end offset (latest) for each partition. This information is forwarded to the storage module for use in consumer
// evaluations.
type KafkaCluster struct {
	// App is a pointer to the application context. This stores the channel to the storage subsystem
	App *protocol.ApplicationContext

	// Log is a logger that has been configured for this module to use. Normally, this means it has been set up with
	// fields that are appropriate to identify this coordinator
	Log *zap.Logger

	name                string
	saramaConfig        *sarama.Config
	servers             []string
	offsetRefresh       int
	topicRefresh        int
	groupsReaperRefresh int

	offsetTicker       *time.Ticker
	metadataTicker     *time.Ticker
	groupsReaperTicker *time.Ticker
	quitChannel        chan struct{}
	running            sync.WaitGroup

	fetchMetadata   bool
	topicPartitions map[string][]int32
}

// Configure validates the configuration for the cluster. At minimum, there must be a list of servers provided for the
// Kafka cluster, of the form host:port. Default values will be set for the intervals to use for refreshing offsets
// (10 seconds) and topics (60 seconds). A missing, or bad, list of servers will cause this func to panic.
func (module *KafkaCluster) Configure(name, configRoot string) { _ = "STUB: not implemented"; return }

// Set defaults for configs if needed

// Start connects to the Kafka cluster using the Shopify/sarama client. Any error connecting to the cluster is returned
// to the caller. Once the client is set up, tickers are started to periodically refresh topics and offsets.
func (module *KafkaCluster) Start() error { _ = "STUB: not implemented"; return nil }

// Connect Kafka client

// Fire off the offset requests once, before we start the ticker, to make sure we start with good data for consumers

// Start main loop that has a timer for offset and topic fetches

// just start and stop a new ticker, the channel will still be active but will not emit ticks
// it'll simplify tick management in the mainLoop func

// Stop causes both the topic and offset refresh tickers to be stopped, and then it closes the Kafka client.
func (module *KafkaCluster) Stop() error { _ = "STUB: not implemented"; return nil }

func (module *KafkaCluster) mainLoop(client helpers.SaramaClient) {
	_ = "STUB: not implemented"
	return
}

// Update metadata on next offset fetch

func (module *KafkaCluster) maybeUpdateMetadataAndDeleteTopics(client helpers.SaramaClient) {
	_ = "STUB: not implemented"
	return
}

// Get the current list of topics and make a map

// We'll use topicPartitions later

// partitionID has a leader
// NOTE: append only happens here
// so cap(topicPartitions[topic]) is the partition count

// Check for deleted topics if we have a previous map to check against

// Topic no longer exists - tell storage to delete it

// Save the new topicPartitions for next time

func (module *KafkaCluster) generateOffsetRequests(client helpers.SaramaClient) (map[int32]*sarama.OffsetRequest, map[int32]helpers.SaramaBroker) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generate an OffsetRequest for each topic:partition and bucket it to the leader broker

// Match the version of the client as sarama's getOffset function does
// https://github.com/IBM/sarama/blob/main/client.go#L863-L876

// Version 4 adds the current leader epoch, which is used for fencing.

// Version 3 is the same as version 2.

// Version 2 adds the isolation level, which is used for transactional reads.

// Version 1 removes MaxNumOffsets.  From this version forward, only a single
// offset can be returned.

// This function performs massively parallel OffsetRequests, which is better than Sarama's internal implementation,
// which does one at a time. Several orders of magnitude faster.
func (module *KafkaCluster) getOffsets(client helpers.SaramaClient) {
	_ = "STUB: not implemented"
	return
}

// Send out the OffsetRequest to each broker for all the partitions it is leader for
// The results go to the offset storage module

// Gather a list of topics that had errors

// If there are any topics that had errors, force a metadata refresh on the next run

func (module *KafkaCluster) reapNonExistingGroups(client helpers.SaramaClient) {
	_ = "STUB: not implemented"
	return
}

// TODO: find how to get reportedConsumerGroup from KafkaClient
