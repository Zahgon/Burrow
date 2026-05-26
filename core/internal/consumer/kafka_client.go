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
	"bytes"
	"regexp"
	"sync"

	"github.com/IBM/sarama"
	"go.uber.org/zap"

	"github.com/linkedin/Burrow/core/internal/helpers"
	"github.com/linkedin/Burrow/core/protocol"
)

// KafkaClient is a consumer module which connects to a single Apache Kafka cluster and reads consumer group information
// from the offsets topic in the cluster, which is typically __consumer_offsets. The messages in this topic are decoded
// and the information is forwarded to the storage subsystem for use in evaluations.
type KafkaClient struct {
	// App is a pointer to the application context. This stores the channel to the storage subsystem
	App *protocol.ApplicationContext

	// Log is a logger that has been configured for this module to use. Normally, this means it has been set up with
	// fields that are appropriate to identify this coordinator
	Log *zap.Logger

	name                  string
	cluster               string
	servers               []string
	offsetsTopic          string
	startLatest           bool
	backfillEarliest      bool
	reportedConsumerGroup string
	saramaConfig          *sarama.Config
	groupAllowlist        *regexp.Regexp
	groupDenylist         *regexp.Regexp

	quitChannel chan struct{}
	running     sync.WaitGroup
}

type offsetKey struct {
	Group     string
	Topic     string
	Partition int32
	ErrorAt   string
}
type offsetValue struct {
	Offset    int64
	Timestamp int64
	ErrorAt   string
}
type metadataHeader struct {
	ProtocolType          string
	Generation            int32
	Protocol              string
	Leader                string
	CurrentStateTimestamp int64
}
type metadataMember struct {
	MemberID         string
	GroupInstanceID  string
	ClientID         string
	ClientHost       string
	RebalanceTimeout int32
	SessionTimeout   int32
	Assignment       map[string][]int32
}
type backfillEndOffset struct {
	Value int64
}

// Configure validates the configuration for the consumer. At minimum, there must be a cluster name to which these
// consumers belong, as well as a list of servers provided for the Kafka cluster, of the form host:port. If not
// explicitly configured, the offsets topic is set to the default for Kafka, which is __consumer_offsets. If the
// cluster name is unknown, or if the server list is missing or invalid, this func will panic.
func (module *KafkaClient) Configure(name, configRoot string) { _ = "STUB: not implemented"; return }

// Set defaults for configs if needed, and get them

// Check for disallowed config values

// Start connects to the Kafka cluster using the Shopify/sarama client. Any error connecting to the cluster is returned
// to the caller. Once the client is set up, the consumers for the configured offsets topic are started.
func (module *KafkaClient) Start() error { _ = "STUB: not implemented"; return nil }

// Connect Kafka client

// Start the consumers

// Stop closes the goroutines that listen to the client consumer.
func (module *KafkaClient) Stop() error { _ = "STUB: not implemented"; return nil }

func (module *KafkaClient) startBackfillPartitionConsumer(partition int32, client helpers.SaramaClient, consumer sarama.Consumer) error {
	_ = "STUB: not implemented"
	return nil
}

// We check for an empty partition after building the consumer, otherwise we
// could be unlucky enough to observe a nonempty partition
// whose only segment expires right after we check.

// GetOffset returns the next (not yet published) offset, but we want the latest published offset.

func (module *KafkaClient) partitionConsumer(consumer sarama.PartitionConsumer, stopAtOffset *backfillEndOffset) {
	_ = "STUB: not implemented"
	return
}

// emulating a consumer which should commit (lastSeenOffset+1)

func (module *KafkaClient) startKafkaConsumer(client helpers.SaramaClient) error {
	_ = "STUB: not implemented"
	// Create the consumer from the client
	return nil
}

// Get a partition count for the consumption topic

// Default to bootstrapping the offsets topic, unless configured otherwise

// Start consumers for each partition with fan in

// Note: since we are consuming each partition twice,
// we need a second consumer instance

func (module *KafkaClient) processConsumerOffsetsMessage(msg *sarama.ConsumerMessage) {
	_ = "STUB: not implemented"
	return
}

func readString(buf *bytes.Buffer) (string, error) {
	_ = "STUB: not implemented" // nolint:interfacer
	return "", nil
}

func readCompactString(buf *bytes.Buffer) (string, error) {
	_ = "STUB: not implemented" // nolint:interfacer
	return "", nil
}

// For compact strings, length 0 means null, length 1 means empty string

// null or empty string

// For actual content, we need to read (length - 1) bytes

func (module *KafkaClient) acceptConsumerGroup(group string) bool {
	_ = "STUB: not implemented"
	return false
}

func (module *KafkaClient) decodeKeyAndOffset(offsetOrder int64, keyBuffer *bytes.Buffer, value []byte, logger *zap.Logger) {
	_ = "STUB: not implemented"
	// Version 0 and 1 keys are decoded the same way
	return
}

// Tombstone message - we don't handle them for now

func (module *KafkaClient) decodeAndSendOffset(offsetOrder int64, offsetKey offsetKey, valueBuffer *bytes.Buffer, logger *zap.Logger, decoder func(*bytes.Buffer) (offsetValue, string)) {
	_ = "STUB: not implemented"
	return
}

func (module *KafkaClient) decodeGroupMetadata(keyBuffer *bytes.Buffer, value []byte, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Tombstone message - group deleted

func (module *KafkaClient) decodeAndSendGroupMetadata(valueVersion int16, group string, valueBuffer *bytes.Buffer, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// If memberCount is zero, clear all ownership

func decodeMetadataValueHeader(buf *bytes.Buffer) (metadataHeader, string) {
	_ = "STUB: not implemented"
	return *new(metadataHeader), ""
}

func decodeMetadataValueHeaderV2(buf *bytes.Buffer) (metadataHeader, string) {
	_ = "STUB: not implemented"
	return *new(metadataHeader), ""
}

func decodeGroupInstanceID(buf *bytes.Buffer, memberVersion int16) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func decodeMetadataMember(buf *bytes.Buffer, memberVersion int16) (metadataMember, string) {
	_ = "STUB: not implemented"
	return *new(metadataMember), ""
}

func decodeMemberAssignmentV0(buf *bytes.Buffer) (map[string][]int32, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func decodeOffsetKeyV0(buf *bytes.Buffer) (offsetKey, string) {
	_ = "STUB: not implemented"
	return *new(offsetKey), ""
}

// Error constants for offset value decoding
const (
	errOffset      = "offset"
	errLeaderEpoch = "leaderEpoch"
	errMetadata    = "metadata"
	errTimestamp   = "timestamp"
)

func decodeOffsetValueV0(valueBuffer *bytes.Buffer) (offsetValue, string) {
	_ = "STUB: not implemented"
	return *new(offsetValue), ""
}

func decodeOffsetValueV3(valueBuffer *bytes.Buffer) (offsetValue, string) {
	_ = "STUB: not implemented"
	return *new(offsetValue), ""
}

func decodeOffsetValueV4(valueBuffer *bytes.Buffer) (offsetValue, string) {
	_ = "STUB: not implemented"
	return *new(offsetValue), ""
}

// Read offset (8 bytes)

// Read leader epoch (4 bytes)

// V4 uses compact strings for metadata

// Read timestamp (8 bytes)

// Skip tagged fields as we only need offset and timestamp
